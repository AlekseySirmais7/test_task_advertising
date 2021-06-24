package models

type Answer struct {
	Body interface{} `json:"body"`
	Code uint        `json:"code"`
	Msg  string      `json:"msg"`
}
