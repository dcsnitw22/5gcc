package udrs

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

	"w5gc.io/wipro5gcore/pkg/udr/udrs/api"

	// "w5gc.io/wipro5gcore/pkg/udr/udrs/sm/nodes"

	// openapiserver "w5gc.io/wipro5gcore/openapi/openapiudrserver"

	"w5gc.io/wipro5gcore/pkg/udr/udrs/config"
	"w5gc.io/wipro5gcore/pkg/udr/udrs/sm/sessions"
	"w5gc.io/wipro5gcore/utils/cache"

	"w5gc.io/wipro5gcore/pkg/udr/udrs/sm"
)

const (
	retransmitInterval = time.Second * 3
	retransmitRetries  = 3
	backoffInterval    = time.Second * 3
	// Period for performing global cleanup tasks.
	housekeepingPeriod = time.Second * 2
)

// UdrsHandler is an interface implemented for testability
type UdrsHandler interface {
	//InitiateSessioReportRequest(sessionMsg *pfcpUdp.Message)
	HandleSessionCleanups() error
}

// Bootstrap is a bootstrapping interface for PDU SMS
type UdrsBootstrap interface {
	//GetConfiguration()
	//GetContext()
	Run(configChannel <-chan config.UdrsConfig)
}

type Udrs struct {
	config          *config.UdrsConfig
	sessionManager  sm.SessionManager
	sessionCache    cache.WorkCache
	apiServer       api.ApiServer
	sessionWorkers  SessionWorkers
	clock           clock.Clock
	backoffInterval time.Duration
	timerT1         time.Duration
	retriesN1       uint8
	context         *UdrsContext
}

type UdrsContext struct {
	startTime   time.Time
	lastRestart time.Time
	restarts    int // If number of restarts in last 10 sec > 3 reset TODO GURU
	//lock                  sync.Mutex
	NodeId string
}

// Initialize Udrs
func NewUdrs(cfg *config.UdrsConfig, time time.Time) (UdrsBootstrap, bool) {

	udrs := &Udrs{
		config:          cfg,
		clock:           clock.New(),
		backoffInterval: backoffInterval,
		timerT1:         retransmitInterval,
		retriesN1:       retransmitRetries,
		context: &UdrsContext{
			startTime: time,
		},
	}

	// Intialize the api handler
	//should both apiClient and apiServer be created in newcsc?

	udrs.apiServer = api.NewApiServer(cfg.NodeInfo)

	// Initialize session manager
	//pdusmsp.sessionManager = sm.NewSessionManager(pdusmsp.config.NodeInfo, pdusmsp.config.n11Nodes, time, pdusmsp.backoffInterval, pdusmsp.timerT1, pdusmsp.retriesN1)

	// s:=(pdusmsp.apiClient).(apiclient.ApiClientInfo)
	// var s sm.SessionManager
	// pdusmsp.sessionManager = sm.NewSessionManager(dbInfo, &pdusmsp.grpc, pdusmsp.apiClient)
	// pdusmsp.sessionManager = sm.NewSessionManager(pdusmsp.apiClient)
	//	s:= pdusmsp.sessionManager.NewSMContextAPIService()
	//	fmt.Println(s)
	// Intialize session cache
	udrs.sessionCache = cache.NewCache(udrs.clock)

	// Intialize session workers
	udrs.sessionWorkers = NewSessionWorkers(udrs.handleSession, udrs.sessionCache, udrs.backoffInterval)

	// resyncInterval, backOffPeriod TODO GURU

	return udrs, true
}

// Run starts the Udrs
func (p *Udrs) Run(configChannel <-chan config.UdrsConfig) {
	// start the session manager
	go p.sessionManager.Start()

	// Start the api handler
	go p.apiServer.Start()

	// Start the pdusmsp event handler
	p.udrsEvents(configChannel, p)

	// return
}

func (p *Udrs) udrsEvents(configChannel <-chan config.UdrsConfig, handler UdrsHandler) {
	klog.Infof("Entered into pdusmspEvents")
	syncTicker := time.NewTicker(time.Second)
	defer syncTicker.Stop()
	housekeepingTicker := time.NewTicker(housekeepingPeriod)
	defer housekeepingTicker.Stop()
	sessionChannel := p.apiServer.WatchApiChannel()
	p.handleUdrsEvents(configChannel, sessionChannel, syncTicker.C, housekeepingTicker.C, handler)
}

// handleUdrsEvents is the main loop for processing events in pdusmsp
func (p *Udrs) handleUdrsEvents(configChannel <-chan config.UdrsConfig, sessionChannel <-chan *api.SessionMessage,
	syncCh <-chan time.Time, housekeepingCh <-chan time.Time,
	handler UdrsHandler) bool {
	klog.Info("Entered into handleUdrsEvents")
	for {
		select {

		case udrsMsg := <-sessionChannel:
			switch udrsMsg.MsgType {

			case sm.NUDR_QUERY_SM_CONTEXT_REQUEST:
				// PDU Session management service - Retrieve SM Context Request
				klog.Infof("handleUdrsEvents (QUERY SM CONTEXT REQUEST)")

				sessionId := sessions.SessionId(udrsMsg.UeId)
				// p.dispatchWork(sessionId, pdusmsMsg, grpcMsg, time.Now())
				p.dispatchWork(
					sessionId,
					udrsMsg,
					udrsMsg.MsgType,
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

func (p *Udrs) HandleSessionCleanups() error {
	deletedSessions := make(map[sessions.SessionId]struct{})
	err := p.sessionWorkers.RemoveSessionWorkers(deletedSessions)
	return err
}

// dispatchWork handles the session in a session worker
func (p *Udrs) dispatchWork(sessionId sessions.SessionId,
	sessionMsg *api.SessionMessage,
	msgType sm.MessageType,
	startTime time.Time) {

	// Run the handle session in an async worker.
	p.sessionWorkers.HandleSessionMessages(&SessionMessageInfo{
		SessionId: sessionId,
		StartTime: startTime,
		UdrsMsg:   *sessionMsg,
		MsgType:   msgType,
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
func (p *Udrs) handleSession(msgInfo SessionMessageInfo) error {
	// Get the required message info
	klog.Info("inside")
	msgType := msgInfo.MsgType
	udrsMsg := msgInfo.UdrsMsg

	switch msgType {

	case sm.NUDR_QUERY_SM_CONTEXT_REQUEST:

		ctx := context.Background()
		resp, err := p.sessionManager.QuerySmData(ctx, udrsMsg.UeId, udrsMsg.ServingPlmnId, udrsMsg.SingleNssai, udrsMsg.Dnn, udrsMsg.Fields, udrsMsg.SupportedFeatures, udrsMsg.IfNoneMatch, udrsMsg.IfModifiedSince)
		//sm.ProcessNsmfRetrieveSmContextRequest(pdusmsMsg)
		chanRec := p.apiServer.WatchRecChannel()
		chanRec <- &api.Receiver{ReceivedResponse: resp, ReceivedErr: err}
		return err

	}

	return nil
}
