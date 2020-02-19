package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

type System struct {
	Appkey    string `json:"app_key"`
	TimeStamp int64  `json:"timestamp"`
	Sign      string `json:"sign"`
	Format    string `json:"format"`
}

type Args struct {
	ModifyTime string `json:"modify_time"`
	Page_no    int32  `json:"page"`  //1
	Page_size  int32  `json:"limit"` //20
}

type Args2 struct {
	ModifyTime string `json:"modify_time"`
	Page_no    int32  `json:"page"`  //1
	Page_size  int32  `json:"limit"` //20
	BillType   int32  `json:"bill_type"`
}

type Msg interface {
	name() string
	RequestData(int32, string, string, string)
}

//---------------------------------------------
type StockDetails struct {
	DetailId string  `json:"detail_id"`
	SkuNo    string  `json:"sku_no"`
	Size     float64 `json:"size"`
	SnValue  string  `json:"sn_value"`
}

type StockResponse struct {
	//销售出库单接
	ExpressCode bool           `json:"express_code"`
	Express     string         `json:"express"`
	CustomCode  string         `json:"custom_code"`
	CustomNick  string         `json:"custom_nick"`
	CustomName  string         `json:"custom_name"`
	StorageCode string         `json:"storage_code"`
	StorageName string         `json:"storage_name"`
	BillDate    string         `json:"bill_date"`
	Provice     string         `json:"provice"`
	City        string         `json:"city"`
	Company     string         `json:"company"`
	ShopNick    string         `json:"shop_nick"`
	Details     []StockDetails `json:"details"`
}

type StockMsg struct {
	Success   bool            `json:"success"`
	ErrorCode string          `json:"error_code"`
	ErrorMsg  string          `json:"error_msg"`
	Response  []StockResponse `json:"response"`
}

func (smsg *StockMsg) name() string {
	return "stock"
}

//---------------------------------------------

type PurchaseDetails struct {
	DetailId string  `json:"detail_id"`
	SkuNo    string  `json:"sku_no"`
	Size     float64 `json:"size"`
	SnValue  string  `json:"sn_value"`
}

type PurchaseResponse struct {
	StorageCode string            `json:"storage_code"`
	StorageName string            `json:"storage_name"`
	BillDate    string            `json:"bill_date"`
	Company     string            `json:"company"`
	Details     []PurchaseDetails `json:"details"`
}

type PurchaseMsg struct {
	Success   bool               `json:"success"`
	ErrorCode string             `json:"error_code"`
	ErrorMsg  string             `json:"error_msg"`
	Response  []PurchaseResponse `json:"response"`
}

func (msg *PurchaseMsg) name() string {
	return "purchase"
}

//---------------------------------------------

type inventoryDetails struct {
	DetailId string  `json:"detail_id"`
	SkuNo    string  `json:"sku_no"`
	Size     float64 `json:"size"`
	SnValue  string  `json:"sn_value"`
}

type inventoryResponse struct {
	StorageCode string            `json:"storage_code"`
	StorageName string            `json:"storage_name"`
	BillDate    string            `json:"bill_date"`
	Company     string            `json:"company"`
	Details     []PurchaseDetails `json:"details"`
}

type inventoryMsg struct {
	Success   bool               `json:"success"`
	ErrorCode string             `json:"error_code"`
	ErrorMsg  string             `json:"error_msg"`
	Response  []PurchaseResponse `json:"response"`
}

func (msg *inventoryMsg) name() string {
	return "inventory"
}

//---------------------------------------------

