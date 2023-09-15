package grpcclient

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"k8s.io/klog"
	"w5gc.io/wipro5gcore/pkg/smf/pdusmsp/config"
	"w5gc.io/wipro5gcore/pkg/smf/pdusmsp/grpc/protos"
)

type GrpcClient struct {
	ConnAddr string                         // Upfgw IP
	Client   protos.SendSmContextDataClient // grpc Client Conn
}

// Initialize with IP data
func NewGrpcClient(cfg config.GrpcClientInfoConfig) *GrpcClient {
	clientAddr := cfg.ClientIP + ":" + cfg.ClientPort
	return &GrpcClient{
		ConnAddr: clientAddr,
	}
}

// Start client with dial
func (g *GrpcClient) Start() {
	// Get UPFGW address and then dial
	conn, err := grpc.Dial(g.ConnAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		klog.Fatalf("Dial error. Did not connect: %v", err)
	}
	g.Client = protos.NewSendSmContextDataClient(conn)
}

// Send SmCreate data request to Upfgw
func (g *GrpcClient) SendSmContextCreateData(createData *protos.SmContextCreateDataRequest) {
	client := g.Client
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	// createData := &SmContextCreateDataRequest{
	// 	Pei:          "Pei",
	// 	Dnn:          "Dnn",
	// 	PduSessionId: 11,
	// }

	// grpc protos generated function
	r, err := client.SendSmContextCreateData(ctx, createData)
	if err != nil {
		klog.Fatalf("Error. Could not get response: %v", err)
	}
	klog.Infof("Response from server: %d", r.GetPduSessionId())
}

// Send Update data
func (g *GrpcClient) SendSmContextUpdateData(updateData *protos.SmContextUpdateDataRequest) {
	client := g.Client
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	// updateData := &SmContextUpdateDataRequest{
	// 	Pei:         "Pei",
	// 	ServingNfId: "ServingNfId",
	// 	N2SmInfo:    &N2SmInformation{},
	// 	Guami: &Guami{
	// 		AmfId: "amfid1",
	// 	},
	// }

	// grpc protos generated function
	res, err := client.SendSmContextUpdateData(ctx, updateData)
	if err != nil {
		klog.Fatalf("Error. Could not get response: %v", err)
	}
	klog.Infof("Response from server: %d", res.GetPduSessionId())
}

// Send Release data
func (g *GrpcClient) SendSmContextReleaseData(releaseData *protos.SmContextReleaseDataRequest) {
	client := g.Client
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	// releaseData := &SmContextUpdateDataRequest{
	// 	Pei:         "Pei",
	// 	ServingNfId: "ServingNfId",
	// 	Guami: &Guami{
	// 		AmfId: "amfid1",
	// 	},
	// }

	// grpc protos generated function
	res, err := client.SendSmContextReleaseData(ctx, releaseData)
	if err != nil {
		klog.Fatalf("Error. Could not get response: %v", err)
	}
	klog.Infof("Response from server: %d", res.GetPduSessionId())
}
