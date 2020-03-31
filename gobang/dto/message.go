package dto

import "gobang/constants"

type Message struct {
	Code int         `json:"code"`
	Info string      `json:"info"`
	Data interface{} `json:"data"`
}

func NewMsg(code int, info string, data interface{}) *Message {
	return &Message{Code: code, Info: info, Data: data}
}

func NewErrMsg(err error) *Message {
	return NewMsg(constants.Fail, "error", err.Error())
}
