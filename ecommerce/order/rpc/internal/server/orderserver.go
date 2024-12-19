// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: order.proto

package server

import (
	"context"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/order/rpc/internal/logic"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/order/rpc/internal/svc"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/order/rpc/order"
)

type OrderServer struct {
	svcCtx *svc.ServiceContext
	order.UnimplementedOrderServer
}

func NewOrderServer(svcCtx *svc.ServiceContext) *OrderServer {
	return &OrderServer{
		svcCtx: svcCtx,
	}
}

func (s *OrderServer) ListOrders(ctx context.Context, in *order.ListOrdersRequest) (*order.ListOrdersResponse, error) {
	l := logic.NewListOrdersLogic(ctx, s.svcCtx)
	return l.ListOrders(in)
}

func (s *OrderServer) SubmitOrderDTM(ctx context.Context, in *order.SubmitOrderDTMRequest) (*order.SubmitOrderDTMResponse, error) {
	l := logic.NewSubmitOrderDTMLogic(ctx, s.svcCtx)
	return l.SubmitOrderDTM(in)
}
