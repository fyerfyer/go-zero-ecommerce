package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CartsModel = (*customCartsModel)(nil)

type (
	// CartsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCartsModel.
	CartsModel interface {
		cartsModel
	}

	customCartsModel struct {
		*defaultCartsModel
	}
)

// NewCartsModel returns a model for the database table.
func NewCartsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) CartsModel {
	return &customCartsModel{
		defaultCartsModel: newCartsModel(conn, c, opts...),
	}
}
