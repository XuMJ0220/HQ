package controllers

import "github.com/gin-gonic/gin"

type ResCode int16

const (
	CodeInvalidParam ResCode = 1000 + iota
	CodeServerBusy
	CodeSignupSuccess
	CodeSignupFailed
	CodeLoginSuccess
	CodeLoginFailed
)

var codeMsgMap = map[ResCode]string{
	CodeInvalidParam:  "请求参数错误",
	CodeServerBusy:    "服务器繁忙",
	CodeSignupSuccess: "注册成功",
	CodeSignupFailed:  "注册失败",
	CodeLoginSuccess:  "登录成功",
	CodeLoginFailed:   "登录失败",
}

// CodeMsgData 生成响应数据
func CodeMsgDetail(rescode ResCode, detail any) gin.H {
	return gin.H{
		"code": rescode,
		"msg":  codeMsgMap[rescode],
		"detail": detail,
	}
}
