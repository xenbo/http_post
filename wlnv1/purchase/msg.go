package purchase

import (
	"encoding/json"
	"github.com/xenbo/http_post/log"
)

type ReturnBillClose struct {
	BillCode        string `json:"bill_code"`         //单据编码
	CloseRemark     string `json:"close_remark"`      //关闭备注
	NotCloseStocked bool   `json:"not_close_stocked"` //已出库不关闭
}

func NewReturnBillClose(BillCode string, CloseRemark string, NotCloseStocked bool) []byte {
	RBillClose = &ReturnBillClose{
		BillCode:        BillCode,
		CloseRemark:     CloseRemark,
		NotCloseStocked: NotCloseStocked,
	}

	b1, err := json.Marshal(RBillClose)
	if err != nil {
		log.DLog.Println(err)
	}

	return b1
}

type RespReturnBillClose struct {
	Code    int    //返回 0 表示执行成功 响应代码
	Data    bool   //响应结果数据
	Message string //仅执行出错时返回 响应异常信息
}

//-------------------------------------------------------------------------

type BillAddBillDetails struct {
	Index    int     `json:"index"`     //行号，查询单据的时候会返回，如果没填，erp给默认
	Price    float64 `json:"price"`     //采购价
	Remark   string  `json:"remark"`    //备注
	Size     float64 `json:"size"`      //数量
	SpecCode string  `json:"spec_code"` //规格编码
	Sum      float64 `json:"sum"`       //总价
}

type BillAddBill struct {
	BillCode     string               `json:"bill_code"`     //单号,如果为空，则使用erp自己规则生成的单号,请保证单号不能重复，如重复报错
	Details      []BillAddBillDetails `json:"details"`       //明细
	Remark       string               `json:"remark"`        //备注
	StorageCode  string               `json:"storage_code"`  //仓库编码
	SupplierCode string               `json:"supplier_code"` //供应商编码
}

type BillAdd struct {
	Bill BillAddBill `json:"bill"`
}

func NewBillAddBillDetail(Index int, Price float64, Remark string, Size float64, SpecCode string, Sum float64) BillAddBillDetails {
	return BillAddBillDetails{
		Index:    Index,
		Price:    Price,
		Remark:   Remark,
		Size:     Size,
		SpecCode: SpecCode,
		Sum:      Sum,
	}
}

func NewBillAdd(
	BillCode string,
	Remark string,
	StorageCode string,
	SupplierCode string,
	Details []BillAddBillDetails) []byte {

	BAdd = &BillAdd{
		Bill: BillAddBill{
			BillCode:     BillCode,
			Remark:       Remark,
			StorageCode:  StorageCode,
			SupplierCode: SupplierCode,
			Details:      Details,
		},
	}

	b1, err := json.Marshal(BAdd)
	if err != nil {
		log.DLog.Println(err)
	}

	return b1
}

type RespBillAdd struct {
	Code    int    `json:"code"`    //返回 0 表示执行成功  响应代码
	Data    string `json:"data"`    //响应结果数据
	Message string `json:"message"` //仅执行出错时返回 响应异常信息
}

//-------------------------------------------------------------------------
type BillQuery struct {
	BillCode   string `json:"bill_code"`   //单据编码
	ModifyTime string `json:"modify_time"` //修改时间，只能查近3个月
	Page       int    `json:"page"`        //当前页码，从1开始
	Limit      int    `json:"limit"`       //每页大小,最大200
}

func NewBillQuery(BillCode string, ModifyTime string, Page int, Limit int) []byte {
	BQuery = &BillQuery{
		BillCode:   BillCode,
		ModifyTime: ModifyTime,
		Page:       Page,
		Limit:      Limit,
	}

	b1, err := json.Marshal(BQuery)
	if err != nil {
		log.DLog.Println(err)
	}

	return b1
}

