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
	repliesFieldNames          = builder.RawFieldNames(&Replies{}, true)
	repliesRows                = strings.Join(repliesFieldNames, ",")
	repliesRowsExpectAutoSet   = strings.Join(stringx.Remove(repliesFieldNames, "id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"), ",")
	repliesRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(repliesFieldNames, "id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"))

	cachePublicRepliesIdPrefix = "cache:public:replies:id:"
)

type (
	repliesModel interface {
		Insert(ctx context.Context, data *Replies) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Replies, error)
		Update(ctx context.Context, data *Replies) error
		Delete(ctx context.Context, id int64) error
	}

	defaultRepliesModel struct {
		sqlc.CachedConn
		table string
	}

	Replies struct {
		Id          int64     `db:"id"`
		CommentId   int64     `db:"comment_id"`
		UserId      int64     `db:"user_id"`
		Content     string    `db:"content"`
		CreatedTime time.Time `db:"created_time"`
	}
)

func newRepliesModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultRepliesModel {
	return &defaultRepliesModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      `"public"."replies"`,
	}
}

func (m *defaultRepliesModel) Delete(ctx context.Context, id int64) error {
	publicRepliesIdKey := fmt.Sprintf("%s%v", cachePublicRepliesIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where id = $1", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, publicRepliesIdKey)
	return err
}

func (m *defaultRepliesModel) FindOne(ctx context.Context, id int64) (*Replies, error) {
	publicRepliesIdKey := fmt.Sprintf("%s%v", cachePublicRepliesIdPrefix, id)
	var resp Replies
	err := m.QueryRowCtx(ctx, &resp, publicRepliesIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", repliesRows, m.table)
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

func (m *defaultRepliesModel) Insert(ctx context.Context, data *Replies) (sql.Result, error) {
	publicRepliesIdKey := fmt.Sprintf("%s%v", cachePublicRepliesIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4)", m.table, repliesRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.CommentId, data.UserId, data.Content, data.CreatedTime)
	}, publicRepliesIdKey)
	return ret, err
}

func (m *defaultRepliesModel) Update(ctx context.Context, data *Replies) error {
	publicRepliesIdKey := fmt.Sprintf("%s%v", cachePublicRepliesIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, repliesRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Id, data.CommentId, data.UserId, data.Content, data.CreatedTime)
	}, publicRepliesIdKey)
	return err
}

func (m *defaultRepliesModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cachePublicRepliesIdPrefix, primary)
}

func (m *defaultRepliesModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", repliesRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultRepliesModel) tableName() string {
	return m.table
}
