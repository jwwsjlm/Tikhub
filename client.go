package Tikhub

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/jwwsjlm/req/v3"
)

const DefaultBaseURL = "https://api.tikhub.io"

// Client is a TikHub API client backed by github.com/jwwsjlm/req/v3.
type Client struct {
	Resources

	r       *req.Client
	apiKey  string
	baseURL string
}

// Option customizes a Client created by NewClient.
type Option func(*Client)

// WithReq lets callers customize the underlying req client.
func WithReq(configure func(*req.Client)) Option {
	return func(c *Client) {
		if configure != nil {
			configure(c.r)
		}
	}
}

// WithBaseURL overrides the TikHub API base URL.
func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.baseURL = strings.TrimRight(baseURL, "/")
		c.r.SetBaseURL(c.baseURL)
	}
}

// WithTimeout configures the request timeout.
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.r.SetTimeout(timeout)
	}
}

// WithUserAgent configures the default User-Agent header.
func WithUserAgent(userAgent string) Option {
	return func(c *Client) {
		if userAgent != "" {
			c.r.SetUserAgent(userAgent)
		}
	}
}

// WithReqClient uses an existing req client. The TikHub base URL and bearer
// token are still applied to it.
func WithReqClient(reqClient *req.Client) Option {
	return func(c *Client) {
		if reqClient != nil {
			c.r = reqClient
			c.r.SetBaseURL(c.baseURL)
			if c.apiKey != "" {
				c.r.SetCommonBearerAuthToken(c.apiKey)
			}
		}
	}
}

// NewClient creates a TikHub API client.
func NewClient(apiKey string, opts ...Option) *Client {
	c := &Client{
		r:       req.C(),
		apiKey:  apiKey,
		baseURL: DefaultBaseURL,
	}
	c.r.SetBaseURL(c.baseURL)
	if apiKey != "" {
		c.r.SetCommonBearerAuthToken(apiKey)
	}
	for _, opt := range opts {
		opt(c)
	}
	c.initResources()
	return c
}

// Ptr returns a pointer to v. It is handy for optional generated parameters.
func Ptr[T any](v T) *T {
	return &v
}

// ReqClient returns the underlying req client for advanced customization.
func (c *Client) ReqClient() *req.Client {
	if c == nil {
		return nil
	}
	return c.r
}

// R creates a raw req request. Use ParseResponse to decode the response back
// into APIResponse when you need req's chainable per-request features.
func (c *Client) R() *req.Request {
	if c == nil || c.r == nil {
		return nil
	}
	return c.r.R()
}

// APIResponse is the common TikHub response envelope.
type APIResponse struct {
	Code           int             `json:"code,omitempty"`
	RequestID      *string         `json:"request_id,omitempty"`
	Message        string          `json:"message,omitempty"`
	MessageZH      string          `json:"message_zh,omitempty"`
	Support        string          `json:"support,omitempty"`
	Time           string          `json:"time,omitempty"`
	TimeStamp      int64           `json:"time_stamp,omitempty"`
	TimeZone       string          `json:"time_zone,omitempty"`
	Docs           *string         `json:"docs,omitempty"`
	CacheMessage   *string         `json:"cache_message,omitempty"`
	CacheMessageZH *string         `json:"cache_message_zh,omitempty"`
	CacheURL       *string         `json:"cache_url,omitempty"`
	Router         string          `json:"router,omitempty"`
	Params         json.RawMessage `json:"params,omitempty"`
	Data           json.RawMessage `json:"data,omitempty"`
	Raw            json.RawMessage `json:"-"`
	StatusCode     int             `json:"-"`
	Header         http.Header     `json:"-"`
}

// DecodeData decodes the response data field into v.
func (r *APIResponse) DecodeData(v any) error {
	if r == nil {
		return fmt.Errorf("tikhub: nil APIResponse")
	}
	if len(r.Data) == 0 || string(r.Data) == "null" {
		return nil
	}
	return json.Unmarshal(r.Data, v)
}

