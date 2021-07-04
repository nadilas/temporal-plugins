package workflow

import (
	"time"

	"github.com/nadilas/temporal-plugins/hashicorp/kv-python"
	"go.temporal.io/sdk/workflow"
)

func KVManipulator(ctx workflow.Context, key string) (string, error) {
	ao := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		ScheduleToCloseTimeout: time.Minute * 1,
	})
	if err := workflow.ExecuteActivity(ao, kv_python.KV.Put, key, []byte("value1")).Get(ctx, nil); err != nil {
		return "", err
	}
	var val1 []byte
	if err := workflow.ExecuteActivity(ao, kv_python.KV.Get, key).Get(ctx, &val1); err != nil {
		return "", err
	}
	if err := workflow.ExecuteActivity(ao, kv_python.KV.Put, key, []byte(string(val1) + "\nvalue 2")).Get(ctx, nil); err != nil {
		return "", err
	}
	var final []byte
	if err := workflow.ExecuteActivity(ao, kv_python.KV.Get, key).Get(ctx, &final); err != nil {
		return "", err
	}
	return string(final), nil
}