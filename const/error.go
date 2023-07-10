package error

import "fmt"

type Error struct {
	code    string   `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details"`
}

func NewError(code string, msg string) *Error {
	return &Error{code: code, msg: msg}
}

func (e *Error) Code() string {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

func (e *Error) Details() []string {
	return e.details
}

var (
	Success                   = NewError("0000", "成功")
	ServerError               = NewError("FFFF", "服务内部错误")
	InvalidParams             = NewError("FFFF", "入参错误")
	NotFound                  = NewError("FFFF", "找不到")
	UnauthorizedAuthNotExist  = NewError("FFFF", "鉴权失败，找不到对应的 AppKey 和 AppSecret")
	UnauthorizedTokenError    = NewError("FFFF", "鉴权失败，Token 错误")
	UnauthorizedTokenTimeout  = NewError("FFFF", "鉴权失败，Token 超时")
	UnauthorizedTokenGenerate = NewError("FFFF", "鉴权失败，Token 生成失败")
	TooManyRequests           = NewError("FFFF", "请求过多")
)
