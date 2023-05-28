gen-order-api:
	goctl api go -api app/order/api/*.api -dir app/order/api/  -style=goZero

gen-order-rpc:
	goctl rpc protoc apps/order/rpc/*.proto --go_out=apps/order/rpc/ --go-grpc_out=apps/order/rpc/  --zrpc_out=apps/order/rpc

gen-product-rpc:
	goctl rpc protoc apps/product/rpc/*.proto --go_out=apps/product/rpc/ --go-grpc_out=apps/product/rpc/  --zrpc_out=apps/product/rpc
