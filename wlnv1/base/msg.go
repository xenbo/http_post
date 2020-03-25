package base

import (
	"encoding/json"
	"github.com/xenbo/http_post/log"
)

//erp/base/supplier/modify
//修改供应商，状态只能传0或1，其他状态返回异常

type SupplierModifySupplier struct {
	Address      string `json:"address"`       //地址
	Area         string `json:"area"`          //区
	City         string `json:"city"`          //市
	Contact      string `json:"contact"`       //联系人
	Mobile       string `json:"mobile"`        //手机
	Province     string `json:"province"`      //省
	Remark       string `json:"remark"`        //备注
	Status       int    `json:"status"`        //状态 1-启用，0-停用
	SupplierCode string `json:"supplier_code"` //供应商编码
	SupplierName string `json:"supplier_name"` //供应商名称
}
type SupplierModify struct {
	Supplier SupplierModifySupplier `json:"supplier"`
}

func NewSupplierModify(
	Address string,
	Area string,
	City string,
	Contact string,
	Mobile string,
	Province string,
	Remark string,
	Status int,
	SupplierCode string,
	SupplierName string,
) []byte {

	SModify = &SupplierModify{
		Supplier: SupplierModifySupplier{
			Address:      Address,
			Area:         Area,
			City:         City,
			Contact:      Contact,
			Mobile:       Mobile,
			Province:     Province,
			Remark:       Remark,
			Status:       Status,
			SupplierCode: SupplierCode,
			SupplierName: SupplierName,
		},
	}

	b1, err := json.Marshal(SModify)
	if err != nil {
		log.DLog.Println(err)
	}
	return b1
}

type RespSupplierModify struct {
	Code    int    `json:"code"`    //返回 0 表示执行成功 响应代码
	Data    bool   `json:"data"`    //响应结果数据
	Message string `json:"message"` //仅执行出错时返回 响应异常信息
}

//erp/b2b/distr/query
//获取多牛经销商,对外开放

type B2bDistrQuery struct {
	ShopNick        string `json:"shop_nick"`        //店铺昵称
	DistributorName string `json:"distributor_name"` //经销商名称
}

func NewB2bDistrQuery(
	ShopNick string,
	DistributorName string,
) []byte {
	BDistrQuery = &B2bDistrQuery{
		ShopNick:        ShopNick,
		DistributorName: DistributorName,
	}

	b1, err := json.Marshal(BDistrQuery)
	if err != nil {
		log.DLog.Println(err)
	}
	return b1
}

type RespB2bDistrQueryData struct {
	DistributorLevel string `json:"distributor_level"` //经销商级别
	DistributorName  string `json:"distributor_name"`  //经销商名称
	DistributorType  string `json:"distributor_type"`  //类型
	Province         string `json:"province"`          //省份
	Region           string `json:"region"`            //经销商所属区域
	RegionAgent      string `json:"region_agent"`      //区代
	SellType         int    `json:"sell_type"`         //销售类型， 1线上 2线下
}

type RespB2bDistrQuery struct {
	Code    int                     `json:"code"`    //返回 0 表示执行成功 响应代码
	Data    []RespB2bDistrQueryData `json:"data"`    //响应结果数据
	Message string                  `json:"message"` //仅执行出错时返回 响应异常信息
}

//erp/base/storage/query
//查询仓库信息
type StorageQuery struct {
	PageNo   int `json:"page_no"`   //分页号
	PageSize int `json:"page_size"` //分页大小 最大允许100
}

func NewStorageQuery(
	PageNo int,
	PageSize int,
) []byte {
	SQuery = &StorageQuery{
		PageNo:   PageNo,
		PageSize: PageSize,
	}

	b1, err := json.Marshal(SQuery)
	if err != nil {
		log.DLog.Println(err)
	}
	return b1
}

type RespStorageQueryData struct {
	Status      int    `json:"status"`       //状态 0禁用，1启用，2删除
	StorageCode string `json:"storage_code"` //仓库编码
	StorageName string `json:"storage_name"` //仓库名称
}

type RespStorageQuery struct {
	Code    int                    `json:"code"`    //返回 0 表示执行成功 响应代码
	Data    []RespStorageQueryData `json:"data"`    //响应结果数据
	Message string                 `json:"message"` //仅执行出错时返回 响应异常信息
}

