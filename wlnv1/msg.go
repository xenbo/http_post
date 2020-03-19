package wlnv1

import (
	"encoding/json"
	"github.com/xenbo/http_post/log"
)

type System struct {
	App string `json:"_app"`   //请求键值
	Time  string `json:"_t"`  //时间戳
	SignKind string `json:"_sign_kind"`  //签名方法,支持 md5 和 hmac，默认 md5
	Sign string  `json:"_sign"`  //请求签名
	S string `json:"_s"`  //授权码
}

var Sys *System

func NewSystem(a string, t string, sik string, si string, s string) []byte {
	Sys = &System{
		App:       a,
		Time:t,
		SignKind: sik,
		Sign:      si,
		S:         s,
	}

	b, err := json.Marshal(Sys)
	log.DLog.Println(err)

	return b
}
