package logic

import (
	"context"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/cart/rpc/cart"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/cart/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveFromCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveFromCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveFromCartLogic {
	return &RemoveFromCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveFromCartLogic) RemoveFromCart(in *cart.RemoveFromCartRequest) (*cart.RemoveFromCartResponse, error) {
	// todo: add your logic here and delete this line

	return &cart.RemoveFromCartResponse{}, nil
}
