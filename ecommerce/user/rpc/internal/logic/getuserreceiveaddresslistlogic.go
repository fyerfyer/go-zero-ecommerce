package logic

import (
	"context"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/user/rpc/internal/svc"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/user/rpc/user"
	"github.com/fyerfyer/go-zero-ecommerce/pkg/e"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserReceiveAddressListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserReceiveAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserReceiveAddressListLogic {
	return &GetUserReceiveAddressListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserReceiveAddressListLogic) GetUserReceiveAddressList(in *user.GetUserReceiveAddressListRequest) (*user.GetUserReceiveAddressListResponse, error) {
	// todo: add your logic here and delete this line
	list, err := l.svcCtx.UserReceiveAddressModel.FindAllByUid(l.ctx, in.GetUserId())
	if err != nil {
		return nil, e.HandleError(
			codes.Internal, 
			err, 
			"failed to get address list", 
			"UserReceiveAddressModel.FindAllByUid",
		)
	}

	var res []*user.UserReceiveAddress
	for _, addr := range list {
		pbAddr := new(user.UserReceiveAddress)
		_ = copier.Copy(&pbAddr, addr)
		res = append(res, pbAddr)
	}

	return &user.GetUserReceiveAddressListResponse{
		List: res,
	}, nil
}
