package svc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/order/rpc/order"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/order/rpc/orderclient"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/product/rpc/product"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/product/rpc/productclient"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/seckill/rmq/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

const (
	chanSize   = 10
	bufferSize = 1024
)

type ServiceContext struct {
	config     config.Config
	ProductRPC productclient.Product
	OrderRPC   orderclient.Order

	wg         sync.WaitGroup
	msgChannel []chan *KafkaMsg
}

type KafkaMsg struct {
	UserID    int64 `json:"user_id"`
	ProductID int64 `json:"product_id"`
}

func NewServiceContext(conf config.Config) *ServiceContext {
	s := &ServiceContext{
		config:     conf,
		ProductRPC: productclient.NewProduct(zrpc.MustNewClient(conf.ProductRPC)),
		OrderRPC:   orderclient.NewOrder(zrpc.MustNewClient(conf.OrderRPC)),
		msgChannel: make([]chan *KafkaMsg, chanSize),
	}

	for i := 0; i < chanSize; i++ {
		ch := make(chan *KafkaMsg, bufferSize)
		s.msgChannel = append(s.msgChannel, ch)
	}

	return s
}

func (s *ServiceContext) consume(ch chan *KafkaMsg) {
	defer s.wg.Done()

	for {
		msg, ok := <-ch
		if !ok {
			log.Fatal("the channel is close")
		}

		fmt.Printf("consume msg:%v", msg)
		_, err := s.ProductRPC.CheckAndUpdateStock(context.Background(),
			&product.CheckAndUpdateStockRequest{
				ProductId: msg.ProductID,
			})
		if err != nil {
			logx.Errorf("[productRPC.CheckAndUpdateStock]:rpc method failed:%v", err)
			return
		}
		_, err = s.OrderRPC.SubmitOrderDTM(context.Background(),
			&order.SubmitOrderDTMRequest{
				UserId:    msg.UserID,
				ProductId: msg.ProductID,
			})
		if err != nil {
			logx.Errorf("[OrderRPC.CreateOrder]:rpc method failed:%v", err)
		}
	}
}

func (s *ServiceContext) Consume(ctx context.Context, _ string, value string) error {
	logx.Infof("Consume value: %s\n", value)
	var data []*KafkaMsg
	if err := json.Unmarshal([]byte(value), &data); err != nil {
		return err
	}

	// 使用分片，这样对某个ProductID的并行请求就可以被串行处理了
	for _, d := range data {
		s.msgChannel[d.ProductID%chanSize] <- d
	}
	return nil
}