type RespBillQueryDataDetails struct {
	ArticleNumber string  `json:"article_number"` //货号
	BasePrice     float64 `json:"base_price"`     //原价
	DiscountRate  float64 `json:"discount_rate"`  //折扣率
	GoodsName     string  `json:"goods_name"`     //商品名称
	Index         int64   `json:"index"`          //行号
	Price         float64 `json:"price"`          //单价
	Receive       int     `json:"receive"`        //已入库数量
	Remark        string  `json:"remark"`         //备注
	Size          int     `json:"size"`           //数量
	SpecCode      string  `json:"spec_code"`      //规格编码
	SpecName      string  `json:"spec_name"`      //规格名称
	Status        int     `json:"status"`         //状态 1:未到货 2:部分到货 3:全部到货 4:已关闭
	Sum           float64 `json:"sum"`            //总价
	TaxRate       float64 `json:"tax_rate"`       //税率
	Unit          string  `json:"unit"`           //单位，默认erp的基本单位
}

type RespBillQueryData struct {
	BillCode     string                     `json:"bill_code"`     //单据编码
	BillDate     string                     `json:"bill_date"`     //单据业务时间
	CreateTime   string                     `json:"create_time"`   //创建时间
	Details      []RespBillQueryDataDetails `json:"details"`       //明细
	ModifiedTime string                     `json:"modified_time"` //修改时间
	Remark       string                     `json:"remark"`        //备注
	Status       int                        `json:"status"`        //状态 1： 未到货 2： 部分到货 3：完成 4 ：关闭 5：待提交 6 ：审核中
	StorageCode  string                     `json:"storage_code"`  //仓库编码
	storageName  string                     `json:"storage_name"`  //仓库名称
	SupplierCode string                     `json:"supplier_code"` //供应商编码
	SupplierName string                     `json:"supplier_name"` //供应商名称
}

type RespBillQuery struct {
	Code    int               `json:"code"`    //返回 0 表示执行成功 响应代码
	Data    []RespBillQueryData `json:"data"`    //响应结果数据
	Message string            `json:"message"` //仅执行出错时返回 响应异常信息
}

//-------------------------------------------------------------------------

type ReturnBillAddDetails struct {
	Index    int64   `json:"index"`     //行号，查询单据的时候会返回，如果没填，erp给默认
	Price    float64 `json:"price"`     //采购价
	Remark   string  `json:"remark"`    //备注
	Size     float64 `json:"size"`      //数量
	SpecCode string  `json:"spec_code"` //规格编码
	Sum      float64 `json:"sum"`       //总价
}

type ReturnBillAddBill struct {
	BillCode     string                 `json:"bill_code"`     //单号,如果为空，则使用erp自己规则生成的单号,请保证单号不能重复，如重复报错
	Details      []ReturnBillAddDetails `json:"details"`       //明细
	Remark       string                 `json:"remark"`        //备注
	StorageCode  string                 `json:"storage_code"`  //仓库编码
	SupplierCode string                 `json:"supplier_code"` //供应商编码
}

type ReturnBillAdd struct {
	Bill ReturnBillAddBill `json:"bill"` //采购退货订单
}

func NewReturnBillAddDetails(
	Index int64,
	Price float64,
	Remark string,
	Size float64,
	SpecCode string,
	Sum float64) ReturnBillAddDetails {

	return ReturnBillAddDetails{
		Index:    Index,
		Price:    Price,
		Remark:   Remark,
		Size:     Size,
		SpecCode: SpecCode,
		Sum:      Sum,
	}
}

func NewReturnBillAdd(
	BillCode string,
	Remark string,
	StorageCode string,
	SupplierCode string,
	Details []ReturnBillAddDetails, ) []byte {

	RBillAdd = &ReturnBillAdd{
		Bill: ReturnBillAddBill{
			BillCode:     BillCode,
			Details:      Details,
			Remark:       Remark,
			StorageCode:  StorageCode,
			SupplierCode: SupplierCode,
		},
	}

	b1, err := json.Marshal(RBillAdd)
	if err != nil {
		log.DLog.Println(err)
	}
	return b1
}

type RespReturnBillAdd struct {
	Code    int    `json:"code"`    //返回 0 表示执行成功 响应代码
	Data    bool   `json:"data"`    //响应结果数据
	Message string `json:"message"` //仅执行出错时返回 响应异常信息
}

//-------------------------------------------------------------------------