func httpPost(ss *System, ag *Args, url string, secret string) string {

	data2, _ := json.Marshal(ss)
	//fmt.Println(string(data2))

	var dat map[string]interface{}
	json.Unmarshal([]byte(string(data2)), &dat)

	data, _ := json.Marshal(ag)
	//fmt.Println(string(data))

	json.Unmarshal([]byte(string(data)), &dat)

	var arr = make([]string, 0)
	for key, _ := range dat {
		arr = append(arr, key)
		//println(key, reflect.TypeOf(key).Name())
	}
	sort.Strings(arr)

	allstring := ""
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

		if key == "sign" {
			continue
		} else {
			//println(key, ":", urlcode)
			allstring += key + urlcode
		}
	}

	res := md5.Sum([]byte(secret + allstring + secret))
	md5str := strings.ToUpper(hex.EncodeToString(res[:]))

	//println("sign:", md5str)

	strbody := ""
	i := 0
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

		if i > 0 {
			strbody += "&"
		}
		strbody += key
		strbody += "="
		if key == "sign" {
			strbody += md5str
		} else {
			strbody += urlcode
		}
		i++
	}

	//fmt.Println("url：", url+"?"+strbody)

	payload := strings.NewReader(strbody)
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/x-www-form-urlencoded; charset=utf-8")
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}

func httpPost2(ss *System, ag *Args2, url string, secret string) string {

	data2, _ := json.Marshal(ss)
	//fmt.Println(string(data2))

	var dat map[string]interface{}
	json.Unmarshal([]byte(string(data2)), &dat)

	data, _ := json.Marshal(ag)
	//fmt.Println(string(data))

	json.Unmarshal([]byte(string(data)), &dat)

	var arr = make([]string, 0)
	for key, _ := range dat {
		arr = append(arr, key)
		//println(key, reflect.TypeOf(key).Name())
	}
	sort.Strings(arr)

	allstring := ""
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

		if key == "sign" {
			continue
		} else {
			//println(key, ":", urlcode)
			allstring += key + urlcode
		}
	}

	res := md5.Sum([]byte(secret + allstring + secret))
	md5str := strings.ToUpper(hex.EncodeToString(res[:]))

	//println("sign:", md5str)

	strbody := ""
	i := 0
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

		if i > 0 {
			strbody += "&"
		}
		strbody += key
		strbody += "="
		if key == "sign" {
			strbody += md5str
		} else {
			strbody += urlcode
		}
		i++
	}

	//fmt.Println("url：", url+"?"+strbody)

	payload := strings.NewReader(strbody)
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/x-www-form-urlencoded; charset=utf-8")
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}

