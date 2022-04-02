package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"wczero/common/cryptx"
	"wczero/services/user/model"

	"wczero/services/user/rpc/internal/svc"
	"wczero/services/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLoginLogic {
	return &AdminLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AdminLoginLogic) AdminLogin(in *user.AdminLoginRequest) (*user.AdminLoginResponse, error) {
	res, err := l.svcCtx.AdminModel.FindOneByMobile(in.Mobile)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "admin user not exists")
		}
		return nil, status.Error(500, err.Error())
	}

	// 判断密码是否正确
	password := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	if password != res.Password {
		return nil, status.Error(100, "密码错误")
	}

	return &user.AdminLoginResponse{
		Id:     res.UserId,
		Name:   res.UserName,
		Gender: res.Gender,
		Mobile: res.Mobile,
	}, nil
}
