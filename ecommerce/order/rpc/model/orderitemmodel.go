package model

import (
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OrderitemModel = (*customOrderitemModel)(nil)

type (
	// OrderitemModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderitemModel.
	OrderitemModel interface {
		orderitemModel
		TxInsert(tx *sql.Tx, data *Orderitem) (sql.Result, error)
	}

	customOrderitemModel struct {
		*defaultOrderitemModel
	}
)

// NewOrderitemModel returns a model for the database table.
func NewOrderitemModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) OrderitemModel {
	return &customOrderitemModel{
		defaultOrderitemModel: newOrderitemModel(conn, c, opts...),
	}
}

func (o *customOrderitemModel) TxInsert(tx *sql.Tx, data *Orderitem) (sql.Result, error) {
	query := fmt.Sprintf(
		"INSERT INTO %s (order_id, user_id, product_id, product_name, product_image, current_price, quantity, total_price, create_time, update_time) "+
			"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), NOW())",
		o.table)
	return tx.Exec(query, data.OrderId, data.UserId, data.ProductId, data.ProductName, data.ProductImage,
		data.CurrentPrice, data.Quantity, data.TotalPrice)
}
