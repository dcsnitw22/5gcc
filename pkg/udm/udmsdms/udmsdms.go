package udmsdms

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

	"w5gc.io/wipro5gcore/pkg/udm/udmsdms/api"
	"w5gc.io/wipro5gcore/pkg/udm/udmsdms/apiclient"

	// "w5gc.io/wipro5gcore/pkg/udm/udmsdms/sm/nodes"

	// openapiserver "w5gc.io/wipro5gcore/openapi/openapiudmserver"

	"w5gc.io/wipro5gcore/pkg/udm/udmsdms/config"
	"w5gc.io/wipro5gcore/pkg/udm/udmsdms/sm/sessions"
	"w5gc.io/wipro5gcore/utils/cache"

	"w5gc.io/wipro5gcore/pkg/udm/udmsdms/sm"
)

const (
	retransmitInterval = time.Second * 3
	retransmitRetries  = 3
	backoffInterval    = time.Second * 3
	// Period for performing global cleanup tasks.
	housekeepingPeriod = time.Second * 2
)

// UdmsdmsHandler is an interface implemented for testability
type UdmsdmsHandler interface {
	//InitiateSessioReportRequest(sessionMsg *pfcpUdp.Message)
	HandleSessionCleanups() error
}

// Bootstrap is a bootstrapping interface for PDU SMS
type UdmsdmsBootstrap interface {
	//GetConfiguration()
	//GetContext()
	Run(configChannel <-chan config.UdmsdmsConfig)
}

type Udmsdms struct {
	config          *config.UdmsdmsConfig
	sessionManager  sm.SessionManager
	sessionCache    cache.WorkCache
	apiClient       apiclient.ApiClient
	apiServer       api.ApiServer
	sessionWorkers  SessionWorkers
	clock           clock.Clock
	backoffInterval time.Duration
	timerT1         time.Duration
	retriesN1       uint8
	context         *UdmsdmsContext
}

type UdmsdmsContext struct {
	startTime   time.Time
	lastRestart time.Time
	restarts    int // If number of restarts in last 10 sec > 3 reset TODO GURU
	//lock                  sync.Mutex
	NodeId string
}

// Initialize Udmsdms
func NewUdmsdms(cfg *config.UdmsdmsConfig, time time.Time) (UdmsdmsBootstrap, bool) {

	udmsdms := &Udmsdms{
		config:          cfg,
		clock:           clock.New(),
		backoffInterval: backoffInterval,
		timerT1:         retransmitInterval,
		retriesN1:       retransmitRetries,
		context: &UdmsdmsContext{
			startTime: time,
		},
	}

	// Intialize the api handler
	//should both apiClient and apiServer be created in newcsc?
	udmsdms.apiClient = apiclient.NewAPIClient(cfg)

	udmsdms.apiServer = api.NewApiServer(cfg.NodeInfo)

	// Initialize session manager
	//pdusmsp.sessionManager = sm.NewSessionManager(pdusmsp.config.NodeInfo, pdusmsp.config.n11Nodes, time, pdusmsp.backoffInterval, pdusmsp.timerT1, pdusmsp.retriesN1)

	// s:=(pdusmsp.apiClient).(apiclient.ApiClientInfo)
	// var s sm.SessionManager
	// pdusmsp.sessionManager = sm.NewSessionManager(dbInfo, &pdusmsp.grpc, pdusmsp.apiClient)
	// pdusmsp.sessionManager = sm.NewSessionManager(pdusmsp.apiClient)
	udmsdms.sessionManager = sm.NewSessionManager(udmsdms.apiClient)
	//	s:= pdusmsp.sessionManager.NewSMContextAPIService()
	//	fmt.Println(s)
	// Intialize session cache
	udmsdms.sessionCache = cache.NewCache(udmsdms.clock)

	// Intialize session workers
	udmsdms.sessionWorkers = NewSessionWorkers(udmsdms.handleSession, udmsdms.sessionCache, udmsdms.backoffInterval)

	// resyncInterval, backOffPeriod TODO GURU

	return udmsdms, true
}

