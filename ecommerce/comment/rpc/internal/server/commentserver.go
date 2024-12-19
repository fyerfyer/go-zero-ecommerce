// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: comment.proto

package server

import (
	"context"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/comment/rpc/comment"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/comment/rpc/internal/logic"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/comment/rpc/internal/svc"
)

type CommentServer struct {
	svcCtx *svc.ServiceContext
	comment.UnimplementedCommentServer
}

func NewCommentServer(svcCtx *svc.ServiceContext) *CommentServer {
	return &CommentServer{
		svcCtx: svcCtx,
	}
}

func (s *CommentServer) AddComment(ctx context.Context, in *comment.AddCommentRequest) (*comment.AddCommentResponse, error) {
	l := logic.NewAddCommentLogic(ctx, s.svcCtx)
	return l.AddComment(in)
}

func (s *CommentServer) DeleteComment(ctx context.Context, in *comment.DeleteCommentRequest) (*comment.DeleteCommentResponse, error) {
	l := logic.NewDeleteCommentLogic(ctx, s.svcCtx)
	return l.DeleteComment(in)
}

func (s *CommentServer) ListComments(ctx context.Context, in *comment.ListCommentsRequest) (*comment.ListCommentsResponse, error) {
	l := logic.NewListCommentsLogic(ctx, s.svcCtx)
	return l.ListComments(in)
}