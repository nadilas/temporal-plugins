package _go

import (
	"context"
	"os"
	"os/exec"

	"github.com/hashicorp/go-plugin"
)

const PluginName = "kv"

// Command is the default command to start up the plugin.
// It uses KV_PLUGIN env var for the name of the executable
var Command = exec.Command("sh", "-c", os.Getenv("KV_PLUGIN"))

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	// This isn't required when using VersionedPlugins
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	PluginName: &Plugin{},
}

//DefaultConfig is the plugin config using sane default
var DefaultConfig = &plugin.ClientConfig{
	HandshakeConfig: Handshake,
	Plugins:         PluginMap,
	Cmd:             Command,
	AllowedProtocols: []plugin.Protocol{
		plugin.ProtocolGRPC,
	},
}

// KV is the interface that we're exposing as a plugin.
// Method signatures are either Activity or Workflow signatures
type KV interface {
	Put(ctx context.Context, key string, value []byte) error
	Get(ctx context.Context, key string) ([]byte, error)
}
