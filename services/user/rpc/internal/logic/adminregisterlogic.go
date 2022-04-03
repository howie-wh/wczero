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

type AdminRegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminRegisterLogic {
	return &AdminRegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AdminRegisterLogic) AdminRegister(in *user.AdminRegisterRequest) (*user.AdminRegisterResponse, error) {
	// 判断手机号是否已经注册
	_, err := l.svcCtx.AdminModel.FindOneByMobile(in.Mobile)
	if err == nil {
		return nil, status.Error(100, "该用户已存在")
	}

	if err == model.ErrNotFound {

		newUser := model.UserAdminTab{
			UserName: in.Name,
			Gender:   in.Gender,
			Mobile:   in.Mobile,
			Password: cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
		}

		res, err := l.svcCtx.AdminModel.Insert(&newUser)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		newUser.UserId, err = res.LastInsertId()
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		return &user.AdminRegisterResponse{
			Id:     newUser.UserId,
			Name:   newUser.UserName,
			Gender: newUser.Gender,
			Mobile: newUser.Mobile,
		}, nil
	}

	return nil, status.Error(500, err.Error())
}
