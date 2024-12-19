// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.7.3

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	paymentsFieldNames          = builder.RawFieldNames(&Payments{}, true)
	paymentsRows                = strings.Join(paymentsFieldNames, ",")
	paymentsRowsExpectAutoSet   = strings.Join(stringx.Remove(paymentsFieldNames, "id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"), ",")
	paymentsRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(paymentsFieldNames, "id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"))

	cachePublicPaymentsIdPrefix            = "cache:public:payments:id:"
	cachePublicPaymentsTransactionIdPrefix = "cache:public:payments:transactionId:"
)

type (
	paymentsModel interface {
		Insert(ctx context.Context, data *Payments) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Payments, error)
		FindOneByTransactionId(ctx context.Context, transactionId string) (*Payments, error)
		Update(ctx context.Context, data *Payments) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPaymentsModel struct {
		sqlc.CachedConn
		table string
	}

	Payments struct {
		Id            int64     `db:"id"`
		UserId        int64     `db:"user_id"`
		OrderId       string    `db:"order_id"`
		Amount        float64   `db:"amount"`
		PaymentMethod string    `db:"payment_method"`
		TransactionId string    `db:"transaction_id"`
		Status        string    `db:"status"`
		CreatedAt     time.Time `db:"created_at"`
	}
)

func newPaymentsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultPaymentsModel {
	return &defaultPaymentsModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      `"public"."payments"`,
	}
}

func (m *defaultPaymentsModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	publicPaymentsIdKey := fmt.Sprintf("%s%v", cachePublicPaymentsIdPrefix, id)
	publicPaymentsTransactionIdKey := fmt.Sprintf("%s%v", cachePublicPaymentsTransactionIdPrefix, data.TransactionId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where id = $1", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, publicPaymentsIdKey, publicPaymentsTransactionIdKey)
	return err
}

func (m *defaultPaymentsModel) FindOne(ctx context.Context, id int64) (*Payments, error) {
	publicPaymentsIdKey := fmt.Sprintf("%s%v", cachePublicPaymentsIdPrefix, id)
	var resp Payments
	err := m.QueryRowCtx(ctx, &resp, publicPaymentsIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", paymentsRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPaymentsModel) FindOneByTransactionId(ctx context.Context, transactionId string) (*Payments, error) {
	publicPaymentsTransactionIdKey := fmt.Sprintf("%s%v", cachePublicPaymentsTransactionIdPrefix, transactionId)
	var resp Payments
	err := m.QueryRowIndexCtx(ctx, &resp, publicPaymentsTransactionIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where transaction_id = $1 limit 1", paymentsRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, transactionId); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPaymentsModel) Insert(ctx context.Context, data *Payments) (sql.Result, error) {
	publicPaymentsIdKey := fmt.Sprintf("%s%v", cachePublicPaymentsIdPrefix, data.Id)
	publicPaymentsTransactionIdKey := fmt.Sprintf("%s%v", cachePublicPaymentsTransactionIdPrefix, data.TransactionId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4, $5, $6)", m.table, paymentsRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.OrderId, data.Amount, data.PaymentMethod, data.TransactionId, data.Status)
	}, publicPaymentsIdKey, publicPaymentsTransactionIdKey)
	return ret, err
}

func (m *defaultPaymentsModel) Update(ctx context.Context, newData *Payments) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	publicPaymentsIdKey := fmt.Sprintf("%s%v", cachePublicPaymentsIdPrefix, data.Id)
	publicPaymentsTransactionIdKey := fmt.Sprintf("%s%v", cachePublicPaymentsTransactionIdPrefix, data.TransactionId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, paymentsRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Id, newData.UserId, newData.OrderId, newData.Amount, newData.PaymentMethod, newData.TransactionId, newData.Status)
	}, publicPaymentsIdKey, publicPaymentsTransactionIdKey)
	return err
}

func (m *defaultPaymentsModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cachePublicPaymentsIdPrefix, primary)
}

func (m *defaultPaymentsModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", paymentsRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultPaymentsModel) tableName() string {
	return m.table
}