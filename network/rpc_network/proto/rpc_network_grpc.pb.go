// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: rpc_network.proto

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
	RpcNetwork_SendBlock_FullMethodName      = "/rpc_network.RpcNetwork/SendBlock"
	RpcNetwork_SendBlockchain_FullMethodName = "/rpc_network.RpcNetwork/SendBlockchain"
	RpcNetwork_GetBlockchain_FullMethodName  = "/rpc_network.RpcNetwork/GetBlockchain"
)

// RpcNetworkClient is the client API for RpcNetwork service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RpcNetworkClient interface {
	SendBlock(ctx context.Context, in *Block, opts ...grpc.CallOption) (*SendBlockResponse, error)
	SendBlockchain(ctx context.Context, in *Blockchain, opts ...grpc.CallOption) (*SendBlockchainResponse, error)
	GetBlockchain(ctx context.Context, in *GetBlockchainMessage, opts ...grpc.CallOption) (*GetBlockchainResponse, error)
}

type rpcNetworkClient struct {
	cc grpc.ClientConnInterface
}

func NewRpcNetworkClient(cc grpc.ClientConnInterface) RpcNetworkClient {
	return &rpcNetworkClient{cc}
}

func (c *rpcNetworkClient) SendBlock(ctx context.Context, in *Block, opts ...grpc.CallOption) (*SendBlockResponse, error) {
	out := new(SendBlockResponse)
	err := c.cc.Invoke(ctx, RpcNetwork_SendBlock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcNetworkClient) SendBlockchain(ctx context.Context, in *Blockchain, opts ...grpc.CallOption) (*SendBlockchainResponse, error) {
	out := new(SendBlockchainResponse)
	err := c.cc.Invoke(ctx, RpcNetwork_SendBlockchain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcNetworkClient) GetBlockchain(ctx context.Context, in *GetBlockchainMessage, opts ...grpc.CallOption) (*GetBlockchainResponse, error) {
	out := new(GetBlockchainResponse)
	err := c.cc.Invoke(ctx, RpcNetwork_GetBlockchain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcNetworkServer is the server API for RpcNetwork service.
// All implementations must embed UnimplementedRpcNetworkServer
// for forward compatibility
type RpcNetworkServer interface {
	SendBlock(context.Context, *Block) (*SendBlockResponse, error)
	SendBlockchain(context.Context, *Blockchain) (*SendBlockchainResponse, error)
	GetBlockchain(context.Context, *GetBlockchainMessage) (*GetBlockchainResponse, error)
	mustEmbedUnimplementedRpcNetworkServer()
}

// UnimplementedRpcNetworkServer must be embedded to have forward compatible implementations.
type UnimplementedRpcNetworkServer struct {
}

func (UnimplementedRpcNetworkServer) SendBlock(context.Context, *Block) (*SendBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendBlock not implemented")
}
func (UnimplementedRpcNetworkServer) SendBlockchain(context.Context, *Blockchain) (*SendBlockchainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendBlockchain not implemented")
}
func (UnimplementedRpcNetworkServer) GetBlockchain(context.Context, *GetBlockchainMessage) (*GetBlockchainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockchain not implemented")
}
func (UnimplementedRpcNetworkServer) mustEmbedUnimplementedRpcNetworkServer() {}

// UnsafeRpcNetworkServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RpcNetworkServer will
// result in compilation errors.
type UnsafeRpcNetworkServer interface {
	mustEmbedUnimplementedRpcNetworkServer()
}

func RegisterRpcNetworkServer(s grpc.ServiceRegistrar, srv RpcNetworkServer) {
	s.RegisterService(&RpcNetwork_ServiceDesc, srv)
}

func _RpcNetwork_SendBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Block)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcNetworkServer).SendBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RpcNetwork_SendBlock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcNetworkServer).SendBlock(ctx, req.(*Block))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcNetwork_SendBlockchain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Blockchain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcNetworkServer).SendBlockchain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RpcNetwork_SendBlockchain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcNetworkServer).SendBlockchain(ctx, req.(*Blockchain))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcNetwork_GetBlockchain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockchainMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcNetworkServer).GetBlockchain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RpcNetwork_GetBlockchain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcNetworkServer).GetBlockchain(ctx, req.(*GetBlockchainMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// RpcNetwork_ServiceDesc is the grpc.ServiceDesc for RpcNetwork service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RpcNetwork_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc_network.RpcNetwork",
	HandlerType: (*RpcNetworkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendBlock",
			Handler:    _RpcNetwork_SendBlock_Handler,
		},
		{
			MethodName: "SendBlockchain",
			Handler:    _RpcNetwork_SendBlockchain_Handler,
		},
		{
			MethodName: "GetBlockchain",
			Handler:    _RpcNetwork_GetBlockchain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc_network.proto",
}
