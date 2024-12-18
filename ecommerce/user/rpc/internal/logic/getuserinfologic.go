package logic

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/order/rpc/model"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/user/rpc/internal/svc"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/user/rpc/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
			return nil, status.Error(codes.Internal,
				fmt.Sprintf("[UserModel.FindOne]:user not found:%v", err))
		} else {
			return nil, status.Error(codes.Internal,
				fmt.Sprintf("[UserModel.FindOne]:%v", err))
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
