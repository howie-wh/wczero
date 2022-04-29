
### 1. "MPBasicSetup"

1. 路由定义

- Url: /api/v1/mp
- Method: GET
- Request: `MPBasicSetupRequest`
- Response: `MPBasicSetupResponse`

2. 请求定义


```golang
type MPBasicSetupRequest struct {
	Signature string `form:"signature"`
	Timestamp string `form:"timestamp"`
	Nonce string `form:"nonce"`
	EchoStr string `form:"echostr"`
}
```


3. 返回定义


```golang
type MPBasicSetupResponse struct {
	Reply string `json:"reply"`
}
```
  


### 2. "MPTextMsg"

1. 路由定义

- Url: /api/v1/mp
- Method: POST
- Request: `MPTextMsgRequest`
- Response: `MPTextMsgResponse`

2. 请求定义


```golang
type MPTextMsgRequest struct {
	ToUserName string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime int64 `xml:"CreateTime"`
	MsgType string `xml:"MsgType"`
	Content string `xml:"Content"`
	MsgId int64 `xml:"MsgId"`
}
```


3. 返回定义


```golang
type MPTextMsgResponse struct {
	ToUserName string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime int64 `xml:"CreateTime"`
	MsgType string `xml:"MsgType"`
	Content string `xml:"Content"`
}
```
  

