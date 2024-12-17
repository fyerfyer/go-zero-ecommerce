package logic

import (
	"context"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/product/rpc/internal/svc"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DecStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDecStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DecStockLogic {
	return &DecStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DecStockLogic) DecStock(in *product.DecStockRequest) (*product.DecStockResponse, error) {
	// todo: add your logic here and delete this line

	return &product.DecStockResponse{}, nil
}
