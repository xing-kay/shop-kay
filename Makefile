gen-order-rpc:
	goctl rpc protoc apps/order/rpc/*.proto --go_out=apps/order/rpc/ --go-grpc_out=apps/order/rpc/  --zrpc_out=apps/order/rpc
