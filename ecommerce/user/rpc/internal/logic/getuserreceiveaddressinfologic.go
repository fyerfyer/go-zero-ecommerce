package logic

import (
	"context"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/user/rpc/internal/svc"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/user/rpc/user"
	"github.com/fyerfyer/go-zero-ecommerce/pkg/e"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserReceiveAddressInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserReceiveAddressInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserReceiveAddressInfoLogic {
	return &GetUserReceiveAddressInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserReceiveAddressInfoLogic) GetUserReceiveAddressInfo(in *user.GetUserReceiveAddressInfoRequest) (*user.UserReceiveAddress, error) {
	// todo: add your logic here and delete this line
	addr, err := l.svcCtx.UserReceiveAddressModel.FindOne(l.ctx, in.GetId())
	if err != nil {
		return nil, e.HandleError(
			codes.Internal,
			err,
			"failed to get address",
			"UserReceiveAddressModel.FindOne",
		)
	}

	res := new(user.UserReceiveAddress)
	_ = copier.Copy(&res, addr)
	res.CreateTime = addr.CreateTime.Unix()
	res.UpdateTime = addr.UpdateTime.Unix()
	return res, nil
}
