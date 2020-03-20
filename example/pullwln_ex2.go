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
	//	bBusiness := purchase.NewClose("CTD202001080001", "", false)
	//	log.DLog.Println(string(bBusiness))
	//
	//	rBody := wln.MakeSign(bSystem, bBusiness, secret)
	//	hBody := wln.PostData(url, rBody)
	//	fmt.Println(hBody)
	//}
	//
	//{
	//	//_app=3123415742&bill={"supplier_code":"aa","storage_code":"001","remark":"123123","details":[{"
	//	//spec_code":"0000001a","size":1.0,"sum":0.0,"price":0.0}]}
	//
	//	Add := purchase.NewAdd("",-1, 0.0, "",1.0,"0000001a",0.0,"123123","001","aa")
	//	log.DLog.Println(string(Add))
	//
	//	rBody := wln.MakeSign(bSystem, Add, secret)
	//	hBody := wln.PostData(url, rBody)
	//	fmt.Println(hBody)
	//}


	{
		//_app=3123415742&_s=&_sign=4b89f9a9b4abe5896be841759065176c
		//&_t=1578477467&bill_code=CD202001060001&limit=100&page=1

		Add := purchase.NewBillQuery("CD202001060001","", 1, 100)
		log.DLog.Println(string(Add))

		rBody := wln.MakeSign(bSystem, Add, secret)
		hBody := wln.PostData(url, rBody)
		fmt.Println(hBody)
	}


}