// DecodeParams decodes the response params field into v.
func (r *APIResponse) DecodeParams(v any) error {
	if r == nil {
		return fmt.Errorf("tikhub: nil APIResponse")
	}
	if len(r.Params) == 0 || string(r.Params) == "null" {
		return nil
	}
	return json.Unmarshal(r.Params, v)
}

// Decode decodes the full response body into v.
func (r *APIResponse) Decode(v any) error {
	if r == nil {
		return fmt.Errorf("tikhub: nil APIResponse")
	}
	if len(r.Raw) == 0 {
		return nil
	}
	return json.Unmarshal(r.Raw, v)
}

// Data decodes resp.Data into a typed value.
func Data[T any](resp *APIResponse) (T, error) {
	var out T
	if resp == nil {
		return out, fmt.Errorf("tikhub: nil APIResponse")
	}
	if err := resp.DecodeData(&out); err != nil {
		return out, err
	}
	return out, nil
}

// DecodeData unwraps a client call and decodes resp.Data into a typed value.
func DecodeData[T any](resp *APIResponse, err error) (T, error) {
	var out T
	if err != nil {
		return out, err
	}
	return Data[T](resp)
}

// Query converts simple values into url.Values for Client.Get/Post/Do.
func Query(values map[string]any) url.Values {
	query := url.Values{}
	for key, value := range values {
		addQueryValue(query, key, value)
	}
	return query
}

// Do calls an arbitrary TikHub endpoint.
func (c *Client) Do(ctx context.Context, method, path string, query url.Values, body any) (*APIResponse, error) {
	return c.do(ctx, method, path, query, body)
}

// Get calls an arbitrary TikHub GET endpoint.
func (c *Client) Get(ctx context.Context, path string, query url.Values) (*APIResponse, error) {
	return c.do(ctx, http.MethodGet, path, query, nil)
}

// Post calls an arbitrary TikHub POST endpoint.
func (c *Client) Post(ctx context.Context, path string, query url.Values, body any) (*APIResponse, error) {
	return c.do(ctx, http.MethodPost, path, query, body)
}

// RequestOption customizes one request in Send.
type RequestOption func(*req.Request)

// WithQuery sets URL query parameters for Send.
func WithQuery(query url.Values) RequestOption {
	return func(r *req.Request) {
		if len(query) > 0 {
			r.SetQueryParamsFromValues(query)
		}
	}
}

// WithQueryMap sets URL query parameters for Send from simple Go values.
func WithQueryMap(query map[string]any) RequestOption {
	return WithQuery(Query(query))
}

// WithBody sets a JSON body for Send.
func WithBody(body any) RequestOption {
	return func(r *req.Request) {
		if body != nil {
			r.SetBody(body)
		}
	}
}

// WithHeader sets one request header for Send.
func WithHeader(key, value string) RequestOption {
	return func(r *req.Request) {
		r.SetHeader(key, value)
	}
}

// WithRequest applies arbitrary req request customization for Send.
func WithRequest(configure func(*req.Request)) RequestOption {
	return func(r *req.Request) {
		if configure != nil {
			configure(r)
		}
	}
}

// Send calls an arbitrary TikHub endpoint while keeping req's per-request
// chainable customization available through RequestOption.
func (c *Client) Send(ctx context.Context, method, path string, opts ...RequestOption) (*APIResponse, error) {
	if c == nil || c.r == nil {
		return nil, fmt.Errorf("tikhub: nil client")
	}
	request := c.r.R()
	if ctx != nil {
		request.SetContext(ctx)
	}
	for _, opt := range opts {
		if opt != nil {
			opt(request)
		}
	}
	return doRequest(request, method, path)
}

func (c *Client) do(ctx context.Context, method, path string, query url.Values, body any) (*APIResponse, error) {
	if c == nil || c.r == nil {
		return nil, fmt.Errorf("tikhub: nil client")
	}
	request := c.r.R()
	if ctx != nil {
		request.SetContext(ctx)
	}
	if len(query) > 0 {
		request.SetQueryParamsFromValues(query)
	}
	if body != nil {
		request.SetBody(body)
	}
	return doRequest(request, method, path)
}

