package apiclient

import (
	"io"
	"strconv"
	"time"

	//"net"
	"context"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/go-redis/redis/v8"
	"k8s.io/klog"
	"w5gc.io/wipro5gcore/openapi/openapiclient"
	"w5gc.io/wipro5gcore/pkg/amf/csp/config"
	"w5gc.io/wipro5gcore/pkg/amf/csp/sm"
)

const (
	ApiChannelCapacity = 100
)

type SessionMessage struct {
	MsgType    sm.MessageType
	SessionMsg sm.SMContextMessage
}

type ApiClient interface {
	Start(RequestType sm.MessageType)
	WatchApiChannel() chan *SessionMessage
}

type ApiClientInfo struct {
	clientStartTime time.Time
	apiChannel      chan *SessionMessage
	nodeInfo        config.AmfNodeInfo
	openApiClient   *openapiclient.APIClient
	sessionDb       *redis.Client
}

// changed this paramter to type config.CspConfig instead of config.AmfNodeInfo
func NewApiClient(cfg *config.CspConfig) ApiClient {
	c := &ApiClientInfo{
		apiChannel: make(chan *SessionMessage, ApiChannelCapacity),
		nodeInfo:   cfg.NodeInfo,
	}
	//need config.N11SmfNodesInfo to create OpenApiCfg with server address and server URL
	//need to change NewApiClient parameters
	N11SmfNodes := cfg.N11SmfNodes
	defaultSmfServerAddress := N11SmfNodes[0].NodeId
	defaultSmfServerPort := N11SmfNodes[0].Port
	klog.Info("Server address and port is ", defaultSmfServerAddress, " ", defaultSmfServerPort)
	var OpenApiCfg *openapiclient.Configuration = openapiclient.NewConfiguration(defaultSmfServerAddress, defaultSmfServerPort)
	c.openApiClient = openapiclient.NewAPIClient(OpenApiCfg)

	ctx := context.Background()
	sessionClient, err := sm.NewRedisClient(ctx, 2)
	if err != nil {
		// Handle error
		klog.Error("unable to connect to session Database")
		klog.Info(err)
	}
	c.sessionDb = sessionClient
	return c
}

func (a *ApiClientInfo) CreateSMContext() {
	var ctx context.Context
	createRequest := a.openApiClient.SMContextsCollectionAPI.PostSmContexts(ctx)
	smContextCreateJsonFile, err := os.Open("/home/ubuntu/wipro5gc/testdata/smContextCreate.json")
	if err != nil {
		klog.Info("Error in opening sm context create file...", err)
	} else {
		klog.Info("Successfully opened sm context json file smcontextcreate.json...")
	}
	defer smContextCreateJsonFile.Close()

	// read opened JSON file as a byte array.
	byteValue, err := io.ReadAll(smContextCreateJsonFile)
	klog.Info("json data as string:", string(byteValue))
	if err != nil {
		klog.Info("Error in reading sm context data...", err)
	} else {
		klog.Info("Successfully read sm context data...")
	}
	var smContextCreateData openapiclient.SmContextCreateData
	json.Unmarshal(byteValue, &smContextCreateData)
	createRequest = createRequest.JsonData(smContextCreateData)

	file, e := os.Open("/home/ubuntu/wipro5gc/testdata/n1msgtest")
	if e == nil {
		createRequest = createRequest.BinaryDataN1SmMessage(file)
	} else {
		klog.Info("Error in reading N1 message file:", e)
	}
	createdData, response, _ := createRequest.Execute()
	smContextRef := strconv.Itoa(int(*createdData.PduSessionId)) + smContextCreateData.Guami.AmfId
	a.sessionDb.Set(ctx, smContextRef, createdData, 0)
	klog.Info(createdData)
	klog.Info(response)
}

