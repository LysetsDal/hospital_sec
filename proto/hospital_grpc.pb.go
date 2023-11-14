// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: hospital.proto

package proto

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

const (
	Hospital_SendToHospital_FullMethodName = "/hospital.Hospital/SendToHospital"
)

// HospitalClient is the client API for Hospital service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HospitalClient interface {
	SendToHospital(ctx context.Context, in *HospitalMessage, opts ...grpc.CallOption) (*HospitalResponse, error)
}

type hospitalClient struct {
	cc grpc.ClientConnInterface
}

func NewHospitalClient(cc grpc.ClientConnInterface) HospitalClient {
	return &hospitalClient{cc}
}

func (c *hospitalClient) SendToHospital(ctx context.Context, in *HospitalMessage, opts ...grpc.CallOption) (*HospitalResponse, error) {
	out := new(HospitalResponse)
	err := c.cc.Invoke(ctx, Hospital_SendToHospital_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HospitalServer is the server API for Hospital service.
// All implementations must embed UnimplementedHospitalServer
// for forward compatibility
type HospitalServer interface {
	SendToHospital(context.Context, *HospitalMessage) (*HospitalResponse, error)
	mustEmbedUnimplementedHospitalServer()
}

// UnimplementedHospitalServer must be embedded to have forward compatible implementations.
type UnimplementedHospitalServer struct {
}

func (UnimplementedHospitalServer) SendToHospital(context.Context, *HospitalMessage) (*HospitalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendToHospital not implemented")
}
func (UnimplementedHospitalServer) mustEmbedUnimplementedHospitalServer() {}

// UnsafeHospitalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HospitalServer will
// result in compilation errors.
type UnsafeHospitalServer interface {
	mustEmbedUnimplementedHospitalServer()
}

func RegisterHospitalServer(s grpc.ServiceRegistrar, srv HospitalServer) {
	s.RegisterService(&Hospital_ServiceDesc, srv)
}

func _Hospital_SendToHospital_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HospitalMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HospitalServer).SendToHospital(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hospital_SendToHospital_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HospitalServer).SendToHospital(ctx, req.(*HospitalMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// Hospital_ServiceDesc is the grpc.ServiceDesc for Hospital service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hospital_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hospital.Hospital",
	HandlerType: (*HospitalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendToHospital",
			Handler:    _Hospital_SendToHospital_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hospital.proto",
}

const (
	Peer2Peer_SendMessageToPeer_FullMethodName     = "/hospital.Peer2Peer/SendMessageToPeer"
	Peer2Peer_InitiateSecretShare_FullMethodName   = "/hospital.Peer2Peer/InitiateSecretShare"
	Peer2Peer_SendAddedOutputToPeer_FullMethodName = "/hospital.Peer2Peer/SendAddedOutputToPeer"
)

// Peer2PeerClient is the client API for Peer2Peer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Peer2PeerClient interface {
	SendMessageToPeer(ctx context.Context, in *PeerMessage, opts ...grpc.CallOption) (*PeerMessage, error)
	InitiateSecretShare(ctx context.Context, in *SecretMessage, opts ...grpc.CallOption) (*SecretMessage, error)
	SendAddedOutputToPeer(ctx context.Context, in *SecretMessage, opts ...grpc.CallOption) (*SecretMessage, error)
}

type peer2PeerClient struct {
	cc grpc.ClientConnInterface
}

func NewPeer2PeerClient(cc grpc.ClientConnInterface) Peer2PeerClient {
	return &peer2PeerClient{cc}
}

func (c *peer2PeerClient) SendMessageToPeer(ctx context.Context, in *PeerMessage, opts ...grpc.CallOption) (*PeerMessage, error) {
	out := new(PeerMessage)
	err := c.cc.Invoke(ctx, Peer2Peer_SendMessageToPeer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peer2PeerClient) InitiateSecretShare(ctx context.Context, in *SecretMessage, opts ...grpc.CallOption) (*SecretMessage, error) {
	out := new(SecretMessage)
	err := c.cc.Invoke(ctx, Peer2Peer_InitiateSecretShare_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peer2PeerClient) SendAddedOutputToPeer(ctx context.Context, in *SecretMessage, opts ...grpc.CallOption) (*SecretMessage, error) {
	out := new(SecretMessage)
	err := c.cc.Invoke(ctx, Peer2Peer_SendAddedOutputToPeer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Peer2PeerServer is the server API for Peer2Peer service.
// All implementations must embed UnimplementedPeer2PeerServer
// for forward compatibility
type Peer2PeerServer interface {
	SendMessageToPeer(context.Context, *PeerMessage) (*PeerMessage, error)
	InitiateSecretShare(context.Context, *SecretMessage) (*SecretMessage, error)
	SendAddedOutputToPeer(context.Context, *SecretMessage) (*SecretMessage, error)
	mustEmbedUnimplementedPeer2PeerServer()
}

// UnimplementedPeer2PeerServer must be embedded to have forward compatible implementations.
type UnimplementedPeer2PeerServer struct {
}

func (UnimplementedPeer2PeerServer) SendMessageToPeer(context.Context, *PeerMessage) (*PeerMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessageToPeer not implemented")
}
func (UnimplementedPeer2PeerServer) InitiateSecretShare(context.Context, *SecretMessage) (*SecretMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitiateSecretShare not implemented")
}
func (UnimplementedPeer2PeerServer) SendAddedOutputToPeer(context.Context, *SecretMessage) (*SecretMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendAddedOutputToPeer not implemented")
}
func (UnimplementedPeer2PeerServer) mustEmbedUnimplementedPeer2PeerServer() {}

// UnsafePeer2PeerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Peer2PeerServer will
// result in compilation errors.
type UnsafePeer2PeerServer interface {
	mustEmbedUnimplementedPeer2PeerServer()
}

func RegisterPeer2PeerServer(s grpc.ServiceRegistrar, srv Peer2PeerServer) {
	s.RegisterService(&Peer2Peer_ServiceDesc, srv)
}

func _Peer2Peer_SendMessageToPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Peer2PeerServer).SendMessageToPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Peer2Peer_SendMessageToPeer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Peer2PeerServer).SendMessageToPeer(ctx, req.(*PeerMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Peer2Peer_InitiateSecretShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SecretMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Peer2PeerServer).InitiateSecretShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Peer2Peer_InitiateSecretShare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Peer2PeerServer).InitiateSecretShare(ctx, req.(*SecretMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Peer2Peer_SendAddedOutputToPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SecretMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Peer2PeerServer).SendAddedOutputToPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Peer2Peer_SendAddedOutputToPeer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Peer2PeerServer).SendAddedOutputToPeer(ctx, req.(*SecretMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// Peer2Peer_ServiceDesc is the grpc.ServiceDesc for Peer2Peer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Peer2Peer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hospital.Peer2Peer",
	HandlerType: (*Peer2PeerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessageToPeer",
			Handler:    _Peer2Peer_SendMessageToPeer_Handler,
		},
		{
			MethodName: "InitiateSecretShare",
			Handler:    _Peer2Peer_InitiateSecretShare_Handler,
		},
		{
			MethodName: "SendAddedOutputToPeer",
			Handler:    _Peer2Peer_SendAddedOutputToPeer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hospital.proto",
}
