package logic

import (
	"context"
	"database/sql"
	"errors"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/product/rpc/internal/svc"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/product/rpc/product"
	"github.com/fyerfyer/go-zero-ecommerce/pkg/e"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/dtm-labs/dtmcli"
	"github.com/dtm-labs/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"
)

type DecStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDecStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DecStockLogic {
	return &DecStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DecStockLogic) DecStock(in *product.DecStockRequest) (*product.DecStockResponse, error) {
	// todo: add your logic here and delete this line
	db, err := l.svcCtx.SqlConn.RawDB()
	if err != nil {
		return nil, e.HandleError(
			codes.Internal,
			err,
			"failed to connect to database",
			"SqlConn.RawDB()",
		)
	}

	// 使用dtm处理分布式事务
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, e.HandleError(
			codes.Internal,
			err,
			"failed to create barrier",
			"dtmgrpc.BarrierFromGrpc",
		)
	}

	err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		//更新库存
		res, err := l.svcCtx.ProductModel.TxUpdateStock(tx,
			in.GetId(), int(-in.GetNum()))
		if err != nil {
			return err
		}

		rows, err := res.RowsAffected()
		if err == nil && rows == 0 {
			return dtmcli.ErrFailure
		}

		return err
	})

	// 如果库存不足，回滚事务
	if errors.Is(err, dtmcli.ErrFailure) {
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}

	// 直接返回，数据库的tx替我们回滚了
	if err != nil {
		return nil, err
	}

	return &product.DecStockResponse{}, nil
}
