// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proof.proto

package l2output

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

// ProofClient is the client API for Proof service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProofClient interface {
	Prove(ctx context.Context, in *ProofRequest, opts ...grpc.CallOption) (*ProofResponse, error)
}

type proofClient struct {
	cc grpc.ClientConnInterface
}

func NewProofClient(cc grpc.ClientConnInterface) ProofClient {
	return &proofClient{cc}
}

func (c *proofClient) Prove(ctx context.Context, in *ProofRequest, opts ...grpc.CallOption) (*ProofResponse, error) {
	out := new(ProofResponse)
	err := c.cc.Invoke(ctx, "/proof.Proof/Prove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProofServer is the server API for Proof service.
// All implementations must embed UnimplementedProofServer
// for forward compatibility
type ProofServer interface {
	Prove(context.Context, *ProofRequest) (*ProofResponse, error)
	mustEmbedUnimplementedProofServer()
}

// UnimplementedProofServer must be embedded to have forward compatible implementations.
type UnimplementedProofServer struct {
}

func (UnimplementedProofServer) Prove(context.Context, *ProofRequest) (*ProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Prove not implemented")
}
func (UnimplementedProofServer) mustEmbedUnimplementedProofServer() {}

// UnsafeProofServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProofServer will
// result in compilation errors.
type UnsafeProofServer interface {
	mustEmbedUnimplementedProofServer()
}

func RegisterProofServer(s grpc.ServiceRegistrar, srv ProofServer) {
	s.RegisterService(&Proof_ServiceDesc, srv)
}

func _Proof_Prove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProofServer).Prove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proof.Proof/Prove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProofServer).Prove(ctx, req.(*ProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Proof_ServiceDesc is the grpc.ServiceDesc for Proof service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Proof_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proof.Proof",
	HandlerType: (*ProofServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Prove",
			Handler:    _Proof_Prove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proof.proto",
}
