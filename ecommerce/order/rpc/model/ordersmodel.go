package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OrdersModel = (*customOrdersModel)(nil)

var cacheOrdersIdPrefix = "cache:orders:id:"

type (
	// OrdersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrdersModel.
	OrdersModel interface {
		ordersModel
		FindOneByID(ctx context.Context, id int64) (*Orders, error)
		Create(ctx context.Context, orderID string, userID, productID int64) error
		UpdateStatus(ctx context.Context, orderID string, status int) error
		TxInsert(tx *sql.Tx, data *Orders) (sql.Result, error)
		TxUpdate(tx *sql.Tx, data *Orders) error
	}

	customOrdersModel struct {
		*defaultOrdersModel
	}
)

// NewOrdersModel returns a model for the database table.
func NewOrdersModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) OrdersModel {
	return &customOrdersModel{
		defaultOrdersModel: newOrdersModel(conn, c, opts...),
	}
}

func (o *customOrdersModel) FindOneByID(ctx context.Context, uid int64) (*Orders, error) {
	var resp Orders

	query := fmt.Sprintf("select %s from %s where \"uid\" = $1 order by create_time desc limit 1",
		ordersRows, o.table)

	err := o.QueryRowNoCacheCtx(ctx, &resp, query, uid)

	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (o *customOrdersModel) Create(ctx context.Context, orderID string, userID, productID int64) error {
	_, err := o.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		// 由于涉及两张表的插入，我们使用事务来保证数据一致性
		err := conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
			_, err := session.ExecCtx(ctx, "INSERT INTO orders(id, userid) VALUES($1, $2)",
				orderID, userID)
			if err != nil {
				return err
			}
			_, err = session.ExecCtx(ctx, "INSERT INTO orderitem(orderid, userid, proid) VALUES($1, $2, $3)",
				orderID, userID, productID)
			return err
		})
		return nil, err
	})
	return err
}

func (o *customOrdersModel) UpdateStatus(ctx context.Context, orderID string, status int) error {
	ordersOrdersIdKey := fmt.Sprintf("%s%v", cacheOrdersIdPrefix, orderID)
	_, err := o.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return conn.ExecCtx(ctx, fmt.Sprintf("UPDATE %s SET status = $1 WHERE id = $2", o.table),
			status, orderID)
	}, ordersOrdersIdKey)
	return err
}

func (o *customOrdersModel) TxUpdate(tx *sql.Tx, data *Orders) error {
	productIdKey := fmt.Sprintf("%s%v", cacheOrdersIdPrefix, data.Id)
	_, err := o.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d",
			o.table, ordersRowsWithPlaceHolder, len(ordersRowsWithPlaceHolder)+1)
		return tx.Exec(query, data.Userid, data.Shoppingid,
			data.Payment, data.Paymenttype, data.Postage, data.Status, data.Id)
	}, productIdKey)
	return err
}

func (o *customOrdersModel) TxInsert(tx *sql.Tx, data *Orders) (sql.Result, error) {
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		o.table, ordersRowsExpectAutoSet)
	ret, err := tx.Exec(query, data.Id, data.Userid, data.Shoppingid,
		data.Payment, data.Paymenttype, data.Postage, data.Status)
	return ret, err
}
