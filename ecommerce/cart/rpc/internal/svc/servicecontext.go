package svc

import (
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/cart/rpc/internal/config"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/cart/rpc/model"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/product/rpc/productclient"
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
	}
}
