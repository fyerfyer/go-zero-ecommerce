package logic

import (
	"context"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/product/rpc/internal/svc"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductsLogic {
	return &ProductsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductsLogic) Products(in *product.GetProductRequest) (*product.GetProductResponse, error) {
	// todo: add your logic here and delete this line
	// products := make(map[int64]*product.ProductItem)
	// productIDs := strings.Split(in.ProductIds, ",")
	// 使用MapReduce并行处理产品
	// 流程: 分割 -> 映射 -> 聚合
	return &product.GetProductResponse{
		Products: []*product.ProductItem{
			{
				ProductId: int64(1),
				Name:      "book",
			},
		},
	}, nil
}
