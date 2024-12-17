package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ OrderitemModel = (*customOrderitemModel)(nil)

type (
	// OrderitemModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderitemModel.
	OrderitemModel interface {
		orderitemModel
		withSession(session sqlx.Session) OrderitemModel
	}

	customOrderitemModel struct {
		*defaultOrderitemModel
	}
)

// NewOrderitemModel returns a model for the database table.
func NewOrderitemModel(conn sqlx.SqlConn) OrderitemModel {
	return &customOrderitemModel{
		defaultOrderitemModel: newOrderitemModel(conn),
	}
}

func (m *customOrderitemModel) withSession(session sqlx.Session) OrderitemModel {
	return NewOrderitemModel(sqlx.NewSqlConnFromSession(session))
}