//erp/batch/billbatch
//根据单号查询批次信息,对外开放
type BatchBillBatch struct {
	Code  string `json:"code"` //ERP单据编码
	Type1 int    `json:"type"` //单据类型:1、采购入库，2、采购退货，3、调拨入库，4、调拨出库，5、其它入库，6、其它出库，7、销售出库，8、售后退货，9、线下销售出库
}

func NewBatchBillBatch(
	Code string,
	Type1 int,
) []byte {
	BBillBatch = &BatchBillBatch{
		Code:  Code,
		Type1: Type1,
	}

	b1, err := json.Marshal(BBillBatch)
	if err != nil {
		log.DLog.Println(err)
	}
	return b1
}

type RespBatchBillBatchDataBatchs struct {
	BatchNo     string `json:"batch_no"`     //批次编码
	ExpiredDate string `json:"expired_date"` //过期日期
	Num         int    `json:"num"`          //数量
	ProduceDate string `json:"produce_date"` //成产日期
}
type RespBatchBillBatchData struct {
	Batchs  []RespBatchBillBatchDataBatchs `json:"batchs"`   //批次信息
	SkuCode string                         `json:"sku_code"` //商品编码
}

type RespBatchBillBatch struct {
	Code    int                      `json:"code"`    //返回 0 表示执行成功 响应代码
	Data    []RespBatchBillBatchData `json:"data"`    //响应结果数据
	Message string                   `json:"message"` //仅执行出错时返回 响应异常信息
}

//erp/goods/update/item
//修改商品信息

type GoodsUpdateItemSkus struct {
	BarCode    string  `json:"bar_code"`    //商品级条码——如果有规格，次条码会忽略，即使规格集中的条码没有传
	PrimePrice float64 `json:"prime_price"` //参考进价——如果有规格，会忽略，即使规格集中的没有传
	SalePrice  float64 `json:"sale_price"`  //标准售价——如果有规格，会忽略，即使规格集中的没有传
	SpecCode   string  `json:"spec_code"`   //规格编码
	SpecHeight int     `json:"spec_height"` //高——如果有规格，会忽略，即使规格集中的没有传
	SpecLength int     `json:"spec_length"` //长——如果有规格，会忽略，即使规格集中的没有传
	SpecPic    string  `json:"spec_pic"`    //规格图片
	SpecValue1 string  `json:"spec_value1"` //规格值1
	SpecValue2 string  `json:"spec_value2"` //规格值2
	SpecWidth  int     `json:"spec_width"`  //宽——如果有规格，会忽略，即使规格集中的没有传
	Volume     float64 `json:"volume"`      //体积——如果有规格，会忽略，即使规格集中的没有传
	Weight     float64 `json:"weight"`      //重量——如果有规格，会忽略，即使规格集中的没有传
}

type GoodsUpdateItem struct {
	ArticleNumber string                `json:"article_number"` //货号
	BarCode       string                `json:"bar_code"`       //商品级条码——如果有规格，次条码会忽略，即使规格集中的条码没有传
	Brand         string                `json:"brand"`          //品牌
	Categroy      string                `json:"categroy"`       //类目
	Height        int                   `json:"height"`         //高——如果有规格，会忽略，即使规格集中的没有传
	ItemCode      string                `json:"item_code"`      //商品编码
	ItemName      string                `json:"item_name"`      //商品名称
	ItemPic       string                `json:"item_pic"`       //图片
	Length        int                   `json:"length"`         //长——如果有规格，会忽略，即使规格集中的没有传
	PrimePrice    float64               `json:"prime_price"`    //参考进价——如果有规格，会忽略，即使规格集中的没有传
	SalePrice     float64               `json:"sale_price"`     //标准售价——如果有规格，会忽略，即使规格集中的没有传
	Skus          []GoodsUpdateItemSkus `json:"skus"`           //规格集
	Unit          string                `json:"unit"`           //单位
	Volume        float64               `json:"volume"`         //体积——如果有规格，会忽略，即使规格集中的没有传
	Weight        float64               `json:"weight"`         //重量——如果有规格，会忽略，即使规格集中的没有传
	Width         int                   `json:"width"`          //宽——如果有规格，会忽略，即使规格集中的没有传
}

