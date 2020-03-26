package main


import "time"
import "github.com/xenbo/http_post/test/wln"

func main() {
	wln.Glc.SetKafkaAddr("192.168.1.146")
	wln.Glc.Init()

	wln.Kc.NewProducer()

	appkey := "QC20201112"
	secret := "6C646AD3AF383B55A07Bgo659E26F741CC"

	imsg := wln.NewInventoryMsg0()
	smsg := wln.NewStockMsg()
	pmsg := wln.NewPurchaseMsg()

	iurl := "http://114.67.231.99/open/api/v1/agent/reduce/invetory/query"
	////////http://103.235.242.21/open/api/v1/inventories/erp
	surl := "http://114.67.231.99/open/api/v1/agent/reduce/stock/query"
	purl := "http://114.67.231.99/open/api/v1/agent/reduce/purchase/query"



	go wln.GetDataLoop(appkey, secret, iurl, &imsg)
	go wln.GetDataLoop(appkey, secret, surl, &smsg)
	go wln.GetDataLoop(appkey, secret, purl, &pmsg)

	time.Sleep(time.Second * 10000)
}

