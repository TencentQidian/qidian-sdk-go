package handler

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractXML(t *testing.T) {
	crypto, err := NewCrypto("KLAvbsjhj2PsozUaQ48k3dKrCx8Hnnxxfy9HlclEeOG", "202019990")
	if err != nil {
		panic(err)
	}

	var data = `<xml><AppId><![CDATA[202019990]]></AppId><Encrypt><![CDATA[ARYpdk4oYm6GaCS+MYcUc6tuOaq0LtHijXNmQWdih41m7ywXwW3tBfc9CR2QUmAlLPjZAVgWtmO5EM9rIQYU5XpBCWQ+BHtb4gUNK+AgpkQsgL9qcKEQqqbLxbs8Cuc/vW+cOy2pEOZRmfGrIRZbijBDO+aFWFKv9jxdiC47Gk3GZr5ypQILL9BdiLCWz+JNgj87OIxCSOFsEbF9ydUI/CGltx2+TaI0Wt1Q4KVhqH4qhSZwk6txSRFp3Z9EhkNyn48LoJKPf9/t1S0qMufGqX9MgDD36GN7ZJgh1g6o2UTMVk6dbgZqtfcZjwvtohzoYwE5HW83hSdKHnVuZDVKjg==]]></Encrypt></xml>`

	var e EncryptedMessage
	err = xml.NewDecoder(bytes.NewReader([]byte(data))).Decode(&e)
	if assert.Nil(t, err) {
		assert.Equal(t, e.AppId.String(), "202019990")

		var m TicketMessage
		err = e.DecryptTo(crypto, &m)
		if assert.Nil(t, err) {
			assert.Equal(t, "component_verify_ticket", m.InfoType)
		}
	}
}
