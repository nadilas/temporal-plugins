# kv example

This example is using a golang plugin implementation.

You can run the example by:
1. Make sure protoc-gen-grpc-go is available 
  
     `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`
2. clone this repo 

      `git clone https://github.com/nadilas/temporal-plugins.git`
3. Build plugin executable

      `cd hashicorp/kv && make build`
5. Starting a local temporal cluster

      e.g. `temporalite start --namespace default --ephemeral`
6. starting the worker:
   ```shell
   KV_PLUGIN=./hashicorp/kv/dist/kv-go-grpc go run ./examples/kv/main.go
   ``` 
7. running the starter:
    ```shell
    go run ./examples/kv/starter/main.go some-key
    ```

Check temporal web (http://localhost:8088 or http://localhost:8233) for the workflow execution history.