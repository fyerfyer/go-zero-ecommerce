package svc

import (
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/cart/rpc/internal/config"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/cart/rpc/model"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/product/rpc/productclient"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	ProductRPC    productclient.Product
	CartModel     model.CartsModel
	CartItemModel model.CartItemsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn("postgres", c.DataSource)

	return &ServiceContext{
		Config: c,
		ProductRPC: productclient.NewProduct(zrpc.MustNewClient(c.ProductRPC)),
		CartModel: model.NewCartsModel(conn, c.CacheRedis),
		CartItemModel: model.NewCartItemsModel(conn, c.CacheRedis),
	}
}
