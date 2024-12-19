package logic

import (
	"context"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/reply/rpc/internal/svc"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/reply/rpc/reply"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListRepliesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListRepliesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListRepliesLogic {
	return &ListRepliesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListRepliesLogic) ListReplies(in *reply.ListRepliesRequest) (*reply.ListRepliesResponse, error) {
	// todo: add your logic here and delete this line

	return &reply.ListRepliesResponse{}, nil
}