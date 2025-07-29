package controllers

import "github.com/gin-gonic/gin"

type ResCode int16

const (
	CodeInvalidParam ResCode = 1000 + iota
	CodeServerBusy
	CodeAdminLogin
	CodeHeaderAuthFailed
	CodeSignupSuccess
	CodeSignupFailed
	CodeLoginSuccess
	CodeLoginFailed
	CodeQuerySuccess
	CodeQueryFailed
	CodeUpdateSuccess
	CodeUpdateFailed
	CodeAddSuccess
	CodeAddFailed
	CodeDeleteSuccess
	CodeDeleteFailed
	CodeCreateNoteSuccess
	CodeCreateNoteFailed
	CodeGetNotesSuccess
	CodeGetNotesFailed
	CodeDelNoteSuccess
	CodeDelNoteFailed
)

var codeMsgMap = map[ResCode]string{
	CodeInvalidParam:      "请求参数错误",
	CodeServerBusy:        "服务器繁忙",
	CodeSignupSuccess:     "注册成功",
	CodeSignupFailed:      "注册失败",
	CodeAdminLogin:        "管理员登录",
	CodeHeaderAuthFailed:  "请求头中头部验证失败",
	CodeLoginSuccess:      "登录成功",
	CodeLoginFailed:       "登录失败",
	CodeQuerySuccess:      "查询成功",
	CodeQueryFailed:       "查询失败",
	CodeUpdateSuccess:     "更新成功",
	CodeUpdateFailed:      "更新失败",
	CodeAddSuccess:        "添加成功",
	CodeAddFailed:         "添加失败",
	CodeDeleteSuccess:     "删除成功",
	CodeDeleteFailed:      "删除失败",
	CodeCreateNoteSuccess: "添加笔记成功",
	CodeCreateNoteFailed:  "添加笔记失败",
	CodeGetNotesSuccess:   "获取笔记成功",
	CodeGetNotesFailed:    "获取笔记失败",
	CodeDelNoteSuccess:    "删除笔记成功",
	CodeDelNoteFailed:     "删除笔记失败",
}

// CodeMsgData 生成响应数据
func CodeMsgDetail(rescode ResCode, detail any) gin.H {
	return gin.H{
		"code":   rescode,
		"msg":    codeMsgMap[rescode],
		"detail": detail,
	}
}
