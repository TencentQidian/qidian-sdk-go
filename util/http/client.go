package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/tencentqidian/qidian-sdk-go/errors"
)

const (
	Version  = "0.1.0"
	Endpoint = "https://api.qidian.qq.com"
)

var (
	Debug = false
)

func init() {
	if os.Getenv("QD_SDK_DEBUG") == "true" {
		Debug = true
	}
}

type CallOption func(c *Client, r *http.Request)
type Option func(c *Client)

// Client is the http client wrapper
type Client struct {
	c *http.Client
}

// NewClient returns a new Client.
func NewClient(opt ...Option) *Client {
	c := &Client{
		c: http.DefaultClient,
	}

	for _, o := range opt {
		o(c)
	}
	return c
}

// Do fire http request and decode body into v
func (c *Client) Do(ctx context.Context, method, url string, params, v interface{}, opt ...CallOption) error {
	var body bytes.Buffer

	if params != nil && (method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch) {
		if err := json.NewEncoder(&body).Encode(params); err != nil {
			return err
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, Endpoint+url, &body)
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "qidian-sdk-go v"+Version)
	if body.Len() > 0 {
		req.Header.Add("Content-Type", "application/json")
	}

	for _, o := range opt {
		o(c, req)
	}

	if Debug {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Print("Request: ", string(dump))
	}

	var resp *http.Response
	if resp, err = http.DefaultClient.Do(req); err != nil {
		return err
	}
	defer resp.Body.Close()

	if Debug {
		dump, _ := httputil.DumpResponse(resp, true)
		log.Print("Response: ", string(dump))
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid response status=%d", resp.StatusCode)
	}

	bodyBuf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// assert SDK Error
	var sdkErr errors.Err
	if err = json.Unmarshal(bodyBuf, &sdkErr); err == nil && sdkErr.Code != 0 {
		return &sdkErr
	}

	if v != nil {
		if err = json.Unmarshal(bodyBuf, v); err == nil {
			return err
		}
	}

	return nil
}

// WithAccessToken set access_token in request query string
func WithAccessToken(s string) CallOption {
	return func(c *Client, r *http.Request) {
		u := r.URL.Query()
		u.Add("access_token", s)
		r.URL.RawQuery = u.Encode()
	}
}

// WithHTTPClient set HTTP client for send request
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) {
		c.c = hc
	}
}
