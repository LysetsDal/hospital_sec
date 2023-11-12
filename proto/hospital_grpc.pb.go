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
	Peer_ReceiveMessageFromPeer_FullMethodName = "/hospital.Peer/ReceiveMessageFromPeer"
)

// PeerClient is the client API for Peer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PeerClient interface {
	// Define RPC methods for communication between peers
	ReceiveMessageFromPeer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PeerMessage, error)
}

type peerClient struct {
	cc grpc.ClientConnInterface
}

func NewPeerClient(cc grpc.ClientConnInterface) PeerClient {
	return &peerClient{cc}
}

func (c *peerClient) ReceiveMessageFromPeer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PeerMessage, error) {
	out := new(PeerMessage)
	err := c.cc.Invoke(ctx, Peer_ReceiveMessageFromPeer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PeerServer is the server API for Peer service.
// All implementations must embed UnimplementedPeerServer
// for forward compatibility
type PeerServer interface {
	// Define RPC methods for communication between peers
	ReceiveMessageFromPeer(context.Context, *Empty) (*PeerMessage, error)
	mustEmbedUnimplementedPeerServer()
}

// UnimplementedPeerServer must be embedded to have forward compatible implementations.
type UnimplementedPeerServer struct {
}

func (UnimplementedPeerServer) ReceiveMessageFromPeer(context.Context, *Empty) (*PeerMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveMessageFromPeer not implemented")
}
func (UnimplementedPeerServer) mustEmbedUnimplementedPeerServer() {}

// UnsafePeerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PeerServer will
// result in compilation errors.
type UnsafePeerServer interface {
	mustEmbedUnimplementedPeerServer()
}

func RegisterPeerServer(s grpc.ServiceRegistrar, srv PeerServer) {
	s.RegisterService(&Peer_ServiceDesc, srv)
}

func _Peer_ReceiveMessageFromPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerServer).ReceiveMessageFromPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Peer_ReceiveMessageFromPeer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerServer).ReceiveMessageFromPeer(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Peer_ServiceDesc is the grpc.ServiceDesc for Peer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Peer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hospital.Peer",
	HandlerType: (*PeerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReceiveMessageFromPeer",
			Handler:    _Peer_ReceiveMessageFromPeer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hospital.proto",
}

const (
	Hospital_SendToHospital_FullMethodName      = "/hospital.Hospital/SendToHospital"
	Hospital_ReceiveFromHospital_FullMethodName = "/hospital.Hospital/ReceiveFromHospital"
)

// HospitalClient is the client API for Hospital service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HospitalClient interface {
	// Define RPC methods for communication with the hospital
	SendToHospital(ctx context.Context, in *PeerMessage, opts ...grpc.CallOption) (*Empty, error)
	ReceiveFromHospital(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HospitalMessage, error)
}

type hospitalClient struct {
	cc grpc.ClientConnInterface
}

func NewHospitalClient(cc grpc.ClientConnInterface) HospitalClient {
	return &hospitalClient{cc}
}

func (c *hospitalClient) SendToHospital(ctx context.Context, in *PeerMessage, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Hospital_SendToHospital_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hospitalClient) ReceiveFromHospital(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HospitalMessage, error) {
	out := new(HospitalMessage)
	err := c.cc.Invoke(ctx, Hospital_ReceiveFromHospital_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HospitalServer is the server API for Hospital service.
// All implementations must embed UnimplementedHospitalServer
// for forward compatibility
type HospitalServer interface {
	// Define RPC methods for communication with the hospital
	SendToHospital(context.Context, *PeerMessage) (*Empty, error)
	ReceiveFromHospital(context.Context, *Empty) (*HospitalMessage, error)
	mustEmbedUnimplementedHospitalServer()
}

// UnimplementedHospitalServer must be embedded to have forward compatible implementations.
type UnimplementedHospitalServer struct {
}

func (UnimplementedHospitalServer) SendToHospital(context.Context, *PeerMessage) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendToHospital not implemented")
}
func (UnimplementedHospitalServer) ReceiveFromHospital(context.Context, *Empty) (*HospitalMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveFromHospital not implemented")
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
	in := new(PeerMessage)
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
		return srv.(HospitalServer).SendToHospital(ctx, req.(*PeerMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hospital_ReceiveFromHospital_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HospitalServer).ReceiveFromHospital(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hospital_ReceiveFromHospital_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HospitalServer).ReceiveFromHospital(ctx, req.(*Empty))
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
		{
			MethodName: "ReceiveFromHospital",
			Handler:    _Hospital_ReceiveFromHospital_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hospital.proto",
}
