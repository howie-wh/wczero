package logic

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
	"os"

	"wczero/services/user/rpc/internal/svc"
	"wczero/services/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	accessKey = os.Getenv("QINIU_ACCESS_KEY1")
	secretKey = os.Getenv("QINIU_SECRET_KEY1")
	bucket    = os.Getenv("WCZERO_BUCKET")
)

type QINIUUploadTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQINIUUploadTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QINIUUploadTokenLogic {
	return &QINIUUploadTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QINIUUploadTokenLogic) QINIUUploadToken(in *user.QINIUUploadTokenRequest) (*user.QINIUUploadTokenResponse, error) {
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := auth.New(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	return &user.QINIUUploadTokenResponse{
		Domain:      l.svcCtx.Config.QiNiu.Domain,
		Zone:        l.svcCtx.Config.QiNiu.Zone,
		UploadToken: upToken,
	}, nil
}
