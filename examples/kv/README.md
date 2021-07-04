# kv example

This example is using a golang plugin implementation.

You can run the example by:
1. starting a local temporal cluster
2. starting the worker:
   ```shell
   KV_PLUGIN=./hashicorp/kv/dist/kv-go-grpc go run ./examples/kv/main.go
   ``` 
3. running the starter:
    ```shell
    go run ./examples/k/starter/main.go
    ```

Check temporal web (http://localhost:8088) for the workflow execution history.