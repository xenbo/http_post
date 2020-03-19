package wlnv1

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	tmap "github.com/liyue201/gostl/ds/map"
	"github.com/xenbo/http_post/log"
	"strings"
)

func createmap(bSys []byte, bBusiness []byte) *tmap.Map {
	var dat map[string]string
	json.Unmarshal([]byte(string(bBusiness)), &dat)
	json.Unmarshal([]byte(string(bSys)), &dat)

	m := tmap.New(tmap.WithGoroutineSafe())
	for key, val := range dat {
		m.Insert(key, val)
	}

	return m
}

func MakeSign(bSys []byte, bBusiness []byte, secret string) string {
	m := createmap(bSys, bBusiness)

	allString := ""
	for iter := m.First(); iter.IsValid(); iter.Next() {

		if strings.Count(strings.ToLower(iter.Key().(string)), "sign") > 0 &&
			strings.Count(strings.ToLower(iter.Key().(string)), "kind") == 0 {
			continue
		} else if iter.Value().(string) == "" {
			continue
		} else {
			allString += iter.Key().(string) + iter.Value().(string)
		}
	}

	//fmt.Println(allString)

	res := md5.Sum([]byte(secret + allString + secret))
	md5str := strings.ToUpper(hex.EncodeToString(res[:]))
	log.DLog.Println("sign:", md5str)

	i := 0
	strBody := ""
	for iter := m.First(); iter.IsValid(); iter.Next() {

		if i > 0 {
			strBody += "&"
		}
		strBody += iter.Key().(string)
		strBody += "="
		if strings.Count(strings.ToLower(iter.Key().(string)), "sign") > 0 &&
		strings.Count(strings.ToLower(iter.Key().(string)), "kind") == 0 {
			strBody += md5str
		} else {
			strBody += iter.Value().(string)
		}
		i++
	}
	log.DLog.Println("strBody:", strBody)

	return strBody
}
