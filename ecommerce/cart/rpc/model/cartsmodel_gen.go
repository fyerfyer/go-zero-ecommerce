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

	"github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	cartsFieldNames          = builder.RawFieldNames(&Carts{}, true)
	cartsRows                = strings.Join(cartsFieldNames, ",")
	cartsRowsExpectAutoSet   = strings.Join(stringx.Remove(cartsFieldNames, "id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"), ",")
	cartsRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(cartsFieldNames, "id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"))

	cachePublicCartsIdPrefix = "cache:public:carts:id:"
)

type (
	cartsModel interface {
		Insert(ctx context.Context, data *Carts) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Carts, error)
		Update(ctx context.Context, data *Carts) error
		Delete(ctx context.Context, id int64) error
	}

	defaultCartsModel struct {
		sqlc.CachedConn
		table string
	}

	Carts struct {
		Id         int64         `db:"id"`
		UserId     int64         `db:"user_id"`
		ProductIds pq.Int64Array `db:"product_ids"`
		CreatedAt  time.Time     `db:"created_at"`
	}
)

func newCartsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultCartsModel {
	return &defaultCartsModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      `"public"."carts"`,
	}
}

func (m *defaultCartsModel) Delete(ctx context.Context, id int64) error {
	publicCartsIdKey := fmt.Sprintf("%s%v", cachePublicCartsIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where id = $1", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, publicCartsIdKey)
	return err
}

func (m *defaultCartsModel) FindOne(ctx context.Context, id int64) (*Carts, error) {
	publicCartsIdKey := fmt.Sprintf("%s%v", cachePublicCartsIdPrefix, id)
	var resp Carts
	err := m.QueryRowCtx(ctx, &resp, publicCartsIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", cartsRows, m.table)
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

func (m *defaultCartsModel) Insert(ctx context.Context, data *Carts) (sql.Result, error) {
	publicCartsIdKey := fmt.Sprintf("%s%v", cachePublicCartsIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2)", m.table, cartsRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.ProductIds)
	}, publicCartsIdKey)
	return ret, err
}

func (m *defaultCartsModel) Update(ctx context.Context, data *Carts) error {
	publicCartsIdKey := fmt.Sprintf("%s%v", cachePublicCartsIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, cartsRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Id, data.UserId, data.ProductIds)
	}, publicCartsIdKey)
	return err
}

func (m *defaultCartsModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cachePublicCartsIdPrefix, primary)
}

func (m *defaultCartsModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", cartsRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultCartsModel) tableName() string {
	return m.table
}
