package open

import (
	"fmt"
	"github.com/xenbo/http_post/log"
	wln "github.com/xenbo/http_post/wlnv1"
	"github.com/xenbo/http_post/wlnv1/b2c"
	"time"
)

func main() {
	log.CreateLog()

	appkey := "QC20201112"
	secret := "6C646AD3AF383B55A07B659E26F741CC"

	url := "http://114.67.231.99/open/api/v1/agent/reduce/invetory/query"
	//surl := "http://114.67.231.99/open/api/v1/agent/reduce/stock/query"
	//purl := "http://114.67.231.99/open/api/v1/agent/reduce/purchase/query"
	//

	now := time.Now()
	var tm = now.UnixNano() / 1e6 //tm = "1578479262"

	bSystem := wln.NewSystemB2C(appkey, "", "json", tm)
	log.DLog.Println(string(bSystem))

	{
		tm := time.Unix(time.Now().Unix()-2*28*24*3600, 0)
		tm1 := tm.Format("2006-01-02 03:04:05")

		Add := b2c.NewStockQuery(tm1, 10, 20)
		log.DLog.Println(string(Add))

		rBody := wln.MakeB2CSign(bSystem, Add, secret)
		hBody := wln.PostData(url, rBody)

		fmt.Println(url)
		fmt.Println(hBody)
	}

	{
		url = "http://114.67.231.99/open/api/v1/agent/reduce/purchase/query"
		tm := time.Unix(time.Now().Unix()-2*28*24*3600, 0)
		tm1 := tm.Format("2006-01-02 03:04:05")

		Add := b2c.NewPurchaseQuery(tm1, 10, 20)
		log.DLog.Println(string(Add))

		rBody := wln.MakeB2CSign(bSystem, Add, secret)
		hBody := wln.PostData(url, rBody)

		fmt.Println(url)
		fmt.Println(hBody)
	}


	{
		url = "http://114.67.231.99/open/api/v1/agent/reduce/stock/query"
		tm := time.Unix(time.Now().Unix()-2*28*24*3600, 0)
		tm1 := tm.Format("2006-01-02 03:04:05")

		Add := b2c.NewReduceInventoryQueryIn(tm1, 10, 20)
		log.DLog.Println(string(Add))

		rBody := wln.MakeB2CSign(bSystem, Add, secret)
		hBody := wln.PostData(url, rBody)

		fmt.Println(url)
		fmt.Println(hBody)
	}

	{
		url = "http://114.67.231.99/open/api/v1/agent/reduce/stock/query"
		tm := time.Unix(time.Now().Unix()-2*28*24*3600, 0)
		tm1 := tm.Format("2006-01-02 03:04:05")

		Add := b2c.NewReduceInventoryQueryOut(tm1, 10, 20)
		log.DLog.Println(string(Add))

		rBody := wln.MakeB2CSign(bSystem, Add, secret)
		hBody := wln.PostData(url, rBody)

		fmt.Println(url)
		fmt.Println(hBody)
	}

	//"","","","","","","","","","","","","","","","","","","","",
}
