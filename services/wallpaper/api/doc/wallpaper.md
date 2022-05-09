
### 1. "Import"

1. 路由定义

- Url: /api/v1/wallpaper/import
- Method: POST
- Request: `ImportRequest`
- Response: `ImportResponse`

2. 请求定义


```golang
type ImportRequest struct {
	List []WallPaperInfo `json:"list"`
}
```


3. 返回定义


```golang
type ImportResponse struct {
}
```
  


### 2. "Remove"

1. 路由定义

- Url: /api/v1/wallpaper/remove
- Method: POST
- Request: `RemoveRequest`
- Response: `RemoveResponse`

2. 请求定义


```golang
type RemoveRequest struct {
	List []string `json:"list"`
}
```


3. 返回定义


```golang
type RemoveResponse struct {
}
```
  


### 3. "Detail"

1. 路由定义

- Url: /api/v1/wallpaper/detail
- Method: GET
- Request: `DetailRequest`
- Response: `DetailResponse`

2. 请求定义


```golang
type DetailRequest struct {
	Wid string `form:"wid"`
}
```


3. 返回定义


```golang
type DetailResponse struct {
	Wid string `json:"wid"`
	Name string `json:"name"`
	Tp string `json:"type"`
	Category string `json:"category"`
	ImageURL string `json:"image_url"`
	Author string `json:"author"`
	Desc string `json:"desc"`
}
```
  


### 4. "List"

1. 路由定义

- Url: /api/v1/wallpaper/list
- Method: GET
- Request: `ListRequest`
- Response: `ListResponse`

2. 请求定义


```golang
type ListRequest struct {
	Start int64 `form:"start"`
	Limit int64 `form:"limit"`
}
```


3. 返回定义


```golang
type ListResponse struct {
	List []WallPaperInfo `json:"list,omitempty"`
	Total int64 `json:"total"`
}
```
  


### 5. "Category"

1. 路由定义

- Url: /api/v1/wallpaper/category
- Method: GET
- Request: `CategoryRequest`
- Response: `CategoryResponse`

2. 请求定义


```golang
type CategoryRequest struct {
	Start int64 `form:"start,optional"`
	Limit int64 `form:"limit,optional"`
}
```


3. 返回定义


```golang
type CategoryResponse struct {
	Tp []string `json:"type"`
	TpTotal int64 `json:"type_total"`
	Category []string `json:"category"`
	CategoryTotal int64 `json:"category_total"`
}
```
  

