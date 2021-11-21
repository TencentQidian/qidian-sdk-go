package handler

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tencentqidian/qidian-sdk-go/types"
)

type cData struct {
	Value string `xml:",cdata"`
}

func (c cData) String() string {
	return c.Value
}

type EncryptedMessage struct {
	XMLName xml.Name `xml:"xml"`
	Encrypt cData    `xml:"Encrypt"`
	AppId   cData    `xml:"AppId"`
}

func (e *EncryptedMessage) DecryptTo(c *Crypto, v interface{}) error {
	plaintext, err := c.Decrypt(e.Encrypt.Value)
	if err != nil {
		return err
	}
	return xml.Unmarshal([]byte(plaintext), v)
}

// MessageHandler a message receiver that implements http.Handler
type MessageHandler struct {
	crypto   *Crypto
	handlers map[string]func(*types.Message)
}

type MessageOption func(*MessageHandler)

// On listen a message type and handler
func On(msgType string, fn func(m *types.Message)) MessageOption {
	return func(h *MessageHandler) {
		h.handlers[msgType] = fn
	}
}

// NewMessageHandler create instance of MessageHandler
func NewMessageHandler(c *Crypto, opt ...MessageOption) *MessageHandler {
	h := &MessageHandler{
		crypto:   c,
		handlers: map[string]func(*types.Message){},
	}

	for _, o := range opt {
		o(h)
	}

	return h
}

// ServeHTTP handle message notify
func (r *MessageHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defer func() {
		// acknowledge 'success'
		w.Write([]byte("success"))
	}()

	var e EncryptedMessage
	topXMLData, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Print(err)
		return
	}

	if err = xml.Unmarshal(topXMLData, &e); err != nil {
		log.Print("invalid xml format: " + string(topXMLData))
		return
	}

	plaintext, err := r.crypto.Decrypt(e.Encrypt.Value)
	if err != nil {
		log.Println("decrypt message err: ", err)
		return
	}

	var m types.Message
	err = xml.Unmarshal([]byte(plaintext), &m)
	if err != nil {
		log.Println("unmarshal message err: ", err)
		return
	}
	log.Println("message: ", m)

	if handler, ok := r.handlers[m.MsgType]; ok {
		handler(&m)
	}
}
