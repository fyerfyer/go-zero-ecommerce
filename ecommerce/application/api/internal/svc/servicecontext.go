package svc

import (
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/application/api/internal/config"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/order/rpc/orderclient"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/product/rpc/productclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	OrderRPC   orderclient.Order
	ProductRPC productclient.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderRPC:   orderclient.NewOrder(zrpc.MustNewClient(c.OrderRPC)),
		ProductRPC: productclient.NewProduct(zrpc.MustNewClient(c.ProductRPC)),
	}
}