// Run starts the Udmsdms
func (p *Udmsdms) Run(configChannel <-chan config.UdmsdmsConfig) {
	// start the session manager
	go p.sessionManager.Start()

	// Start the api handler
	p.apiClient.Start()
	go p.apiServer.Start()

	// Start the pdusmsp event handler
	p.udmsdmsEvents(configChannel, p)

	// return
}

func (p *Udmsdms) udmsdmsEvents(configChannel <-chan config.UdmsdmsConfig, handler UdmsdmsHandler) {
	klog.Infof("Entered into udmsdmsEvents")
	syncTicker := time.NewTicker(time.Second)
	defer syncTicker.Stop()
	housekeepingTicker := time.NewTicker(housekeepingPeriod)
	defer housekeepingTicker.Stop()
	sessionChannel := p.apiServer.WatchApiChannel()
	p.handleUdmsdmsEvents(configChannel, sessionChannel, syncTicker.C, housekeepingTicker.C, handler)
}

// handleUdmsdmsEvents is the main loop for processing events in pdusmsp
func (p *Udmsdms) handleUdmsdmsEvents(configChannel <-chan config.UdmsdmsConfig, sessionChannel <-chan *api.SessionMessage,
	syncCh <-chan time.Time, housekeepingCh <-chan time.Time,
	handler UdmsdmsHandler) bool {
	klog.Info("Entered into handleUdmsdmsEvents")
	for {
		select {

		case udmsdmsMsg := <-sessionChannel:
			switch udmsdmsMsg.MsgType {

			case sm.NUDM_RETRIEVE_SM_CONTEXT_REQUEST:
				// PDU Session management service - Retrieve SM Context Request
				klog.Infof("handleUdmsdmsEvents (RETRIEVE SM CONTEXT REQUEST)")

				sessionId := sessions.SessionId(udmsdmsMsg.Supi)
				// p.dispatchWork(sessionId, pdusmsMsg, grpcMsg, time.Now())
				p.dispatchWork(
					sessionId,
					udmsdmsMsg,
					udmsdmsMsg.MsgType,
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

func (p *Udmsdms) HandleSessionCleanups() error {
	deletedSessions := make(map[sessions.SessionId]struct{})
	err := p.sessionWorkers.RemoveSessionWorkers(deletedSessions)
	return err
}

// dispatchWork handles the session in a session worker
func (p *Udmsdms) dispatchWork(sessionId sessions.SessionId,
	sessionMsg *api.SessionMessage,
	msgType sm.MessageType,
	startTime time.Time) {

	// Run the handle session in an async worker.
	p.sessionWorkers.HandleSessionMessages(&SessionMessageInfo{
		SessionId:  sessionId,
		StartTime:  startTime,
		UdmsdmsMsg: *sessionMsg,
		MsgType:    msgType,
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
func (p *Udmsdms) handleSession(msgInfo SessionMessageInfo) error {
	// Get the required message info
	klog.Info("inside")
	msgType := msgInfo.MsgType
	udmsdmsMsg := msgInfo.UdmsdmsMsg

	switch msgType {

	case sm.NUDM_RETRIEVE_SM_CONTEXT_REQUEST:

		ctx := context.Background()
		resp, err := p.sessionManager.GetSmData(ctx, udmsdmsMsg.Supi, udmsdmsMsg.SupportedFeatures, udmsdmsMsg.SingleNssai, udmsdmsMsg.Dnn, udmsdmsMsg.PlmnId, udmsdmsMsg.IfNoneMatch, udmsdmsMsg.IfModifiedSince)
		//sm.ProcessNsmfRetrieveSmContextRequest(pdusmsMsg)
		chanRec := p.apiServer.WatchRecChannel()
		chanRec <- &api.Receiver{ReceivedResponse: resp, ReceivedErr: err}
		return err

	}

	return nil
}
