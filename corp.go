package qidian_sdk_go

import (
	"context"
	"net/http"

	"github.com/tencentqidian/qidian-sdk-go/types"
)

// GetCorpInfo 获取企业信息
func (c *Client) GetCorpInfo(ctx context.Context) (*types.GetCorpInfoRsp, error) {
	var resp struct {
		ErrorV1
		types.GetCorpInfoRsp `json:"data"`
	}

	err := c.Do(ctx, http.MethodGet, "/cgi-bin/v1/account/profile/corp/info", nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.GetCorpInfoRsp, nil
}
