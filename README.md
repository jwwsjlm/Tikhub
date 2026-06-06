# Tikhub

TikHub API 的 Go SDK，基于 [`github.com/jwwsjlm/req/v3`](https://github.com/jwwsjlm/req)。

这个库按 TikHub OpenAPI 生成，主推和官方 Python SDK 接近的资源分组写法：

```go
client.TikTokWeb.FetchPostDetail(ctx, tikhub.TikTokWebFetchPostDetailRequest{
	ItemID: "7218694761253735723",
})
```

## 安装

```bash
go get github.com/jwwsjlm/Tikhub
```

指定版本：

```bash
go get github.com/jwwsjlm/Tikhub@v0.2.1
```

## 快速开始

```go
package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jwwsjlm/req/v3"
	tikhub "github.com/jwwsjlm/Tikhub"
)

func main() {
	apiKey := os.Getenv("TIKHUB_API_KEY")
	if apiKey == "" {
		log.Fatal("missing TIKHUB_API_KEY")
	}

	client := tikhub.NewClient(apiKey,
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

	log.Println(resp.Code)
	log.Println(string(resp.Data))
}
```

## 命名规则

官方 Python SDK 是这种形式：

```python
client.tiktok_web.fetch_post_detail(itemId="7218694761253735723")
```

Go 版对应：

```go
client.TikTokWeb.FetchPostDetail(ctx, tikhub.TikTokWebFetchPostDetailRequest{
	ItemID: "7218694761253735723",
})
```

规则：

- OpenAPI tag 变成资源名：`TikTok-Web-API` -> `client.TikTokWeb`
- API path 最后一段变成方法名：`fetch_post_detail` -> `FetchPostDetail`
- 参数放进对应的 request struct：`TikTokWebFetchPostDetailRequest`

常见对照：

| API | Go 调用 |
| --- | --- |
| `/api/v1/tiktok/web/fetch_post_detail` | `client.TikTokWeb.FetchPostDetail(...)` |
| `/api/v1/tiktok/web/fetch_explore_post` | `client.TikTokWeb.FetchExplorePost(...)` |
| `/api/v1/douyin/web/fetch_one_video` | `client.DouyinWeb.FetchOneVideo(...)` |
| `/api/v1/health/check` | `client.HealthCheck.Check(...)` |

旧的文档标题式方法也保留，可以继续用：

```go
client.TikTokWebGetSingleVideoData(ctx, tikhub.TikTokWebGetSingleVideoDataRequest{
	ItemID: "7218694761253735723",
})
```

## 参数规则

每个接口都有自己的 request struct，字段注释来自 TikHub 文档。

必填参数直接传：

```go
resp, err := client.TikTokWeb.FetchPostDetail(ctx, tikhub.TikTokWebFetchPostDetailRequest{
	ItemID: "7218694761253735723",
})
```

可选的 `string`、`int`、`bool` 字段也直接传普通值：

```go
resp, err := client.TikTokWeb.FetchExplorePost(ctx, tikhub.TikTokWebFetchExplorePostRequest{
	CategoryType: "120",
	Count:        20,
})
```

零值会被当作“不传”，例如 `""`、`0`、`false` 不会进入 query/body。如果某个接口确实需要显式传 `0` 或 `false`，可以用 `Send`：

```go
resp, err := client.Send(ctx, "GET", "/api/v1/example/path",
	tikhub.WithQueryMap(map[string]any{
		"count": 0,
		"flag":  false,
	}),
)
```

## 返回值

所有生成方法都返回 `*tikhub.APIResponse`：

```go
resp.Code       // TikHub code
resp.Message    // 英文 message
resp.MessageZH  // 中文 message
resp.Router     // router
resp.Params     // 原始 params，json.RawMessage
resp.Data       // 原始 data，json.RawMessage
resp.Raw        // 完整响应 body
resp.StatusCode // HTTP status code
resp.Header     // HTTP response header
```

只想拿 `data`，可以用泛型解包：

```go
data, err := tikhub.DecodeData[map[string]any](
	client.TikTokWeb.FetchPostDetail(ctx, tikhub.TikTokWebFetchPostDetailRequest{
		ItemID: "7218694761253735723",
	}),
)
if err != nil {
	log.Fatal(err)
}

log.Printf("%#v\n", data)
```

也可以解包到自己的结构体：

```go
type VideoData struct {
	ItemInfo map[string]any `json:"itemInfo"`
}

data, err := tikhub.DecodeData[VideoData](
	client.TikTokWeb.FetchPostDetail(ctx, tikhub.TikTokWebFetchPostDetailRequest{
		ItemID: "7218694761253735723",
	}),
)
```

## 常见调用

### TikTok 获取单个作品

```go
resp, err := client.TikTokWeb.FetchPostDetail(ctx, tikhub.TikTokWebFetchPostDetailRequest{
	ItemID: "7218694761253735723",
})
```

### Douyin 获取单个作品

```go
resp, err := client.DouyinWeb.FetchOneVideo(ctx, tikhub.DouyinWebFetchOneVideoRequest{
	AwemeID:        "7369956465575087400",
	NeedAnchorInfo: true,
})
```

### TikTok 探索作品

```go
resp, err := client.TikTokWeb.FetchExplorePost(ctx, tikhub.TikTokWebFetchExplorePostRequest{
	CategoryType: "120",
	Count:        20,
})
```

### 没有参数的接口

```go
resp, err := client.TikTokWeb.FetchTrendingPost(ctx)
```

### POST 接口

```go
resp, err := client.TikTokWeb.GenerateXBogus(ctx, tikhub.TikTokWebGenerateXBogusRequest{
	URL:       "https://www.tiktok.com/api/item/detail/?itemId=7218694761253735723",
	UserAgent: "Mozilla/5.0 ...",
})
```

## 高级用法

文档更新但 SDK 还没重新生成，或者需要临时加 query/header/body，可以用 `Send`：

```go
resp, err := client.Send(ctx, "GET", "/api/v1/health/check",
	tikhub.WithHeader("X-Trace-ID", "trace-id"),
	tikhub.WithQueryMap(map[string]any{
		"ids": []string{"1", "2"},
	}),
)
```

如果需要直接使用底层 `req` 链式能力：

```go
rawResp, err := client.R().
	SetHeader("X-Trace-ID", "trace-id").
	EnableDump().
	Get("/api/v1/health/check")
if err != nil {
	log.Fatal(err)
}

resp, err := tikhub.ParseResponse(rawResp)
```

## 覆盖范围

SDK 覆盖 TikHub OpenAPI 中的资源分组和 GET/POST 接口，并额外保留旧版本中仍可用的少量兼容接口。当前生成了两套入口：

- 官方风格资源入口：`client.TikTokWeb.FetchPostDetail(...)`
- 兼容的标题式入口：`client.TikTokWebGetSingleVideoData(...)`

## 测试

默认测试使用本地 mock server：

```bash
go test ./...
```

真实调用示例可以自己设置 API key：

```bash
set TIKHUB_API_KEY=你的 TikHub API Key
go test ./...
```
