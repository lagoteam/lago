package result

type ErrorCode interface {
	GetCode() string
	GetMessage() string
	GetData() interface{}
}

/*var Results = [...]Result{
	{Code: "-1", Message: "系统内部错误"},
	{Code: "404", Message: "未登录"},
}*/

const (
	SuccessCode             = "200"
	FailCode                = "400"
	UnauthorizedCode        = "401"
	NotFoundCode            = "404"
	MethodNotAllowedCode    = "405"
	InternalServerErrorCode = "500"
)

var ResultCodeMessage = map[string]string{
	SuccessCode:             "成功",
	FailCode:                "失败",
	UnauthorizedCode:        "未授权",
	NotFoundCode:            "未找到",
	MethodNotAllowedCode:    "方法不被允许",
	InternalServerErrorCode: "服务器内部错误",
}
