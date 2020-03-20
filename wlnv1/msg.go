package wlnv1

import (
	"encoding/json"
	"github.com/xenbo/http_post/log"
)

type System struct {
	App      string `json:"_app"`       //请求键值
	Time     string `json:"_t"`         //时间戳
	SignKind string `json:"_sign_kind"` //签名方法,支持 md5 和 hmac，默认 md5
	Sign     string `json:"_sign"`      //请求签名
	S        string `json:"_s"`         //授权码
}

var Sys *System

func NewSystem(App string, Time string, SignKind string, Sign string, S string) []byte {
	Sys = &System{
		App:      App,
		Time:     Time,
		SignKind: SignKind,
		Sign:     Sign,
		S:        S,
	}

	b, err := json.Marshal(Sys)
	log.DLog.Println(err)

	return b
}
