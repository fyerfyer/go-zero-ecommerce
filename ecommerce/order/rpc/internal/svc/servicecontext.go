package svc

import (
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/order/rpc/internal/config"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/order/rpc/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	OrdersModel model.OrdersModel
	OrderItem   model.OrderitemModel
	Shipping    model.ShippingModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)

	return &ServiceContext{
		Config:      c,
		OrdersModel: model.NewOrdersModel(conn, c.CacheRedis),
		OrderItem:   model.NewOrderitemModel(conn, c.CacheRedis),
		Shipping:    model.NewShippingModel(conn, c.CacheRedis),
	}
}
