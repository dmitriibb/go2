// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: kitchen/service/orders.proto

package service

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

// KitchenServiceClient is the client API for KitchenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KitchenServiceClient interface {
	NewOrderEvent(ctx context.Context, in *OrderEvent, opts ...grpc.CallOption) (*OrderEventResponse, error)
}

type kitchenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKitchenServiceClient(cc grpc.ClientConnInterface) KitchenServiceClient {
	return &kitchenServiceClient{cc}
}

func (c *kitchenServiceClient) NewOrderEvent(ctx context.Context, in *OrderEvent, opts ...grpc.CallOption) (*OrderEventResponse, error) {
	out := new(OrderEventResponse)
	err := c.cc.Invoke(ctx, "/go2.kitchen.service.KitchenService/NewOrderEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KitchenServiceServer is the server API for KitchenService service.
// All implementations must embed UnimplementedKitchenServiceServer
// for forward compatibility
type KitchenServiceServer interface {
	NewOrderEvent(context.Context, *OrderEvent) (*OrderEventResponse, error)
	mustEmbedUnimplementedKitchenServiceServer()
}

// UnimplementedKitchenServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKitchenServiceServer struct {
}

func (UnimplementedKitchenServiceServer) NewOrderEvent(context.Context, *OrderEvent) (*OrderEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewOrderEvent not implemented")
}
func (UnimplementedKitchenServiceServer) mustEmbedUnimplementedKitchenServiceServer() {}

// UnsafeKitchenServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KitchenServiceServer will
// result in compilation errors.
type UnsafeKitchenServiceServer interface {
	mustEmbedUnimplementedKitchenServiceServer()
}

func RegisterKitchenServiceServer(s grpc.ServiceRegistrar, srv KitchenServiceServer) {
	s.RegisterService(&KitchenService_ServiceDesc, srv)
}

func _KitchenService_NewOrderEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KitchenServiceServer).NewOrderEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go2.kitchen.service.KitchenService/NewOrderEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KitchenServiceServer).NewOrderEvent(ctx, req.(*OrderEvent))
	}
	return interceptor(ctx, in, info, handler)
}

// KitchenService_ServiceDesc is the grpc.ServiceDesc for KitchenService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KitchenService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "go2.kitchen.service.KitchenService",
	HandlerType: (*KitchenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewOrderEvent",
			Handler:    _KitchenService_NewOrderEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kitchen/service/orders.proto",
}
