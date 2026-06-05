# Tikhub

TikHub API 的 Go SDK，基于 `github.com/jwwsjlm/req/v3`。

## 安装

```bash
go get github.com/jwwsjlm/Tikhub
```

## 基本用法

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

	data, err := tikhub.DecodeData[map[string]any](client.TikTokWebGetSingleVideoData(context.Background(), tikhub.TikTokWebGetSingleVideoDataRequest{
		ItemID: "7218694761253735723",
	}))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v\n", data)
}
```

可选参数使用指针，不传则不会出现在 query 里：

```go
resp, err := client.TikTokWebGetExploreVideoData(context.Background(), tikhub.TikTokWebGetExploreVideoDataRequest{
	Count: tikhub.Ptr(20),
})
```

POST 接口直接传 JSON body：

```go
resp, err := client.TikTokWebGenerateXBogus(context.Background(), tikhub.TikTokWebGenerateXBogusRequest{
	URL:       "https://www.tiktok.com/",
	UserAgent: "Mozilla/5.0 ...",
})
```

## 三种调用方式

推荐优先使用生成方法，参数类型清楚，路径不容易写错：

```go
resp, err := client.TikTokWebGetSingleVideoData(context.Background(), tikhub.TikTokWebGetSingleVideoDataRequest{
	ItemID: "7218694761253735723",
})
```

只关心 `data` 时，用泛型直接解包：

```go
type Result map[string]any

data, err := tikhub.DecodeData[Result](client.TikTokWebGetSingleVideoData(context.Background(), tikhub.TikTokWebGetSingleVideoDataRequest{
	ItemID: "7218694761253735723",
}))
```

文档更新但 SDK 还没重新生成，或者要加临时 header/query/body，用 `Send`：

```go
resp, err := client.Send(context.Background(), "GET", "/api/v1/health/check",
	tikhub.WithQueryMap(map[string]any{
		"ids": []string{"1", "2"},
	}),
	tikhub.WithHeader("X-Trace-ID", "trace-id"),
)
```

需要完整使用 `req` 的链式能力时，直接拿原生 request：

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

## 直播弹幕 WS 链接

旧版 helper 仍然保留：

```go
ua := tikhub.RandUserAgent()
ws, err := tikhub.GenerateWsLink("你的 TikHub API Key", ua, "直播间号")
if err != nil {
	log.Fatal(err)
}
log.Println(ws.Url, ws.Ttwid)
```

## 集成测试

默认测试只走本地 mock server：

```bash
go test ./...
```

真实 TikHub 调用需要设置环境变量：

```bash
set TIKHUB_API_KEY=你的 TikHub API Key
set TIKHUB_WEBCAST_ID=直播间号
go test ./...
```
