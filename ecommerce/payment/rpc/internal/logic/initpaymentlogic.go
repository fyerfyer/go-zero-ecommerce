package logic

import (
	"context"
	"fmt"
	"os"
	"sync/atomic"
	"time"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/payment/rpc/internal/svc"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/payment/rpc/model"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/payment/rpc/payment"
	"github.com/fyerfyer/go-zero-ecommerce/pkg/e"
	"google.golang.org/grpc/codes"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitPaymentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitPaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitPaymentLogic {
	return &InitPaymentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InitPaymentLogic) InitPayment(in *payment.InitPaymentRequest) (*payment.InitPaymentResponse, error) {
	// todo: add your logic here and delete this line
	txID := generateTransactionID(time.Now())
	if _, err := l.svcCtx.PaymentsModel.Insert(l.ctx, &model.Payments{
		UserId:        in.GetUserId(),
		OrderId:       in.GetOrderId(),
		Amount:        float64(in.GetAmount()),
		TransactionId: txID,
		PaymentMethod: in.GetPaymentMethod(),
		Status:        "unpaid",
		CreatedAt:     time.Now(),
	}); err != nil {
		return nil, e.HandleError(
			codes.Internal,
			err,
			"failed to insert payment",
			"PaymentsModel.Insert",
		)
	}

	return &payment.InitPaymentResponse{TransactionId: txID}, nil
}

var num int64

func generateTransactionID(t time.Time) string {
	s := t.Format("20060102150405")
	m := t.UnixNano()/1e6 - t.UnixNano()/1e9*1e3
	ms := sup(m, 3)
	p := os.Getpid() % 1000
	ps := sup(int64(p), 3)
	i := atomic.AddInt64(&num, 1)
	r := i % 10000
	rs := sup(r, 4)
	n := fmt.Sprintf("%s%s%s%s", s, ms, ps, rs)
	return n
}

func sup(i int64, n int) string {
	m := fmt.Sprintf("%d", i)
	for len(m) < n {
		m = fmt.Sprintf("0%s", m)
	}
	return m
}
