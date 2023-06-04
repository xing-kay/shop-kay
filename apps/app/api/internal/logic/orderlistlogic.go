package logic

import (
	"context"
	"shop-kay/apps/app/api/internal/svc"
	"shop-kay/apps/app/api/internal/types"
	"shop-kay/apps/order/rpc/order"
	"shop-kay/apps/product/rpc/product"
	"shop-kay/pkg/otel"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderListLogic {
	return &OrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderListLogic) OrderList(req *types.OrderListRequest) (resp *types.OrderListResponse, err error) {
	ctx, span := otel.StartSpan(l.ctx, "Custom.BFF.OrderList")
	defer span.End()

	l.ctx = ctx
	orderRet, err := l.svcCtx.OrderRPC.Orders(l.ctx, &order.OrdersRequest{UserId: req.UID})
	if err != nil {
		return nil, err
	}
	var pids []int64
	for _, o := range orderRet.Orders {
		pids = append(pids, o.Id)
	}
	productRet, err := l.svcCtx.ProductRPC.Products(l.ctx, &product.ProductRequest{ProductIds: pids})
	if err != nil {
		return nil, err
	}
	var orders []*types.Order
	for _, o := range orderRet.Orders {
		if p, ok := productRet.Products[o.Id]; ok {
			orders = append(orders, &types.Order{
				OrderID:     o.Orderid,
				ProductName: p.Name,
			})
		}
	}
	return &types.OrderListResponse{Orders: orders}, nil
}