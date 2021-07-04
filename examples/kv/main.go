package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/go-plugin"
	"github.com/nadilas/temporal-plugins/examples/kv/workflow"
	"github.com/nadilas/temporal-plugins/hashicorp/kv"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// fetch plugin
	// We're a host. Start by launching the plugin process.
	pm := plugin.NewClient(kv.DefaultConfig)
	defer pm.Kill()
	kvPlugin, err := kvPlugin(pm)

	// init temporal
	cli, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln(err)
	}
	w := worker.New(cli, "KV_TASKS", worker.Options{})

	w.RegisterWorkflow(workflow.KVManipulator)

	// register plugin methods
	w.RegisterActivity(kvPlugin.Get)
	w.RegisterActivity(kvPlugin.Put)

	// start worker
	if err = w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln(err)
	}
}

func kvPlugin(pm *plugin.Client) (kv.KV, error) {
	// Connect via RPC
	c, err := pm.Client()
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
	// Request the plugin
	raw, err := c.Dispense(kv.PluginName)
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	return raw.(kv.KV), err
}
