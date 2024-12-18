package logic

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/product/rpc/internal/svc"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/product/rpc/product"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/dtm-labs/dtmgrpc"
)

type DecStockRevertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDecStockRevertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DecStockRevertLogic {
	return &DecStockRevertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DecStockRevertLogic) DecStockRevert(in *product.DecStockRequest) (*product.DecStockResponse, error) {
	// todo: add your logic here and delete this line
	db, err := l.svcCtx.SqlConn.RawDB()
	if err != nil {
		return nil, status.Error(codes.Internal,
			fmt.Sprintf("[svcCtx.SqlConn.RawDB]:failed to get db:%s",
				err.Error()))
	}

	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, status.Error(codes.Internal,
			fmt.Sprintf("[dtmgrpc.BarrierFromGrpc]:failed to create barrier:%s",
				err.Error()))
	}

	err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		_, err := l.svcCtx.ProductModel.TxUpdateStock(tx,
			in.GetId(), int(in.GetNum()))
		return err
	})

	if err != nil {
		return nil, err
	}
	return &product.DecStockResponse{}, nil
}
