package Tikhub

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClientGeneratedEndpoint(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/tiktok/web/fetch_post_detail" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("Authorization"); got != "Bearer test-key" {
			t.Fatalf("unexpected authorization header: %s", got)
		}
		if got := r.URL.Query().Get("itemId"); got != "123" {
			t.Fatalf("unexpected itemId: %s", got)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code":200,"router":"/api/v1/tiktok/web/fetch_post_detail","data":{"ok":true}}`))
	}))
	defer server.Close()

	client := NewClient("test-key", WithBaseURL(server.URL))
	resp, err := client.TikTokWebGetSingleVideoData(context.Background(), TikTokWebGetSingleVideoDataRequest{
		ItemID: "123",
	})
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	if resp.Code != 200 {
		t.Fatalf("unexpected code: %d", resp.Code)
	}

	var data struct {
		OK bool `json:"ok"`
	}
	if err := resp.DecodeData(&data); err != nil {
		t.Fatalf("decode data: %v", err)
	}
	if !data.OK {
		t.Fatal("expected decoded data ok=true")
	}

	decoded, err := DecodeData[struct {
		OK bool `json:"ok"`
	}](client.TikTokWebGetSingleVideoData(context.Background(), TikTokWebGetSingleVideoDataRequest{ItemID: "123"}))
	if err != nil {
		t.Fatalf("decode generated response: %v", err)
	}
	if !decoded.OK {
		t.Fatal("expected decoded generated response ok=true")
	}
}

func TestClientPostBody(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		var body map[string]string
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode request body: %v", err)
		}
		if body["url"] != "https://example.com" || body["user_agent"] != "test-agent" {
			t.Fatalf("unexpected body: %#v", body)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code":200,"data":{"received":true}}`))
	}))
	defer server.Close()

	client := NewClient("test-key", WithBaseURL(server.URL))
	resp, err := client.TikTokWebGenerateXBogus(context.Background(), TikTokWebGenerateXBogusRequest{
		URL:       "https://example.com",
		UserAgent: "test-agent",
	})
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code: %d", resp.StatusCode)
	}
}

func TestClientSendWithReqOptions(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.URL.Query()["ids"]; len(got) != 2 || got[0] != "1" || got[1] != "2" {
			t.Fatalf("unexpected ids query: %#v", got)
		}
		if got := r.Header.Get("X-Test"); got != "ok" {
			t.Fatalf("unexpected X-Test header: %s", got)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code":200,"data":{"name":"tikhub"}}`))
	}))
	defer server.Close()

	client := NewClient("test-key", WithBaseURL(server.URL))
	resp, err := client.Send(context.Background(), http.MethodGet, "/anything",
		WithQueryMap(map[string]any{"ids": []string{"1", "2"}}),
		WithHeader("X-Test", "ok"),
	)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	data, err := Data[struct {
		Name string `json:"name"`
	}](resp)
	if err != nil {
		t.Fatalf("decode data: %v", err)
	}
	if data.Name != "tikhub" {
		t.Fatalf("unexpected data: %#v", data)
	}
}

func TestRawReqRequestCanBeParsed(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("X-Raw"); got != "yes" {
			t.Fatalf("unexpected X-Raw header: %s", got)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code":200,"data":{"raw":true}}`))
	}))
	defer server.Close()

	client := NewClient("test-key", WithBaseURL(server.URL))
	rawResp, err := client.R().
		SetHeader("X-Raw", "yes").
		Get("/raw")
	if err != nil {
		t.Fatalf("raw request failed: %v", err)
	}

	resp, err := ParseResponse(rawResp)
	if err != nil {
		t.Fatalf("parse response: %v", err)
	}
	if resp.Code != 200 {
		t.Fatalf("unexpected code: %d", resp.Code)
	}
}
