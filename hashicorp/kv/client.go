package kv

import (
	"github.com/nadilas/temporal-plugins/hashicorp/kv/kvpb"
	"golang.org/x/net/context"
)

// GRPCClient is an implementation of KV that talks over RPC.
//
// This client destructures the grpc response and thereby conform to the KV interface
type GRPCClient struct{ client kvpb.KVClient }

func (m *GRPCClient) Put(
	ctx context.Context,
	key string,
	value []byte,
) error {
	_, err := m.client.Put(ctx, &kvpb.PutRequest{
		Key:   key,
		Value: value,
	})
	return err
}

func (m *GRPCClient) Get(
	ctx context.Context,
	key string,
) ([]byte, error) {
	resp, err := m.client.Get(ctx, &kvpb.GetRequest{
		Key: key,
	})
	if err != nil {
		return nil, err
	}

	return resp.Value, nil
}
