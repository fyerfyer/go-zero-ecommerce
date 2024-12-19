package logic

import (
	"context"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/cart/rpc/cart"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/cart/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCartItemsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListCartItemsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCartItemsLogic {
	return &ListCartItemsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListCartItemsLogic) ListCartItems(in *cart.ListCartItemsRequest) (*cart.ListCartItemsResponse, error) {
	// todo: add your logic here and delete this line

	return &cart.ListCartItemsResponse{}, nil
}