type StockInQuery struct {
	BillCode      string `json:"bill_code"`       //采购订单单据编码
	StockBillCode string `json:"stock_bill_code"` //采购入库单单据编码
	ModifyTime    string `json:"modify_time"`     //修改时间，只能查近3个月
	Page          int    `json:"page"`            //当前页码，从1开始
	Limit         int    `json:"limit"`           //每页大小,最大200
}

func NewStockInQuery(BillCode string, StockBillCode string, ModifyTime string, Page int, Limit int) []byte {
	SInQuery = &StockInQuery{
		BillCode:      BillCode,
		StockBillCode: StockBillCode,
		ModifyTime:    ModifyTime,
		Page:          Page,
		Limit:         Limit,
	}

	b1, err := json.Marshal(SInQuery)
	if err != nil {
		log.DLog.Println(err)
	}
	return b1
}

type RespStockInQueryDatadetails struct {
	BasePrice    float64 `json:"base_price"`     //原价
	DGoodsName   string  `json:"goods_name"`     //商品名称
	Index        int64   `json:"index"`          //行号
	Nums         int     `json:"nums"`           //数量
	PchsBillCode string  `json:"pchs_bill_code"` //采购订单编码
	Price        float64 `json:"price"`          //单价
	Remark       string  `json:"remark"`         //备注
	SpecCode     string  `json:"spec_code"`      //规格编码
	SpecName     string  `json:"spec_name"`      //规格名称
	TaxRate      float64 `json:"tax_rate"`       //税率
	TotalMoney   float64 `json:"total_money"`    //总金额
	Unit         string  `json:"unit"`           //单位
}

type RespStockInQueryData struct {
	BillDate     string                        `json:"bill_date"`     //业务日期
	CreateTime   string                        `json:"create_time"`   //创建时间
	Debt         float64                       `json:"debt"`          //总商品金额
	Details      []RespStockInQueryDatadetails `json:"details"`       //明细
	ModifiedTime string                        `json:"modified_time"` //单据修改时间
	Remark       string                        `json:"remark"`        //备注
	Saleman      string                        `json:"saleman"`       //业务员
	StockCode    string                        `json:"stock_code"`    //出入库单编码
	StorageCode  string                        `json:"storage_code"`  //仓库编码
	StorageName  string                        `json:"storage_name"`  //仓库名称
	Sumprice     float64                       `json:"sumprice"`      //总商品金额
	SupplierCode string                        `json:"supplier_code"` //供应商编码
	SupplierName string                        `json:"supplier_name"` //供应商名称
}

type RespStockInQuery struct {
	Code    int                    `json:"code"`    //返回 0 表示执行成功  响应代码
	Data    []RespStockInQueryData `json:"data"`    //响应结果数据
	Message string                 `json:"message"` //仅执行出错时返回 响应异常信息
}

//-------------------------------------------------------------------------

type TurnBillQuery struct {
	BillCode   string `json:"bill_code"`   //采购订单单据编码
	ModifyTime string `json:"modify_time"` //修改时间，只能查近3个月
	Page       int    `json:"page"`        //当前页码，从1开始
	Limit      int    `json:"limit"`       //每页大小,最大200
}

func NewTurnBillQuery(BillCode string, ModifyTime string, Page int, Limit int) []byte {
	TBillQuery = &TurnBillQuery{
		BillCode:   BillCode,
		ModifyTime: ModifyTime,
		Page:       Page,
		Limit:      Limit,
	}

	b1, err := json.Marshal(TBillQuery)
	if err != nil {
		log.DLog.Println(err)
	}

	return b1
}

type RespTurnBillQueryDataDetails struct {
	ArticleNumber string  `json:"article_number"` //货号
	BasePrice     float64 `json:"base_price"`     //原价
	DiscountRate  float64 `json:"discount_rate"`  //折扣率
	GoodsName     string  `json:"goods_name"`     //商品名称
	Index         int64   `json:"index"`          //行号
	Price         float64 `json:"price"`          //单价
	Receive       int     `json:"receive"`        //已入库数量
	Remark        string  `json:"remark"`         //备注
	Size          int     `json:"size"`           //数量
	SpecCode      string  `json:"spec_code"`      //规格编码
	SpecName      string  `json:"spec_name"`      //规格名称
	Status        int     `json:"status"`         //状态 1:未到货 2:部分到货 3:全部到货 4:已关闭
	Sum           float64 `json:"sum"`            //总价
	TaxRate       float64 `json:"tax_rate"`       //税率
	Unit          string  `json:"unit"`           //单位，默认erp的基本单位
}