func doRequest(request *req.Request, method, path string) (*APIResponse, error) {
	if request == nil {
		return nil, fmt.Errorf("tikhub: nil request")
	}
	var (
		resp *req.Response
		err  error
	)
	switch strings.ToUpper(method) {
	case http.MethodGet:
		resp, err = request.Get(path)
	case http.MethodPost:
		resp, err = request.Post(path)
	default:
		return nil, fmt.Errorf("tikhub: unsupported method %s", method)
	}
	if err != nil {
		return nil, err
	}
	return ParseResponse(resp)
}

// ParseResponse decodes a raw req response into APIResponse.
func ParseResponse(resp *req.Response) (*APIResponse, error) {
	if resp == nil {
		return nil, fmt.Errorf("tikhub: nil response")
	}
	raw, err := resp.ToBytes()
	if err != nil {
		return nil, err
	}
	out := &APIResponse{
		Raw:        append(json.RawMessage(nil), raw...),
		StatusCode: resp.GetStatusCode(),
	}
	if resp.Response != nil {
		out.Header = resp.Header.Clone()
	}
	if shouldDecodeJSON(resp.GetContentType(), raw) {
		if err := json.Unmarshal(raw, out); err != nil {
			return nil, fmt.Errorf("tikhub: decode response: %w", err)
		}
	}
	if !resp.IsSuccessState() {
		if out.Message != "" {
			return out, fmt.Errorf("tikhub: http %d: %s", out.StatusCode, out.Message)
		}
		return out, fmt.Errorf("tikhub: http %d: %s", out.StatusCode, string(raw))
	}
	return out, nil
}

func shouldDecodeJSON(contentType string, body []byte) bool {
	if len(body) == 0 {
		return false
	}
	if strings.Contains(strings.ToLower(contentType), "json") {
		return true
	}
	for _, b := range body {
		switch b {
		case ' ', '\n', '\r', '\t':
			continue
		case '{', '[':
			return true
		default:
			return false
		}
	}
	return false
}

func addQueryValue(values url.Values, key string, value any) {
	if key == "" || value == nil {
		return
	}

	rv := reflect.ValueOf(value)
	for rv.IsValid() && (rv.Kind() == reflect.Pointer || rv.Kind() == reflect.Interface) {
		if rv.IsNil() {
			return
		}
		rv = rv.Elem()
	}
	if !rv.IsValid() {
		return
	}

	switch rv.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < rv.Len(); i++ {
			addQueryValue(values, key, rv.Index(i).Interface())
		}
	default:
		values.Add(key, fmt.Sprint(rv.Interface()))
	}
}

func addOptionalQueryValue(values url.Values, key string, value any) {
	if isZeroValue(value) {
		return
	}
	addQueryValue(values, key, value)
}

func addBodyValue(body map[string]any, key string, value any) {
	if key == "" || value == nil {
		return
	}

	rv := reflect.ValueOf(value)
	for rv.IsValid() && (rv.Kind() == reflect.Pointer || rv.Kind() == reflect.Interface) {
		if rv.IsNil() {
			return
		}
		rv = rv.Elem()
	}
	if !rv.IsValid() {
		return
	}
	if (rv.Kind() == reflect.Slice || rv.Kind() == reflect.Map) && rv.IsNil() {
		return
	}

	body[key] = rv.Interface()
}

func addOptionalBodyValue(body map[string]any, key string, value any) {
	if isZeroValue(value) {
		return
	}
	addBodyValue(body, key, value)
}

func isZeroValue(value any) bool {
	if value == nil {
		return true
	}

	rv := reflect.ValueOf(value)
	for rv.IsValid() && (rv.Kind() == reflect.Pointer || rv.Kind() == reflect.Interface) {
		if rv.IsNil() {
			return true
		}
		rv = rv.Elem()
	}
	if !rv.IsValid() {
		return true
	}
	if (rv.Kind() == reflect.Slice || rv.Kind() == reflect.Map) && rv.IsNil() {
		return true
	}
	return rv.IsZero()
}
