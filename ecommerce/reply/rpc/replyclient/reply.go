// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: reply.proto

package replyclient

import (
	"context"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/reply/rpc/reply"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddReplyRequest     = reply.AddReplyRequest
	AddReplyResponse    = reply.AddReplyResponse
	DeleteReplyRequest  = reply.DeleteReplyRequest
	DeleteReplyResponse = reply.DeleteReplyResponse
	ListRepliesRequest  = reply.ListRepliesRequest
	ListRepliesResponse = reply.ListRepliesResponse
	ReplyItem           = reply.ReplyItem

	Reply interface {
		AddReply(ctx context.Context, in *AddReplyRequest, opts ...grpc.CallOption) (*AddReplyResponse, error)
		DeleteReply(ctx context.Context, in *DeleteReplyRequest, opts ...grpc.CallOption) (*DeleteReplyResponse, error)
		ListReplies(ctx context.Context, in *ListRepliesRequest, opts ...grpc.CallOption) (*ListRepliesResponse, error)
	}

	defaultReply struct {
		cli zrpc.Client
	}
)

func NewReply(cli zrpc.Client) Reply {
	return &defaultReply{
		cli: cli,
	}
}

func (m *defaultReply) AddReply(ctx context.Context, in *AddReplyRequest, opts ...grpc.CallOption) (*AddReplyResponse, error) {
	client := reply.NewReplyClient(m.cli.Conn())
	return client.AddReply(ctx, in, opts...)
}

func (m *defaultReply) DeleteReply(ctx context.Context, in *DeleteReplyRequest, opts ...grpc.CallOption) (*DeleteReplyResponse, error) {
	client := reply.NewReplyClient(m.cli.Conn())
	return client.DeleteReply(ctx, in, opts...)
}

func (m *defaultReply) ListReplies(ctx context.Context, in *ListRepliesRequest, opts ...grpc.CallOption) (*ListRepliesResponse, error) {
	client := reply.NewReplyClient(m.cli.Conn())
	return client.ListReplies(ctx, in, opts...)
}
