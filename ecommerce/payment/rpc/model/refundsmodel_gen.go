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
	refundsFieldNames          = builder.RawFieldNames(&Refunds{}, true)
	refundsRows                = strings.Join(refundsFieldNames, ",")
	refundsRowsExpectAutoSet   = strings.Join(stringx.Remove(refundsFieldNames, "id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"), ",")
	refundsRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(refundsFieldNames, "id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"))

	cachePublicRefundsIdPrefix = "cache:public:refunds:id:"
)

type (
	refundsModel interface {
		Insert(ctx context.Context, data *Refunds) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Refunds, error)
		Update(ctx context.Context, data *Refunds) error
		Delete(ctx context.Context, id int64) error
	}

	defaultRefundsModel struct {
		sqlc.CachedConn
		table string
	}

	Refunds struct {
		Id            int64     `db:"id"`
		TransactionId string    `db:"transaction_id"`
		Amount        float64   `db:"amount"`
		RefundedAt    time.Time `db:"refunded_at"`
	}
)

func newRefundsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultRefundsModel {
	return &defaultRefundsModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      `"public"."refunds"`,
	}
}

func (m *defaultRefundsModel) Delete(ctx context.Context, id int64) error {
	publicRefundsIdKey := fmt.Sprintf("%s%v", cachePublicRefundsIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where id = $1", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, publicRefundsIdKey)
	return err
}

func (m *defaultRefundsModel) FindOne(ctx context.Context, id int64) (*Refunds, error) {
	publicRefundsIdKey := fmt.Sprintf("%s%v", cachePublicRefundsIdPrefix, id)
	var resp Refunds
	err := m.QueryRowCtx(ctx, &resp, publicRefundsIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", refundsRows, m.table)
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

func (m *defaultRefundsModel) Insert(ctx context.Context, data *Refunds) (sql.Result, error) {
	publicRefundsIdKey := fmt.Sprintf("%s%v", cachePublicRefundsIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3)", m.table, refundsRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.TransactionId, data.Amount, data.RefundedAt)
	}, publicRefundsIdKey)
	return ret, err
}

func (m *defaultRefundsModel) Update(ctx context.Context, data *Refunds) error {
	publicRefundsIdKey := fmt.Sprintf("%s%v", cachePublicRefundsIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, refundsRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Id, data.TransactionId, data.Amount, data.RefundedAt)
	}, publicRefundsIdKey)
	return err
}

func (m *defaultRefundsModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cachePublicRefundsIdPrefix, primary)
}

func (m *defaultRefundsModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", refundsRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultRefundsModel) tableName() string {
	return m.table
}