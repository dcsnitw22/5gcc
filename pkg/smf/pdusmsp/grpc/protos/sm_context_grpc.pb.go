// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: sm_context.proto

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SendSmContextDataClient is the client API for SendSmContextData service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SendSmContextDataClient interface {
	SendSmContextCreateData(ctx context.Context, in *SmContextCreateDataRequest, opts ...grpc.CallOption) (*SmContextCreateDataResponse, error)
	SendSmContextUpdateData(ctx context.Context, in *SmContextUpdateDataRequest, opts ...grpc.CallOption) (*SmContextUpdateDataResponse, error)
	SendSmContextReleaseData(ctx context.Context, in *SmContextReleaseDataRequest, opts ...grpc.CallOption) (*SmContextReleaseDataResponse, error)
}

type sendSmContextDataClient struct {
	cc grpc.ClientConnInterface
}

func NewSendSmContextDataClient(cc grpc.ClientConnInterface) SendSmContextDataClient {
	return &sendSmContextDataClient{cc}
}

func (c *sendSmContextDataClient) SendSmContextCreateData(ctx context.Context, in *SmContextCreateDataRequest, opts ...grpc.CallOption) (*SmContextCreateDataResponse, error) {
	out := new(SmContextCreateDataResponse)
	err := c.cc.Invoke(ctx, "/grpc.SendSmContextData/SendSmContextCreateData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sendSmContextDataClient) SendSmContextUpdateData(ctx context.Context, in *SmContextUpdateDataRequest, opts ...grpc.CallOption) (*SmContextUpdateDataResponse, error) {
	out := new(SmContextUpdateDataResponse)
	err := c.cc.Invoke(ctx, "/grpc.SendSmContextData/SendSmContextUpdateData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sendSmContextDataClient) SendSmContextReleaseData(ctx context.Context, in *SmContextReleaseDataRequest, opts ...grpc.CallOption) (*SmContextReleaseDataResponse, error) {
	out := new(SmContextReleaseDataResponse)
	err := c.cc.Invoke(ctx, "/grpc.SendSmContextData/SendSmContextReleaseData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SendSmContextDataServer is the server API for SendSmContextData service.
// All implementations must embed UnimplementedSendSmContextDataServer
// for forward compatibility
type SendSmContextDataServer interface {
	SendSmContextCreateData(context.Context, *SmContextCreateDataRequest) (*SmContextCreateDataResponse, error)
	SendSmContextUpdateData(context.Context, *SmContextUpdateDataRequest) (*SmContextUpdateDataResponse, error)
	SendSmContextReleaseData(context.Context, *SmContextReleaseDataRequest) (*SmContextReleaseDataResponse, error)
	mustEmbedUnimplementedSendSmContextDataServer()
}

// UnimplementedSendSmContextDataServer must be embedded to have forward compatible implementations.
type UnimplementedSendSmContextDataServer struct {
}

func (UnimplementedSendSmContextDataServer) SendSmContextCreateData(context.Context, *SmContextCreateDataRequest) (*SmContextCreateDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSmContextCreateData not implemented")
}
func (UnimplementedSendSmContextDataServer) SendSmContextUpdateData(context.Context, *SmContextUpdateDataRequest) (*SmContextUpdateDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSmContextUpdateData not implemented")
}
func (UnimplementedSendSmContextDataServer) SendSmContextReleaseData(context.Context, *SmContextReleaseDataRequest) (*SmContextReleaseDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSmContextReleaseData not implemented")
}
func (UnimplementedSendSmContextDataServer) mustEmbedUnimplementedSendSmContextDataServer() {}

// UnsafeSendSmContextDataServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SendSmContextDataServer will
// result in compilation errors.
type UnsafeSendSmContextDataServer interface {
	mustEmbedUnimplementedSendSmContextDataServer()
}

func RegisterSendSmContextDataServer(s grpc.ServiceRegistrar, srv SendSmContextDataServer) {
	s.RegisterService(&SendSmContextData_ServiceDesc, srv)
}

func _SendSmContextData_SendSmContextCreateData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmContextCreateDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendSmContextDataServer).SendSmContextCreateData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.SendSmContextData/SendSmContextCreateData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendSmContextDataServer).SendSmContextCreateData(ctx, req.(*SmContextCreateDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SendSmContextData_SendSmContextUpdateData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmContextUpdateDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendSmContextDataServer).SendSmContextUpdateData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.SendSmContextData/SendSmContextUpdateData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendSmContextDataServer).SendSmContextUpdateData(ctx, req.(*SmContextUpdateDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SendSmContextData_SendSmContextReleaseData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmContextReleaseDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendSmContextDataServer).SendSmContextReleaseData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.SendSmContextData/SendSmContextReleaseData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendSmContextDataServer).SendSmContextReleaseData(ctx, req.(*SmContextReleaseDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SendSmContextData_ServiceDesc is the grpc.ServiceDesc for SendSmContextData service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SendSmContextData_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.SendSmContextData",
	HandlerType: (*SendSmContextDataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendSmContextCreateData",
			Handler:    _SendSmContextData_SendSmContextCreateData_Handler,
		},
		{
			MethodName: "SendSmContextUpdateData",
			Handler:    _SendSmContextData_SendSmContextUpdateData_Handler,
		},
		{
			MethodName: "SendSmContextReleaseData",
			Handler:    _SendSmContextData_SendSmContextReleaseData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sm_context.proto",
}
