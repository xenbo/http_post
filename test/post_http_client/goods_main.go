package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xenbo/go_kfk_client/rdkfk"
	"github.com/xenbo/http_post/postwln"
	"gitlab.com/eosforce/vbbirdworker/event"
	"time"
)

type System struct {
	App string `json:"_app"`
	T   int64  `json:"_t"`
	//SignKind string `json:"_sign_kind"`
	Sign string `json:"_sign"`
	S    string `json:"_s"`
}

func (sys *System) SetNewTime() {
	sys.T = time.Now().Unix()
}

type Args struct {
	//<arg name="spec_code" title="规格编码" type="String"/>
	//<arg name="item_code" title="商品编码" type="String"/>
	//<arg name="bar_code" title="条码" type="String"/>

	//Key string `json:"keyword"`
	Code      string `json:"item_code"`
	Page_no   int32  `json:"page"`  //1
	Page_size int32  `json:"limit"` //20
}

type GoodsInfo struct {
	ArticleNumber string  `json:"article_number"`
	BarCode       string  `json:"bar_code"`
	Brand         string  `json:"brand"`
	Catagory      string  `json:"catagory"`
	Color         string  `json:"color"`
	ItemName      string  `json:"item_name"`
	ItemCode      string  `json:"item_code"`
	OtherProp     string  `json:"pther_prop"`
	Price         float32 `json:"price"`
	SpecCode      string  `json:"spec_code"`
	UInit         string  `json:"uinit"`
}

type HttpData struct {
	Code int32       `json:"code"`
	Data []GoodsInfo `json:"data"`
}

func GetRetMsg(httpGoodsInfo *GoodsInfo) []byte {

	goodsInfo := event.DataGoodspec{}
	goodsInfo.ArticleNumber = httpGoodsInfo.ArticleNumber
	goodsInfo.BarCode = httpGoodsInfo.BarCode
	goodsInfo.Brand = httpGoodsInfo.Brand
	goodsInfo.Catagory = httpGoodsInfo.Catagory
	goodsInfo.Color = httpGoodsInfo.Color
	goodsInfo.Item_code = httpGoodsInfo.ItemCode
	goodsInfo.OtherProp = httpGoodsInfo.OtherProp
	goodsInfo.Price = httpGoodsInfo.Price
	goodsInfo.Prop1 = "Prop1"
	goodsInfo.Prop2 = "Prop2"
	goodsInfo.Prop3 = "Prop3"
	goodsInfo.SpecCode = httpGoodsInfo.SpecCode
	goodsInfo.Unit = httpGoodsInfo.UInit

	jsonb, _ := json.Marshal(goodsInfo)

	event1 := event.Event{}
	event1.Header.Typ = event.TypeGoodsSpec
	event1.Datas = jsonb

	evb, _ := json.Marshal(event1)

	return evb
}

func httpPost(kc rdkfk.KafkaClient) {
	ss := &System{}
	ss.App = "3123415742"
	//1573181240
	//ss.SignKind = "md5"
	ss.Sign = ""
	ss.S = ""

	for i := 0; i < 10000; i++ {
		ag := &Args{}
		//ag.Code = "000029"
		//ag.Key = "围巾"
		ag.Page_no = int32(i)
		ag.Page_size = 1

		ss.T = time.Now().Unix()

		var msg = make([]byte, 0, 4096)
		cystr := "c3b5fee170b52b8397852c8ba03ef109"
		url := "http://114.67.231.162/api/erp/goods/spec/open/query"
		hmsg := postwln.HttpPost(cystr, ss, ag, url)

		//hmsg := []byte(`{"code":0,"data":[{"bar_code":"","catagory":"","brand":"","item_name":"△枸橼酸铋钾片(Y)","spec_code":"115011044","color":"","other_prop":"","price":14.7,"article_number":"115011044","unit":"","item_code":"115011044"}]}`)
		//fmt.Println(string(hmsg), cystr, url)

		var buffer bytes.Buffer
		buffer.Write(hmsg)
		buffer.Write(msg)
		hdata := &HttpData{}
		//hdata.data = make([]GoodsInfo, 0, 100)

		err := json.Unmarshal(buffer.Bytes(), hdata)
		if err == nil {
			for id := 0; id < len(hdata.Data); id++ {
				msg = GetRetMsg(&hdata.Data[id])
				fmt.Println(string(msg))

				kc.SendMsgWithCache("goods_info_test", string(msg))
				fmt.Println("-------------------------------------------------------------")
			}
		}
		time.Sleep(time.Second * 1)
	}
}

func main() {

	glc := rdkfk.GlobeCleaner{}
	glc.SetKafkaAddr("192.168.1.172")
	glc.Init()

	kc3 := rdkfk.KafkaClient{}
	kc3.NewProducer()
	kc3.AddProduceTopic("goods_info_test")

	httpPost(kc3)

}
