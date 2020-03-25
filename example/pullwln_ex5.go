package main

import (
	"fmt"
	"github.com/xenbo/http_post/log"
	wln "github.com/xenbo/http_post/wlnv1"
	"github.com/xenbo/http_post/wlnv1/base"
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
	//	url = "http://114.67.231.162/api/erp/base/supplier/modify"
	//	//_app=3123415742&_s=&_sign=fa12a26f1c131e6abc5b02f3ecdccd56&_t=1574235347&
	//	// supplier={"supplier_code":"code_1","supplier_name":"name",
	//	// "status":1,"province":"bnm","city":"sdaf",
	//	// "area":"123_1","mobile":"13456789098","address":"111111111111111111",
	//	// "remark":"remark","contact":"basvsda"}
	//
	//	Add := base.NewSupplierModify("111111111111111111", "123_1", "sdaf", "basvsda", "13456789098", "bnm",
	//		"remark", 1, "code_1", "name")
	//	log.DLog.Println(string(Add))
	//
	//	rBody := wln.MakeSign(bSystem, Add, secret)
	//	hBody := wln.PostData(url, rBody)
	//
	//	fmt.Println(url)
	//	fmt.Println(hBody)
	//}

	//{
	//	url = "http://114.67.231.162/api/erp/b2b/distr/query"
	//	//_app=3123415742&_s=&_sign=d376929e08beb62124dd4110a8f9c5e9&_t=1558952036
	//
	//	Add := base.NewB2bDistrQuery("", "")
	//	log.DLog.Println(string(Add))
	//
	//	rBody := wln.MakeSign(bSystem, Add, secret)
	//	hBody := wln.PostData(url, rBody)
	//
	//	fmt.Println(url)
	//	fmt.Println(hBody)
	//}
	//
	//
	//{
	//	url = "http://114.67.231.162/api/erp/base/storage/query"
	//	//_app=3123415742&_s=&_sign=4a5e4dd9baad910057af03914f2164fa&_t=1574236561&page_no=1&page_size=10
	//
	//	Add := base.NewStorageQuery(1, 10)
	//	log.DLog.Println(string(Add))
	//
	//	rBody := wln.MakeSign(bSystem, Add, secret)
	//	hBody := wln.PostData(url, rBody)
	//
	//	fmt.Println(url)
	//	fmt.Println(hBody)
	//}

	//{
	//	url = "http://114.67.231.162/api/erp/batch/billbatch"
	//	// _app=3123415742&_s=&_sign=a3c4a55313b8904f778aff5598b66446&_t=1578553959&code=CG202001090001&type=1
	//
	//	Add := base.NewBatchBillBatch("CG202001090001", 1)
	//	log.DLog.Println(string(Add))
	//
	//	rBody := wln.MakeSign(bSystem, Add, secret)
	//	hBody := wln.PostData(url, rBody)
	//
	//	fmt.Println(url)
	//	fmt.Println(hBody)
	//}

	//{
	//	url = "http://114.67.231.162/api/erp/goods/update/item"
	//	// _app=3123415742&_s=&_sign=aeafc6a18fc9f3c5eefcc200bfaceaf3&_t=1574651808&
	//
	//	// item={"item_code":"openApiTestItem3","article_number":"123","item_name":"openApiTestItem3","unit":"个",
	//	// "categroy":"分类4","brand":"004","sale_price":3.3,"prime_price":3.2,"weight":2.3,"volume":2.2,
	//	// "length":2,"width":4,"height":1,"bar_code":"123123","item_pic":"asdasd",
	//
	//	// "skus":[{"spec_code":"skuCode3","spec_value1":"123123123121312321","spec_value2":"22","sale_price":3.3,
	//	// "prime_price":3.2,"weight":3.3,"volume":2.2,"spec_length":3,"spec_width":3,"spec_height":1,
	//	// "bar_code":"skubarcode3","spec_pic":"3123"}]}
	//
	//	var Vec1 []base.GoodsUpdateItemSkus
	//	Vec1 = append(Vec1,base.NewGoodsUpdateItemSkus("skubarcode3", 3.2, 3.3, "skuCode3",
	//		1, 3, "3123","123123123121312321","22",3,2.2,3.3))
	//
	//	Add := base.NewGoodsUpdate("123", "123123", "004", "分类4", 1,
	//		"openApiTestItem3", "openApiTestItem3",
	//		"asdasd", 2, 3.5, 3.3, Vec1, "个", 2.2, 2.3, 4)
	//	log.DLog.Println(string(Add))
	//
	//	rBody := wln.MakeSign(bSystem, Add, secret)
	//	hBody := wln.PostData(url, rBody)
	//
	//	fmt.Println(url)
	//	fmt.Println(hBody)
	//}

	//{
	//	url = "http://114.67.231.162/api/erp/goods/spec/open/query/goodswithspeclist"
	//	// _app=3123415742&_s=&_sign=36b47a1943457c650998d8683e10187b&_t=1571651968&
	//	// limit=10&page=1&spec_code=ltskucode
	//
	//	Add := base.NewQueryGoodsWithSpecList("ltskucode","","","",1,10)
	//	log.DLog.Println(string(Add))
	//
	//	rBody := wln.MakeSign(bSystem, Add, secret)
	//	hBody := wln.PostData(url, rBody)
	//
	//	fmt.Println(url)
	//	fmt.Println(hBody)
	//
	//}


	//{
	//	url = "http://114.67.231.162/api/erp/goods/add/item"
	//	// _app=3123415742&_s=&_sign=aeafc6a18fc9f3c5eefcc200bfaceaf3&_t=1574651808&
	//
	//	// item={"item_code":"openApiTestItem3","article_number":"123","item_name":"openApiTestItem3","unit":"个",
	//	// "categroy":"分类4","brand":"004","sale_price":3.3,"prime_price":3.2,"weight":2.3,"volume":2.2,
	//	// "length":2,"width":4,"height":1,"bar_code":"123123","item_pic":"asdasd",
	//
	//	// "skus":[{"spec_code":"skuCode3","spec_value1":"123123123121312321","spec_value2":"22","sale_price":3.3,
	//	// "prime_price":3.2,"weight":3.3,"volume":2.2,"spec_length":3,"spec_width":3,"spec_height":1,
	//	// "bar_code":"skubarcode3","spec_pic":"3123"}]}
	//
	//	var Vec1 []base.GoodsAddItemSkus
	//	Vec1 = append(Vec1, base.NewGoodsAddItemSkus("skubarcodeda3", 3.2, 3.3, "skuCode3",
	//		4, 3, "3123", "1231da23123121312321", "22", 3, 2.2, 3.3))
	//
	//	Add := base.NewGoodsAdd("1d23", "123d123", "004", "分类4", 1,
	//		"openApdiTestItem3", "openApiTestItem3",
	//		"asdadsd", 2, 3.5, 3.3, Vec1, "个", 2.2, 2.3, 4)
	//	log.DLog.Println(string(Add))
	//
	//	rBody := wln.MakeSign(bSystem, Add, secret)
	//	hBody := wln.PostData(url, rBody)
	//
	//	fmt.Println(url)
	//	fmt.Println(hBody)
	//
	//}

	{
		//url = "http://114.67.231.162/api/erp/base/supplier/add"
		////_app=3123415742&_s=&_sign=fa12a26f1c131e6abc5b02f3ecdccd56&_t=1574235347&
		//// supplier={"supplier_code":"code_1","supplier_name":"name",
		//// "status":1,"province":"bnm","city":"sdaf",
		//// "area":"123_1","mobile":"13456789098","address":"111111111111111111",
		//// "remark":"remark","contact":"basvsda"}
		//
		//Add := base.NewSupplierAdd("2111111111111111111", "1231_1", "1sdaf", "basvsda", "13456789098", "bnm",
		//	"remark", 1, "code_12", "name")
		//log.DLog.Println(string(Add))
		//
		//rBody := wln.MakeSign(bSystem, Add, secret)
		//hBody := wln.PostData(url, rBody)
		//
		//fmt.Println(url)
		//fmt.Println(hBody)
	}


	{
		url = "http://114.67.231.162/api/erp/base/supplier/query"
		//_app=3123415742&_s=&_sign=a197bd77dc142d3c89dcd6242f14ff9d&_t=1573180267&
		// page_no=1&page_size=100

		Add := base.NewSupplierQuery("", 1, 10)
		log.DLog.Println(string(Add))

		rBody := wln.MakeSign(bSystem, Add, secret)
		hBody := wln.PostData(url, rBody)

		fmt.Println(url)
		fmt.Println(hBody)

	}



	//	"","","","","","","","","","","","","","","","","","","","","","","","","","","","",
}
