package logic

import (
	"context"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/comment/rpc/comment"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/comment/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCommentLogic {
	return &AddCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCommentLogic) AddComment(in *comment.AddCommentRequest) (*comment.AddCommentResponse, error) {
	// todo: add your logic here and delete this line

	return &comment.AddCommentResponse{}, nil
}
