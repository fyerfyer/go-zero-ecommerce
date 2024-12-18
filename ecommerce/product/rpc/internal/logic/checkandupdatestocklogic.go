package logic

import (
	"context"
	"fmt"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/product/rpc/internal/svc"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/product/rpc/product"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckAndUpdateStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckAndUpdateStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckAndUpdateStockLogic {
	return &CheckAndUpdateStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 使用Lua脚本确保并发安全
// HINCRBY是Redis的原子操作，是并发安全的
const (
	luaCheckAndUpdateScript = `
local counts = redis.call("HMGET", KEYS[1], "total", "seckill")
local total = tonumber(counts[1])
local seckill = tonumber(counts[2])
if seckill + 1 <= total then
	redis.call("HINCRBY", KEYS[1], "seckill", 1)
	return 1
end
return 0
`
)

func (l *CheckAndUpdateStockLogic) CheckAndUpdateStock(in *product.CheckAndUpdateStockRequest) (*product.CheckAndUpdateStockResponse, error) {
	// todo: add your logic here and delete this line

	// 在Product的运行中缓存中执行脚本
	val, err := l.svcCtx.ProductRedis.EvalCtx(l.ctx,
		luaCheckAndUpdateScript,
		[]string{fmt.Sprintf("stock:%d", in.GetProductId())})
	if err != nil {
		return nil, err
	}

	if val.(int64) == 0 {
		return nil, status.Errorf(codes.ResourceExhausted,
			fmt.Sprintf("insufficient stock: %d", in.ProductId))
	}

	return &product.CheckAndUpdateStockResponse{}, nil
}
