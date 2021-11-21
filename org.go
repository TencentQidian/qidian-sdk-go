package qidian_sdk_go

import (
	"context"
	"net/http"

	"github.com/tencentqidian/qidian-sdk-go/types"
)

// Departments 读取部门列表
// Reference https://api.qidian.qq.com/wiki/doc/open/eldtzp0xgllhpmn33wxi
func (c *Client) Departments(ctx context.Context) (*types.GetDepartmentsRsp, error) {
	var resp struct {
		ErrorV1
		types.GetDepartmentsRsp `json:"data"`
	}

	err := c.Do(ctx, http.MethodPost, "/cgi-bin/v1/org/basic/GetDepList", nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.GetDepartmentsRsp, nil
}
