package xerr

import (
	"fmt"
)

/**
常用通用固定错误
*/

type CodeError struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
}

type CodeErrorResponse struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
}

// 返回给前端的错误码
func (e *CodeError) GetErrCode() uint32 {
	return e.Code
}

// 返回给前端显示端错误信息
func (e *CodeError) GetErrMsg() string {
	return e.Msg
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d,ErrMsg:%s", e.Code, e.Msg)
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}

func NewRpcError(code uint32, msg string) *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: code,
		Msg:  msg,
	}
}

func NewErrCodeMsg(code uint32, msg string) *CodeError {
	return &CodeError{Code: code, Msg: msg}
}
func NewErrCode(code uint32) *CodeError {
	return &CodeError{Code: code, Msg: MapErrMsg(code)}
}

func NewErrMsg(msg string) *CodeError {
	return &CodeError{Code: SERVER_COMMON_ERROR, Msg: msg}
}
