package logic

import (
	"context"
	"encoding/json"
	"wczero/services/user/rpc/userclient"

	"wczero/services/user/api/internal/svc"
	"wczero/services/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) AdminUserInfoLogic {
	return AdminUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminUserInfoLogic) AdminUserInfo() (*types.AdminUserInfoResponse, error) {
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	res, err := l.svcCtx.UserRpc.AdminUserInfo(l.ctx, &userclient.AdminUserInfoRequest{
		Id: uid,
	})
	if err != nil {
		return nil, err
	}

	return &types.AdminUserInfoResponse{
		Id:     res.Id,
		Name:   res.Name,
		Gender: res.Gender,
		Mobile: res.Mobile,
	}, nil
}
