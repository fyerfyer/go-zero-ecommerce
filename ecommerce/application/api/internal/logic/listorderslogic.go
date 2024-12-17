package logic

import (
	"context"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/application/api/internal/svc"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/application/api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type ListOrdersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListOrdersLogic {
	return &ListOrdersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListOrdersLogic) ListOrders(req *types.ListOrdersRequest) (resp *types.ListOrdersResponse, err error) {
	// todo: add your logic here and delete this line
	return &types.ListOrdersResponse{}, nil
}
