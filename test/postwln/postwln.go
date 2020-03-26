package postwln

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

type SystemArg interface {
	SetNewTime()
}

func tryHttpPost(encrypstr string, SysMsg interface{}, ArgMsg interface{}, strurl string, body *[]byte, times int32) int32 {

	tryTimes := int32(0)
	retFlag := int32(0)

	sysdata, _ := json.Marshal(SysMsg)
	//fmt.Println(string(sysdata))

	var dat map[string]interface{}
	json.Unmarshal([]byte(string(sysdata)), &dat)

	argdata, _ := json.Marshal(ArgMsg)
	//fmt.Println(string(argdata))

	json.Unmarshal([]byte(string(argdata)), &dat)

	var arr = make([]string, 0)
	for key, _ := range dat {
		arr = append(arr, key)
	}
	sort.Strings(arr)

	v1 := url.Values{}
	for _, key := range arr {
		value := dat[key]
		var urlcode string
		if reflect.TypeOf(value).Name() == "string" {
			urlcode = reflect.ValueOf(value).String()
		} else if reflect.TypeOf(value).Name() == "float64" { //fuck bugs
			n := int64(reflect.ValueOf(value).Float())
			urlcode = strconv.FormatInt(n, 10)
		} else if reflect.TypeOf(value).Name() == "int" {
			n := int64(reflect.ValueOf(value).Int())
			urlcode = string(n)
		}

		//println(key, "  ", urlcode)

		if key == "_sign" {
			continue
		} else {
			v1.Add(key, urlcode)
		}
	}

	cystr := encrypstr + v1.Encode() + encrypstr
	res := md5.Sum([]byte(cystr))
	md5str := hex.EncodeToString(res[:])

	//println("Sign:", md5str)

	v := url.Values{}
	for _, key := range arr {
		value := dat[key]
		var urlcode string
		if reflect.TypeOf(value).Name() == "string" {
			urlcode = reflect.ValueOf(value).String()
		} else if reflect.TypeOf(value).Name() == "float64" { //fuck bugs
			n := int64(reflect.ValueOf(value).Float())
			urlcode = strconv.FormatInt(n, 10)
		} else if reflect.TypeOf(value).Name() == "int" {
			n := int64(reflect.ValueOf(value).Int())
			urlcode = string(n)
		}

		if key == "_sign" {
			v.Add(key, md5str)
		} else {
			v.Add(key, urlcode)
		}
	}

	strbody := v.Encode()
	//fmt.Println("url：", strurl, "\nPOST:", strbody)

	var resp *http.Response

	{
	agin:
		payload := strings.NewReader(strbody)
		req, _ := http.NewRequest("POST", strurl, payload)
		req.Header.Add("content-type", "application/x-www-form-urlencoded")

		resp, _ = http.DefaultClient.Do(req)
		*body, _ = ioutil.ReadAll(resp.Body)

		i3 := strings.Index(string(*body), `{"code":0,`)
		if i3 < 0 {
			time.Sleep(time.Second * 2)
			if tryTimes < times {
				fmt.Println(string(*body))

				tryTimes++
				goto agin
			}
			retFlag = -1
		}
	}
	resp.Body.Close()

	return retFlag
}

func HttpPost(encrypstr string, SysMsg SystemArg, ArgMsg interface{}, strurl string) []byte {

	var body = make([]byte, 4096)

	for t := 0; t < 5; t++ {
		f := tryHttpPost(encrypstr, SysMsg, ArgMsg, strurl, &body, 3)

		if f == 0 {
			return body
		} else if f < 0 {
			SysMsg.SetNewTime()
		}
	}

	return []byte("")
}
