package qidian_sdk_go

import (
	"context"

	httputil "github.com/tencentqidian/qidian-sdk-go/util/http"
)

// Client SDK client
type Client struct {
	client *httputil.Client

	accessToken string
}

// NewClient create new instance of Client with given ClientOption
func NewClient(opts ...ClientOption) *Client {
	c := Client{
		client: &httputil.Client{},
	}
	for _, o := range opts {
		o(&c)
	}
	return &c
}

// Do fire http request with access token
func (c *Client) Do(ctx context.Context, method, url string, params, v interface{}, opt ...httputil.CallOption) error {
	opt = append(opt, httputil.WithAccessToken(c.accessToken))
	return c.client.Do(ctx, method, url, params, v, opt...)
}
