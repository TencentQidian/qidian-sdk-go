package qidian_sdk_go

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tencentqidian/qidian-sdk-go/types"
)

func TestUnmarshal(t *testing.T) {
	var resp struct {
		ErrorV1
		types.GetDepartmentsRsp `json:"data"`
	}

	data := `{
	"errcode":0,
	"errmsg":"success",
	"data":{
		"departments_total":"17",
		"departments_list":[
			{"dep_id":"1","dep_name":"企点服务_testy","parent_id":"0","order":0,"leader_id":""},
			{"dep_id":"51","dep_name":"哈哈哈","parent_id":"1","order":2,"leader_id":""},
			{"dep_id":"52","dep_name":"主号客户轨迹","parent_id":"1","order":1,"leader_id":""},
			{"dep_id":"60","dep_name":"客服部","parent_id":"1","order":3,"leader_id":"7C0C33FFFCB96A88792FD40E32155843"},
			{"dep_id":"144","dep_name":"TEST部","parent_id":"1","order":0,"leader_id":""},
			{"dep_id":"175","dep_name":"test","parent_id":"1","order":0,"leader_id":""},
			{"dep_id":"176","dep_name":"test","parent_id":"175","order":0,"leader_id":""}
		]
	}
}`

	err := json.Unmarshal([]byte(data), &resp)
	if assert.Nil(t, err) {
		assert.Equal(t, 0, resp.ErrCode)
		assert.Equal(t, "success", resp.ErrMsg)
		assert.Equal(t, "17", resp.GetDepartmentsRsp.Total)
	}

}
