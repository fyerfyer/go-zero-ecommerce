package model

import (
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ShippingModel = (*customShippingModel)(nil)

type (
	// ShippingModel is an interface to be customized, add more methods here,
	// and implement the added methods in customShippingModel.
	ShippingModel interface {
		shippingModel
		TxInsert(tx *sql.Tx, data *Shipping) (sql.Result, error)
	}

	customShippingModel struct {
		*defaultShippingModel
	}
)

// NewShippingModel returns a model for the database table.
func NewShippingModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ShippingModel {
	return &customShippingModel{
		defaultShippingModel: newShippingModel(conn, c, opts...),
	}
}

// todo: 加上withctx
func (s *customShippingModel) TxInsert(tx *sql.Tx, data *Shipping) (sql.Result, error) {
	query := fmt.Sprintf(
		"INSERT INTO %s (orderid, userid, receiver_name, receiver_phone, receiver_mobile, receiver_province, receiver_city, receiver_district, receiver_address, create_time, update_time) "+
			"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW(), NOW())",
		s.table)

	return tx.Exec(query, data.Orderid, data.Userid, data.ReceiverName, data.ReceiverPhone, data.ReceiverMobile,
		data.ReceiverProvince, data.ReceiverCity, data.ReceiverDistrict, data.ReceiverAddress)
}
