package codes

import (
	"fmt"
	"strings"
)

type CodeError struct {
	ErrCode int
	ErrMsg  string
}

func (e *CodeError) String() string {
	return e.ErrMsg
}

func (e *CodeError) GetErrCode() int {
	return e.ErrCode
}

func (e *CodeError) GetErrMsg() string {
	// 去除rpc层错误信息前缀
	return strings.ReplaceAll(e.ErrMsg, "rpc error: code = Unknown desc = ", "")
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.ErrCode, e.ErrMsg)
}

func NewErrCodeMsg(errCode int, errMsg string) *CodeError {
	return &CodeError{ErrCode: errCode, ErrMsg: errMsg}
}

func NewErrCode(errCode int) *CodeError {
	return &CodeError{ErrCode: errCode, ErrMsg: GetMsg(errCode)}
}

func NewErrMsg(errMsg string) *CodeError {
	return &CodeError{ErrCode: ERROR, ErrMsg: errMsg}
}
