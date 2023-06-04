package logic

import (
	"context"

	"github.com/pkg/errors"
	"shop-kay/apps/user/rpc/internal/svc"
	"shop-kay/apps/user/rpc/model"
	"shop-kay/apps/user/rpc/user"
	"shop-kay/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserCollectionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserCollectionLogic {
	return &AddUserCollectionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加收藏
func (l *AddUserCollectionLogic) AddUserCollection(in *user.UserCollectionAddReq) (*user.UserCollectionAddRes, error) {
	dbCollection := new(model.UserCollection)
	dbCollection.Uid = in.Uid
	dbCollection.ProductId = in.ProductId
	_, err := l.svcCtx.UserCollectionModel.Insert(l.ctx, dbCollection)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "AddUserCollection Database Exception : %+v , err: %v", dbCollection, err)
	}
	return &user.UserCollectionAddRes{}, nil
}
