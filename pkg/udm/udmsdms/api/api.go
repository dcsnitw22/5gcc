/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

import (
	"net/http"

	"encoding/json"
	"fmt"
	"time"

	"k8s.io/klog"
	"w5gc.io/wipro5gcore/openapi/openapiserver"
	openapi "w5gc.io/wipro5gcore/openapi/openapiudmserver"
	"w5gc.io/wipro5gcore/pkg/udm/udmsdms/config"

	"strings"

	"github.com/gorilla/mux"
	"w5gc.io/wipro5gcore/pkg/udm/udmsdms/sm"
)

const (
	ApiChannelCapacity = 100
)

type SessionMessage struct {
	MsgType           sm.MessageType
	SessionMsg        sm.SMContextMessage
	Supi              string
	SupportedFeatures string
	SingleNssai       openapi.Snssai
	Dnn               string
	PlmnId            openapi.PlmnId
	IfNoneMatch       string
	IfModifiedSince   string
	Writer            http.ResponseWriter
	Request           *http.Request
}

type Receiver struct {
	ReceivedResponse openapi.ImplResponse
	ReceivedErr      error
}
type ApiServer interface {
	Start()
	WatchApiChannel() chan *SessionMessage
	WatchRecChannel() chan *Receiver
}

type ApiServerInfo struct {
	serverStartTime time.Time
	ApiChannel      chan *SessionMessage
	ApiReceiver     chan *Receiver
	//	router          http.Handler
	nodeInfo        config.UdmNodeInfo
	RequestResponse openapiserver.ImplResponse
	ErrorResponse   error
}

func NewApiServer(cfg config.UdmNodeInfo) ApiServer {
	return &ApiServerInfo{
		nodeInfo:    cfg,
		ApiChannel:  make(chan *SessionMessage, ApiChannelCapacity),
		ApiReceiver: make(chan *Receiver),
	}
}

// func GetResponse(resp openapiserver.ImplResponse,err error)

func (a *ApiServerInfo) Start() {
	klog.Infof("Started SMF pdusmsp API server")
	router := NewRouter(a.Routes())
	klog.Infof("Started the server on Port: %v", a.nodeInfo.ApiPort)
	klog.Fatal(http.ListenAndServe(a.nodeInfo.ApiPort, router))

}

// Routes returns all the api routes for the SessionManagementSubscriptionDataRetrievalAPIController
func (c *ApiServerInfo) Routes() Routes {
	return Routes{
		"GetSmData": Route{
			strings.ToUpper("Get"),
			"/nudm-sdm/v2/{supi}/sm-data",
			c.GetSmData,
		},
	}
}

// GetSmData - retrieve a UE's Session Management Subscription Data
func (c *ApiServerInfo) GetSmData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("coming on route side")
	params := mux.Vars(r)
	query := r.URL.Query()
	supiParam := params["supi"]
	if supiParam == "" {
		openapi.DefaultErrorHandler(w, r, &openapi.RequiredError{"supi"}, nil)
		return
	}
	var supportedFeaturesParam string
	if query.Has("supported-features") {
		supportedFeaturesParam = query.Get("supported-features")
	}
	var singleNssaiParam openapi.Snssai
	if query.Has("single-nssai") {
		singleNssaiString := query.Get("single-nssai")
		json.Unmarshal([]byte(singleNssaiString), &singleNssaiParam)
	}
	var dnnParam string
	if query.Has("dnn") {
		fmt.Println("COMING IN DNN ROUTE")
		dnnParam = query.Get("dnn")
	}
	var plmnIdParam openapi.PlmnId
	if query.Has("plmn-id") {
		plmnIdstring := query.Get("plmn-id")
		json.Unmarshal([]byte(plmnIdstring), &plmnIdParam)
	}
	ifNoneMatchParam := r.Header.Get("If-None-Match")
	ifModifiedSinceParam := r.Header.Get("If-Modified-Since")

	// Create the message to send to the ApiChannel
	smDataRequest := &SessionMessage{MsgType: sm.NUDM_RETRIEVE_SM_CONTEXT_REQUEST,
		Supi:              supiParam,
		SupportedFeatures: supportedFeaturesParam,
		SingleNssai:       singleNssaiParam,
		Dnn:               dnnParam,
		PlmnId:            plmnIdParam,
		IfNoneMatch:       ifNoneMatchParam,
		IfModifiedSince:   ifModifiedSinceParam,
	}

	// Send the request to the ApiChannel
	c.ApiChannel <- smDataRequest

	// Wait for the response from the ApiReceiver
	response := <-c.ApiReceiver

	// If an error occurred, encode the error with the status code
	if response.ReceivedErr != nil {
		openapi.DefaultErrorHandler(w, r, response.ReceivedErr, nil)
		return
	}

	// If no error, encode the body and the result code
	EncodeJSONResponse(response.ReceivedResponse.Body, &response.ReceivedResponse.Code, w)
}

func (a *ApiServerInfo) WatchApiChannel() chan *SessionMessage {
	return a.ApiChannel
}

func (a *ApiServerInfo) WatchRecChannel() chan *Receiver {
	return a.ApiReceiver
}
