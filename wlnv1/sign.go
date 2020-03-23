package wlnv1

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	tmap "github.com/liyue201/gostl/ds/map"
	"github.com/xenbo/http_post/log"
	"net/url"
	"reflect"
	"strings"
)

func assign(value interface{}) (string, bool) {

	tmpValue := fmt.Sprint(value)
	isString := false

	t := reflect.ValueOf(value).Type()
	kind := t.Kind()

	switch kind {
	case reflect.Bool:
		if value.(bool) {
			tmpValue = "true"
		} else {
			tmpValue = "false"
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if value.(int64) < 0 {
			tmpValue = ""
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if value.(int64) < 0 {
			tmpValue = ""
		}
	case reflect.Float32:
		if value.(float32) < 0.0 {
			tmpValue = ""
		} else {
			tmpValue = fmt.Sprintf("%f", value)
		}
	case reflect.Float64:
		if value.(float64) < 0.0 {
			tmpValue = ""
		} else {
			tmpValue = fmt.Sprintf("%f", value)
		}
	case reflect.String:
		isString = true
	case reflect.Slice: //TODO...

	case reflect.Map: //TODO...
		dat := value.(map[string]interface{})
		tmpValue = "{"
		for key, value2 := range dat {
			log.LOGI(key, ":", value2)

			t2 := reflect.ValueOf(value2).Type()
			kind2 := t2.Kind()

			if kind2 == reflect.Slice {
				for _, vSlice := range value2.([]interface{}) {
					v, b := assign(vSlice)
					if len(v) > 0 {
						tmpValue += `"` + key + `"` + ":["
						if b {
							tmpValue += `"` + v + `"`
						} else {
							tmpValue += v
						}
						tmpValue += "],"
					}
				}
			} else {
				v, b := assign(value2)

				if len(v) > 0 {
					tmpValue += `"` + key + `"` + ":"
					if b {
						tmpValue += `"` + v + `"`
					} else {
						tmpValue += v
					}
					tmpValue += ","
				}
			}

		}
		tmpValue = strings.TrimRight(tmpValue, ",")
		tmpValue += "}"

	case reflect.Struct:

	case reflect.Ptr:

	}

	return tmpValue, isString
}

func createMap(bSys []byte, bBusiness []byte) *tmap.Map {

	var FromJSON interface{}
	decoder := json.NewDecoder(bytes.NewReader(bSys))
	decoder.UseNumber()
	decoder.Decode(&FromJSON)
	datSys := FromJSON.(map[string]interface{})

	decoder = json.NewDecoder(bytes.NewReader(bBusiness))
	//decoder.UseNumber()
	decoder.Decode(&FromJSON)
	datBus := FromJSON.(map[string]interface{})

	m := tmap.New(tmap.WithGoroutineSafe())
	for key, value := range datSys {
		tmpValue, _ := assign(value)
		m.Insert(key, tmpValue)
	}

	for key, value := range datBus {
		log.LOGI(key, ":", value)

		tmpValue, _ := assign(value)
		m.Insert(key, tmpValue)
	}

	for iter := m.First(); iter.IsValid(); iter.Next() {
		log.LOGI(iter.Key().(string), ":", iter.Value().(string))
	}

	return m
}

func MakeSign(bSys []byte, bBusiness []byte, secret string) string {
	m := createMap(bSys, bBusiness)

	//i := 0
	//allString := ""
	//for iter := m.First(); iter.IsValid(); iter.Next() {
	//
	//	if len(iter.Value().(string)) <= 0 {
	//		continue
	//	}
	//
	//	if i > 0 {
	//		allString += "&"
	//	}
	//	allString += iter.Key().(string)
	//	allString += "="
	//	allString += iter.Value().(string)
	//	i++
	//}

	signAllString := url.Values{}
	for iter := m.First(); iter.IsValid(); iter.Next() {

		if len(iter.Value().(string)) <= 0 {
			continue
		}
		signAllString.Add(iter.Key().(string), iter.Value().(string))
	}



	strSign := secret + signAllString.Encode() + secret
	log.DLog.Println("strSign:", strSign)
	fmt.Println(strSign)

	res := md5.Sum([]byte(strSign))
	md5str := strings.ToUpper(hex.EncodeToString(res[:]))
	log.DLog.Println("finish sign:", md5str)

	//i = 0
	//strBody := ""
	//for iter := m.First(); iter.IsValid(); iter.Next() {
	//	isSign := false
	//	if strings.Count(strings.ToLower(iter.Key().(string)), "sign") > 0 &&
	//		strings.Count(strings.ToLower(iter.Key().(string)), "kind") == 0 {
	//		isSign = true
	//	}
	//
	//	if len(iter.Value().(string)) <= 0 && !isSign {
	//		continue
	//	}
	//
	//	if i > 0 {
	//		strBody += "&"
	//	}
	//	strBody += iter.Key().(string)
	//	strBody += "="
	//	if isSign {
	//		strBody += md5str
	//	} else {
	//		strBody += iter.Value().(string)
	//	}
	//	i++
	//}


	strBody := url.Values{}
	for iter := m.First(); iter.IsValid(); iter.Next() {
		isSign := false
		if strings.Count(strings.ToLower(iter.Key().(string)), "sign") > 0 &&
			strings.Count(strings.ToLower(iter.Key().(string)), "kind") == 0 {
			isSign = true
		}

		if len(iter.Value().(string)) <= 0 && !isSign {
			continue
		}

		if isSign {
			strBody.Add(iter.Key().(string), md5str)
		} else {
			strBody.Add(iter.Key().(string), iter.Value().(string))
		}
	}

	log.DLog.Println("strBody:", strBody.Encode())
	fmt.Println(strBody.Encode())


	return strBody.Encode()
}

func MakeB2CSign(bSys []byte, bBusiness []byte, secret string) string {
	return ""
}
