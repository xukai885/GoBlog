package controllers

type ResCode int

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidToken
	CodeNeedLogin
	CodeServerBusy
	CodeInvalidParam
	//CodeUserNotExist
	//CodeUserExist
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:      "success",
	CodeInvalidToken: "无效token",
	CodeNeedLogin:    "需要登陆",
	CodeServerBusy:   "系统繁忙",
	CodeInvalidParam: "请求参数有误",
	//CodeUserNotExist: "用户不存在",
	//CodeUserExist:    "用户已存在",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
