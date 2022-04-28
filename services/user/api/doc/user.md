
### 1. "WeChatLogin"

1. 路由定义

- Url: /api/v1/user/wechat_login
- Method: POST
- Request: `WeChatLoginRequest`
- Response: `WeChatLoginResponse`

2. 请求定义


```golang
type WeChatLoginRequest struct {
	AppID string `json:"appid"`
	Code string `json:"code"`
}
```


3. 返回定义


```golang
type WeChatLoginResponse struct {
	AccessToken string `json:"access_token"`
	AccessExpire int64 `json:"access_expire"`
}
```
  


### 2. "AdminRegister"

1. 路由定义

- Url: /api/v1/user/admin_register
- Method: POST
- Request: `AdminRegisterRequest`
- Response: `AdminRegisterResponse`

2. 请求定义


```golang
type AdminRegisterRequest struct {
	Name string `json:"name"`
	Gender int64 `json:"gender"`
	Mobile string `json:"mobile"`
	Password string `json:"password"`
}
```


3. 返回定义


```golang
type AdminRegisterResponse struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Gender int64 `json:"gender"`
	Mobile string `json:"mobile"`
}
```
  


### 3. "AdminLogin"

1. 路由定义

- Url: /api/v1/user/admin_login
- Method: POST
- Request: `AdminLoginRequest`
- Response: `AdminLoginResponse`

2. 请求定义


```golang
type AdminLoginRequest struct {
	Mobile string `json:"mobile"`
	Password string `json:"password"`
}
```


3. 返回定义


```golang
type AdminLoginResponse struct {
	AccessToken string `json:"access_token"`
	AccessExpire int64 `json:"access_expire"`
}
```
  


### 4. "AdminUserInfo"

1. 路由定义

- Url: /api/v1/user/admin_userinfo
- Method: GET
- Request: `-`
- Response: `AdminUserInfoResponse`

2. 请求定义


3. 返回定义


```golang
type AdminUserInfoResponse struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Gender int64 `json:"gender"`
	Mobile string `json:"mobile"`
}
```
  


### 5. "QINIUUploadToken"

1. 路由定义

- Url: /api/v1/user/qiniu/upload_token
- Method: GET
- Request: `QINIUUploadTokenRequest`
- Response: `QINIUUploadTokenResponse`

2. 请求定义


```golang
type QINIUUploadTokenRequest struct {
}
```


3. 返回定义


```golang
type QINIUUploadTokenResponse struct {
	Domain string `json:"domain"`
	Zone string `json:"zone"`
	UploadToken string `json:"upload_token"`
	Desc string `json:"desc"`
}
```
  

