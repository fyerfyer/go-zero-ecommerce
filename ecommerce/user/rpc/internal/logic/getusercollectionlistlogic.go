package logic

import (
	"context"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/user/rpc/internal/svc"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/user/rpc/user"
	"github.com/fyerfyer/go-zero-ecommerce/pkg/e"
	"google.golang.org/grpc/codes"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserCollectionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserCollectionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserCollectionListLogic {
	return &GetUserCollectionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserCollectionListLogic) GetUserCollectionList(in *user.GetUserCollectionListRequest) (*user.GetUserCollectionListResponse, error) {
	// todo: add your logic here and delete this line
	list, err := l.svcCtx.UserCollectionModel.FindAllByUid(l.ctx, in.GetUserId())
	if err != nil {
		return nil, e.HandleError(
			codes.Internal,
			err,
			"failed to get collection list",
			"UserCollectionModel.FindAllByUid",
		)
	}

	var ids []int64
	for _, c := range list {
		ids = append(ids, c.ProductId)
	}
	return &user.GetUserCollectionListResponse{
		ProductIds: ids,
	}, nil
}
