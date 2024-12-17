package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ProductModel = (*customProductModel)(nil)

type (
	// ProductModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductModel.
	ProductModel interface {
		productModel
		withSession(session sqlx.Session) ProductModel
	}

	customProductModel struct {
		*defaultProductModel
	}
)

// NewProductModel returns a model for the database table.
func NewProductModel(conn sqlx.SqlConn) ProductModel {
	return &customProductModel{
		defaultProductModel: newProductModel(conn),
	}
}

func (m *customProductModel) withSession(session sqlx.Session) ProductModel {
	return NewProductModel(sqlx.NewSqlConnFromSession(session))
}
