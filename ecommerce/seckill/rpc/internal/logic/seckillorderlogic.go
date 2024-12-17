package logic

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/product/rpc/product"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/seckill/rpc/internal/svc"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/seckill/rpc/seckill"
	"github.com/fyerfyer/go-zero-ecommerce/pkg/batcher"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	limitPeriod       = 10
	limitQuota        = 1
	seckillUserPrefix = "seckill#u#"
	localCacheExpire  = time.Second * 60

	batcherSize     = 100
	batcherBuffer   = 100
	batcherWorker   = 10
	batcherInterval = time.Second
)

type SeckillOrderLogic struct {
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	limiter    *limit.PeriodLimit
	localCache *collection.Cache
	batcher    *batcher.Batcher
	logx.Logger
}

type KafkaMsg struct {
	UserID     int64 `json:"user_id"`
	Product_ID int64 `json:"product_id"`
}

func NewSeckillOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SeckillOrderLogic {
	localCache, err := collection.NewCache(localCacheExpire)
	if err != nil {
		panic("failed to set up local cache:" + err.Error())
	}

	logic := &SeckillOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		limiter: limit.NewPeriodLimit(limitPeriod, limitQuota,
			svcCtx.SeckillRedis, seckillUserPrefix),

		Logger:     logx.WithContext(ctx),
		localCache: localCache,
	}

	b := batcher.New(
		batcher.WithSize(batcherSize),
		batcher.WithBuffer(batcherBuffer),
		batcher.WithWorker(batcherWorker),
		batcher.WithInterval(batcherInterval),
	)

	b.Sharding = func(key string) int {
		pid, _ := strconv.ParseInt(key, 10, 64)
		return int(pid) % batcherWorker
	}

	// 发送消息到消息队列中
	b.Do = func(ctx context.Context, val map[string][]interface{}) {
		var msgs []*KafkaMsg
		for _, kv := range val {
			for _, v := range kv {
				msgs = append(msgs, v.(*KafkaMsg))
			}
		}

		data, err := json.Marshal(msgs)
		if err != nil {
			logx.Errorf("[batcher do] failed to marshal data:%v", err)
		}

		if err = logic.svcCtx.KafkaPusher.Push(ctx, string(data)); err != nil {
			logx.Errorf("[batcher do] failed to push message into kafka:%v", err)
		}
	}

	logic.batcher = b
	logic.batcher.Start()
	return logic
}

func (l *SeckillOrderLogic) SeckillOrder(in *seckill.SeckillOrderRequest) (*seckill.SeckillOrderResponse, error) {
	// todo: add your logic here and delete this line
	code, _ := l.limiter.Take(strconv.FormatInt(in.UserId, 10))
	if code == limit.OverQuota {
		return nil, status.Errorf(codes.OutOfRange, "requests number exceeded")
	}

	res, err := l.svcCtx.ProductRPC.GetProductByID(l.ctx,
		&product.GetProductByIDRequest{ProductId: in.ProductId})
	if err != nil {
		return nil, err
	}
	if res.Product.Stock <= 0 {
		return nil, status.Errorf(codes.OutOfRange, "insufficient stock")
	}

	// 提交订单到消息队列中
	if err = l.batcher.Add(strconv.FormatInt(in.ProductId, 10),
		&KafkaMsg{UserID: in.UserId, Product_ID: in.ProductId}); err != nil {
			logx.Errorf("[batcher.Add] failed to add:%v", err)
	}

	return &seckill.SeckillOrderResponse{}, nil
}
