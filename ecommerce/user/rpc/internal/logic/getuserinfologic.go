package logic

import (
	"context"

	"github.com/jinzhu/copier"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/order/rpc/model"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/user/rpc/internal/svc"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/user/rpc/user"
	"github.com/fyerfyer/go-zero-ecommerce/pkg/e"
	"google.golang.org/grpc/codes"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error) {
	// todo: add your logic here and delete this line
	u, err := l.svcCtx.UsersModel.FindOne(l.ctx, in.GetId())
	if err != nil {
		if err == model.ErrNotFound {
			return nil, e.HandleError(
				codes.NotFound,
				err,
				"user not found",
				"UserModel.FindOne",
			)
		} else {
			return nil, e.HandleError(
				codes.Internal,
				err,
				"",
				"UserModel.FindOne",
			)
		}
	}

	var res user.UserInfo
	_ = copier.Copy(&res, u)
	res.CreateTime = u.CreateTime.Unix()
	res.UpdateTime = u.UpdateTime.Unix()

	return &user.GetUserInfoResponse{
		User: &res,
	}, nil
}
