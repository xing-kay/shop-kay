package svc

import (
	"shop-kay/apps/app/api/internal/config"
	"shop-kay/apps/order/rpc/order"
	"shop-kay/apps/product/rpc/product"
	"shop-kay/apps/reply/rpc/reply"
	"shop-kay/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	OrderRPC   order.Order
	ProductRPC product.Product
	ReplyRPC   reply.Reply
	UserRPC    user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderRPC:   order.NewOrder(zrpc.MustNewClient(c.OrderRPC)),
		ProductRPC: product.NewProduct(zrpc.MustNewClient(c.ProductRPC)),
		ReplyRPC:   reply.NewReply(zrpc.MustNewClient(c.ReplyRPC)),
		UserRPC:    user.NewUser(zrpc.MustNewClient(c.UserRPC)),
	}
}
