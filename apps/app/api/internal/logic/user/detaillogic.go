package user

import (
	"context"
	"encoding/json"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"

	"shop-kay/apps/app/api/internal/svc"
	"shop-kay/apps/app/api/internal/types"
	"shop-kay/apps/user/rpc/user"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	//l.ctx.Value("uid")  get jwt token uid
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	userInfo, err := l.svcCtx.UserRPC.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: uid,
	})
	if err != nil {
		return nil, err
	}

	var user types.UserInfo
	_ = copier.Copy(&user, userInfo.User)

	return &types.UserInfoResp{
		UserInfo: user,
	}, nil
}
