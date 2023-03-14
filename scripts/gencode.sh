# service
goctl rpc protoc proto/jocker.proto  -I./proto --go_out=./jockerd --go-grpc_out=./jockerd  --zrpc_out=./jockerd --style=goZero

# # grpc gateway
protoc proto/jocker.proto -I./proto/ --go_out=./jockerd --go-grpc_out=./jockerd --grpc-gateway_out=./jockerd