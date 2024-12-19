// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: payment.proto

package paymentclient

import (
	"context"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/payment/rpc/payment"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	InitPaymentRequest    = payment.InitPaymentRequest
	InitPaymentResponse   = payment.InitPaymentResponse
	RefundPaymentRequest  = payment.RefundPaymentRequest
	RefundPaymentResponse = payment.RefundPaymentResponse
	VerifyPaymentRequest  = payment.VerifyPaymentRequest
	VerifyPaymentResponse = payment.VerifyPaymentResponse

	Payment interface {
		InitPayment(ctx context.Context, in *InitPaymentRequest, opts ...grpc.CallOption) (*InitPaymentResponse, error)
		VerifyPayment(ctx context.Context, in *VerifyPaymentRequest, opts ...grpc.CallOption) (*VerifyPaymentResponse, error)
		RefundPayment(ctx context.Context, in *RefundPaymentRequest, opts ...grpc.CallOption) (*RefundPaymentResponse, error)
	}

	defaultPayment struct {
		cli zrpc.Client
	}
)

func NewPayment(cli zrpc.Client) Payment {
	return &defaultPayment{
		cli: cli,
	}
}

func (m *defaultPayment) InitPayment(ctx context.Context, in *InitPaymentRequest, opts ...grpc.CallOption) (*InitPaymentResponse, error) {
	client := payment.NewPaymentClient(m.cli.Conn())
	return client.InitPayment(ctx, in, opts...)
}

func (m *defaultPayment) VerifyPayment(ctx context.Context, in *VerifyPaymentRequest, opts ...grpc.CallOption) (*VerifyPaymentResponse, error) {
	client := payment.NewPaymentClient(m.cli.Conn())
	return client.VerifyPayment(ctx, in, opts...)
}

func (m *defaultPayment) RefundPayment(ctx context.Context, in *RefundPaymentRequest, opts ...grpc.CallOption) (*RefundPaymentResponse, error) {
	client := payment.NewPaymentClient(m.cli.Conn())
	return client.RefundPayment(ctx, in, opts...)
}
