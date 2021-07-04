package kv

import (
	"github.com/nadilas/temporal-plugins/hashicorp/kv/kvpb"
	"golang.org/x/net/context"
)

// GRPCServer Here is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	// This is the real implementation
	Impl KV
	kvpb.UnimplementedKVServer
}

func (m *GRPCServer) Put(
	ctx context.Context,
	req *kvpb.PutRequest) (*kvpb.Empty, error) {
	return &kvpb.Empty{}, m.Impl.Put(ctx, req.Key, req.Value)
}

func (m *GRPCServer) Get(
	ctx context.Context,
	req *kvpb.GetRequest) (*kvpb.GetResponse, error) {
	v, err := m.Impl.Get(ctx, req.Key)
	return &kvpb.GetResponse{Value: v}, err
}
