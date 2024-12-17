package logic

import (
	"context"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/product/rpc/internal/svc"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductByIDLogic {
	return &GetProductByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductByIDLogic) GetProductByID(in *product.GetProductByIDRequest) (*product.GetProductByIDResponse, error) {
	// todo: add your logic here and delete this line

	return &product.GetProductByIDResponse{}, nil
}
