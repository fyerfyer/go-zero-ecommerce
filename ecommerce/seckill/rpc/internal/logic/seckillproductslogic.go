package logic

import (
	"context"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/seckill/rpc/internal/svc"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/seckill/rpc/seckill"

	"github.com/zeromicro/go-zero/core/logx"
)

type SeckillProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSeckillProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SeckillProductsLogic {
	return &SeckillProductsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SeckillProductsLogic) SeckillProducts(in *seckill.SeckillProductsRequest) (*seckill.SeckillProductsResponse, error) {
	// todo: add your logic here and delete this line
	return &seckill.SeckillProductsResponse{}, nil
}
