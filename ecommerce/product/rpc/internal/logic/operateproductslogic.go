package logic

import (
	"context"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/product/rpc/internal/svc"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type OperateProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOperateProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OperateProductsLogic {
	return &OperateProductsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OperateProductsLogic) OperateProducts(in *product.OperateProductsRequest) (*product.OperateProductsResponse, error) {
	// todo: add your logic here and delete this line

	return &product.OperateProductsResponse{}, nil
}
