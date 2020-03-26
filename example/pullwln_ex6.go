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
	url := "http://114.67.231.162/api/erp/base/supplier/modify"

	tm := time.Unix(time.Now().Unix()-2*28*24*3600, 0).Format("2006-01-02 03:04:05")
	tm = fmt.Sprint(time.Now().Unix())
	//tm = "1578479262"

	bSystem := wln.NewSystem(appkey, tm, "", "", "")
	log.DLog.Println(string(bSystem))

	//{
	//	url = "http://114.67.231.162/api/erp/opentrade/query/trades"
	//	//_app=3123415742&_s=&_sign=e77277e2901194192cbf86ba3d345354 &_t=1578472121&
	//	// bill_code=XD191204000547&limit=100&page=1
	//
	//	Add := sale.NewQueryTrades("XD191204000547","","","","",
	//		"","","",1,10,false,"","")
	//	log.DLog.Println(string(Add))
	//
	//	rBody := wln.MakeSign(bSystem, Add, secret)
	//	hBody := wln.PostData(url, rBody)
	//
	//	fmt.Println(url)
	//	fmt.Println(hBody)
	//}

	//{
	//
	//	url = "http://114.67.231.162/api/erp/sale/stock/out/query"
	//	//_app=3123415742&_s=&_sign=312ae236063ff8771e64d03d09492e0c&_t=1558951412&
	//	// bill_code=XC190429000001&limit=1&page=1
	//
	//	Add := sale.NewStockOutQuery("XC190429000001", "", 1, 1, true)
	//	log.DLog.Println(string(Add))
	//
	//	rBody := wln.MakeSign(bSystem, Add, secret)
	//	hBody := wln.PostData(url, rBody)
	//
	//	fmt.Println(url)
	//	fmt.Println(hBody)
	//}

	//{
	//
	//	url = "http://114.67.231.162/api/erp/opentrade/modify/remark"
	//	//_app=3123415742&_s=&_sign=f8062e212f7eba191091d7bce8e92611&_t=1572576388&
	//	// bill_code=XD191025000015&remark=测试修改备注
	//
	//	var vec []sale.ModifyRemarkOrders
	//	vec = append(vec, sale.NewModifyRemarkOrders("",""))
	//
	//	Add := sale.NewModifyRemark("XD191025000015", vec, `测试修改备注xxxxxx`)
	//	log.DLog.Println(string(Add))
	//
	//	rBody := wln.MakeSign(bSystem, Add, secret)
	//	hBody := wln.PostData(url, rBody)
	//
	//	fmt.Println(url)
	//	fmt.Println(hBody)
	//}
	//
	//{
	//	url = "http://114.67.231.162/api/erp/opentrade/modify/mark"
	//	//_app=3123415742&_s=&_sign=f8062e212f7eba191091d7bce8e92611&_t=1572576388&
	//	// bill_code=XD191025000015&mark_name=绿的
	//
	//	Add := sale.NewModifyMark("XD191025000015", `绿的xxx`)
	//	log.DLog.Println(string(Add))
	//
	//	rBody := wln.MakeSign(bSystem, Add, secret)
	//	hBody := wln.PostData(url, rBody)
	//
	//	fmt.Println(url)
	//	fmt.Println(hBody)
	//}
	//

	//{
	//	url = "http://114.67.231.162/api/erp/opentrade/reply/exception/trades"
	//	//bill_codes=["XD190507000011"]
	//
	//	var vec []string
	//	vec = append(vec,"XD190507000011")
	//
	//	Add := sale.NewReplyExceptionTrades(vec, `绿的xxx`)
	//	log.DLog.Println(string(Add))
	//
	//	rBody := wln.MakeSign(bSystem, Add, secret)
	//	hBody := wln.PostData(url, rBody)
	//
	//	fmt.Println(url)
	//	fmt.Println(hBody)
	//}

	//{
	//	url = "http://114.67.231.162/api/erp/opentrade/send/trades"
	//	//&send_trade={"bill_code":"XD190516000018","part_delivered":false,
	//	// "packages":[{"erp_express_code":"0002","way_bill":"VE52421702535"}]}
	//
	//	var vecd []sale.SendTradesSendTradesPackagesDetails
	//	vecd = append(vecd, sale.NewSendTradesSendTradesPackagesDetails(-1, ""))
	//
	//	var vec []sale.SendTradesSendTradesPackages
	//	vec = append(vec, sale.NewSendTradesSendTradesPackages(
	//		vecd, "0002", "", -1.0, "VE52421702535", -1.0))
	//
	//	Add := sale.NewSendTrades("XD190516000018", vec, false, "")
	//	log.DLog.Println(string(Add))
	//
	//	rBody := wln.MakeSign(bSystem, Add, secret)
	//	hBody := wln.PostData(url, rBody)
	//
	//	fmt.Println(url)
	//	fmt.Println(hBody)
	//}

	//{
	//	url = "http://114.67.231.162/api/erp/sale/stock/out/add"
	//	//_app=3523466433&_s=&_sign=e4e8fa7522f305f22619e624e869346e&_t=1577325553&
	//	// bill={"customer_nick":"9999999","bill_date":1577325553982,"create_time":1577325553982,"sum_sale":316.81,"post_fee":0.0,
	//	// "paid_fee":316.81,"discount_fee":0.0,"service_fee":0.0,"shop_nick":"WMF电器自营旗舰店","remark":"sgaasg","storage_code":"0101",
	//
	//	// "details":[{"sku_no":"C030-0128","nums":2,"sum_sale":159.29},{"sku_no":"C030-0128","nums":2,"sum_sale":159.29}]}
	//
	//	var vec []sale.StockOutAddBillDetails
	//	vec = append(vec, sale.NewStockOutAddBillDetails(2, "C030-0128", 159.29))
	//
	//	Add := sale.NewStockOutAdd("1577325553982", "1577325553982", "9999999", vec, 0.0, 0.0,
	//		0.0, "sgaasg", 0.0, "WMF电器自营旗舰店", "0101", 316.81)
	//	log.DLog.Println(string(Add))
	//
	//	rBody := wln.MakeSign(bSystem, Add, secret)
	//	hBody := wln.PostData(url, rBody)
	//
	//	fmt.Println(url)
	//	fmt.Println(hBody)
	//}

	//{
	//	url = "http://114.67.231.162/api/erp/opentrade/reply/approve/trades"
	//	//_app=3123415742&_s=&_sign=d23eb91e00a2ce580e2eec1cdea06880&_t=1578472491&bill_code=XD190516000010
	//
	//	Add := sale.NewApproveTrades("XD190516000010")
	//	log.DLog.Println(string(Add))
	//
	//	rBody := wln.MakeSign(bSystem, Add, secret)
	//	hBody := wln.PostData(url, rBody)
	//
	//	fmt.Println(url)
	//	fmt.Println(hBody)
	//}

	//
	//	{ //没有开放的接口
	////
	//			url = "http://114.67.231.162/api/erp/opentrade/trade/commit"
	//
	//
	//			Add := sale.NewTradeCommit("XD190516000018",0,"","","","","","","","")
	//			log.DLog.Println(string(Add))
	//
	//			rBody := wln.MakeSign(bSystem, Add, secret)
	//			hBody := wln.PostData(url, rBody)
	//
	//			fmt.Println(url)
	//			fmt.Println(hBody)
	//
	//	}

	//{
	//	url = "http://114.67.231.162/api/erp/opentrade/query/mark"
	//
	//	var a []byte
	//
	//	rBody := wln.MakeSign(bSystem, a, secret)
	//	hBody := wln.PostData(url, rBody)
	//
	//	fmt.Println(url)
	//	fmt.Println(hBody)
	//	//_app=3123415742&_s=&_sign=6ff82caf71aa2b2b844bc0ef8f73f3ab&_t=1577325685
	//}

	{

		url = "http://114.67.231.162/api/erp/sale/stock/count/query"
		//_app=3123415742&_s=&_sign=d23eb91e00a2ce580e2eec1cdea06880&_t=1578472491&bill_code=XD190516000010

		Add := sale.NewStockCountQuery(1,"123123123123")
		log.DLog.Println(string(Add))

		rBody := wln.MakeSign(bSystem, Add, secret)
		hBody := wln.PostData(url, rBody)

		fmt.Println(url)
		fmt.Println(hBody)
	}
	//"","","","","","","","","","","","","","","","","","","","",
}
