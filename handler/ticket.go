package handler

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
)

type TicketMessage struct {
	XMLName               xml.Name `xml:"xml"`
	AppId                 string   `xml:"AppId"`
	CreateTime            int      `xml:"CreateTime"`
	InfoType              string   `xml:"InfoType"`
	ComponentVerifyTicket string   `xml:"ComponentVerifyTicket"`
}

type TicketHandler struct {
	crypto *Crypto
	fn     func(*TicketMessage)
}

func NewTicketHandler(c *Crypto, f func(*TicketMessage)) http.Handler {
	return &TicketHandler{crypto: c, fn: f}
}

func (h *TicketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var e EncryptedMessage

	xmlData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
		return
	}

	if err = xml.Unmarshal(xmlData, &e); err != nil {
		log.Print("invalid xml format: " + string(xmlData))
		return
	}

	var m TicketMessage
	err = e.DecryptTo(h.crypto, &m)
	if err != nil {
		log.Print(err)
		return
	}

	log.Printf("Tikcet: %+v", &m)

	if h.fn != nil {
		h.fn(&m)
	}
}
