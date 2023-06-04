package svc

import (
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/sync/singleflight"
	"shop-kay/apps/product/rpc/internal/config"
	"shop-kay/apps/product/rpc/internal/model"
	"time"
)

const localCacheExpire = time.Second * 60

type ServiceContext struct {
	Config         config.Config
	ProductModel   model.ProductModel
	CategoryModel  model.CategoryModel
	OperationModel model.ProductOperationModel
	LocalCache     *collection.Cache
	BizRedis       *redis.Redis
	SingleGroup    singleflight.Group
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	localCache, err := collection.NewCache(localCacheExpire)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:         c,
		ProductModel:   model.NewProductModel(conn, c.CacheConf),
		CategoryModel:  model.NewCategoryModel(conn, c.CacheConf),
		OperationModel: model.NewProductOperationModel(conn, c.CacheConf),
		LocalCache:     localCache,
		BizRedis:       redis.New(c.BizRedis.Host, redis.WithPass(c.BizRedis.Pass)),
	}
}
