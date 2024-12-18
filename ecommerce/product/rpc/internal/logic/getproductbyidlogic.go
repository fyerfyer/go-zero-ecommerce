package logic

import (
	"context"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/product/rpc/internal/svc"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/product/rpc/product"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/user/rpc/model"
	"github.com/fyerfyer/go-zero-ecommerce/pkg/e"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"

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
	p, err := l.svcCtx.ProductModel.FindOne(l.ctx, in.GetProductId())
	if err != nil {
		if err == model.ErrNotFound {
			return nil, e.HandleError(
				codes.NotFound,
				err,
				"product not found",
				"ProductModel.FindOne",
			)
		} else {
			return nil, e.HandleError(
				codes.Internal,
				err,
				"",
				"ProductModel.FindOne",
			)
		}
	}

	var res product.GetProductByIDResponse
	_ = copier.Copy(&res, p)
	res.Product.CreateTime = p.CreateTime.Unix()
	return &res, nil
}
