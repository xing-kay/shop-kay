package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"shop-kay/apps/product/rpc/productclient"

	"shop-kay/apps/order/rpc/internal/config"
	"shop-kay/apps/order/rpc/model"
	"shop-kay/apps/user/rpc/user"
)

type ServiceContext struct {
	Config         config.Config
	OrderModel     model.OrdersModel
	OrderitemModel model.OrderitemModel
	ShippingModel  model.ShippingModel
	UserRpc        user.User
	ProductRpc     productclient.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:         c,
		OrderModel:     model.NewOrdersModel(conn, c.CacheRedis),
		OrderitemModel: model.NewOrderitemModel(conn, c.CacheRedis),
		ShippingModel:  model.NewShippingModel(conn, c.CacheRedis),
		UserRpc:        user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		ProductRpc:     productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
