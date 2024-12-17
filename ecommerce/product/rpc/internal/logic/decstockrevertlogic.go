package logic

import (
	"context"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/product/rpc/internal/svc"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DecStockRevertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDecStockRevertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DecStockRevertLogic {
	return &DecStockRevertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DecStockRevertLogic) DecStockRevert(in *product.DecStockRequest) (*product.DecStockResponse, error) {
	// todo: add your logic here and delete this line

	return &product.DecStockResponse{}, nil
}
