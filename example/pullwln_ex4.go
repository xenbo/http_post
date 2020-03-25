package main

import (
	"fmt"
	"github.com/xenbo/http_post/log"
	wln "github.com/xenbo/http_post/wlnv1"
	"github.com/xenbo/http_post/wlnv1/sale"
	"time"
)

func main() {
	log.CreateLog()

	appkey := "3123415742"
	secret := "c3b5fee170b52b8397852c8ba03ef109"
	url := "http://114.67.231.99/api/erp/purchase/purchasereturnbill/close"

	tm := time.Unix(time.Now().Unix()-2*28*24*3600, 0).Format("2006-01-02 03:04:05")
	tm = fmt.Sprint(time.Now().Unix())
	//tm = "1578479262"

	bSystem := wln.NewSystem(appkey, tm, "", "", "")
	log.DLog.Println(string(bSystem))

	{
		url = "http://114.67.231.162/api/erp/open/return/order/query"
		//_app=3123415742&_s=&_sign=adab07417c15ece3167cf2a4a4c3b9b0&_t=1558951263&
		// bill_code=SH201905140005&limit=20&page=1

		Add := sale.NewReturnOrderQuery("SH201905140005", "", "", 1, "", 1, 20, "")
		log.DLog.Println(string(Add))

		rBody := wln.MakeSign(bSystem, Add, secret)
		hBody := wln.PostData(url, rBody)

		fmt.Println(url)
		fmt.Println(hBody)
	}

	//"","","","","","","","","","","","","","","","","","","","","","","",

	//{
	//	url = "http://114.67.231.162/api/erp/sale/stock/in/add"
	//	//_app=3123415742&_s=&_sign=f8841c27e8216bbbe3b4913930f1bd97&_t=1578475359&
	//	// bill={"customer_nick":"123","bill_date":1578475359511,"create_time":1578475359511,"post_fee":1.0,"paid_fee":112.0,"discount_fee":0.0,
	//	// "service_fee":2.0,"shop_nick":"bbgillian",
	//	// "remark":"#FHTZD001889#自提1、本单金额447.5*97=43407.5USD;\r\n2、07.09客户付款8950USD;\r\n3、07.25客户付款34457.50USD",
	//	// "storage_code":"00003","bill_code":"33333333333",
	//	// "details":[{"sku_no":"00001411","nums":2,"sum_sale":43.0}]}
	//
	//	var Vec []sale.StockInAddBillDetails
	//	Vec = append(Vec, sale.NewStockInAddBillDetails(2, "00001411", 43.0))
	//	Add := sale.NewStockInAdd("1578475359511", "1578475359511", "123", Vec, 0.0, 112.0, 1.0,
	//		`#FHTZD001889#自提1、\r\n2、USD`,
	//		2.0, "bbgillian", "00003", -1.0, )
	//	log.DLog.Println(string(Add))
	//
	//	rBody := wln.MakeSign(bSystem, Add, secret)
	//	hBody := wln.PostData(url, rBody)
	//
	//	fmt.Println(url)
	//	fmt.Println(hBody)
	//}

	//{
	//	url = "http://114.67.231.99/api/erp/sale/stock/in/query"
	//	//_app=3123415742&_s=&_sign=0b61a959830600146ad6b88911526b72&_t=1558951263&
	//	// bill_code=XT190411000003&limit=200&page=1
	//
	//	Add := sale.NewStockInQuery("XT190411000003", "", 1, 200, true)
	//	log.DLog.Println(string(Add))
	//
	//	rBody := wln.MakeSign(bSystem, Add, secret)
	//	hBody := wln.PostData(url, rBody)
	//
	//	fmt.Println(url)
	//	fmt.Println(hBody)
	//}
	{
		url = "http://114.67.231.162/api/erp/open/return/order/stock/in"
		//_app=3123415742&_s=&_sign=c3a09976cfe3740f74521de49a95a28a&_t=1558951263&
		// bill_code=SH201904280001 &items=[{"size":1,"sku_code":"001"}]&remark=123

		var Vec []sale.OrderStockInItems
		Vec = append(Vec, sale.NewOrderStockInItems(1, "001"))

		Add := sale.NewOrderStockIn("SH201904280001", "123", Vec)
		log.DLog.Println(string(Add))

		rBody := wln.MakeSign(bSystem, Add, secret)
		hBody := wln.PostData(url, rBody)

		fmt.Println(url)
		fmt.Println(hBody)
	}

}
