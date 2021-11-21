package types

import "encoding/xml"

type Message struct {
	XMLName xml.Name `xml:"Msg"`

	Appid           string `xml:"Appid"`
	MsgType         string `xml:"MsgType"`
	ApplicationId   string `xml:"ApplicationId"`
	UnauthorizeTime string `xml:"UnauthorizeTime"`
	SelfDefine1     string `xml:"selfDefine1"`
	SelfDefine2     string `xml:"selfDefine2"`
	SelfDefine3     string `xml:"selfDefine3"`
	SelfDefine4     string `xml:"selfDefine4"`
	SelfDefine5     string `xml:"selfDefine5"`
}

// Command ?code=895666ac72553985af7479fc536d439a&state=705531fb080fe6377bec4f6fe8a1346d&app_id=1300000983&sid=1300000983
type Command struct {
	Code  string `json:"code"`
	State string `json:"state"`
	AppID int    `json:"app_id"`
	SID   int    `json:"sid"`
}
