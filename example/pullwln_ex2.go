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
	url := "http://114.67.231.162/api/erp/purchase/purchasereturnbill/close"

	tm := time.Unix(time.Now().Unix()-2*28*24*3600, 0).Format("2006-01-02 03:04:05")
	tm = fmt.Sprint(time.Now().Unix())
	tm = "1578479262"

	bSystem := wln.NewSystem(appkey, tm, "", "", "")
	log.DLog.Println(string(bSystem))

	bBusiness := purchase.NewClose("CTD202001080001", "", true)
	log.DLog.Println(string(bBusiness))

	rBody := wln.MakeSign(bSystem, bBusiness, secret)
	hBody := wln.PostData(url, rBody)

	fmt.Println(hBody)
}