type RespTurnBillQueryData struct {
	BillCode     string                         `json:"bill_code"`     //单据编码
	BillDate     string                         `json:"bill_date"`     //单据业务时间
	CreateTime   string                         `json:"create_time"`   //创建时间
	Details      []RespTurnBillQueryDataDetails `json:"details"`       //明细
	ModifiedTime string                         `json:"modified_time"` //修改时间
	Remark       string                         `json:"remark"`        //备注
	Status       int                            `json:"status"`        //状态 1： 未到货 2： 部分到货 3：完成 4 ：关闭 5：待提交 6 ：审核中
	StorageCode  string                         `json:"storage_code"`  //仓库编码
	StorageName  string                         `json:"storage_name"`  //仓库名称
	SupplierCode string                         `json:"supplier_code"` //供应商编码
	SupplierName string                         `json:"supplier_name"` //供应商名称
}

type RespTurnBillQuery struct {
	Code    int                     `json:"code"`  //返回 0 表示执行成功 响应代码
	Data    []RespTurnBillQueryData `json:"data"`  //响应结果数据
	Message string                  `json:"limit"` //仅执行出错时返回 响应异常信息
}

//-------------------------------------------------------------------------

type StockInAddDetails struct {
	Index         int64   `json:"index"`       //行号，查询单据的时候会返回，如果没填，erp给默认
	Nums          int     `json:"nums"`        //明细商品数量（必填）
	PchsBillCode2 string  `json:"bill_code"`   //采购单据编号
	Remark2       string  `json:"remark"`      //备注
	SpecCode      string  `json:"spec_code"`   //规格编码（必填）
	TaxRate       float64 `json:"tax_rate"`    //税率
	TotalMoney    float64 `json:"total_money"` //明细商品总价格
	TotalPrice    float64 `json:"total_price"` //明细实际支付总金额
}

type StockInAddBill struct {
	BillDate     string              `json:"bill_date"`     //业务日期（必填）
	Carriage     float64             `json:"carriage"`      //采购运费
	Debt         float64             `json:"debt"`          //欠款
	Details      []StockInAddDetails `json:"details"`       //采购明细（必填）
	PchsBillCode string              `json:"bill_code"`     //采购单据编号,如果万里牛中不存在此单号会报错，如不传此单号，则不会关联采购订单，直接做采购入库单
	Remark       string              `json:"remark"`        //备注
	SaleMan      string              `json:"sale_man"`      //业务员
	StorageCode  string              `json:"storage_code"`  //仓库编码（必填）
	SupplierCode string              `json:"supplier_code"` //供应商编码（必填）
}

type StockInAdd struct {
	Bill StockInAddBill `json:"bill"` //采购入库单
}

func NewStockInAdd(BillDate string, Carriage float64, Debt float64, PchsBillCode string, Remark string, SaleMan string,
	StorageCode string, SupplierCode string, Details []StockInAddDetails) []byte {

	SInAdd = &StockInAdd{
		Bill: StockInAddBill{
			BillDate:     BillDate,
			Carriage:     Carriage,
			Debt:         Debt,
			PchsBillCode: PchsBillCode,
			Remark:       Remark,
			SaleMan:      SaleMan,
			StorageCode:  StorageCode,
			SupplierCode: SupplierCode,
		},
	}

	for _, v := range Details {
		SInAdd.Bill.Details = append(SInAdd.Bill.Details, v)
	}

	b1, err := json.Marshal(SInAdd)
	if err != nil {
		log.DLog.Println(err)
	}

	return b1
}

