package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CartItemsModel = (*customCartItemsModel)(nil)

type (
	// CartItemsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCartItemsModel.
	CartItemsModel interface {
		cartItemsModel
	}

	customCartItemsModel struct {
		*defaultCartItemsModel
	}
)

// NewCartItemsModel returns a model for the database table.
func NewCartItemsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) CartItemsModel {
	return &customCartItemsModel{
		defaultCartItemsModel: newCartItemsModel(conn, c, opts...),
	}
}
