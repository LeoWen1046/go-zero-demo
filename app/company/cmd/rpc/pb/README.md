1、生成service：goctl rpc protoc office.proto --go_out=../office --go-grpc_out=../office --zrpc_out=../office
2、go run office.go -f etc/office.yaml
3、调试rpc: grpcui -plaintext localhost:8080(localhost：8080为第二步启动的grpc服务对应起来)