package result

import (
	"fmt"
	"github.com/goravel/framework/facades"
	"testing"
)

func ResultTest(t *testing.T) {
	// 自定义数据结构
	data := map[string]interface{}{
		"name":    facades.Config().GetString("app.name"),
		"version": facades.Config().GetString("app.version"),
		"url":     facades.Config().GetString("app.url"),
	}

	// 返回1 预置状态码和消息 [规范]
	response := Success().SetCode(CodeSuccess).SetMessage(CodeMessage[CodeSuccess]).SetData(data)
	fmt.Println(response)

	// 返回2 自定义状态码和消息 [灵活]
	res := Success().SetCode("200").SetMessage("OK").SetData(data)
	fmt.Println(res)

	// 返回错误1 预置状态码和消息 [规范]
	resultError := NewResultError(CodeUnauthorized, CodeMessage[CodeUnauthorized])
	fmt.Println(resultError)

	// 返回错误2 自定义状态码和消息 [灵活]
	responseError := Fail().SetCode("500").SetMessage("Fail")
	fmt.Println(responseError)

	// panic
	panic(NewResultError("401", "unauthorized"))
}
