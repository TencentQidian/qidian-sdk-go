package qidian_sdk_go

type Error interface {
	Code() int
	Message() string
}

type ErrorV1 struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
