package logic

import (
	"context"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/reply/rpc/internal/svc"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/reply/rpc/reply"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteReplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteReplyLogic {
	return &DeleteReplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteReplyLogic) DeleteReply(in *reply.DeleteReplyRequest) (*reply.DeleteReplyResponse, error) {
	// todo: add your logic here and delete this line

	return &reply.DeleteReplyResponse{}, nil
}
