package logic

import (
	"context"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/user/rpc/internal/svc"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/user/rpc/user"
	"github.com/fyerfyer/go-zero-ecommerce/pkg/e"
	"github.com/fyerfyer/go-zero-ecommerce/pkg/encrypt"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	// todo: add your logic here and delete this line
	// 确认用户是否存在
	u, err := l.svcCtx.UsersModel.FindOneByUsername(l.ctx, in.GetUsername())
	if err != nil {
		return nil, e.HandleError(
			codes.NotFound,
			err,
			"user not found",
			"UsersModel.FindOneByUsername",
		)
	}

	// 验证密码
	md5String, err := encrypt.Md5String(in.Password)
	if md5String != u.Password {
		return nil, e.HandleError(
			codes.Unauthenticated,
			err,
			"wrong password",
			"encrypt.Md5String",
		)
	}

	res := new(user.LoginResponse)
	_ = copier.Copy(res, u)
	return res, nil
}
