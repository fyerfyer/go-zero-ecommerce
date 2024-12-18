package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UsersReceiveAddressModel = (*customUsersReceiveAddressModel)(nil)

type (
	// UsersReceiveAddressModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersReceiveAddressModel.
	UsersReceiveAddressModel interface {
		usersReceiveAddressModel
		UpdateIsDelete(ctx context.Context, ra *UsersReceiveAddress) error
		FindAllByUid(ctx context.Context, userID int64) ([]*UsersReceiveAddress, error)
	}

	customUsersReceiveAddressModel struct {
		*defaultUsersReceiveAddressModel
	}
)

// NewUsersReceiveAddressModel returns a model for the database table.
func NewUsersReceiveAddressModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UsersReceiveAddressModel {
	return &customUsersReceiveAddressModel{
		defaultUsersReceiveAddressModel: newUsersReceiveAddressModel(conn, c, opts...),
	}
}

func (u *customUsersReceiveAddressModel) UpdateIsDelete(ctx context.Context, ra *UsersReceiveAddress) error {
	userAddressIdKey := fmt.Sprintf("%s%v", cachePublicUsersReceiveAddressIdPrefix, ra.Id)
	_, err := u.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf(`UPDATE %s SET is_delete = 1 WHERE "id" = $1`,
			u.table)
		return conn.ExecCtx(ctx, query, ra.Id)
	}, userAddressIdKey)

	return err
}

func (u *customUsersReceiveAddressModel) FindAllByUid(ctx context.Context, userID int64) ([]*UsersReceiveAddress, error) {
	var res []*UsersReceiveAddress
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