func NewStockInAddDetails(Index int64, Nums int, PchsBillCode2 string, Remark2 string, SpecCode string, TaxRate float64,
	TotalMoney float64, TotalPrice float64) StockInAddDetails {

	return StockInAddDetails{
		Index:         Index,
		Nums:          Nums,
		PchsBillCode2: PchsBillCode2,
		Remark2:       Remark2,
		SpecCode:      SpecCode,
		TaxRate:       TaxRate,
		TotalMoney:    TotalMoney,
		TotalPrice:    TotalPrice,
	}
}

type RespStockInAdd struct {
	Code    int                     `json:"code"`    //返回 0 表示执行成功 //响应代码
	Data    []RespStockOutQueryData `json:"data"`    //响应结果数据
	Message string                  `json:"message"` //仅执行出错时返回 响应异常信息
}

//-------------------------------------------------------------------------

type StockOutQuery struct {
	BillCode      string `json:"bill_code"`       //采购退货订单单据编码
	StockBillCode string `json:"stock_bill_code"` //采购退货出库单单据编码
	ModifyTime    string `json:"modify_time"`     //修改时间，只能查近3个月
	Page          int    `json:"page"`            //当前页码，从1开始
	Limit         int    `json:"limit"`           //每页大小,最大200
}

func NewStockOutQuery(BillCode string, StockBillCode string, ModifyTime string, Page int, Limit int) []byte {
	SOutQuery = &StockOutQuery{
		BillCode:      BillCode,
		StockBillCode: StockBillCode,
		ModifyTime:    ModifyTime,
		Page:          Page,
		Limit:         Limit,
	}

	b1, err := json.Marshal(SOutQuery)
	if err != nil {
		log.DLog.Println(err)
	}

	return b1
}

type RespStockOutQueryDataDetails struct {
	BasePrice    float64 `json:"base_price"`     //原价
	DiscountRate float64 `json:"discount_rate"`  //折扣率
	GoodsName    string  `json:"goods_name"`     //商品名称
	Index        int64   `json:"index"`          //行号
	Nums         int64   `json:"nums"`           //数量
	PchsBillCode string  `json:"pchs_bill_code"` //采购订单编码
	Price        float64 `json:"price"`          //单价
	Remark       string  `json:"remark"`         //备注
	SpecCode     string  `json:"spec_code"`      //规格编码
	SpecName     string  `json:"spec_name"`      //规格名称
	TaxRate      float64 `json:"tax_rate"`       //税率
	TotalMoney   float64 `json:"total_money"`    //总金额
	Unit         string  `json:"unit"`           //单位
}

type RespStockOutQueryData struct {
	BillDate   string                         `json:"bill_date"`   //业务日期
	CreateTime string                         `json:"create_time"` //创建时间
	Debt       float64                        `json:"debt"`        //总商品金额
	Details    []RespStockOutQueryDataDetails `json:"details"`     //明细
}

type RespStockOutQuery struct {
	Code    int                     `json:"code"`    //返回 0 表示执行成功 //响应代码
	Data    []RespStockOutQueryData `json:"data"`    //响应结果数据
	Message string                  `json:"message"` //仅执行出错时返回 响应异常信息
}

//-------------------------------------------------------------------------

type StockBackProductBillAddBillDetails struct {
	Index        int64   `json:"index"`          //行号，查询单据的时候会返回，如果没填，erp给默认
	Num          int     `json:"num"`            //明细商品数量（必填）
	PchsBillCode string  `json:"pchs_bill_code"` //采购单据编号
	Remark       string  `json:"remark"`         //备注
	SpecCode     string  `json:"spec_code"`      //规格编码（必填）
	TaxRate      float64 `json:"tax_rate"`       //税率
	TotalMoney   float64 `json:"total_money"`    //明细商品总价格
	TotalPrice   float64 `json:"total_price"`    //明细实际支付总金额
}

