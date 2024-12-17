package logic

import (
	"context"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/order/rpc/internal/svc"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderByIDLogic {
	return &GetOrderByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderByIDLogic) GetOrderByID(in *order.GetOrderRequest) (*order.GetOrderResponse, error) {
	// todo: add your logic here and delete this line

	return &order.GetOrderResponse{}, nil
}
