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
	usersReceiveAddressFieldNames          = builder.RawFieldNames(&UsersReceiveAddress{}, true)
	usersReceiveAddressRows                = strings.Join(usersReceiveAddressFieldNames, ",")
	usersReceiveAddressRowsExpectAutoSet   = strings.Join(stringx.Remove(usersReceiveAddressFieldNames, "id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"), ",")
	usersReceiveAddressRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(usersReceiveAddressFieldNames, "id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"))

	cachePublicUsersReceiveAddressIdPrefix = "cache:public:usersReceiveAddress:id:"
)

type (
	usersReceiveAddressModel interface {
		Insert(ctx context.Context, data *UsersReceiveAddress) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UsersReceiveAddress, error)
		Update(ctx context.Context, data *UsersReceiveAddress) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUsersReceiveAddressModel struct {
		sqlc.CachedConn
		table string
	}

	UsersReceiveAddress struct {
		Id            int64     `db:"id"`
		Uid           int64     `db:"uid"`
		Name          string    `db:"name"`
		Phone         string    `db:"phone"`
		IsDefault     bool      `db:"is_default"`
		PostCode      string    `db:"post_code"`
		Province      string    `db:"province"`
		City          string    `db:"city"`
		Region        string    `db:"region"`
		DetailAddress string    `db:"detail_address"`
		IsDelete      bool      `db:"is_delete"`
		CreateTime    time.Time `db:"create_time"`
		UpdateTime    time.Time `db:"update_time"`
	}
)

func newUsersReceiveAddressModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultUsersReceiveAddressModel {
	return &defaultUsersReceiveAddressModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      `"public"."users_receive_address"`,
	}
}

func (m *defaultUsersReceiveAddressModel) Delete(ctx context.Context, id int64) error {
	publicUsersReceiveAddressIdKey := fmt.Sprintf("%s%v", cachePublicUsersReceiveAddressIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where id = $1", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, publicUsersReceiveAddressIdKey)
	return err
}

func (m *defaultUsersReceiveAddressModel) FindOne(ctx context.Context, id int64) (*UsersReceiveAddress, error) {
	publicUsersReceiveAddressIdKey := fmt.Sprintf("%s%v", cachePublicUsersReceiveAddressIdPrefix, id)
	var resp UsersReceiveAddress
	err := m.QueryRowCtx(ctx, &resp, publicUsersReceiveAddressIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", usersReceiveAddressRows, m.table)
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

func (m *defaultUsersReceiveAddressModel) Insert(ctx context.Context, data *UsersReceiveAddress) (sql.Result, error) {
	publicUsersReceiveAddressIdKey := fmt.Sprintf("%s%v", cachePublicUsersReceiveAddressIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", m.table, usersReceiveAddressRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Uid, data.Name, data.Phone, data.IsDefault, data.PostCode, data.Province, data.City, data.Region, data.DetailAddress, data.IsDelete)
	}, publicUsersReceiveAddressIdKey)
	return ret, err
}

func (m *defaultUsersReceiveAddressModel) Update(ctx context.Context, data *UsersReceiveAddress) error {
	publicUsersReceiveAddressIdKey := fmt.Sprintf("%s%v", cachePublicUsersReceiveAddressIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, usersReceiveAddressRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Id, data.Uid, data.Name, data.Phone, data.IsDefault, data.PostCode, data.Province, data.City, data.Region, data.DetailAddress, data.IsDelete)
	}, publicUsersReceiveAddressIdKey)
	return err
}

func (m *defaultUsersReceiveAddressModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cachePublicUsersReceiveAddressIdPrefix, primary)
}

func (m *defaultUsersReceiveAddressModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", usersReceiveAddressRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUsersReceiveAddressModel) tableName() string {
	return m.table
}