type StockBackProductBillAddBill struct {
	BillDate     string                               `json:"bill_date"`      //业务日期（必填）
	Carriage     float64                              `json:"carriage"`       //采购运费
	Debt         float64                              `json:"debt"`           //欠款
	Details      []StockBackProductBillAddBillDetails `json:"details"`        //采购明细（必填）
	PchsBillCode string                               `json:"pchs_bill_code"` //采购单据编号,如果万里牛中不存在此单号会报错，如不传此单号，则不会关联采购订单，直接做采购入库单
	Remark       string                               `json:"remark"`         //备注
	SaleMan      string                               `json:"sale_man"`       //业务员
	StorageCode  string                               `json:"storage_code"`   //仓库编码（必填）
	SupplierCode string                               `json:"supplier_code"`  //供应商编码（必填）
}

type StockBackProductBillAdd struct {
	Bill StockBackProductBillAddBill `json:"bill"` //采购退货出库单
}

func NewStockBackProductBillAddBillDetails(
	Index int64,
	Num int,
	PchsBillCode string,
	Remark string,
	SpecCode string,
	TaxRate float64,
	TotalMoney float64,
	TotalPrice float64) StockBackProductBillAddBillDetails {

	return StockBackProductBillAddBillDetails{
		Index:        Index,
		Num:          Num,
		PchsBillCode: PchsBillCode,
		Remark:       Remark,
		SpecCode:     SpecCode,
		TaxRate:      TaxRate,
		TotalMoney:   TotalMoney,
		TotalPrice:   TotalPrice,
	}
}

func NewStockBackProductBillAdd(
	BillDate string,
	Carriage float64,
	Debt float64,
	PchsBillCode string,
	Remark string,
	SaleMan string,
	StorageCode string,
	SupplierCode string,
	Details []StockBackProductBillAddBillDetails) []byte {

	SBackProductBillAdd = &StockBackProductBillAdd{
		Bill: StockBackProductBillAddBill{
			BillDate:     BillDate,
			Carriage:     Carriage,
			Debt:         Debt,
			PchsBillCode: PchsBillCode,
			Remark:       Remark,
			SaleMan:      SaleMan,
			StorageCode:  StorageCode,
			SupplierCode: SupplierCode,
			Details:      Details,
		},
	}

	b1, err := json.Marshal(SBackProductBillAdd)
	if err != nil {
		log.DLog.Println(err)
	}
	return b1
}

type RespStockBackProductBillAdd struct {
	Code    int    `json:"code"`    //返回 0 表示执行成功 响应代码
	Data    bool   `json:"data"`    //响应结果数据
	Message string `json:"message"` //仅执行出错时返回 响应异常信息
}

//-------------------------------------------------------------------------

type BillClose struct {
	BillCode        string `json:"bill_code"`         //单据编码
	CloseRemark     string `json:"close_remark"`      //关闭备注
	NotCloseStocked bool   `json:"not_close_stocked"` //已入库不关闭
}

func NewBillClose(
	BillCode string,
	CloseRemark string,
	NotCloseStocked bool) []byte {

	BClose = &BillClose{
		BillCode:        BillCode,
		CloseRemark:     CloseRemark,
		NotCloseStocked: NotCloseStocked,
	}

	b1, err := json.Marshal(BClose)
	if err != nil {
		log.DLog.Println(err)
	}

	return b1
}

type RespBillClose struct {
	Code    int    `json:"code"`    //返回 0 表示执行成功 响应代码
	Data    bool   `json:"data"`    //响应结果数据
	Message string `json:"message"` //仅执行出错时返回 响应异常信息
}

//-------------------------------------------------------------------------
//-------------------------------------------------------------------------
//-------------------------------------------------------------------------
//-------------------------------------------------------------------------
//-------------------------------------------------------------------------
//-------------------------------------------------------------------------
//-------------------------------------------------------------------------
//-------------------------------------------------------------------------
//-------------------------------------------------------------------------
//-------------------------------------------------------------------------
//-------------------------------------------------------------------------
//-------------------------------------------------------------------------
//-------------------------------------------------------------------------

var RBillClose *ReturnBillClose
var BAdd *BillAdd
var BQuery *BillQuery
var RBillAdd *ReturnBillAdd
var SInQuery *StockInQuery
var TBillQuery *TurnBillQuery
var SInAdd *StockInAdd
var SOutQuery *StockOutQuery
var SBackProductBillAdd *StockBackProductBillAdd
var BClose *BillClose
