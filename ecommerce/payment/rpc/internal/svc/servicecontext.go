package svc

import (
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/payment/rpc/internal/config"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/payment/rpc/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	PaymentsModel model.PaymentsModel
	RefundModel   model.RefundsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn("postgres", c.DataSource)
	return &ServiceContext{
		Config:        c,
		PaymentsModel: model.NewPaymentsModel(conn, c.CacheRedis),
		RefundModel:   model.NewRefundsModel(conn, c.CacheRedis),
	}
}
