package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	//"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

type System struct {
	App string `json:"_app"`
	T   int64  `json:"_t"`
	//SignKind string `json:"_sign_kind"`
	Sign string `json:"_sign"`
	S    string `json:"_s"`
}

type Args struct {

	//<arg name="spec_code" title="规格编码" type="String"/>
	//<arg name="item_code" title="商品编码" type="String"/>
	//<arg name="bar_code" title="条码" type="String"/>

	//Key string `json:"keyword"`
	Code string `json:"item_code"`
	Page_no   int32 `json:"page"`   //1
	Page_size int32 `json:"limit"` //20
}

func httpPost() {
	ss := &System{}
	ss.App = "3123415742"
	ss.T =  time.Now().Unix()   //1573181240
	//ss.SignKind = "md5"
	ss.Sign = ""
	ss.S = ""

	ag := &Args{}
	ag.Code = "000029"
	//ag.Key = "围巾"
	ag.Page_no = 1
	ag.Page_size = 10

	data2, _ := json.Marshal(ss)
	fmt.Println(string(data2))

	var dat map[string]interface{}
	json.Unmarshal([]byte(string(data2)), &dat)

	data, _ := json.Marshal(ag)
	fmt.Println(string(data))

	json.Unmarshal([]byte(string(data)), &dat)

	var arr = make([]string, 0)
	for key, _ := range dat {
		arr = append(arr, key)
		//println(key, reflect.TypeOf(key).Name())
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

	cystr := "c3b5fee170b52b8397852c8ba03ef109" + v1.Encode() + "c3b5fee170b52b8397852c8ba03ef109"
	res := md5.Sum([]byte(cystr))
	md5str := hex.EncodeToString(res[:])

	println("Sign:", md5str)

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
		}  else {
			v.Add(key, urlcode)
		}
	}

	strbody := v.Encode()

	url := "http://114.67.231.162/api/erp/goods/spec/open/query"
	fmt.Println("url：",url,"\nPOST:", strbody)

	payload :=  strings.NewReader(strbody)
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()


	//
	//body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println("err:",err)


	datab := make([]byte, 30000)
	r := bufio.NewReader(resp.Body)
	r.Read(datab)

	fmt.Println(string(datab))

}

func main() {

	httpPost()

}
