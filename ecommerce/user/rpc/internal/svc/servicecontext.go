package svc

import (
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/user/rpc/internal/config"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/user/rpc/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config                  config.Config
	UsersModel              model.UsersModel
	UserReceiveAddressModel model.UsersReceiveAddressModel
	UserCollectionModel     model.UsersCollectionModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewSqlConn("postgres", c.DataSource)
	return &ServiceContext{
		Config:                  c,
		UsersModel:              model.NewUsersModel(sqlConn, c.CacheRedis),
		UserReceiveAddressModel: model.NewUsersReceiveAddressModel(sqlConn, c.CacheRedis),
		UserCollectionModel:     model.NewUsersCollectionModel(sqlConn, c.CacheRedis),
	}
}
