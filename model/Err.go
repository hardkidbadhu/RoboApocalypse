package model

import "fmt"

type ErrCode string

const (
	ErrInternalSRVError ErrCode = "ERR_INTERNAL_SERVER_ERROR"
	ErrInvalidFilter ErrCode = "ERR_INVALID_FILTER"
	ErrDB               ErrCode = "ERR_PG"
	ErrInvalidUSERID    ErrCode = "ERR_INVALID_USERID"
	ErrInvalidJSON      ErrCode = "ERR_INVALID_PAYLOAD"
)

type ErrResp struct {
	Message string  `json:"message"`
	ErrCode ErrCode `json:"err_code"`
	Err     error   `json:"error"`
}

func (e ErrResp) Error() string {
	return fmt.Sprintf("%s:%s - (%v)", e.Message, e.ErrCode, e.Err)
}
