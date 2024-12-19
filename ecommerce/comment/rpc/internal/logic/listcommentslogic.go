package logic

import (
	"context"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/comment/rpc/comment"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/comment/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCommentsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCommentsLogic {
	return &ListCommentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListCommentsLogic) ListComments(in *comment.ListCommentsRequest) (*comment.ListCommentsResponse, error) {
	// todo: add your logic here and delete this line

	return &comment.ListCommentsResponse{}, nil
}
