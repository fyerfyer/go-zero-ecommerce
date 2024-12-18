package svc

import (
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/order/rpc/internal/config"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/order/rpc/model"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/product/rpc/productclient"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/user/rpc/userclient"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	OrdersModel    model.OrdersModel
	OrderItemModel model.OrderitemModel
	ShippingModel  model.ShippingModel
	UserRPC        userclient.User
	ProductRPC     productclient.Product
	SqlConn        sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn("postgres", c.DataSource)

	return &ServiceContext{
		Config:         c,
		OrdersModel:    model.NewOrdersModel(conn, c.CacheRedis),
		OrderItemModel: model.NewOrderitemModel(conn, c.CacheRedis),
		ShippingModel:  model.NewShippingModel(conn, c.CacheRedis),
		UserRPC:        userclient.NewUser(zrpc.MustNewClient(c.UserRPC)),
		ProductRPC:     productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
		SqlConn:        conn,
	}
}