type GoodsUpdate struct {
	Item GoodsUpdateItem `json:"item"` //商品信息
}

func NewGoodsUpdateItemSkus(
	BarCode string,
	PrimePrice float64,
	SalePrice float64,
	SpecCode string,
	SpecHeight int,
	SpecLength int,
	SpecPic string,
	SpecValue1 string,
	SpecValue2 string,
	SpecWidth int,
	Volume float64,
	Weight float64,
) GoodsUpdateItemSkus {
	return GoodsUpdateItemSkus{
		BarCode:    BarCode,
		PrimePrice: PrimePrice,
		SalePrice:  SalePrice,
		SpecCode:   SpecCode,
		SpecHeight: SpecHeight,
		SpecLength: SpecLength,
		SpecPic:    SpecPic,
		SpecValue1: SpecValue1,
		SpecValue2: SpecValue2,
		SpecWidth:  SpecWidth,
		Volume:     Volume,
		Weight:     Weight,
	}
}

func NewGoodsUpdate(
	ArticleNumber string,
	BarCode string,
	Brand string,
	Categroy string,
	Height int,
	ItemCode string,
	ItemName string,
	ItemPic string,
	Length int,
	PrimePrice float64,
	SalePrice float64,
	Skus []GoodsUpdateItemSkus,
	Unit string,
	Volume float64,
	Weight float64,
	Width int,
) []byte {
	GUpdate = &GoodsUpdate{
		Item: GoodsUpdateItem{
			ArticleNumber: ArticleNumber,
			BarCode:       BarCode,
			Brand:         Brand,
			Categroy:      Categroy,
			Height:        Height,
			ItemCode:      ItemCode,
			ItemName:      ItemName,
			ItemPic:       ItemPic,
			Length:        Length,
			PrimePrice:    PrimePrice,
			SalePrice:     SalePrice,
			Skus:          Skus,
			Unit:          Unit,
			Volume:        Volume,
			Weight:        Weight,
			Width:         Width,
		},
	}

	b1, err := json.Marshal(GUpdate)
	if err != nil {
		log.DLog.Println(err)
	}
	return b1
}

type RespGoodsUpdate struct {
	Code    int    `json:"code"`    //返回 0 表示执行成功 响应代码
	Data    bool   `json:"data"`    //响应结果数据
	Message string `json:"message"` //仅执行出错时返回 响应异常信息
}

//erp/goods/spec/open/query/goodswithspeclist
//查询出商品并带有规格列表,spec_code,item_code,modify_time,bar_code至少一个不能为空——对外开放

type QueryGoodsWithSpecList struct {
	SpecCode   string `json:"spec_code"`   //规格编码
	ItemCode   string `json:"item_code"`   //商品编码
	ModifyTime string `json:"modify_time"` //修改时间
	BarCode    string `json:"bar_code"`    //条码
	Page       int    `json:"page"`        //当前页码
	Limit      int    `json:"limit"`       //每页大小
}

func NewQueryGoodsWithSpecList(
	SpecCode string,
	ItemCode string,
	ModifyTime string,
	BarCode string,
	Page int,
	Limit int,
) []byte {
	QGoodsWithSpecList = &QueryGoodsWithSpecList{
		SpecCode:   SpecCode,
		ItemCode:   ItemCode,
		ModifyTime: ModifyTime,
		BarCode:    BarCode,
		Page:       Page,
		Limit:      Limit,
	}

	b1, err := json.Marshal(QGoodsWithSpecList)
	if err != nil {
		log.DLog.Println(err)
	}
	return b1
}

type RespQueryGoodsWithSpecListDataSpecs struct {
	Barcode        string  `json:"barcode"`         //条码
	Height         float64 `json:"height"`          //高度
	Length         float64 `json:"length"`          //长度
	Pic            string  `json:"pic"`             //规格图片
	PrimePrice     float64 `json:"prime_price"`     //参考进价
	SalePrice      float64 `json:"sale_price"`      //标准售价
	Spec1          string  `json:"spec1"`           //规格1值
	Spec2          string  `json:"spec2"`           //规格2值
	SpecCode       string  `json:"spec_code"`       //规格编码
	Weight         float64 `json:"weight"`          //重量
	WholesalePrice float64 `json:"wholesale_price"` //批发价
	Width          float64 `json:"width"`           //宽度
}

