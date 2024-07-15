package apiclient

import (
	"k8s.io/klog"
	// "w5gc.io/wipro5gcore/openapi/openapi_commn_client"
	openapi_udr_cli "w5gc.io/wipro5gcore/openapi/openapiudrclient"
	"w5gc.io/wipro5gcore/pkg/udm/udmsdms/config"
)

type ApiClient interface {
	Start()
	GetSessionManagementSubscriptionDataAPI() *openapi_udr_cli.SessionManagementSubscriptionDataAPIService
	SubsToNotifySubscriptionDataAPI() *openapi_udr_cli.SubsToNotifyCollectionAPIService
}

type ApiClientInfo struct {
	nodeInfo      config.UdmNodeInfo
	openApiClient *openapi_udr_cli.APIClient
}

func NewAPIClient(cfg *config.UdmsdmsConfig) ApiClient {
	c := &ApiClientInfo{
		nodeInfo: cfg.NodeInfo,
	}

	udrcfg := openapi_udr_cli.NewConfiguration()
	c.openApiClient = openapi_udr_cli.NewAPIClient(udrcfg)

	return c
}

func (a *ApiClientInfo) Start() {
	klog.Infof("Started UDR API client")
}

func (a *ApiClientInfo) GetSessionManagementSubscriptionDataAPI() *openapi_udr_cli.SessionManagementSubscriptionDataAPIService {
	return a.openApiClient.SessionManagementSubscriptionDataAPI
}
func (a *ApiClientInfo) SubsToNotifySubscriptionDataAPI() *openapi_udr_cli.SubsToNotifyCollectionAPIService {
	return a.openApiClient.SubsToNotifyCollectionAPI
}