func (a *ApiClientInfo) ReleaseSMContext() {
	var ctx context.Context
	smContextRef := "123"
	if a.sessionDb.Get(ctx, smContextRef).Err() == redis.Nil {
		klog.Info("SM Context Reference does not exist.")
		return
	}
	releaseSmContextRequest := a.openApiClient.IndividualSMContextAPI.ReleaseSmContext(ctx, smContextRef)
	//var smContextReleaseData openapi.SmContextReleaseData
	smContextReleaseJsonFile, err := os.Open("/home/ubuntu/wipro5gc/testdata/smContextRelease.json")
	if err != nil {
		klog.Info("Error in opening sm context release file...", err)
	} else {
		klog.Info("Successfully opened sm context json file smcontextrelease.json...")
	}
	defer smContextReleaseJsonFile.Close()
	// read opened JSON file as a byte array.
	byteValue, err := ioutil.ReadAll(smContextReleaseJsonFile)
	klog.Info("json data as string:", string(byteValue))
	if err != nil {
		klog.Info("Error in reading sm context data...", err)
	} else {
		klog.Info("Successfully read sm context data...")
	}
	var smContextReleaseData openapiclient.SmContextReleaseData
	json.Unmarshal(byteValue, &smContextReleaseData)
	klog.Info("Sm context release data:", smContextReleaseData)
	releaseSmContextRequest = releaseSmContextRequest.SmContextReleaseData(smContextReleaseData)
	binaryDataN2SmInformationFile, e := os.Open("/home/ubuntu/wipro5gc/testdata/n2infotest")
	if e == nil {
		releaseSmContextRequest = releaseSmContextRequest.BinaryDataN2SmInformation(binaryDataN2SmInformationFile)
	} else {
		klog.Info("Error in reading N2 information file:", e)
	}
	defer binaryDataN2SmInformationFile.Close()
	response, err := releaseSmContextRequest.Execute()
	klog.Info("Release SM Context response : ", response)
	klog.Info("Release Sm Context err : ", err)
	if err == nil {
		a.sessionDb.Del(ctx, smContextRef)
	}
}

func (a *ApiClientInfo) RetrieveSMContext() {
	var ctx context.Context
	smContextRef := "123"
	if a.sessionDb.Get(ctx, smContextRef).Err() == redis.Nil {
		klog.Info("SM Context Reference does not exist.")
		return
	}
	retrieveSmContextRequest := a.openApiClient.IndividualSMContextAPI.RetrieveSmContext(ctx, smContextRef)
	//var smContextRetrieveData openapi.SmContextRetrieveData
	// nonIpSupported := false
	// var smContextRetrieveData = openapiclient.SmContextRetrieveData{
	// 	TargetMmeCap: &openapiclient.MmeCapabilities{
	// 		NonIpSupported: &nonIpSupported,
	// 	},
	// }
	smContextRetrieveJsonFile, err := os.Open("/home/ubuntu/wipro5gc/testdata/smContextRetrieve.json")
	if err != nil {
		klog.Info("Error in opening sm context retrieve file...", err)
	} else {
		klog.Info("Successfully opened sm context json file smcontextretrieve.json...")
	}
	defer smContextRetrieveJsonFile.Close()
	// read opened JSON file as a byte array.
	byteValue, err := ioutil.ReadAll(smContextRetrieveJsonFile)
	klog.Info("json data as string:", string(byteValue))
	if err != nil {
		klog.Info("Error in reading sm context data...", err)
	} else {
		klog.Info("Successfully read sm context data...")
	}
	var smContextRetrieveData openapiclient.SmContextRetrieveData
	json.Unmarshal(byteValue, &smContextRetrieveData)
	klog.Info("Sm context retrieve data:", smContextRetrieveData)
	retrieveSmContextRequest = retrieveSmContextRequest.SmContextRetrieveData(smContextRetrieveData)
	smContextRetrievedData, response, err := retrieveSmContextRequest.Execute()
	klog.Info("Retrieved SM Context data : ", smContextRetrievedData)
	klog.Info("Retrieve SM Context response : ", response)
	klog.Info("Retrieve Sm Context err : ", err)
}

