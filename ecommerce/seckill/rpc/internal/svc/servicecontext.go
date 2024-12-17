package svc

import (
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/product/rpc/productclient"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/seckill/rpc/internal/config"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	SeckillRedis *redis.Redis
	ProductRPC   productclient.Product
	KafkaPusher  *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		SeckillRedis: redis.New(c.SeckillRedis.Host,
			redis.WithPass(c.SeckillRedis.Pass)),
		ProductRPC:  productclient.NewProduct(zrpc.MustNewClient(c.ProductRPC)),
		KafkaPusher: kq.NewPusher(c.Kafka.Addrs, c.Kafka.SeckillTopic),
	}
}
