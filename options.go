package qidian_sdk_go

import "net/http"

type ComponentOption func(d *Component)

type ClientOption func(c *Client)

type DoOption func(c *Client, r *http.Request)

// WithAccessToken 设置 access_token
func WithAccessToken(s string) ClientOption {
	return func(c *Client) {
		c.accessToken = s
	}
}

// WithComponentAppID set app id
func WithComponentAppID(s string) ComponentOption {
	return func(c *Component) {
		c.appID = s
	}
}

// WithComponentAppSecret set secret
func WithComponentAppSecret(s string) ComponentOption {
	return func(c *Component) {
		c.appSecret = s
	}
}

// WithComponentAccessToken set component access token
func WithComponentAccessToken(s string) ComponentOption {
	return func(c *Component) {
		c.componentAccessToken = s
	}
}
