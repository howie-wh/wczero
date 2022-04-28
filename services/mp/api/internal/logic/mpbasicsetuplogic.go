package logic

import (
	"context"
	"crypto/sha1"
	"fmt"
	"sort"

	"wczero/services/mp/api/internal/svc"
	"wczero/services/mp/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MPBasicSetupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMPBasicSetupLogic(ctx context.Context, svcCtx *svc.ServiceContext) MPBasicSetupLogic {
	return MPBasicSetupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MPBasicSetupLogic) MPBasicSetup(req types.MPBasicSetupRequest) (resp *types.MPBasicSetupResponse, err error) {
	paramList := []string{l.svcCtx.Config.MP.Token, req.Timestamp, req.Nonce}
	sort.Strings(paramList)

	h := sha1.New()
	for _, msg := range paramList {
		h.Write([]byte(msg))
	}
	hashCode := h.Sum(nil)
	signature := fmt.Sprintf("%x", hashCode)

	var reply int64
	if signature == req.Signature {
		reply = req.EchoStr
	}

	return &types.MPBasicSetupResponse{
		Reply: reply,
	}, nil
}
