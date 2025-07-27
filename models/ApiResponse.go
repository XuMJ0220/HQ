package models

type APIResponse struct{
	Code int `json:"code"` // 状态码
	Msg  string `json:"msg"` // 状态消息
	Detail interface{} `json:"detail"` //额外携带的详细信息
}