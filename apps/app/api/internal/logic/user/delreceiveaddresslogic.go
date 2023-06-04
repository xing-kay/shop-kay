package user

import (
	"context"

	"shop-kay/apps/app/api/internal/svc"
	"shop-kay/apps/app/api/internal/types"
	"shop-kay/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelReceiveAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelReceiveAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelReceiveAddressLogic {
	return &DelReceiveAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelReceiveAddressLogic) DelReceiveAddress(req *types.UserReceiveAddressDelReq) (resp *types.UserReceiveAddressDelRes, err error) {
	_, err = l.svcCtx.UserRPC.DelUserReceiveAddress(l.ctx, &user.UserReceiveAddressDelReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
