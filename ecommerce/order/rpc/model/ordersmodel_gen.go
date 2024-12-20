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
	ordersFieldNames          = builder.RawFieldNames(&Orders{}, true)
	ordersRows                = strings.Join(ordersFieldNames, ",")
	ordersRowsExpectAutoSet   = strings.Join(stringx.Remove(ordersFieldNames, "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"), ",")
	ordersRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(ordersFieldNames, "id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"))

	cachePublicOrdersIdPrefix = "cache:public:orders:id:"
)

type (
	ordersModel interface {
		Insert(ctx context.Context, data *Orders) (sql.Result, error)
		FindOne(ctx context.Context, id string) (*Orders, error)
		Update(ctx context.Context, data *Orders) error
		Delete(ctx context.Context, id string) error
	}

	defaultOrdersModel struct {
		sqlc.CachedConn
		table string
	}

	Orders struct {
		Id          string    `db:"id"`
		Userid      int64     `db:"userid"`
		Shoppingid  int64     `db:"shoppingid"`
		Payment     int64     `db:"payment"`
		Paymenttype int64     `db:"paymenttype"`
		Postage     int64     `db:"postage"`
		Status      int64     `db:"status"`
		CreateTime  time.Time `db:"create_time"`
		UpdateTime  time.Time `db:"update_time"`
	}
)

func newOrdersModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultOrdersModel {
	return &defaultOrdersModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      `"public"."orders"`,
	}
}

func (m *defaultOrdersModel) Delete(ctx context.Context, id string) error {
	publicOrdersIdKey := fmt.Sprintf("%s%v", cachePublicOrdersIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where id = $1", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, publicOrdersIdKey)
	return err
}

func (m *defaultOrdersModel) FindOne(ctx context.Context, id string) (*Orders, error) {
	publicOrdersIdKey := fmt.Sprintf("%s%v", cachePublicOrdersIdPrefix, id)
	var resp Orders
	err := m.QueryRowCtx(ctx, &resp, publicOrdersIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", ordersRows, m.table)
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

func (m *defaultOrdersModel) Insert(ctx context.Context, data *Orders) (sql.Result, error) {
	publicOrdersIdKey := fmt.Sprintf("%s%v", cachePublicOrdersIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4, $5, $6, $7)", m.table, ordersRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.Userid, data.Shoppingid, data.Payment, data.Paymenttype, data.Postage, data.Status)
	}, publicOrdersIdKey)
	return ret, err
}

func (m *defaultOrdersModel) Update(ctx context.Context, data *Orders) error {
	publicOrdersIdKey := fmt.Sprintf("%s%v", cachePublicOrdersIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, ordersRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Id, data.Userid, data.Shoppingid, data.Payment, data.Paymenttype, data.Postage, data.Status)
	}, publicOrdersIdKey)
	return err
}

func (m *defaultOrdersModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cachePublicOrdersIdPrefix, primary)
}

func (m *defaultOrdersModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", ordersRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultOrdersModel) tableName() string {
	return m.table
}
