package pcfsmrpcr

import (
	//"encoding/binary"
	//"net"

	//"sync"

	"context"
	"time"

	//	"fmt"
	"k8s.io/klog"
	// "github.com/gin-gonic/gin"
	"github.com/benbjohnson/clock"

	"w5gc.io/wipro5gcore/pkg/pcf/pcfsmrpcr/api"
	"w5gc.io/wipro5gcore/pkg/pcf/pcfsmrpcr/apiclient"

	// "w5gc.io/wipro5gcore/pkg/pcf/pcfsmrpcr/sm/nodes"

	"w5gc.io/wipro5gcore/pkg/pcf/pcfsmrpcr/config"
	"w5gc.io/wipro5gcore/pkg/pcf/pcfsmrpcr/sm/sessions"
	"w5gc.io/wipro5gcore/utils/cache"

	"w5gc.io/wipro5gcore/pkg/pcf/pcfsmrpcr/sm"
)

const (
	retransmitInterval = time.Second * 3
	retransmitRetries  = 3
	backoffInterval    = time.Second * 3
	// Period for performing global cleanup tasks.
	housekeepingPeriod = time.Second * 2
)

// PcfsmrpcrHandler is an interface implemented for testability
type PcfsmrpcrHandler interface {
	//InitiateSessioReportRequest(sessionMsg *pfcpUdp.Message)
	HandleSessionCleanups() error
}

// Bootstrap is a bootstrapping interface for PDU SMS
type PcfsmrpcrBootstrap interface {
	//GetConfiguration()
	//GetContext()
	Run(configChannel <-chan config.PcfsmrpcrConfig)
}

type Pcfsmrpcr struct {
	config          *config.PcfsmrpcrConfig
	sessionManager  sm.SessionManager
	sessionCache    cache.WorkCache
	apiClient       apiclient.ApiClient
	apiServer       api.ApiServer
	sessionWorkers  SessionWorkers
	clock           clock.Clock
	backoffInterval time.Duration
	timerT1         time.Duration
	retriesN1       uint8
	context         *PcfsmrpcrContext
}

type PcfsmrpcrContext struct {
	startTime   time.Time
	lastRestart time.Time
	restarts    int // If number of restarts in last 10 sec > 3 reset TODO GURU
	//lock                  sync.Mutex
	NodeId string
}

// Initialize Pcfsmrpcr
func NewPcfsmrpcr(cfg *config.PcfsmrpcrConfig, time time.Time) (PcfsmrpcrBootstrap, bool) {

	pcfsmrpcr := &Pcfsmrpcr{
		config:          cfg,
		clock:           clock.New(),
		backoffInterval: backoffInterval,
		timerT1:         retransmitInterval,
		retriesN1:       retransmitRetries,
		context: &PcfsmrpcrContext{
			startTime: time,
		},
	}

	// Intialize the api handler
	//should both apiClient and apiServer be created in newcsc?
	pcfsmrpcr.apiClient = apiclient.NewAPIClient(cfg)

	pcfsmrpcr.apiServer = api.NewApiServer(cfg.NodeInfo)

	// Initialize session manager
	//pdusmsp.sessionManager = sm.NewSessionManager(pdusmsp.config.NodeInfo, pdusmsp.config.n11Nodes, time, pdusmsp.backoffInterval, pdusmsp.timerT1, pdusmsp.retriesN1)

	// s:=(pdusmsp.apiClient).(apiclient.ApiClientInfo)
	// var s sm.SessionManager
	// pdusmsp.sessionManager = sm.NewSessionManager(dbInfo, &pdusmsp.grpc, pdusmsp.apiClient)
	// pdusmsp.sessionManager = sm.NewSessionManager(pdusmsp.apiClient)
	pcfsmrpcr.sessionManager = sm.NewSessionManager(pcfsmrpcr.apiClient)
	//	s:= pdusmsp.sessionManager.NewSMContextAPIService()
	//	fmt.Println(s)
	// Intialize session cache
	pcfsmrpcr.sessionCache = cache.NewCache(pcfsmrpcr.clock)

	// Intialize session workers
	pcfsmrpcr.sessionWorkers = NewSessionWorkers(pcfsmrpcr.handleSession, pcfsmrpcr.sessionCache, pcfsmrpcr.backoffInterval)

	// resyncInterval, backOffPeriod TODO GURU

	return pcfsmrpcr, true
}

