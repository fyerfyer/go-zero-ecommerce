package logic

import (
	"context"
	"fmt"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/user/rpc/internal/svc"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/user/rpc/model"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/user/rpc/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserReceiveAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserReceiveAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserReceiveAddressLogic {
	return &DeleteUserReceiveAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserReceiveAddressLogic) DeleteUserReceiveAddress(in *user.DeleteUserReceiveAddressRequest) (*user.DeleteUserReceiveAddressResponse, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.UserReceiveAddressModel.FindOne(l.ctx, in.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal,
			fmt.Sprintf("[UserReceiveAddressModel.FindOne]:user not found:%v", err))
	}

	err = l.svcCtx.UserReceiveAddressModel.UpdateIsDelete(l.ctx,
		&model.UsersReceiveAddress{
			Uid:      in.GetId(),
			IsDelete: true,
		})

	if err != nil {
		return nil, status.Error(codes.Internal,
			fmt.Sprintf("[UserReceiveAddressModel.UpdateIsDelete]:failed to update:%v", err))
	}
	return &user.DeleteUserReceiveAddressResponse{}, nil
}
