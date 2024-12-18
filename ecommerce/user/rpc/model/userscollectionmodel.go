package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UsersCollectionModel = (*customUsersCollectionModel)(nil)

type (
	// UsersCollectionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersCollectionModel.
	UsersCollectionModel interface {
		usersCollectionModel
		UpdateIsDelete(ctx context.Context, c *UsersCollection) error
		FindAllByUid(ctx context.Context, userID int64) ([]*UsersCollection, error)
	}

	customUsersCollectionModel struct {
		*defaultUsersCollectionModel
	}
)

// NewUsersCollectionModel returns a model for the database table.
func NewUsersCollectionModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UsersCollectionModel {
	return &customUsersCollectionModel{
		defaultUsersCollectionModel: newUsersCollectionModel(conn, c, opts...),
	}
}

func (u *customUsersCollectionModel) UpdateIsDelete(ctx context.Context, c *UsersCollection) error {
	userCollectionIdKey := fmt.Sprintf("%s%v", cachePublicUsersCollectionIdPrefix, c.Id)
	_, err := u.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf(`UPDATE %s SET is_delete = 1 WHERE "id" = $1`,
			u.table)
		return conn.ExecCtx(ctx, query, c.Id)
	}, userCollectionIdKey)

	return err
}

func (u *customUsersCollectionModel) FindAllByUid(ctx context.Context, userID int64) ([]*UsersCollection, error) {
	var res []*UsersCollection
	query := fmt.Sprintf(`SELECT %v FROM %s WHERE "uid" = $1 AND is_delete = 0`, userID, u.table)
	err := u.QueryRowsNoCacheCtx(ctx, &res, query, userID)
	switch err {
	case nil:
		return res, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
