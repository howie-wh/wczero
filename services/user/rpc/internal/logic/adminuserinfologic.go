package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"wczero/services/user/model"

	"wczero/services/user/rpc/internal/svc"
	"wczero/services/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminUserInfoLogic {
	return &AdminUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AdminUserInfoLogic) AdminUserInfo(in *user.AdminUserInfoRequest) (*user.AdminUserInfoResponse, error) {
	// 查询用户是否存在
	res, err := l.svcCtx.AdminModel.FindOne(in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	return &user.AdminUserInfoResponse{
		Id:     res.UserId,
		Name:   res.UserName,
		Gender: res.Gender,
		Mobile: res.Mobile,
	}, nil
}
