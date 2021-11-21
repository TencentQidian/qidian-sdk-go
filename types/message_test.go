package types

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessage(t *testing.T) {
	data := `<Msg>
	<Appid>202019990</Appid>
	<MsgType>UnauthorizeApplication</MsgType>
	<ApplicationId>1300000983</ApplicationId>
	<UnauthorizeTime>1637481437</UnauthorizeTime>
	<selfDefine1></selfDefine1>
	<selfDefine2></selfDefine2>
	<selfDefine3></selfDefine3>
	<selfDefine4></selfDefine4>
	<selfDefine5></selfDefine5>
</Msg>`

	var m Message
	err := xml.Unmarshal([]byte(data), &m)
	if assert.Nil(t, err) {
		assert.Equal(t, "202019990", m.Appid)
		assert.Equal(t, "UnauthorizeApplication", m.MsgType)
		assert.Equal(t, "1300000983", m.ApplicationId)
		assert.Equal(t, "1637481437", m.UnauthorizeTime)
		assert.Equal(t, "", m.SelfDefine1)
		assert.Equal(t, "", m.SelfDefine2)
		assert.Equal(t, "", m.SelfDefine3)
		assert.Equal(t, "", m.SelfDefine4)
		assert.Equal(t, "", m.SelfDefine5)
	}
}
