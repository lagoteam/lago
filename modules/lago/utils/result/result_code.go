package result

type ResultCode struct {
	Code    interface{} `json:"code"`
	Message string      `json:"message"`
}

func NewResultCode() *ResultCode {
	return &ResultCode{}
}

func (r *ResultCode) Success() *ResultCode {
	r.Code = CodeSuccess
	r.Message = CodeMessage[CodeSuccess]
	return r
}

func (r *ResultCode) Fail() *ResultCode {
	r.Code = CodeFail
	r.Message = CodeMessage[CodeFail]
	return r
}

/*var Codes = [...]ResultCode{
	{Code: "-1", Message: "系统内部错误"},
	{Code: "404", Message: "未登录"},
}*/

const (
	CodeSuccess             = "200"
	CodeFail                = "400"
	CodeUnauthorized        = "401"
	CodeNotFound            = "404"
	CodeMethodNotAllowed    = "405"
	CodeInternalServerError = "500"
)

var CodeMessage = map[interface{}]string{
	CodeSuccess:             "成功",
	CodeFail:                "失败",
	CodeUnauthorized:        "未授权",
	CodeNotFound:            "未找到",
	CodeMethodNotAllowed:    "方法不被允许",
	CodeInternalServerError: "服务器内部错误",
}
