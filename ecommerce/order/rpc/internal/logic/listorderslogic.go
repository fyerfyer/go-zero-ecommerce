package logic

import (
	"context"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/order/rpc/internal/svc"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListOrdersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListOrdersLogic {
	return &ListOrdersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListOrdersLogic) ListOrders(in *order.ListOrdersRequest) (*order.ListOrdersResponse, error) {
	// todo: add your logic here and delete this line
	uid := int64(123)
	if in.UserId == uid {
		orders := []*order.OrderItem{
			{
				OrderId:   "20220609123456",
				UserId:    uid,
				ProductId: 1,
				Quantity:  1,
			},
		}
		return &order.ListOrdersResponse{Orders: orders}, nil
	}
	return &order.ListOrdersResponse{}, nil
}
