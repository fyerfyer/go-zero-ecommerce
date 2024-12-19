package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PaymentsModel = (*customPaymentsModel)(nil)

type (
	// PaymentsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPaymentsModel.
	PaymentsModel interface {
		paymentsModel
		DeleteByTxID(ctx context.Context, txID string) (sql.Result, error)
	}

	customPaymentsModel struct {
		*defaultPaymentsModel
	}
)

// NewPaymentsModel returns a model for the database table.
func NewPaymentsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PaymentsModel {
	return &customPaymentsModel{
		defaultPaymentsModel: newPaymentsModel(conn, c, opts...),
	}
}

func (p *customPaymentsModel) DeleteByTxID(ctx context.Context, txID string) (sql.Result, error) {
	paymentsPaymentTxIdKey := fmt.Sprintf("%s%v", cachePublicPaymentsTransactionIdPrefix, txID)
	return p.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return conn.ExecCtx(ctx, fmt.Sprintf("DELETE FROM %s WHERE transaction_id = $1",
			p.table), txID)
	}, paymentsPaymentTxIdKey)
}
