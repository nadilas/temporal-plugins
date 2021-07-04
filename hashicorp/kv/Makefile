protoc:
	protoc -I. kv.proto --go-grpc_out=./kvpb --go_out=./kvpb

build:
	go build -o ./dist/kv-go-grpc ./plugin