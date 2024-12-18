// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: user.proto

package server

import (
	"context"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/user/rpc/internal/logic"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/user/rpc/internal/svc"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/user/rpc/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Login(ctx context.Context, in *user.LoginRequest) (*user.LoginResponse, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServer) GetUserInfo(ctx context.Context, in *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error) {
	l := logic.NewGetUserInfoLogic(ctx, s.svcCtx)
	return l.GetUserInfo(in)
}

func (s *UserServer) AddUserReceiveAddress(ctx context.Context, in *user.AddUserReceiveAddressRequest) (*user.AddUserReceiveAddressResponse, error) {
	l := logic.NewAddUserReceiveAddressLogic(ctx, s.svcCtx)
	return l.AddUserReceiveAddress(in)
}

func (s *UserServer) EditUserReceiveAddress(ctx context.Context, in *user.EditUserReceiveAddressRequest) (*user.EditUserReceiveAddressResponse, error) {
	l := logic.NewEditUserReceiveAddressLogic(ctx, s.svcCtx)
	return l.EditUserReceiveAddress(in)
}

func (s *UserServer) DeleteUserReceiveAddress(ctx context.Context, in *user.DeleteUserReceiveAddressRequest) (*user.DeleteUserReceiveAddressResponse, error) {
	l := logic.NewDeleteUserReceiveAddressLogic(ctx, s.svcCtx)
	return l.DeleteUserReceiveAddress(in)
}

func (s *UserServer) GetUserReceiveAddressList(ctx context.Context, in *user.GetUserReceiveAddressListRequest) (*user.GetUserReceiveAddressListResponse, error) {
	l := logic.NewGetUserReceiveAddressListLogic(ctx, s.svcCtx)
	return l.GetUserReceiveAddressList(in)
}

func (s *UserServer) AddUserCollection(ctx context.Context, in *user.AddUserCollectionRequest) (*user.AddUserCollectionResponse, error) {
	l := logic.NewAddUserCollectionLogic(ctx, s.svcCtx)
	return l.AddUserCollection(in)
}

func (s *UserServer) DeleteUserCollection(ctx context.Context, in *user.DeleteUserCollectionRequest) (*user.DeleteUserCollectionResponse, error) {
	l := logic.NewDeleteUserCollectionLogic(ctx, s.svcCtx)
	return l.DeleteUserCollection(in)
}

func (s *UserServer) GetUserCollectionList(ctx context.Context, in *user.GetUserCollectionListRequest) (*user.GetUserCollectionListResponse, error) {
	l := logic.NewGetUserCollectionListLogic(ctx, s.svcCtx)
	return l.GetUserCollectionList(in)
}

func (s *UserServer) GetUserReceiveAddressInfo(ctx context.Context, in *user.GetUserReceiveAddressInfoRequest) (*user.UserReceiveAddress, error) {
	l := logic.NewGetUserReceiveAddressInfoLogic(ctx, s.svcCtx)
	return l.GetUserReceiveAddressInfo(in)
}