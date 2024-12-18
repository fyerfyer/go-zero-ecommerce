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

type DeleteUserCollectionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserCollectionLogic {
	return &DeleteUserCollectionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserCollectionLogic) DeleteUserCollection(in *user.DeleteUserCollectionRequest) (*user.DeleteUserCollectionResponse, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.UserCollectionModel.FindOne(l.ctx, in.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal,
			fmt.Sprintf("[UserCollectionModel.FindOne]:user not found:%v", err))
	}

	// 通过更新isDelete标签的方式进行懒删除
	err = l.svcCtx.UserCollectionModel.UpdateIsDelete(l.ctx, &model.UsersCollection{
		Uid:      in.GetId(),
		IsDelete: true,
	})

	if err != nil {
		return nil, status.Error(codes.Internal,
			fmt.Sprintf("[UserCollectionModel.UpdateIsDelete]:failed to update:%v", err))
	}
	return &user.DeleteUserCollectionResponse{}, nil
}
