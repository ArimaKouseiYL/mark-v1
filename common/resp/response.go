package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

const (
	CodeSuccess = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
	CodeRegistrySuccess
	CodeGenTokenError
)

var msgMap = map[int]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户已存在",
	CodeUserNotExist:    "用户不存在",
	CodeInvalidPassword: "密码错误",
	CodeServerBusy:      "服务繁忙",
	CodeRegistrySuccess: "用户注册成功",
	CodeGenTokenError:   "token生成失败",
}

func GetMsg(code int) string {
	msg, ok := msgMap[code]
	if !ok {
		msg = msgMap[CodeServerBusy]
	}
	return msg
}

func ResponseError(c *gin.Context, code int) {
	c.JSON(http.StatusOK, &Result{
		Code: code,
		Msg:  GetMsg(code),
		Data: nil,
	})
}

func ResponseErrorWithMsg(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, &Result{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Result{
		Code: CodeSuccess,
		Msg:  GetMsg(CodeSuccess),
		Data: data,
	})
}

func ResponseSuccessWithCode(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, &Result{
		Code: code,
		Msg:  GetMsg(code),
		Data: data,
	})
}
