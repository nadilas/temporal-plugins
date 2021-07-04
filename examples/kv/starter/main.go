package main

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/nadilas/temporal-plugins/examples/kv/workflow"
	"go.temporal.io/sdk/client"
)

func main() {
	newClient, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln(err)
	}

	if len(os.Args) < 2 {
		log.Fatalln(errors.New("key is not provided"))
	}

	we, err := newClient.ExecuteWorkflow(context.Background(), client.StartWorkflowOptions{
		TaskQueue: "KV_TASKS",
	}, workflow.KVManipulator, os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Workflow started", we.GetID(), we.GetRunID())
}
