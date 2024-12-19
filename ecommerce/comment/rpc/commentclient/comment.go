// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: comment.proto

package commentclient

import (
	"context"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/comment/rpc/comment"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddCommentRequest     = comment.AddCommentRequest
	AddCommentResponse    = comment.AddCommentResponse
	CommentItem           = comment.CommentItem
	DeleteCommentRequest  = comment.DeleteCommentRequest
	DeleteCommentResponse = comment.DeleteCommentResponse
	ListCommentsRequest   = comment.ListCommentsRequest
	ListCommentsResponse  = comment.ListCommentsResponse

	Comment interface {
		AddComment(ctx context.Context, in *AddCommentRequest, opts ...grpc.CallOption) (*AddCommentResponse, error)
		DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error)
		ListComments(ctx context.Context, in *ListCommentsRequest, opts ...grpc.CallOption) (*ListCommentsResponse, error)
	}

	defaultComment struct {
		cli zrpc.Client
	}
)

func NewComment(cli zrpc.Client) Comment {
	return &defaultComment{
		cli: cli,
	}
}

func (m *defaultComment) AddComment(ctx context.Context, in *AddCommentRequest, opts ...grpc.CallOption) (*AddCommentResponse, error) {
	client := comment.NewCommentClient(m.cli.Conn())
	return client.AddComment(ctx, in, opts...)
}

func (m *defaultComment) DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error) {
	client := comment.NewCommentClient(m.cli.Conn())
	return client.DeleteComment(ctx, in, opts...)
}

func (m *defaultComment) ListComments(ctx context.Context, in *ListCommentsRequest, opts ...grpc.CallOption) (*ListCommentsResponse, error) {
	client := comment.NewCommentClient(m.cli.Conn())
	return client.ListComments(ctx, in, opts...)
}