// Run starts the Pcfsmrpcr
func (p *Pcfsmrpcr) Run(configChannel <-chan config.PcfsmrpcrConfig) {
	// start the session manager
	go p.sessionManager.Start()

	// Start the api handler
	p.apiClient.Start()
	go p.apiServer.Start()

	// Start the pdusmsp event handler
	p.pcfsmrpcrEvents(configChannel, p)

	// return
}

func (p *Pcfsmrpcr) pcfsmrpcrEvents(configChannel <-chan config.PcfsmrpcrConfig, handler PcfsmrpcrHandler) {
	klog.Infof("Entered into pcfsmrpcrEvents")
	syncTicker := time.NewTicker(time.Second)
	defer syncTicker.Stop()
	housekeepingTicker := time.NewTicker(housekeepingPeriod)
	defer housekeepingTicker.Stop()
	sessionChannel := p.apiServer.WatchApiChannel()
	p.handlePcfsmrpcrEvents(configChannel, sessionChannel, syncTicker.C, housekeepingTicker.C, handler)
}

// handlePcfsmrpcrEvents is the main loop for processing events in pdusmsp
func (p *Pcfsmrpcr) handlePcfsmrpcrEvents(configChannel <-chan config.PcfsmrpcrConfig, sessionChannel <-chan *api.SessionMessage,
	syncCh <-chan time.Time, housekeepingCh <-chan time.Time,
	handler PcfsmrpcrHandler) bool {
	klog.Info("Entered into handlePcfsmrpcrEvents")
	for {
		select {

		case pcfsmrpcrMsg := <-sessionChannel:
			switch pcfsmrpcrMsg.MsgType {

			case sm.NPCF_POST_SM_CONTEXT_REQUEST:
				// PDU Session management service - Retrieve SM Context Request
				klog.Infof("handlePcfsmrpcrEvents (RETRIEVE SM CONTEXT REQUEST)")

				sessionId := sessions.SessionId(pcfsmrpcrMsg.SmPolicyId)
				// p.dispatchWork(sessionId, pdusmsMsg, grpcMsg, time.Now())
				p.dispatchWork(
					sessionId,
					pcfsmrpcrMsg,
					pcfsmrpcrMsg.MsgType,
					time.Now())

			case sm.NPCF_UPDATE_SM_CONTEXT_REQUEST:
				// PDU Session management service - Update SM Context Request
				klog.Infof("handlePdusmspEvents (UPDATE SM CONTEXT  REQUEST)")
				// fmt.Println(reflect.TypeOf(pdusmsMsg.SessionMsg))
				//have to change in future
				// id, err := strconv.Atoi(pdusmsMsg.SmContextRefID)
				// if err != nil {
				// 	klog.Error(err.Error())
				// }
				// klog.Info(id)
				sessionId := sessions.SessionId(pcfsmrpcrMsg.SmPolicyId)
				// p.dispatchWork(sessionId, pdusmsMsg, grpcMsg, time.Now())
				p.dispatchWork(
					sessionId,
					pcfsmrpcrMsg,
					pcfsmrpcrMsg.MsgType,
					time.Now())

			case sm.NPCF_DELETE_SM_CONTEXT_REQUEST:
				// PDU Session management service - Release SM Context Request
				klog.Infof("handlePdusmspEvents (RELEASE SM CONTEXT REQUEST)")
				// id, err := strconv.Atoi(pdusmsMsg.SmContextRefID)
				// if err != nil {
				// 	klog.Error(err.Error())
				// }
				sessionId := sessions.SessionId(pcfsmrpcrMsg.SmPolicyId)
				// p.dispatchWork(sessionId, pdusmsMsg, grpcMsg, time.Now())
				p.dispatchWork(
					sessionId,
					pcfsmrpcrMsg,
					pcfsmrpcrMsg.MsgType,
					time.Now())

			case sm.NPCF_RETRIEVE_SM_CONTEXT_REQUEST:
				// PDU Session management service - Retrieve SM Context Request
				klog.Infof("handlePdusmspEvents (RETRIEVE SM CONTEXT REQUEST)")
				// id, err := strconv.Atoi(pdusmsMsg.SmContextRefID)
				// if err != nil {
				// 	klog.Error(err.Error())
				// }
				sessionId := sessions.SessionId(pcfsmrpcrMsg.SmPolicyId)
				// p.dispatchWork(sessionId, pdusmsMsg, grpcMsg, time.Now())
				p.dispatchWork(
					sessionId,
					pcfsmrpcrMsg,
					pcfsmrpcrMsg.MsgType,
					time.Now())
			}

		case <-housekeepingCh:
			klog.V(4).Infof("SyncLoop (housekeeping)")
			if err := handler.HandleSessionCleanups(); err != nil {
				klog.Errorf("Failed cleaning session: %v", err)
			}

		}
	}

}

