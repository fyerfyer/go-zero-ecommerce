package logic

import (
	"context"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/cart/rpc/cart"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/cart/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddToCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddToCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddToCartLogic {
	return &AddToCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddToCartLogic) AddToCart(in *cart.AddToCartRequest) (*cart.AddToCartResponse, error) {
	// todo: add your logic here and delete this line

	return &cart.AddToCartResponse{}, nil
}
