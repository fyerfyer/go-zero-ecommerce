package logic

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dtm-labs/dtmgrpc"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/order/rpc/internal/svc"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/order/rpc/model"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/order/rpc/order"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/product/rpc/product"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/user/rpc/user"
	"github.com/fyerfyer/go-zero-ecommerce/pkg/e"
	"github.com/fyerfyer/go-zero-ecommerce/pkg/snowflake"
	"google.golang.org/grpc/codes"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

type SubmitOrderDTMLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubmitOrderDTMLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitOrderDTMLogic {
	return &SubmitOrderDTMLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SubmitOrderDTMLogic) SubmitOrderDTM(in *order.SubmitOrderDTMRequest) (*order.SubmitOrderDTMResponse, error) {
	// todo: add your logic here and delete this line

	// 并行处理对product和user的检查
	var (
		userRes *user.GetUserInfoResponse
		itemRes *product.GetProductByIDResponse
		addrRes *user.UserReceiveAddress
	)

	checkProduct := func() error {
		var err error
		itemRes, err = l.svcCtx.ProductRPC.GetProductByID(l.ctx,
			&product.GetProductByIDRequest{
				ProductId: in.ProductId,
			})
		return err
	}

	checkUser := func() error {
		var err error
		userRes, err = l.svcCtx.UserRPC.GetUserInfo(l.ctx,
			&user.GetUserInfoRequest{
				Id: in.GetUserId(),
			})

		return err
	}

	checkAddr := func() error {
		var err error
		addrRes, err = l.svcCtx.UserRPC.GetUserReceiveAddressInfo(l.ctx,
			&user.GetUserReceiveAddressInfoRequest{
				Id: in.GetUserId(),
			})

		return err
	}

	// 并行调用
	err := mr.Finish(checkProduct, checkUser, checkAddr)
	if itemRes == nil {
		return nil, e.HandleError(
			codes.Internal,
			err,
			"failed to get product",
			"checkProduct",
		)
	}
	if userRes == nil {
		return nil, e.HandleError(
			codes.Internal,
			err,
			"failed to get user",
			"checkUser",
		)
	}
	if addrRes == nil {
		return nil, e.HandleError(
			codes.Internal,
			err,
			"failed to get user address",
			"checkAddress",
		)
	}

	if itemRes.Product.Stock <= 0 {
		return nil, e.HandleError(
			codes.ResourceExhausted,
			errors.New("not having enough products"),
			"not having enough products",
			"Product.Stock",
		)
	}

	// 生成唯一order id
	orderID := snowflake.GenIDString()
	db, err := l.svcCtx.SqlConn.RawDB()
	if err != nil {
		return nil, e.HandleError(
			codes.Internal,
			err,
			"failed to connect to database",
			"SqlConn.RawDB()",
		)
	}

	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, e.HandleError(
			codes.Internal,
			err,
			"failed to create barrier",
			"dtmgrpc.BarrierFromGrpc",
		)
	}

	if err := barrier.CallWithDB(db, func(tx *sql.Tx) error {
		res, err := l.svcCtx.ShippingModel.TxInsert(tx,
			&model.Shipping{
				Orderid:          orderID,
				Userid:           in.GetUserId(),
				ReceiverName:     addrRes.GetName(),
				ReceiverPhone:    addrRes.GetPhone(),
				ReceiverMobile:   addrRes.GetPhone(),
				ReceiverProvince: addrRes.GetProvince(),
				ReceiverDistrict: addrRes.GetRegion(),
				ReceiverCity:     addrRes.GetCity(),
				ReceiverAddress:  addrRes.GetDetailedAddress(),
			})
		if err != nil {
			return err
		}

		// 读取创建操作的id
		shippingID, err := res.LastInsertId()
		if err != nil {
			return err
		}

		res, err = l.svcCtx.OrderItemModel.TxInsert(tx,
			&model.Orderitem{
				OrderId:      orderID,
				UserId:       in.GetUserId(),
				ProductId:    in.GetProductId(),
				ProductName:  itemRes.Product.Name,
				ProductImage: itemRes.Product.ImageUrl,
				CurrentPrice: itemRes.Product.GetPrice(),
				Quantity:     in.Quantity,
				TotalPrice:   float64(in.Quantity) * itemRes.Product.GetPrice(),
			})

		if err != nil {
			return err
		}

		// 获取插入item的id
		_, err = res.LastInsertId()
		if err != nil {
			return err
		}

		_, err = l.svcCtx.OrdersModel.TxInsert(tx,
			&model.Orders{
				Id:         orderID,
				Userid:     in.GetUserId(),
				Shoppingid: shippingID,
				Postage:    in.Postage,
			})

		return err
	}); err != nil {
		return nil, e.HandleError(
			codes.Internal,
			err,
			"submit order transaction failed",
			"barrier.CallWithDB",
		)
	}

	return &order.SubmitOrderDTMResponse{}, nil
}