func (p *Pcfsmrpcr) HandleSessionCleanups() error {
	deletedSessions := make(map[sessions.SessionId]struct{})
	err := p.sessionWorkers.RemoveSessionWorkers(deletedSessions)
	return err
}

// dispatchWork handles the session in a session worker
func (p *Pcfsmrpcr) dispatchWork(sessionId sessions.SessionId,
	sessionMsg *api.SessionMessage,
	msgType sm.MessageType,
	startTime time.Time) {

	// Run the handle session in an async worker.
	p.sessionWorkers.HandleSessionMessages(&SessionMessageInfo{
		SessionId:    sessionId,
		StartTime:    startTime,
		PcfsmrpcrMsg: *sessionMsg,
		MsgType:      msgType,
		OnCompleteFunc: func(err error) {
			// Handle on completion of session update
			if err == nil {
				klog.Infof("Successfully handled pdusms  for session %s", sessionId)
				//metrics.SessionWorkerDuration.WithLabelValues(syncType.String()).Observe(metrics.SinceInSeconds(start))
			} else {
				// Log error and update cause with request rejected
				klog.Info(err)
				klog.Errorf("Unable to handle pdusms for session %s", sessionId)
			}
		},
	})

}

// handleSession handles the pdu session message in a session worker
func (p *Pcfsmrpcr) handleSession(msgInfo SessionMessageInfo) error {
	// Get the required message info
	klog.Info("inside")
	msgType := msgInfo.MsgType
	pcfsmrpcrMsg := msgInfo.PcfsmrpcrMsg

	switch msgType {

	case sm.NPCF_POST_SM_CONTEXT_REQUEST:

		ctx := context.Background()
		resp, err := p.sessionManager.SmPoliciesPost(ctx, pcfsmrpcrMsg.SmPolicyContextData)
		chanRec := p.apiServer.WatchRecChannel()
		chanRec <- &api.Receiver{ReceivedResponse: resp, ReceivedErr: err}
		return err

	case sm.NPCF_UPDATE_SM_CONTEXT_REQUEST:
		// Process pdusms update request
		ctx := context.Background()

		resp, err := p.sessionManager.SmPoliciesSmPolicyIdUpdatePost(ctx, pcfsmrpcrMsg.SmPolicyId, pcfsmrpcrMsg.SmPolicyUpdateContextData)

		chanRec := p.apiServer.WatchRecChannel()
		chanRec <- &api.Receiver{ReceivedResponse: resp, ReceivedErr: err}
		return err
	case sm.NPCF_DELETE_SM_CONTEXT_REQUEST:
		// Process pdusms release request
		ctx := context.Background()
		resp, err := p.sessionManager.SmPoliciesSmPolicyIdDeletePost(ctx, pcfsmrpcrMsg.SmPolicyId, pcfsmrpcrMsg.SmPolicyDeleteData)
		chanRec := p.apiServer.WatchRecChannel()
		chanRec <- &api.Receiver{ReceivedResponse: resp, ReceivedErr: err}
		return err
	case sm.NPCF_RETRIEVE_SM_CONTEXT_REQUEST:
		// Process pdusms retrieve request
		ctx := context.Background()
		resp, err := p.sessionManager.SmPoliciesSmPolicyIdGet(ctx, pcfsmrpcrMsg.SmPolicyId)
		chanRec := p.apiServer.WatchRecChannel()
		chanRec <- &api.Receiver{ReceivedResponse: resp, ReceivedErr: err}
		return err
	}
	return nil

}
