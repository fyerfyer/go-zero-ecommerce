package logic

import (
	"context"
	"errors"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/user/rpc/internal/svc"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/user/rpc/model"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/user/rpc/user"
	"github.com/fyerfyer/go-zero-ecommerce/pkg/e"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditUserReceiveAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditUserReceiveAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditUserReceiveAddressLogic {
	return &EditUserReceiveAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EditUserReceiveAddressLogic) EditUserReceiveAddress(in *user.EditUserReceiveAddressRequest) (*user.EditUserReceiveAddressResponse, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.UserReceiveAddressModel.FindOne(l.ctx, in.GetId())
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, e.HandleError(
				codes.Internal,
				err,
				"address not found",
				"UserReceiveAddressModel.FindOne")
		} else {
			return nil, e.HandleError(codes.Internal,
				err,
				"",
				"UserReceiveAddressModel.FindOne")
		}
	}

	if in.IsDefault {
		// 查找该用户是否已有默认地址
		addrs, err := l.svcCtx.UserReceiveAddressModel.FindAllByUid(l.ctx, in.GetId())
		if err != nil {
			return nil, e.HandleError(
				codes.Internal,
				err,
				"",
				"UserReceiveAddressModel.FindAllByUid",
			)
		}

		for _, addr := range addrs {
			if addr.IsDefault {
				return nil, e.HandleError(
					codes.AlreadyExists,
					err,
					"default address already exists",
					"EditUserReceiveAddress",
				)
			}
		}
	}

	addr := new(model.UsersReceiveAddress)
	_ = copier.Copy(&addr, in)
	err = l.svcCtx.UserReceiveAddressModel.Update(l.ctx, addr)
	if err != nil {
		return nil, e.HandleError(
			codes.Internal,
			err,
			"failed to update address",
			"UserReceiveAddressModel.Update",
		)
	}
	return &user.EditUserReceiveAddressResponse{}, nil
}
