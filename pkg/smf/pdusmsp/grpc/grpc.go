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
	MsgType sm.MessageType
	GrpcMsg *GrpcMessageInfo
}
type Grpc interface {
	Start()
	WatchGrpcChannel() chan *GrpcMessage
}

type GrpcInfo struct {
	GrpcStartTime time.Time
	GrpcChannel   chan *GrpcMessage
}

func NewGrpc() Grpc {
	return &GrpcInfo{
		GrpcChannel: make(chan *GrpcMessage, GrpcChannelCapacity),
	}
}

func (g *GrpcInfo) Start() {
	klog.Infof("Started pdusmsp grpc server")
}

func (g *GrpcInfo) WatchGrpcChannel() chan *GrpcMessage {
	return g.GrpcChannel
}
