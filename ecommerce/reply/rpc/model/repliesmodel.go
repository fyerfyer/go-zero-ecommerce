package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RepliesModel = (*customRepliesModel)(nil)

type (
	// RepliesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRepliesModel.
	RepliesModel interface {
		repliesModel
	}

	customRepliesModel struct {
		*defaultRepliesModel
	}
)

// NewRepliesModel returns a model for the database table.
func NewRepliesModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) RepliesModel {
	return &customRepliesModel{
		defaultRepliesModel: newRepliesModel(conn, c, opts...),
	}
}
