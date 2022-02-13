package global

const (
	SUCCESS                        = 200
	INTERNAL_ERROR                 = 500
	INVALID_PARAMS                 = 400
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 401
)

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	INTERNAL_ERROR:                 "内部错误",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "token超时",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[INTERNAL_ERROR]
}
