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

func (l *AddUserCollectionLogic) AddUserCollection(in *user.AddUserCollectionRequest) (*user.AddUserCollectionResponse, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.UserCollectionModel.Insert(l.ctx, &model.UsersCollection{
		Uid:       in.GetUserId(),
		ProductId: in.GetProductId(),
	})
	if err != nil {
		return nil, status.Error(codes.Internal,
			fmt.Sprintf("[UserCollectionModel.Insert]:failed to insert:%v", err))
	}

	return &user.AddUserCollectionResponse{}, nil
}
