package codes

// 错误码定义
const (
	SUCCESS       = 200
	ERROR         = 500
	InvalidParams = 400

	ErrAuthCheckTokenFail    = 20001
	ErrAuthCheckTokenTimeout = 20002
	ErrAuthToken             = 20003
	ErrAuth                  = 20004

	ErrExistUser    = 30001
	ErrNotFoundUser = 30002

	PageNotFound = 40001
)

// MsgFlags 错误信息
var MsgFlags = map[int]string{
	SUCCESS:                  "ok",
	ERROR:                    "fail",
	InvalidParams:            "请求参数错误",
	ErrAuthCheckTokenFail:    "Token鉴权失败",
	ErrAuthCheckTokenTimeout: "Token已超时",
	ErrAuthToken:             "Token生成失败",
	ErrAuth:                  "Token错误",
	PageNotFound:             "Page not found",
	ErrExistUser:             "用户已存在",
	ErrNotFoundUser:          "用户不存在",
}

// GetMsg 获取错误信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
