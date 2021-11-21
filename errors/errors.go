package errors

import "errors"

var (
	ErrInvalidCallbackRequest = errors.New("invalid callback request")
)

type Err struct {
	Code    int    `json:"errcode"`
	Message string `json:"errmsg"`
}

// Error implements the error interface.
func (e *Err) Error() string {
	return e.Message
}
