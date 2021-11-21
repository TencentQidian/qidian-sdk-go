package handler

import (
	"crypto/sha1"
	"encoding/hex"
	"log"
	"net/http"
	"sort"
	"strings"
)

// VerifyHandler verify request handler
type VerifyHandler struct {
	Token string
}

// NewVerifyHandler create verify handler
func NewVerifyHandler(token string) http.Handler {
	return &VerifyHandler{Token: token}
}

// ServerHTTP handle verify request
// Verify incoming request via GET method. So should handle GET method
// Reference: https://api.qidian.qq.com/wiki/doc/open/epko939s7aq8br19gz0i
func (h *VerifyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	signature := q.Get("signature")
	if signature == "" {
		log.Print("invalid signature")
		return
	}

	timestamp := q.Get("timestamp")
	if timestamp == "" {
		log.Print("invalid timestamp")
		return
	}

	nonce := q.Get("nonce")
	if nonce == "" {
		log.Print("invalid nonce")
		return
	}

	echoStr := q.Get("echostr")

	// GET /message_receiver?signature=f208bdcde5cd6b1c83e911446b9a318e7d59242c&timestamp=1637478915&nonce=kavlyvgg&echostr=cnopzpro
	s := Sign(h.Token, timestamp, nonce)
	if s != signature {
		log.Print("invalid sign")
		return
	}

	log.Println("echo: ", echoStr)
	// echo
	w.Write([]byte(echoStr))
}

// Sign return signature
func Sign(token, timestamp, nonce string) string {
	params := []string{token, timestamp, nonce}
	sort.Strings(params)
	signStr := strings.Join(params, "")

	h := sha1.New()
	h.Write([]byte(signStr))
	return hex.EncodeToString(h.Sum(nil))
}
