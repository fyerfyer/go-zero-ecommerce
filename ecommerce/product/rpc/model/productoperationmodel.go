package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductOperationModel = (*customProductOperationModel)(nil)

type (
	// ProductOperationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductOperationModel.
	ProductOperationModel interface {
		productOperationModel
	}

	customProductOperationModel struct {
		*defaultProductOperationModel
	}
)

// NewProductOperationModel returns a model for the database table.
func NewProductOperationModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ProductOperationModel {
	return &customProductOperationModel{
		defaultProductOperationModel: newProductOperationModel(conn, c, opts...),
	}
}
