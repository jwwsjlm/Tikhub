# Tikhub

TikHub API 的 Go SDK，基于 `github.com/jwwsjlm/req/v3`。

## 安装

```bash
go get github.com/jwwsjlm/Tikhub
```

## 基本用法

推荐使用和官方 Python SDK 接近的资源分组写法：

```go
client.资源.接口(ctx, tikhub.资源接口Request{})
```

规则和官方 Python SDK 基本一致：

- 资源名来自 OpenAPI tag：`TikTok-Web-API` -> `client.TikTokWeb`
- 方法名来自接口 path 最后一段：`fetch_post_detail` -> `FetchPostDetail`
- 参数使用 Go struct：`TikTokWebFetchPostDetailRequest`

比如官方 Python SDK 是：

```python
client.tiktok_web.fetch_post_detail(itemId="7218694761253735723")
```

Go 里对应：

```go
resp, err := client.TikTokWeb.FetchPostDetail(ctx, tikhub.TikTokWebFetchPostDetailRequest{
	ItemID: "7218694761253735723",
})
```

旧的文档标题式方法也保留：

```go
client.TikTokWebGetSingleVideoData(ctx, tikhub.TikTokWebGetSingleVideoDataRequest{})
```

返回值统一是 `*tikhub.APIResponse`，其中：

```go
resp.Code   // TikHub 返回 code
resp.Router // TikHub 返回 router
resp.Data   // 原始 data，类型是 json.RawMessage
```

## 创建客户端

```go
package main

import (
	"context"
	"log"
	"time"

	"github.com/jwwsjlm/req/v3"
	tikhub "github.com/jwwsjlm/Tikhub"
)

func main() {
	client := tikhub.NewClient("你的 TikHub API Key",
		tikhub.WithTimeout(20*time.Second),
		tikhub.WithReq(func(r *req.Client) {
			r.SetCommonRetryCount(2)
		}),
	)

	resp, err := client.TikTokWeb.FetchPostDetail(context.Background(), tikhub.TikTokWebFetchPostDetailRequest{
		ItemID: "7218694761253735723",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.Code, string(resp.Data))
}
```

## 常见例子

### 1. TikTok 获取单个作品

```go
resp, err := client.TikTokWeb.FetchPostDetail(context.Background(), tikhub.TikTokWebFetchPostDetailRequest{
	ItemID: "7218694761253735723",
})
if err != nil {
	log.Fatal(err)
}

log.Println(resp.Code)
log.Println(string(resp.Data))
```

只想拿 `data`，可以直接解包成 `map`：

```go
data, err := tikhub.DecodeData[map[string]any](client.TikTokWeb.FetchPostDetail(context.Background(), tikhub.TikTokWebFetchPostDetailRequest{
	ItemID: "7218694761253735723",
}))
if err != nil {
	log.Fatal(err)
}

log.Printf("%#v\n", data)
```

### 2. Douyin 获取单个作品

```go
resp, err := client.DouyinWeb.FetchOneVideo(context.Background(), tikhub.DouyinWebFetchOneVideoRequest{
	AwemeID:        "7369956465575087400",
	NeedAnchorInfo: true,
})
if err != nil {
	log.Fatal(err)
}

log.Println(string(resp.Data))
```

### 3. 可选参数怎么传

生成的 `Request` 里，文档标记为可选的 `string`、`int`、`bool` 字段可以直接传普通值。零值会被当作不传，比如 `0`、`false`、空字符串不会放到 query/body 里。

```go
resp, err := client.TikTokWeb.FetchExplorePost(context.Background(), tikhub.TikTokWebFetchExplorePostRequest{
	Count:        20,
	CategoryType: "120",
})
if err != nil {
	log.Fatal(err)
}

log.Println(string(resp.Data))
```

### 4. 没有参数的接口

没有参数的接口只传 `ctx`，不需要传空 `Request`。

```go
resp, err := client.TikTokWeb.FetchTrendingPost(context.Background())
if err != nil {
	log.Fatal(err)
}

log.Println(string(resp.Data))
```

### 5. POST 接口

POST 接口也是传对应的 `Request`，SDK 会自动作为 JSON body 发送。

```go
resp, err := client.TikTokWeb.GenerateXBogus(context.Background(), tikhub.TikTokWebGenerateXBogusRequest{
	URL:       "https://www.tiktok.com/api/item/detail/?itemId=7218694761253735723",
	UserAgent: "Mozilla/5.0 ...",
})
if err != nil {
	log.Fatal(err)
}

log.Println(string(resp.Data))
```

### 6. 解包成自己的结构体

TikHub 每个接口的 `data` 结构可能不同，所以 SDK 不强行写死深层 response。你可以按自己需要定义结构体解包。

```go
type VideoData struct {
	ItemInfo map[string]any `json:"itemInfo"`
}

data, err := tikhub.DecodeData[VideoData](client.TikTokWeb.FetchPostDetail(context.Background(), tikhub.TikTokWebFetchPostDetailRequest{
	ItemID: "7218694761253735723",
}))
if err != nil {
	log.Fatal(err)
}

log.Printf("%#v\n", data.ItemInfo)
```

### 7. 临时调用一个路径

文档更新但 SDK 还没重新生成，或者你想临时加 header/query/body，可以用 `Send`。

```go
resp, err := client.Send(context.Background(), "GET", "/api/v1/health/check",
	tikhub.WithQueryMap(map[string]any{
		"ids": []string{"1", "2"},
	}),
	tikhub.WithHeader("X-Trace-ID", "trace-id"),
)
if err != nil {
	log.Fatal(err)
}

log.Println(resp.Code)
```

### 8. 直接使用 req

需要完整使用 `req` 的链式能力时，可以从 client 里拿原生 request。

```go
rawResp, err := client.R().
	SetHeader("X-Trace-ID", "trace-id").
	EnableDump().
	Get("/api/v1/health/check")
if err != nil {
	log.Fatal(err)
}

resp, err := tikhub.ParseResponse(rawResp)
if err != nil {
	log.Fatal(err)
}

log.Println(resp.Code)
```

## 集成测试

默认测试只走本地 mock server：

```bash
go test ./...
```

真实 TikHub 调用需要设置环境变量：

```bash
set TIKHUB_API_KEY=你的 TikHub API Key
go test ./...
```