func (a *ApiClientInfo) UpdateSMContext() {
	var ctx context.Context
	smContextRef := "123"
	sessionInfo, err := a.sessionDb.Get(ctx, smContextRef).Result()
	if err == redis.Nil {
		klog.Info("SM Context Reference does not exist.")
		return
	}
	if err != nil {
		klog.Info("Internal Server Error : Redis")
		return
	}
	updateSmContextRequest := a.openApiClient.IndividualSMContextAPI.UpdateSmContext(ctx, smContextRef)

	smContextUpdateJsonFile, err := os.Open("/home/ubuntu/wipro5gc/testdata/smContextUpdate.json")
	if err != nil {
		klog.Info("Error in opening sm context update file...", err)
	} else {
		klog.Info("Successfully opened sm context json file smcontextupdate.json...")
	}
	defer smContextUpdateJsonFile.Close()
	// read opened JSON file as a byte array.
	byteValue, err := ioutil.ReadAll(smContextUpdateJsonFile)
	klog.Info("json data as string:", string(byteValue))
	if err != nil {
		klog.Info("Error in reading sm context data...", err)
	} else {
		klog.Info("Successfully read sm context data...")
	}
	var smContextUpdateData openapiclient.SmContextUpdateData
	json.Unmarshal(byteValue, &smContextUpdateData)
	klog.Info("Sm context update data:", smContextUpdateData)
	updateSmContextRequest = updateSmContextRequest.SmContextUpdateData(smContextUpdateData)

	binaryDataN1SmMessageFile, e := os.Open("/home/ubuntu/wipro5gc/testdata/n1msgtest")
	if e == nil {
		updateSmContextRequest = updateSmContextRequest.BinaryDataN1SmMessage(binaryDataN1SmMessageFile)
	} else {
		klog.Info("Error in reading N1 message file:", e)
	}
	defer binaryDataN1SmMessageFile.Close()

	binaryDataN2SmInformationFile, e := os.Open("/home/ubuntu/wipro5gc/testdata/n2infotest")
	if e == nil {
		updateSmContextRequest = updateSmContextRequest.BinaryDataN2SmInformation(binaryDataN2SmInformationFile)
	} else {
		klog.Info("Error in reading N2 information file:", e)
	}
	defer binaryDataN2SmInformationFile.Close()

	binaryDataN2SmInformationExt1File, e := os.Open("/home/ubuntu/wipro5gc/testdata/n2infoext1test")
	if e == nil {
		updateSmContextRequest = updateSmContextRequest.BinaryDataN2SmInformationExt1(binaryDataN2SmInformationExt1File)
	} else {
		klog.Info("Error in reading N2 information ext1 file:", e)
	}
	defer binaryDataN2SmInformationExt1File.Close()
	smContextUpdatedData, response, err := updateSmContextRequest.Execute()
	klog.Info("Updated SM Context data : ", smContextUpdatedData)
	klog.Info("Update SM Context response : ", response)
	klog.Info("Update Sm Context err : ", err)
	if err == nil {
		var session openapiclient.SmContextCreatedData
		e := json.Unmarshal([]byte(sessionInfo), &session)
		if e != nil {
			klog.Info("Error in unmarshalling session data")
		}
		session.UpCnxState = smContextUpdatedData.UpCnxState
		session.AllocatedEbiList = smContextUpdatedData.AllocatedEbiList
		a.sessionDb.Set(ctx, smContextRef, session, 0)
	}
}

func (a *ApiClientInfo) Start(RequestType sm.MessageType) {
	klog.Infof("Started SMF csp API client")
	switch RequestType {
	case 1:
		a.CreateSMContext()
	case 3:
		a.UpdateSMContext()
	case 5:
		a.ReleaseSMContext()
	case 7:
		a.RetrieveSMContext()
	}
}

func (a *ApiClientInfo) WatchApiChannel() chan *SessionMessage {
	return a.apiChannel
}
