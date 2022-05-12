package logic

import (
	"context"
	"wczero/services/user/rpc/userclient"

	"wczero/services/user/api/internal/svc"
	"wczero/services/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QINIUUploadTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQINIUUploadTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) QINIUUploadTokenLogic {
	return QINIUUploadTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QINIUUploadTokenLogic) QINIUUploadToken(req types.QINIUUploadTokenRequest) (resp *types.QINIUUploadTokenResponse, err error) {
	res, err := l.svcCtx.UserRpc.QINIUUploadToken(l.ctx, &userclient.QINIUUploadTokenRequest{})
	if err != nil {
		return nil, err
	}

	return &types.QINIUUploadTokenResponse{
		Domain:      res.Domain,
		Zone:        res.Zone,
		UploadToken: res.UploadToken,
		Desc:        res.Desc,
	}, nil
}