type RespQueryGoodsWithSpecListData struct {
	BrandName        string                                `json:"brand_name"`         //品牌
	CatagoryName     string                                `json:"catagory_name"`      //分类
	Expiration       int                                   `json:"expiration"`         //保质期
	GoodsCode        string                                `json:"goods_code"`         //商品编码
	GoodsName        string                                `json:"goods_name"`         //商品名称
	ManufacturerName string                                `json:"manufacturer_name"`  //生产商
	Pic              string                                `json:"pic"`                //商品图片
	PurchaseNum      float64                               `json:"purchase_num"`       //采购数量
	PurchaseTypeName string                                `json:"purchase_type_name"` //采购类型
	Remark           string                                `json:"remark"`             //备注
	Specs            []RespQueryGoodsWithSpecListDataSpecs `json:"specs"`              //规格集
	TagPrice         float64                               `json:"tag_price"`          //吊牌价
	UnitName         string                                `json:"unit_name"`          //单位
}

type RespQueryGoodsWithSpecList struct {
	Code    int                              `json:"code"`    //返回 0 表示执行成功 响应代码
	Data    []RespQueryGoodsWithSpecListData `json:"data"`    //响应结果数据
	Message string                           `json:"message"` //仅执行出错时返回 响应异常信息
}

//erp/goods/spec/open/query
//查询商品规格集合编码,其中spec_code,item_code,modify_time,bar_code不能同时为空——对外开放

type GoodsSpecOpenQuery struct {
	SpecCode   string `json:"spec_code"`   //规格编码
	ItemCode   string `json:"item_code"`   //商品编码
	ModifyTime string `json:"modify_time"` //修改时间
	BarCode    string `json:"bar_code"`    //条码
	Page       int    `json:"page"`        //当前页码
	Limit      int    `json:"limit"`       //每页大小
}

func NewGoodsSpecOpenQuery(SpecCode string,
	ItemCode string,
	ModifyTime string,
	BarCode string,
	Page int,
	Limit int,
) []byte {
	GSpecOpenQuery = &GoodsSpecOpenQuery{
		SpecCode:   SpecCode,
		ItemCode:   ItemCode,
		ModifyTime: ModifyTime,
		BarCode:    BarCode,
		Page:       Page,
		Limit:      Limit,
	}

	b1, err := json.Marshal(GSpecOpenQuery)
	if err != nil {
		log.DLog.Println(err)
	}
	return b1
}

type RespGoodsSpecOpenQueryData struct {
	ArticleNumber string  `json:"article_number"` //货号
	BarCode       string  `json:"bar_code"`       //条码
	Brand         string  `json:"brand"`          //品牌
	Catagory      string  `json:"catagory"`       //类目
	Color         string  `json:"color"`          //颜色
	ItemCode      string  `json:"item_code"`      //商品编码
	ItemName      string  `json:"item_name"`      //商品名称
	OtherProp     string  `json:"other_prop"`     //其它属性
	Price         float64 `json:"price"`          //标准售价
	Prop1         string  `json:"prop1"`          //自定义属性
	Prop2         string  `json:"prop2"`          //自定义属性
	Prop3         string  `json:"prop3"`          //自定义属性
	SpecCode      string  `json:"spec_code"`      //规格编码
	Unit          string  `json:"unit"`           //单位
}
type RespGoodsSpecOpenQuery struct {
	Code    int                          `json:"code"`    //返回 0 表示执行成功 响应代码
	Data    []RespGoodsSpecOpenQueryData `json:"data"`    //响应结果数据
	Message string                       `json:"message"` //仅执行出错时返回 响应异常信息
}

//erp/goods/add/item
//保存商品信息——对外开放

