package logic

import (
	"context"
	"fmt"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/user/rpc/internal/svc"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/user/rpc/model"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/user/rpc/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserReceiveAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserReceiveAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserReceiveAddressLogic {
	return &AddUserReceiveAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserReceiveAddressLogic) AddUserReceiveAddress(in *user.AddUserReceiveAddressRequest) (*user.AddUserReceiveAddressResponse, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.UserReceiveAddressModel.Insert(l.ctx, &model.UsersReceiveAddress{
		Uid:           in.GetUserId(),
		Name:          in.GetName(),
		Phone:         in.GetPhone(),
		Province:      in.GetProvince(),
		City:          in.GetCity(),
		IsDefault:     in.GetIsDefault(),
		PostCode:      in.GetPostCode(),
		Region:        in.GetRegion(),
		DetailAddress: in.GetDetailedAddress(),
	})

	if err != nil {
		return nil, status.Error(codes.Internal,
			fmt.Sprintf("[UserReceiveAddressModel.Insert]:failed to insert:%v", err))
	}
	return &user.AddUserReceiveAddressResponse{}, nil
}
