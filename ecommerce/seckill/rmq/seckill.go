package rmq

import (
	"flag"
	"fmt"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/seckill/rmq/internal/config"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/seckill/rmq/internal/svc"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/seckill.yaml", "the etc file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	srv := svc.NewServiceContext(c)
	queue := kq.MustNewQueue(c.Kafka, kq.WithHandle(srv.Consume))
	defer queue.Stop()

	fmt.Println("seckill started")
	queue.Start()
}
