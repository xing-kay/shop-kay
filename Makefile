gen-order-api:
	goctl api go -api app/order/api/*.api -dir app/order/api/  -style=goZero

gen-order-rpc:
	goctl rpc protoc apps/order/rpc/*.proto --go_out=apps/order/rpc/ --go-grpc_out=apps/order/rpc/  --zrpc_out=apps/order/rpc

gen-product-rpc:
	goctl rpc protoc apps/product/rpc/*.proto --go_out=apps/product/rpc/ --go-grpc_out=apps/product/rpc/  --zrpc_out=apps/product/rpc

go-run-user-rpc:
	go run apps/user/rpc/user.go -f apps/user/rpc/etc/user.yaml

gen-product-model:
	goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/product" -table="*"  -dir="apps/product/rpc/internal/model" -c