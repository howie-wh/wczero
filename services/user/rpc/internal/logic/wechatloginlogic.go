package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"wczero/services/user/rpc/internal/svc"
	"wczero/services/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type WeChatLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWeChatLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WeChatLoginLogic {
	return &WeChatLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WeChatLoginLogic) WeChatLogin(in *user.WeChatLoginRequest) (*user.WeChatLoginResponse, error) {
	params := url.Values{}
	Url, err := url.Parse("https://api.weixin.qq.com/sns/jscode2session")
	if err != nil {
		return nil, err
	}
	params.Set("appid", in.Appid)
	params.Set("secret", l.svcCtx.Config.Salt)
	params.Set("js_code", in.Code)
	params.Set("grant_type", "authorization_code")

	Url.RawQuery = params.Encode() // 如果参数中有中文参数,这个方法会进行URLEncode
	urlPath := Url.String()

	resp, err := http.Get(urlPath)
	defer resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("url: %s, err: %s", urlPath, err)
	}
	body, _ := ioutil.ReadAll(resp.Body)

	var authResp user.WeChatLoginResponse
	err = json.Unmarshal(body, &authResp)
	if err != nil {
		return nil, fmt.Errorf("url: %s, err: %s", body, err)
	}

	return &authResp, nil
}
