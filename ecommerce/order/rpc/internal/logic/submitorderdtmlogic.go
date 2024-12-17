package logic

import (
	"context"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/order/rpc/internal/svc"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubmitOrderDTMLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubmitOrderDTMLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitOrderDTMLogic {
	return &SubmitOrderDTMLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SubmitOrderDTMLogic) SubmitOrderDTM(in *order.SubmitOrderDTMRequest) (*order.SubmitOrderDTMResponse, error) {
	// todo: add your logic here and delete this line

	return &order.SubmitOrderDTMResponse{}, nil
}
