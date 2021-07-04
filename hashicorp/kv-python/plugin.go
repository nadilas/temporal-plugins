package kv_python

import (
	"context"

	"github.com/hashicorp/go-plugin"
	"github.com/nadilas/temporal-plugins/hashicorp/kv-python/kvpb"
	"google.golang.org/grpc"
)

// Plugin is the implementation of plugin.GRPCPlugin so we can serve/consume this.
type Plugin struct {
	// Plugin must still implement the Plugin interface
	plugin.Plugin
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl KV
}

func (p *Plugin) GRPCServer(
	broker *plugin.GRPCBroker,
	s *grpc.Server,
) error {
	// no-op
	return nil
}

func (p *Plugin) GRPCClient(
	ctx context.Context,
	broker *plugin.GRPCBroker,
	c *grpc.ClientConn,
) (interface{}, error) {
	return &GRPCClient{client: kvpb.NewKVClient(c)}, nil
}
