package models

type APIResponse struct {
	Code   int         `json:"code"`   // 状态码
	Msg    string      `json:"msg"`    // 状态消息
	Detail interface{} `json:"detail"` //额外携带的详细信息
}

type APINoteCreateSuccessResponse struct {
	Code   int          `json:"code"`   //状态码
	Msg    string       `json:"msg"`    //状态信息
	Detail NoteResponse `json:"detail"` //详细信息
}

type APINotesGetSuccessResponse struct {
	Code   int             `json:"code"`   //状态码
	Msg    string          `json:"msg"`    //状态信息
	Detail []NoteResponse `json:"detail"` //详细信息
}
type APINoteFailed struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Delete string `json:"delete"`
}
