package main

import (
	"fmt"
	"github.com/xenbo/http_post/log"
	wln "github.com/xenbo/http_post/wlnv1"
	"github.com/xenbo/http_post/wlnv1/stock"
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

	//{
	//	url = "http://114.67.231.99/api/erp/stock/out/requestbill/add"
	//
	//	//_app=3123415742&_s=&_sign=88e2ffe8c472b80ce1e6d712f8c69ff0&_t=1577084053&bill={"storage_code":"00003","details":[{"spec_code":"00001411","size":1.0}]}
	//
	//	var RequestBillAddBillDetailsVec []stock.RequestBillAddBillDetails
	//	RequestBillAddBillDetailsVec = append(RequestBillAddBillDetailsVec, stock.NewRequestBillAddBillDetails(-1, "", 1.0, "00001411"))
	//
	//	Add := stock.NewRequestBillAdd("", "123123", "001", "aa","" ,
	//		"","","","","","00003",RequestBillAddBillDetailsVec)
	//	log.DLog.Println(string(Add))
	//
	//	rBody := wln.MakeSign(bSystem, Add, secret)
	//	hBody := wln.PostData(url, rBody)
	//	fmt.Println(hBody)
	//}
	//
	//{
	//	url = "http://114.67.231.99/api/erp/sn/querytrace"
	//	//_app=3123415742&_s=&_sign=106379b0a270da021ae11782acb1ce95&_t=1558951138&
	//	// end=1552300487000&limit=200&page=1&start=1552217687000
	//
	//	Add := stock.NewSNQueryTrace("","1552217687000","1552300487000",1,10)
	//	log.DLog.Println(string(Add))
	//
	//	rBody := wln.MakeSign(bSystem, Add, secret)
	//	hBody := wln.PostData(url, rBody)
	//	fmt.Println(hBody)
	//}
	//
	//{
	//	//_app=3123415742&_s=&_sign=5b7064bff0374c6c1c28a47f2f1e03c7&_t=1577084005&
	//	// bill={"storage_code":"00003","details":[{"spec_code":"00001411","size":1.0}]}
	//
	//	url = "http://114.67.231.99/api/erp/stock/in/requestbill/add"
	//
	//	var InRequestBillAddBillDetailsVec []stock.InRequestBillAddBillDetails
	//	InRequestBillAddBillDetailsVec = append(InRequestBillAddBillDetailsVec,
	//		stock.NewInRequestBillAddBillDetails(-1,"",1.0,"00001411"))
	//
	//	Add := stock.NewInRequestBillAdd("","","","","","","",
	//		"","","","00003",InRequestBillAddBillDetailsVec)
	//	log.DLog.Println(string(Add))
	//
	//	rBody := wln.MakeSign(bSystem, Add, secret)
	//	hBody := wln.PostData(url, rBody)
	//	fmt.Println(hBody)
	//}

	//{
	//	//_app=3123415742&_s=&_sign=a54a7389f07a8cfd416d6bf08feda82f&_t=1570674550&
	//	// bill={"bill_type":1, "storage_in_code":"00003","storage_out_code":"002","remark": "12312321321",
	//	// "details":[{"spec_code":"00001411","size":10.0,"remark ":"111111111"},
	//	// {"spec_code":"0000200102","size":10.0,"remark": "111111111"}]}
	//	//
	//	url = "http://114.67.231.99/api/erp/allocation/changebill/add"
	//
	//	var Vec []stock.CHangeBillAddBillDetails
	//	Vec = append(Vec, stock.NewCHangeBillAddBillDetails(-1, "111111111", 10.0, "0000200102"))
	//	Add := stock.NewCHangeBillAdd(1, "", "12312321321", "00003", "002", Vec)
	//
	//	log.DLog.Println(string(Add))
	//
	//	rBody := wln.MakeSign(bSystem, Add, secret)
	//	hBody := wln.PostData(url, rBody)
	//	fmt.Println(hBody)
	//}

	//{
	//	//_app=3123415742&_s=&_sign=c98f24a487d784a8eea0faeb2ff4d541&_t=1578542302&
	//	// limit=2&modify_time=1525449600000&page=1
	//
	//	url = "http://114.67.231.99/api/erp/stock/in/requestbill/query"
	//	Add := stock.NewInRequestBillQuery("", "1525449600000", 1, 2)
	//
	//	log.DLog.Println(string(Add))
	//
	//	rBody := wln.MakeSign(bSystem, Add, secret)
	//	hBody := wln.PostData(url, rBody)
	//	fmt.Println(hBody)
	//}

	//{
	//	//_app=3123415742&_s=&_sign=37846dc5eb28095d467c7f8101a9429b&_t=1578471871
	//	//&bill_code=12345678
	//
	//	url = "http://114.67.231.99/api/erp/allocation/changebill/close"
	//	Add := stock.NewChangeBillClose("12345678")
	//
	//	log.DLog.Println(string(Add))
	//
	//	rBody := wln.MakeSign(bSystem, Add, secret)
	//	hBody := wln.PostData(url, rBody)
	//	fmt.Println(hBody)
	//}
	//
	//{
	//	//_app=3123415742&_s=&_sign=cdca301b131e5bd1a75096adfb4b986f&_t=1558952036&
	//	// sku_invs=[{"sku_no":"FJXM063-0460B ","amount":20.0,"storage_code":"001"}]
	//
	//	url = "http://114.67.231.99/api/erp/open/inventory/syn"
	//	var Vec1 []stock.InventorySynSkuInvsBatchInvs
	//	Vec1 = append(Vec1, stock.NewInventorySynSkuInvsBatchInvs("", "", -1, ""))
	//
	//	var Vec []stock.InventorySynSkuInvs
	//	Vec = append(Vec, stock.NewInventorySynSkuInvs(20.0, Vec1, "FJXM063-0460B", "001"))
	//
	//	Add := stock.NewInventorySyn(Vec)
	//
	//	log.DLog.Println(string(Add))
	//
	//	rBody := wln.MakeSign(bSystem, Add, secret)
	//	hBody := wln.PostData(url, rBody)
	//	fmt.Println(hBody)
	//}
	//
	//{
	//	//_app=3423466116&_s=&_sign=ac15dc6b1b1366303cf1a3040ec783e2&_t=1578541623&
	//	// bill_code=XD190716000016&bill_type=1
	//
	//	url = "http://114.67.231.99/api/erp/sn/querysnbybillcode"
	//	Add := stock.NewQuerySnbyBillCode("XD190716000016",1)
	//	log.DLog.Println(string(Add))
	//
	//	rBody := wln.MakeSign(bSystem, Add, secret)
	//	hBody := wln.PostData(url, rBody)
	//	fmt.Println(hBody)
	//}

	//{
	//	//bill={"reason":"其他","storage_code":"A01",
	//	// "remark":"#FHTZD001889#自提1、本单金额447.5*97=43407.5USD;\r\n2、07.09客户付款8950USD;\r\n3、07.25客户付款34457.50USD",
	//	// "details":[{"spec_code":"0000100102","size":1.0,"sum_money":100.0}]}
	//
	//	url = "http://114.67.231.99/api/erp/stock/in/stockbill/add"
	//	var vec []stock.StockBillAddBillDetails
	//	vec = append(vec, stock.NewStockBillAddBillDetails("", 1.0, "0000100102", 100.0))
	//
	//	Add := stock.NewStockBillAdd("其他", `XXXXX`,	"A01", vec)
	//	log.DLog.Println(string(Add))
	//
	//	rBody := wln.MakeSign(bSystem, Add, secret)
	//	hBody := wln.PostData(url, rBody)
	//	fmt.Println(hBody)
	//	//remark 不能太长
	//}

	//{
	//	//_app=3123415742&_s=&_sign=dbaa38ef02d7f0cc45a50f3fd109891d&_t=1578471870&limit=100
	//	//&modify_time=1578385013000&page=1
	//
	//	url = "http://114.67.231.99/api/erp/allocation/changebill/query"
	//	Add := stock.NewChangeBillQuery("", "1578385013000", 1, 10)
	//	log.DLog.Println(string(Add))
	//
	//	rBody := wln.MakeSign(bSystem, Add, secret)
	//	hBody := wln.PostData(url, rBody)
	//	fmt.Println(hBody)
	//}

	//{
	//	//_app=3123415742&_s=&_sign=dbaa38ef02d7f0cc45a50f3fd109891d&_t=1578471870&limit=100
	//	//&modify_time=1578385013000&page=1
	//
	//	Add := stock.NewInChangeBillQuery("","","",1,10)
	//	log.DLog.Println(string(Add))
	//
	//	rBody := wln.MakeSign(bSystem, Add, secret)
	//	hBody := wln.PostData(url, rBody)
	//	fmt.Println(hBody)
	//
	//	//{"code":1005,"message":"缺少参数 [bill]"}
	//}

	{
		//_app=3123415742&_s=&_sign=77f2cdd4ee47f096268d3e1462aa5093&_t=1558952035&limit=100&modify_time=1525449600&page=1

		url = "http://114.67.231.99/api/erp/stock/in/stockbill/query"
		Add := stock.NewInStockBillQuery("", "", "1525449600", 1, 10)
		log.DLog.Println(string(Add))

		rBody := wln.MakeSign(bSystem, Add, secret)
		hBody := wln.PostData(url, rBody)
		fmt.Println(hBody)

	}

	//"","","","","","","","","","","","","","","","","","","","","","","","","","","","",

}
