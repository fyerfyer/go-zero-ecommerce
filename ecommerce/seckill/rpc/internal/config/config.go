package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	ProductRPC   zrpc.RpcClientConf
	SeckillRedis redis.RedisConf
	Kafka        struct {
		Addrs        []string
		SeckillTopic string
	}
}