type GoodsAddItemSkus struct {
	BarCode    string  `json:"bar_code"`    //商品级条码——如果有规格，次条码会忽略，即使规格集中的条码没有传
	PrimePrice float64 `json:"prime_price"` //参考进价——如果有规格，会忽略，即使规格集中的没有传
	SalePrice  float64 `json:"sale_price"`  //标准售价——如果有规格，会忽略，即使规格集中的没有传
	SpecCode   string  `json:"spec_code"`   //规格编码
	SpecHeight int     `json:"spec_height"` //高——如果有规格，会忽略，即使规格集中的没有传
	SpecLength int     `json:"spec_length"` //长——如果有规格，会忽略，即使规格集中的没有传
	SpecPic    string  `json:"spec_pic"`    //规格图片
	SpecValue1 string  `json:"spec_value1"` //规格值1
	SpecValue2 string  `json:"spec_value2"` //规格值2
	SpecWidth  int     `json:"spec_width"`  //宽——如果有规格，会忽略，即使规格集中的没有传
	Volume     float64 `json:"volume"`      //体积——如果有规格，会忽略，即使规格集中的没有传
	Weight     float64 `json:"weight"`      //重量——如果有规格，会忽略，即使规格集中的没有传
}

type GoodsAddItem struct {
	ArticleNumber string             `json:"article_number"` //货号
	BarCode       string             `json:"bar_code"`       //商品级条码——如果有规格，次条码会忽略，即使规格集中的条码没有传
	Brand         string             `json:"brand"`          //品牌
	Categroy      string             `json:"categroy"`       //类目
	Height        int                `json:"height"`         //高——如果有规格，会忽略，即使规格集中的没有传
	ItemCode      string             `json:"item_code"`      //商品编码
	ItemName      string             `json:"item_name"`      //商品名称
	ItemPic       string             `json:"item_pic"`       //图片
	Length        int                `json:"length"`         //长——如果有规格，会忽略，即使规格集中的没有传
	PrimePrice    float64            `json:"prime_price"`    //参考进价——如果有规格，会忽略，即使规格集中的没有传
	SalePrice     float64            `json:"sale_price"`     //标准售价——如果有规格，会忽略，即使规格集中的没有传
	Skus          []GoodsAddItemSkus `json:"skus"`           //规格集
	Unit          string             `json:"unit"`           //单位
	Volume        float64            `json:"volume"`         //体积——如果有规格，会忽略，即使规格集中的没有传
	Weight        float64            `json:"weight"`         //重量——如果有规格，会忽略，即使规格集中的没有传
	Width         int                `json:"width"`          //宽——如果有规格，会忽略，即使规格集中的没有传
}

type GoodsAdd struct {
	Item GoodsAddItem `json:"item"` //商品信息
}

func NewGoodsAddItemSkus(
	BarCode string,
	PrimePrice float64,
	SalePrice float64,
	SpecCode string,
	SpecHeight int,
	SpecLength int,
	SpecPic string,
	SpecValue1 string,
	SpecValue2 string,
	SpecWidth int,
	Volume float64,
	Weight float64,
) GoodsAddItemSkus {
	return GoodsAddItemSkus{
		BarCode:    BarCode,
		PrimePrice: PrimePrice,
		SalePrice:  SalePrice,
		SpecCode:   SpecCode,
		SpecHeight: SpecHeight,
		SpecLength: SpecLength,
		SpecPic:    SpecPic,
		SpecValue1: SpecValue1,
		SpecValue2: SpecValue2,
		SpecWidth:  SpecWidth,
		Volume:     Volume,
		Weight:     Weight,
	}
}

func NewGoodsAdd(
	ArticleNumber string,
	BarCode string,
	Brand string,
	Categroy string,
	Height int,
	ItemCode string,
	ItemName string,
	ItemPic string,
	Length int,
	PrimePrice float64,
	SalePrice float64,
	Skus []GoodsAddItemSkus,
	Unit string,
	Volume float64,
	Weight float64,
	Width int,
) []byte {
	GAdd = &GoodsAdd{
		Item: GoodsAddItem{
			ArticleNumber: ArticleNumber,
			BarCode:       BarCode,
			Brand:         Brand,
			Categroy:      Categroy,
			Height:        Height,
			ItemCode:      ItemCode,
			ItemName:      ItemName,
			ItemPic:       ItemPic,
			Length:        Length,
			PrimePrice:    PrimePrice,
			SalePrice:     SalePrice,
			Skus:          Skus,
			Unit:          Unit,
			Volume:        Volume,
			Weight:        Weight,
			Width:         Width,
		},
	}

	b1, err := json.Marshal(GAdd)
	if err != nil {
		log.DLog.Println(err)
	}
	return b1
}