func (data *inventoryMsg) RequestData(pid int32, appkey string, secret string, url string) {

	ss := &System{}
	ss.Appkey = appkey
	ss.TimeStamp = time.Now().UnixNano() / 1e6 //1581415369242
	ss.Sign = ""
	ss.Format = "json"

	//ag := &Args{}
	//ag.ModifyTime = "2020-02-12 00:00:00" // time.Now().Format("2006-01-02 15:04:11")  //"2019-12-01 00:00:00"
	//ag.Page_no = 1
	//ag.Page_size =1

	ag2 := &Args2{}
	ag2.ModifyTime = "2019-12-01 00:00:00" // time.Now().Format("2006-01-02 15:04:11")  //"2019-12-01 00:00:00"
	ag2.Page_no = pid
	ag2.Page_size = 1
	ag2.BillType = 1

	msg := httpPost2(ss, ag2, url, secret)
	time.Sleep(time.Microsecond)

	msg = strings.Replace(msg, `]"`, "]", -1)
	msg = strings.Replace(msg, `"[`, "[", -1)
	msg = strings.Replace(msg, `\`, "", -1)

	var buffer bytes.Buffer
	buffer.Write([]byte(msg))

	err := json.Unmarshal(buffer.Bytes(), data)
	if err == nil {
		//	for id := 0; id < len(data.Response); id++ {
		//		fmt.Println(data.Response[id].Details[0].SkuNo)
		//		fmt.Println("-------------------------------------------------------------")
		//	}
	}
}

func (data *StockMsg) RequestData(pid int32, appkey string, secret string, url string) {

	ss := &System{}
	ss.Appkey = appkey
	ss.TimeStamp = time.Now().UnixNano() / 1e6 //1581415369242
	ss.Sign = ""
	ss.Format = "json"

	ag := &Args{}
	ag.ModifyTime = "2020-02-12 00:00:00" // time.Now().Format("2006-01-02 15:04:11")  //"2019-12-01 00:00:00"
	ag.Page_no = pid
	ag.Page_size = 1

	//ag2 := &Args2{}
	//ag2.ModifyTime = "2019-12-01 00:00:00" // time.Now().Format("2006-01-02 15:04:11")  //"2019-12-01 00:00:00"
	//ag2.Page_no = 1
	//ag2.Page_size = 10
	//ag2.BillType = 0

	msg := httpPost(ss, ag, url, secret)
	time.Sleep(time.Microsecond)

	msg = strings.Replace(msg, `]"`, "]", -1)
	msg = strings.Replace(msg, `"[`, "[", -1)
	msg = strings.Replace(msg, `\`, "", -1)

	var buffer bytes.Buffer
	buffer.Write([]byte(msg))

	err := json.Unmarshal(buffer.Bytes(), data)
	if err == nil {
		//	for id := 0; id < len(data.Response); id++ {
		//		fmt.Println(data.Response[id].Details[0].SkuNo)
		//		fmt.Println("-------------------------------------------------------------")
		//	}
	}
}

func (data *PurchaseMsg) RequestData(pid int32, appkey string, secret string, url string) {

	ss := &System{}
	ss.Appkey = appkey
	ss.TimeStamp = time.Now().UnixNano() / 1e6 //1581415369242
	ss.Sign = ""
	ss.Format = "json"

	ag := &Args{}
	ag.ModifyTime = "2019-12-01 00:00:00" // time.Now().Format("2006-01-02 15:04:11")  //"2019-12-01 00:00:00"
	ag.Page_no = pid
	ag.Page_size = 1

	//ag2 := &Args2{}
	//ag2.ModifyTime = "2019-12-01 00:00:00" // time.Now().Format("2006-01-02 15:04:11")  //"2019-12-01 00:00:00"
	//ag2.Page_no = 1
	//ag2.Page_size = 10
	//ag2.BillType = 0

	msg := httpPost(ss, ag, url, secret)
	time.Sleep(time.Microsecond)

	msg = strings.Replace(msg, `]"`, "]", -1)
	msg = strings.Replace(msg, `"[`, "[", -1)
	msg = strings.Replace(msg, `\`, "", -1)

	var buffer bytes.Buffer
	buffer.Write([]byte(msg))

	err := json.Unmarshal(buffer.Bytes(), data)
	if err == nil {
		//	for id := 0; id < len(data.Response); id++ {
		//		fmt.Println(data.Response[id].Details[0].SkuNo)
		//		fmt.Println("-------------------------------------------------------------")
		//	}
	}

}

func GetDataLoop(appkey string, secret string, url string, msg Msg) {

	for i := 0; i < 10000; i++ {
		msg.RequestData(int32(i), appkey, secret, url)
		fmt.Print(msg.name(), "   ")
		fmt.Println(msg)
	}
}

func main() {
	appkey := "QC20201112"
	secret := "6C646AD3AF383B55A07B659E26F741CC"

	imsg := &inventoryMsg{}
	//pmsg := &PurchaseMsg{}
	//smsg := &StockMsg{}

	iurl := "http://114.67.231.99/open/api/v1/agent/reduce/invetory/query"
	//surl := "http://114.67.231.99/open/api/v1/agent/reduce/stock/query"
	//purl := "http://114.67.231.99/open/api/v1/agent/reduce/purchase/query"

	go GetDataLoop(appkey, secret, iurl, imsg)
	//go GetDataLoop(appkey, secret, surl, smsg)
	//go GetDataLoop(appkey, secret, purl, pmsg)

	time.Sleep(time.Second * 10000)
}

/*

get data now

while 「
	get data check database
		if find it  drop it
		if not find it save in database send kfk

	if（now_unixno % 60 == 0）
		check database  1st-20th  timeout del it
」

*/
