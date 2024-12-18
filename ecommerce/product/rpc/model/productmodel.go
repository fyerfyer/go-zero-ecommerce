package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductModel = (*customProductModel)(nil)

var (
	cacheProductProductIdPrefix = "cache:product:product:id:"
)

type (
	// ProductModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductModel.
	ProductModel interface {
		productModel
		ProductCategory(ctx context.Context, time string, cateID, limit int64) ([]*Product, error)
		UpdateProductStock(ctx context.Context, productID, num int64) error
		TxUpdateStock(tx *sql.Tx, id int64, num int) (sql.Result, error)
	}

	customProductModel struct {
		*defaultProductModel
	}
)

// NewProductModel returns a model for the database table.
func NewProductModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ProductModel {
	return &customProductModel{
		defaultProductModel: newProductModel(conn, c, opts...),
	}
}

func (p *customProductModel) ProductCategory(ctx context.Context, time string, cateID, limit int64) ([]*Product, error) {
	var products []*Product
	err := p.QueryRowsNoCacheCtx(ctx, &products, fmt.Sprintf(
		"SELECT %s FROM %s WHERE cateid=$1 AND status=1 AND create_time<$2 ORDER BY create_time DESC LIMIT $3",
		productRows, p.table), cateID, time, limit)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *customProductModel) UpdateProductStock(ctx context.Context, productID, num int64) error {
	productProductIdKey := fmt.Sprintf("%s%v", cacheProductProductIdPrefix, productID)
	_, err := p.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return conn.ExecCtx(ctx, fmt.Sprintf(
			"UPDATE %s SET stock = stock - $1 WHERE id = $2 AND stock > 0",
			p.table), num, productID)
	}, productProductIdKey)
	return err
}

func (p *customProductModel) TxUpdateStock(tx *sql.Tx, id int64, num int) (sql.Result, error) {
	productIdKey := fmt.Sprintf("%s%v", cacheProductProductIdPrefix, id)
	return p.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf(
			"UPDATE %s SET stock = stock + $1 WHERE stock >= -$2 AND id=$3",
			p.table)
		return tx.Exec(query, num, num, id)
	}, productIdKey)
}
