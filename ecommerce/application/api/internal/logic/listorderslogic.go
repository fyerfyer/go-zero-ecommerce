package logic

import (
	"context"
	"time"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/application/api/internal/svc"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/application/api/internal/types"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/order/rpc/order"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListOrdersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListOrdersLogic {
	return &ListOrdersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListOrdersLogic) ListOrders(req *types.ListOrdersRequest) (resp *types.ListOrdersResponse, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.OrderRPC.ListOrders(l.ctx, &order.ListOrdersRequest{
		UserId: req.UID,
	})

	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.ProductRPC.GetProducts(l.ctx, &product.GetProductRequest{
		ProductIds: "1,2,3",
	})

	if err != nil {
		return nil, err
	}

	return &types.ListOrdersResponse{
		Orders: []*types.Order{
			{
				OrderID:  "123",
				Status:   int32(1),
				Quantity: int64(1),
			},
		},
		IsEnd:     true,
		OrderTime: time.Now().Unix(),
	}, nil
}
