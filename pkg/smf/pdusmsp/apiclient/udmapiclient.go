package apiclient

import (
	"context"
	"fmt"
	"log"

	"k8s.io/klog"
	openapi_udm_cli "w5gc.io/wipro5gcore/openapi/openapiudmclient"
	"w5gc.io/wipro5gcore/pkg/smf/pdusmsp/config"
)

type UdmApiClient interface {
	Udm_Start()
	GetSessionManagementSubscriptionDataRetrievalAPI()
}

type UdmApiClientInfo struct {
	nodeInfo      config.SmfNodeInfo
	openApiClient *openapi_udm_cli.APIClient
}

func NewUdmAPIClient(cfg *config.PdusmspConfig) UdmApiClient {
	c := &UdmApiClientInfo{
		nodeInfo: cfg.NodeInfo,
	}

	udmcfg := openapi_udm_cli.NewConfiguration()
	c.openApiClient = openapi_udm_cli.NewAPIClient(udmcfg)

	return c
}

// TODO We need to call this function ask guru
func (a *UdmApiClientInfo) GetSessionManagementSubscriptionDataRetrievalAPI() {
	supi := "121"
	dnn := "internet"
	sd := "7C7e8b"
	singleNssai := openapi_udm_cli.Snssai{
		Sst: 25,
		Sd:  &sd,
	}
	plmnId := openapi_udm_cli.PlmnId{
		Mcc: "2",
		Mnc: "15",
	}

	c := a.openApiClient.SessionManagementSubscriptionDataRetrievalAPI.GetSmData(context.Background(), supi).
		Dnn(dnn).
		SingleNssai(singleNssai).
		PlmnId(plmnId)

	// Now you can execute the request
	response, _, err := c.Execute()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("responseBody:%+v", response.ExtendedSmSubsData.IndividualSmSubsData)
}

func (a *UdmApiClientInfo) Udm_Start() {
	klog.Infof("Started UDM API client")
}
