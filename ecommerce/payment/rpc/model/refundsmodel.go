package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RefundsModel = (*customRefundsModel)(nil)

type (
	// RefundsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRefundsModel.
	RefundsModel interface {
		refundsModel
	}

	customRefundsModel struct {
		*defaultRefundsModel
	}
)

// NewRefundsModel returns a model for the database table.
func NewRefundsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) RefundsModel {
	return &customRefundsModel{
		defaultRefundsModel: newRefundsModel(conn, c, opts...),
	}
}
