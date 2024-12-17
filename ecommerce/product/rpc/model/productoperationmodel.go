package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ProductOperationModel = (*customProductOperationModel)(nil)

type (
	// ProductOperationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductOperationModel.
	ProductOperationModel interface {
		productOperationModel
		withSession(session sqlx.Session) ProductOperationModel
	}

	customProductOperationModel struct {
		*defaultProductOperationModel
	}
)

// NewProductOperationModel returns a model for the database table.
func NewProductOperationModel(conn sqlx.SqlConn) ProductOperationModel {
	return &customProductOperationModel{
		defaultProductOperationModel: newProductOperationModel(conn),
	}
}

func (m *customProductOperationModel) withSession(session sqlx.Session) ProductOperationModel {
	return NewProductOperationModel(sqlx.NewSqlConnFromSession(session))
}
