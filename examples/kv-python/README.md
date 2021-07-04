# kv example

This example is using a golang plugin implementation.

You can run the example by:
1. starting a local temporal cluster
2. starting the worker:
   ```shell
   KV_PLUGIN="python3 ./hashicorp/kv-python/plugin.py" go run ./examples/kv-python/main.go
   ``` 
3. running the starter:
    ```shell
    go run ./examples/kv/starter/main.go
    ```

Check temporal web (http://localhost:8088) for the workflow execution history.


