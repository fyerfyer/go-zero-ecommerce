package logic

import (
	"context"
	"errors"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/payment/rpc/internal/svc"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/payment/rpc/payment"
	"github.com/fyerfyer/go-zero-ecommerce/pkg/e"
	"google.golang.org/grpc/codes"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyPaymentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyPaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyPaymentLogic {
	return &VerifyPaymentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifyPaymentLogic) VerifyPayment(in *payment.VerifyPaymentRequest) (*payment.VerifyPaymentResponse, error) {
	// todo: add your logic here and delete this line
	p, err := l.svcCtx.PaymentsModel.FindOneByTransactionId(l.ctx, in.GetTransactionId())
	if err != nil {
		return nil, e.HandleError(
			codes.NotFound,
			err,
			"payment not found",
			"PaymentsModel.FindOneByTransactionId",
		)
	}

	if p.Status != "unpaid" {
		return nil, e.HandleError(
			codes.Unavailable,
			errors.New("payment has done"),
			"payment has done",
			"p.Status",
		)
	}
	return &payment.VerifyPaymentResponse{}, nil
}
