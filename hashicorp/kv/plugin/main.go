package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/hashicorp/go-plugin"
	kv "github.com/nadilas/temporal-plugins/hashicorp/kv/go"
)

// KV Here is a real implementation of KV that writes to a local file with
// the key name and the contents are the value of the key.
type KV struct{}

func (KV) Put(ctx context.Context, key string, value []byte) error {
	value = []byte(fmt.Sprintf("%s", string(value)))
	return ioutil.WriteFile("kv_"+key, value, 0644)
}

func (KV) Get(ctx context.Context, key string) ([]byte, error) {
	return ioutil.ReadFile("kv_" + key)
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: kv.Handshake,
		Plugins: map[string]plugin.Plugin{
			kv.PluginName: &kv.Plugin{Impl: &KV{}},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
