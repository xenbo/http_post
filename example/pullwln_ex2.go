package main

import (
	"fmt"
	"github.com/xenbo/http_post/log"
	wln "github.com/xenbo/http_post/wlnv1"
	"github.com/xenbo/http_post/wlnv1/purchase"
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
	//	url := "http://114.67.231.99/api/erp/purchase/purchasereturnbill/close"
	//
	//	bBusiness := purchase.NewReturnBillClose("CTD202001080001", "", false)
	//	log.DLog.Println(string(bBusiness))
	//
	//	rBody := wln.MakeSign(bSystem, bBusiness, secret)
	//	hBody := wln.PostData(url, rBody)
	//	fmt.Println(hBody)
	//}

	//{
	//	//_app=3123415742&_s=&_sign=ee7c702721711900f74529bb8179d207&_t=1558951263&
	//	// bill={"supplier_code":"aa","storage_code":"001","remark":"123123",
	//	// "details":[{" spec_code":"0000001a","size":1.0,"sum":0.0,"price":0.0}]}
	//
	//	var BillAddBillDetailVec []purchase.BillAddBillDetails
	//	BillAddBillDetailVec = append(BillAddBillDetailVec, purchase.NewBillAddBillDetail(-1, 0.0, "", 1.0, "0000001a", 0.0))
	//
	//	Add := purchase.NewBillAdd("CECGTHDH04001", "123123", "001", "aa", BillAddBillDetailVec)
	//	log.DLog.Println(string(Add))
	//
	//	rBody := wln.MakeSign(bSystem, Add, secret)
	//	hBody := wln.PostData(url, rBody)
	//	fmt.Println(hBody)
	//}


	{
		//_app=3123415742&_s=&_sign=a638208efb74fd655bdb60cf7e69d5c0
		//&_t=1578479597&
		// bill={"bill_code":"CECGDH00401", "supplier_code":"0002","storage_code":"011 ","remark":"123123",
		// "details":[{ "spec_code":"00001411","size":1.0,"sum ":0.0,"price":0.0,"index":123123}]}

		url = "http://114.67.231.99/api/erp/purchase/purchasereturnbill/add"

		var ReturnBillAddDetailsVec []purchase.ReturnBillAddDetails
		ReturnBillAddDetailsVec = append(ReturnBillAddDetailsVec, purchase.NewReturnBillAddDetails(123123, 0.0, "", 1.0, "00001411", 0.0))

		Add := purchase.NewReturnBillAdd("CECGTHDH04001", "123123", "011", "0002", ReturnBillAddDetailsVec)
		log.DLog.Println(string(Add))

		rBody := wln.MakeSign(bSystem, Add, secret)
		hBody := wln.PostData(url, rBody)
		fmt.Println(hBody)

	}

	//{
	//	//_app=3123415742&_s=&_sign=4b89f9a9b4abe5896be841759065176c
	//	//&_t=1578477467&bill_code=CD202001060001&limit=100&page=1
	//
	//	Add := purchase.NewBillQuery("CD202001060001","", 1, 100)
	//	log.DLog.Println(string(Add))
	//
	//	rBody := wln.MakeSign(bSystem, Add, secret)
	//	hBody := wln.PostData(url, rBody)
	//	fmt.Println(hBody)
	//}

	//{
	//	//_app=3123415742&_s=&_sign=0c2f478cb1c3fdf1aee044d7a62d8f97&_t=1558951963&limit=100&modify_time=1525449600&page=1
	//	url = "http://114.67.231.99/api/erp/purchase/purchasebill/stockin/query"
	//	tm = fmt.Sprint(time.Now().Unix())
	//
	//	Add := purchase.NewStockInQuery("","",tm, 1, 10)
	//	log.DLog.Println(string(Add))
	//
	//	rBody := wln.MakeSign(bSystem, Add, secret)
	//	hBody := wln.PostData(url, rBody)
	//	fmt.Println(hBody)
	//}

	//{
	//	//_app=3123415742&_s=&_sign=c8ec471a36f27f9298a91d27472465ca
	//	//&_t=1578479183&bill_code=CTD202001080001&limit=100&page=1
	//
	//	url = "http://114.67.231.99/api/erp/purchase/purchasereturnbill/query"
	//	tm = fmt.Sprint(time.Now().Unix())
	//
	//	Add := purchase.NewTurnBillQuery("CTD202001080001", tm, 1, 10)
	//	log.DLog.Println(string(Add))
	//
	//	rBody := wln.MakeSign(bSystem, Add, secret)
	//	hBody := wln.PostData(url, rBody)
	//	fmt.Println(hBody)
	//}

	//{
	//	//{
	//	//"remark":"123123",
	//	//"supplier_code":"0002",
	//	//"storage_code":"011",
	//	//"bill_code":"CECGTHDH04003",
	//	//"details":[
	//	//{
	//	//"spec_code":"00001411 ",
	//	//"nums":1.0,
	//	//"remark":"",
	//	//"index":1
	//	//},
	//	//]
	//	//}
	//
	//	url = "http://114.67.231.99/api/erp/purchase/purchasestockbill/add"
	//	tm = fmt.Sprint(time.Now().Unix())
	//
	//	var StockInAddDetailsVec []purchase.StockInAddDetails
	//	StockInAddDetailsVec = append(StockInAddDetailsVec, purchase.NewStockInAddDetails(1,1.0,"","","00001411",-1.0,-1.0,-1.0))
	//	Add := purchase.NewStockInAdd("", -1.0, -1.0, "CECGTHDH04003", "123123", "", "011", "0002", StockInAddDetailsVec)
	//	log.DLog.Println(string(Add))
	//
	//	rBody := wln.MakeSign(bSystem, Add, secret)
	//	hBody := wln.PostData(url, rBody)
	//	fmt.Println(hBody)
	//}

	//{
	//	//_app=3123415742&_s=&_sign=d248d3b42b7758a6c1ceb8bc537d53c0&_t=1558951964&limit=100&modify_time=1554825600000&page=1
	//	url = "http://114.67.231.99/api/erp/purchase/purchasebill/stockout/query"
	//	tm = fmt.Sprint(time.Now().Unix())
	//
	//	Add := purchase.NewStockOutQuery("", "", tm, 1, 10)
	//	log.DLog.Println(string(Add))
	//
	//	rBody := wln.MakeSign(bSystem, Add, secret)
	//	hBody := wln.PostData(url, rBody)
	//	fmt.Println(hBody)
	//}

	//{
	//	//_app=3123415742&_s=&_sign=8c56b58d18db19a89d257d3b7793594a
	//	//&_t=1578474956&
	//	// bill={"remark":"123123"," supplier_code":"0002","storage_code":"011 ","bill_code":"CECGTHDH04002",
	//	//
	//	//
	//	// "details
	//	//":[{"spec_code":"00001411","nums ":1.0,"remark":"","index":1}
	//	//,{"spec_code":"00001411","nums":1.0 ,"remark":"","index":12341234}]}
	//
	//	url = "http://114.67.231.99/api/erp/purchase/purchasestockbackproductbill/add"
	//	tm = fmt.Sprint(time.Now().Unix())
	//
	//	var StockBackProductBillAddBillDetailsVec []purchase.StockBackProductBillAddBillDetails
	//	StockBackProductBillAddBillDetailsVec = append(StockBackProductBillAddBillDetailsVec,
	//		purchase.NewStockBackProductBillAddBillDetails(1, 1.0, "", "", "00001411", -1.0, -1.0, -1.0))
	//
	//	StockBackProductBillAddBillDetailsVec = append(StockBackProductBillAddBillDetailsVec,
	//		purchase.NewStockBackProductBillAddBillDetails(12341234, 1.0, "", "", "00001411", -1.0, -1.0, -1.0))
	//
	//	Add := purchase.NewStockBackProductBillAdd("", -1.0, -1.0, "CECGTHDH04002", "123123", "", "011", "0002", StockBackProductBillAddBillDetailsVec)
	//	log.DLog.Println(string(Add))
	//
	//	rBody := wln.MakeSign(bSystem, Add, secret)
	//	hBody := wln.PostData(url, rBody)
	//	fmt.Println(hBody)
	//}

	//{
	//	//_app=3123415742&_s=&_sign=4a608572ea88119f7120e3b498e8326d &_t=1578478903&bill_code=CD202001080002
	//	url = "http://114.67.231.99/api/erp/purchase/purchasebill/close"
	//	tm = fmt.Sprint(time.Now().Unix())
	//
	//	Add := purchase.NewBillClose("CD202001080002", "", false)
	//	log.DLog.Println(string(Add))
	//
	//	rBody := wln.MakeSign(bSystem, Add, secret)
	//	hBody := wln.PostData(url, rBody)
	//	fmt.Println(hBody)
	//}

}
