gen-order-api:
	goctl api go -api app/order/api/*.api -dir app/order/api/  -style=goZero

gen-order-rpc:
	goctl rpc protoc apps/order/rpc/*.proto --go_out=apps/order/rpc/ --go-grpc_out=apps/order/rpc/  --zrpc_out=apps/order/rpc

gen-product-rpc:
	goctl rpc protoc apps/product/rpc/*.proto --go_out=apps/product/rpc/ --go-grpc_out=apps/product/rpc/  --zrpc_out=apps/product/rpc

gen-product-model:
	goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/product" -table="*"  -dir="apps/product/rpc/internal/model" -c


go-run-user-rpc:
	go run apps/user/rpc/user.go -f apps/user/rpc/etc/user.yaml

go-run-product-rpc:
	go run apps/product/rpc/product.go -f apps/product/rpc/etc/product.yaml

debug-product:
	grpcurl -plaintext -d '{"product_id":1}' 127.0.0.1:9001 product.Product.Product

debug-products:
	grpcurl -plaintext -d '{"product_ids":[1,2,3,4,5]}' 127.0.0.1:9001 product.Product.Products

debug-ProductList:
	grpcurl -plaintext -d '{"category_id":8}' 127.0.0.1:9001 product.Product.ProductList

debug-ProductList:
	grpcurl -plaintext -d '{}' 127.0.0.1:9001 product.Product.OperationProducts
