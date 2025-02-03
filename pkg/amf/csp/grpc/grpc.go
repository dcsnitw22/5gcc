package grpc

import (
	"time"
	//"net"
	"k8s.io/klog"

	//"w5gc.io/wipro5gcore/openapi"
	"w5gc.io/wipro5gcore/pkg/smf/pdusmsp/sm"
)

const (
	GrpcChannelCapacity = 100
)

type GrpcMessageInfo interface{}

type GrpcMessage struct {
	msgType sm.MessageType
	grpcMsg *GrpcMessageInfo
}
type Grpc interface {
	Start()
	WatchGrpcChannel() chan *GrpcMessage
}

type GrpcInfo struct {
	grpcStartTime time.Time
	grpcChannel   chan *GrpcMessage
}

func NewGrpc() Grpc {
	return &GrpcInfo{
		grpcChannel: make(chan *GrpcMessage, GrpcChannelCapacity),
	}
}

func (g *GrpcInfo) Start() {
	klog.Infof("Started csp grpc server")
}

func (g *GrpcInfo) WatchGrpcChannel() chan *GrpcMessage {
	return g.grpcChannel
}
