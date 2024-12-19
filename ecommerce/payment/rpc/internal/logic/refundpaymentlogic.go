package logic

import (
	"context"
	"errors"
	"time"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/payment/rpc/internal/svc"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/payment/rpc/model"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/payment/rpc/payment"
	"github.com/fyerfyer/go-zero-ecommerce/pkg/e"
	"google.golang.org/grpc/codes"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefundPaymentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRefundPaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefundPaymentLogic {
	return &RefundPaymentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RefundPaymentLogic) RefundPayment(in *payment.RefundPaymentRequest) (*payment.RefundPaymentResponse, error) {
	// todo: add your logic here and delete this line
	
	// 先看这个payment是否存在
	res, err := l.svcCtx.PaymentsModel.FindOneByTransactionId(l.ctx, in.GetTransactionId())
	if err != nil {
		return nil, e.HandleError(
			codes.NotFound,
			err,
			"payment not found",
			"PaymentsModel.FindOneByTransactionId",
		)
	}

	if res.Status != "completed" {
		return nil, e.HandleError( 
			codes.InvalidArgument,
			errors.New("not completed payment"),
			"not completed payment",
			"res.Status",
		)
	}

	if _, err := l.svcCtx.PaymentsModel.DeleteByTxID(l.ctx, in.GetTransactionId()); err != nil {
		return nil, e.HandleError(
			codes.Internal,
			err,
			"failed to delete payment",
			"PaymentsModel.DeleteByTxID",
		)
	}

	if _, err := l.svcCtx.RefundModel.Insert(l.ctx, &model.Refunds {
		TransactionId: in.GetTransactionId(),
		Amount: res.Amount,
		RefundedAt: time.Now(),
	}); err != nil {
		return nil, e.HandleError(
			codes.Internal,
			err,
			"failed to insert refund",
			"RefundModel.Insert",
		)
	}
	return &payment.RefundPaymentResponse{}, nil
}
