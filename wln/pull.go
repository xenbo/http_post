package wln

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/xenbo/go_kfk_client/rdkfk"
	"gitlab.com/eosforce/vbbirdworker/event"
	"io/ioutil"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

var Glc rdkfk.GlobeCleaner
var Kc rdkfk.KafkaClient

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

func (data *InventoryMsg) RequestData(pid int32, appkey string, secret string, url string) {

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
	Kc.AddProduceTopic("bullmsg")

	for pi := 0; pi < 1; pi++ {
		msg.RequestData(int32(pi), appkey, secret, url)
		ExtStr, _ := json.Marshal(msg)
		//fmt.Println(msg.name(), "   ", ExtStr)

		RespN := msg.RespN()

		for i := 0; i < RespN; i++ {
			DetailN := msg.ContainerN(i)
			for j := 0; j < DetailN; j++ {
				info0 := event.NodeInfo{ //fachu zhe
					NodeID:        msg.GetNodeID0(i),
					NodeRole:      msg.GetNodeRole0(i),
					NodeActorID:   "", //msg.NodeActorID0(i),
					NodeActorRole: msg.NodeActorRole0(i),
				}

				//info1 := event.NodeInfo{
				//	NodeID:        msg.GetNodeID1(i, j),
				//	NodeRole:      msg.GetNodeRole1(i, j),
				//	NodeActorID:   msg.NodeActorID1(i, j),
				//	NodeActorRole: msg.NodeActorRole1(i, j),
				//}

				info1 := event.NodeInfo{
					NodeID:        msg.GetNodeID1(i,j),
					NodeRole:      "",
					NodeActorID:   "",
					NodeActorRole: "",
				}

				evt, err := event.NewEvent(
					msg.GetType(), &info0, msg.GetContainerId(i, j), time.Now(),
					event.WithDatas(event.DataOutbound{
						ShipNum:     msg.GetExpressId(),
						Destination: info1, //consumer
					}),
					event.WithExtStrInfo(string(ExtStr)),
					event.WithSubTyp(int(msg.GetType())))

				if err == nil {
					emsg, _ := json.Marshal(evt)
					fmt.Println(string(emsg))
					Kc.SendMsgWithCache("bullmsg", string(emsg))
				}
			}
		}
	}
}
