package qidian_sdk_go

import (
	"context"
	"net/http"

	"github.com/tencentqidian/qidian-sdk-go/types"
	httputil "github.com/tencentqidian/qidian-sdk-go/util/http"
)

// Component .
type Component struct {
	client *httputil.Client

	appID                string
	appSecret            string
	token                string
	encodingAESKey       string
	componentAccessToken string
}

// NewComponent .
func NewComponent(opts ...ComponentOption) *Component {
	c := &Component{
		client: httputil.NewClient(),
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// Init component with ComponentOptions
func (c *Component) Init(opts ...ComponentOption) {
	for _, opt := range opts {
		opt(c)
	}
}

// GetOAuthAPPToken 使用应用授权code换取应用授权token
// Reference https://api.qidian.qq.com/wiki/doc/open/enudsepks7pq90r54frh
func (c *Component) GetOAuthAPPToken(ctx context.Context, req *types.GetOAuthAPPTokenReq) (*types.GetOAuthAPPTokenRsp, error) {
	req.ComponentAppID = c.appID

	var resp types.GetOAuthAPPTokenRsp
	err := c.client.Do(ctx, http.MethodPost, "/cgi-bin/component/oauth_app_token?component_access_token="+c.componentAccessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetComponentToken 获取应用开发商 token
func (c *Component) GetComponentToken(ctx context.Context, req *types.GetComponentTokenReq) (*types.GetComponentTokenRsp, error) {
	req.ComponentAppID = c.appID
	req.ComponentAppSecret = c.appSecret

	var resp types.GetComponentTokenRsp
	err := c.client.Do(ctx, http.MethodPost, "/cgi-bin/component/api_component_token", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RefreshAuthorizerToken 应用授权刷新token的使用方法
func (c *Component) RefreshAuthorizerToken(ctx context.Context, req *types.RefreshAuthorizerTokenReq) (*types.RefreshAuthorizerTokenRsp, error) {
	req.ComponentAppID = c.appID

	var resp types.RefreshAuthorizerTokenRsp
	err := c.client.Do(ctx, http.MethodPost, "/cgi-bin/component/api_authorizer_token?component_access_token="+c.componentAccessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