type RespGoodsAdd struct {
	Code    int    `json:"code"`    //返回 0 表示执行成功 响应代码
	Data    bool   `json:"data"`    //响应结果数据
	Message string `json:"message"` //仅执行出错时返回 响应异常信息
}

//erp/base/supplier/add
//添加供应商
type SupplierAddSupplier struct {
	Address      string `json:"address"`       //地址
	Area         string `json:"area"`          //区
	City         string `json:"city"`          //市
	Contact      string `json:"contact"`       //联系人
	Mobile       string `json:"mobile"`        //手机
	Province     string `json:"province"`      //省
	Remark       string `json:"remark"`        //备注
	Status       int    `json:"status"`        //状态 1-启用，0-停用
	SupplierCode string `json:"supplier_code"` //供应商编码
	SupplierName string `json:"supplier_name"` //供应商名称
}

type SupplierAdd struct {
	Supplier SupplierAddSupplier `json:"supplier"`
}

func NewSupplierAdd(
	Address string,
	Area string,
	City string,
	Contact string,
	Mobile string,
	Province string,
	Remark string,
	Status int,
	SupplierCode string,
	SupplierName string,
) []byte {

	SAdd = &SupplierAdd{
		Supplier: SupplierAddSupplier{
			Address:      Address,
			Area:         Area,
			City:         City,
			Contact:      Contact,
			Mobile:       Mobile,
			Province:     Province,
			Remark:       Remark,
			Status:       Status,
			SupplierCode: SupplierCode,
			SupplierName: SupplierName,
		},
	}

	b1, err := json.Marshal(SAdd)
	if err != nil {
		log.DLog.Println(err)
	}
	return b1
}

type RespSupplierAdd struct {
	Code    int    `json:"code"`    //返回 0 表示执行成功 响应代码
	Data    bool   `json:"data"`    //响应结果数据
	Message string `json:"message"` //仅执行出错时返回 响应异常信息
}

//erp/base/supplier/query
//查询供应商信息

type SupplierQuery struct {
	Keyword  string `json:"keyword"`   //供应商编码
	PageNo   int    `json:"page_no"`   //分页号
	PageSize int    `json:"page_size"` //分页大小 最大允许100
}

func NewSupplierQuery(
	Keyword string,
	PageNo int,
	PageSize int) []byte {

	SpQuery = &SupplierQuery{
		Keyword:  Keyword,
		PageNo:   PageNo,
		PageSize: PageSize,
	}

	b1, err := json.Marshal(SpQuery)
	if err != nil {
		log.DLog.Println(err)
	}
	return b1
}

type RespSupplierQueryData struct {
	Address      string `json:"address"`       //地址
	Area         string `json:"area"`          //区
	City         string `json:"city"`          //市
	Contact      string `json:"contact"`       //联系人
	Mobile       string `json:"mobile"`        //手机
	Province     string `json:"province"`      //省
	Remark       string `json:"remark"`        //备注
	Status       int    `json:"status"`        //状态 1-启用，0-停用
	SupplierCode string `json:"supplier_code"` //供应商编码
	SupplierName string `json:"supplier_name"` //供应商名称
}

type RespSupplierQuery struct {
	Code    int    `json:"code"`    //返回 0 表示执行成功 响应代码
	Data    []RespSupplierQueryData   `json:"data"`    //响应结果数据
	Message string `json:"message"` //仅执行出错时返回 响应异常信息
}



var SModify *SupplierModify
var BDistrQuery *B2bDistrQuery
var SQuery *StorageQuery
var BBillBatch *BatchBillBatch
var GUpdate *GoodsUpdate
var QGoodsWithSpecList *QueryGoodsWithSpecList
var GSpecOpenQuery *GoodsSpecOpenQuery
var GAdd *GoodsAdd
var SAdd *SupplierAdd
var SpQuery *SupplierQuery
