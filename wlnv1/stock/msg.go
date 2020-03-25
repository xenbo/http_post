package stock

import (
	"encoding/json"
	"github.com/xenbo/http_post/log"
)

//erp/stock/out/requestbill/add
//新增其他出库订单,明细按照基本单位保存,对外开放

type RequestBillAddBillDetails struct {
	Index    int     `json:"index"`     //行号，查询单据的时候会返回，如果没填，erp给默认
	Remark   string  `json:"remark"`    //备注
	Size     float64 `json:"size"`      //数量
	SpecCode string  `json:"spec_code"` //规格编码
}

type RequestBillAddBill struct {
	Address       string                      `json:"address"`         //详细地址——其它出库单使用
	BillCode      string                      `json:"bill_code"`       //单号,如果为空，则使用erp自己规则生成的单号,请保证单号不能重复，如重复报错
	City          string                      `json:"city"`            //市——其它出库单使用
	Details       []RequestBillAddBillDetails `json:"details"`         //明细
	District      string                      `json:"district"`        //区——其它出库单使用
	Mobile        string                      `json:"mobile"`          //手机——其它出库单使用
	OuterBillCode string                      `json:"outer_bill_code"` //外部bill_code,用来存储对接系统的一些特别的值，万里牛中无作用,可以不传
	Province      string                      `json:"province"`        //省——其它出库单使用
	Reason        string                      `json:"reason"`          //业务原因
	Receiver      string                      `json:"receiver"`        //收件人——其它出库单使用
	Remark        string                      `json:"remark"`          //备注
	StorageCode   string                      `json:"storage_code"`    //仓库编码
}

type RequestBillAdd struct {
	Bill RequestBillAddBill `json:"bill"` //其他出库订单
}

func NewRequestBillAddBillDetails(
	Index int,
	Remark string,
	Size float64,
	SpecCode string,
) RequestBillAddBillDetails {

	return RequestBillAddBillDetails{
		Index:    Index,
		Remark:   Remark,
		Size:     Size,
		SpecCode: SpecCode,
	}
}

func NewRequestBillAdd(Address string,
	BillCode string,
	City string,
	District string,
	Mobile string,
	OuterBillCode string,
	Province string,
	Reason string,
	Receiver string,
	Remark string,
	StorageCode string,
	Details []RequestBillAddBillDetails,
) []byte {

	RBillAdd = &RequestBillAdd{
		Bill: RequestBillAddBill{
			Address:       Address,
			BillCode:      BillCode,
			City:          City,
			Details:       Details,
			District:      District,
			Mobile:        Mobile,
			OuterBillCode: OuterBillCode,
			Province:      Province,
			Reason:        Reason,
			Receiver:      Receiver,
			Remark:        Remark,
			StorageCode:   StorageCode,
		},
	}

	b1, err := json.Marshal(RBillAdd)
	if err != nil {
		log.DLog.Println(err)
	}

	return b1

}

type RespRequestBillAdd struct {
	Code    int    //返回 0 表示执行成功 响应代码
	Data    bool   //响应结果数据
	Message string //仅执行出错时返回 响应异常信息
}

//erp/sn/querytrace
//查询sn轨迹，开始时间可以任意输入，不能为空，单结束时间必须是开始时间之后的24小时内,对外开放

type SNQueryTrace struct {
	SnCode string `json:"sn_code"` //sn编码
	Start  string `json:"start"`   //开始时间
	End    string `json:"end"`     //结束时间
	Page   int    `json:"page"`    //当前页码
	Limit  int    `json:"limit"`   //每页大小
}

func NewSNQueryTrace(
	SnCode string,
	Start string,
	End string,
	Page int,
	Limit int,
) []byte {

	SNQTrace = &SNQueryTrace{
		SnCode: SnCode,
		Start:  Start,
		End:    End,
		Page:   Page,
		Limit:  Limit,
	}

	b1, err := json.Marshal(SNQTrace)
	if err != nil {
		log.DLog.Println(err)
	}

	return b1
}

type RespSNQueryTraceData struct {
	BillCode    string `json:"bill_code"`    //单据编码
	BillType    int    `json:"bill_type"`    //单据类型,1: 销售出库 2： 退货入库 3：采购入库 4：其他出库 5：其他入库 6：采购退货出库 7：调拨出库 8：调拨入库 9：其他
	CustomName  string `json:"custom_name"`  //客户名称
	CustomNick  string `json:"custom_nick"`  //客户昵称
	CustomType  string `json:"custom_type"`  //客户类型
	ItemName    string `json:"item_name"`    //商品名称
	ShopType    string `json:"shop_type"`    //店铺来源（bill_type=1时会有）
	SnCode      string `json:"sn_code"`      //sn码
	SpecCode    string `json:"spec_code"`    //规格编码
	StorageCode string `json:"storage_code"` //仓库编码
	StorageName string `json:"storage_name"` //仓库名称
	Supplier    string `json:"supplier"`     //供应商
	Time        int64  `json:"time"`         //sn操作时间
}

type RespSNQueryTrace struct {
	Code    int                  `json:"code"`    //返回 0 表示执行成功 响应代码
	Data    RespSNQueryTraceData `json:"data"`    //响应结果数据
	Message string               `json:"message"` //仅执行出错时返回 响应异常信息
}

//erp/stock/in/requestbill/add
//新增其他入库订单,明细按照基本单位保存,对外开放

type InRequestBillAddBillDetails struct {
	Index    int64   `json:"index"`     //行号，查询单据的时候会返回，如果没填，erp给默认
	Remark   string  `json:"remark"`    //备注
	Size     float64 `json:"size"`      //数量
	SpecCode string  `json:"spec_code"` //规格编码
}

type InRequestBillAddBill struct {
	Address       string                        `json:"address"`         //详细地址——其它出库单使用
	BillCode      string                        `json:"bill_code"`       //单号,如果为空，则使用erp自己规则生成的单号,请保证单号不能重复，如重复报错
	City          string                        `json:"city"`            //市——其它出库单使用
	Details       []InRequestBillAddBillDetails `json:"details"`         //明细
	District      string                        `json:"district"`        ///区——其它出库单使用
	Mobile        string                        `json:"mobile"`          //手机——其它出库单使用
	OuterBillCode string                        `json:"outer_bill_code"` //外部bill_code,用来存储对接系统的一些特别的值，万里牛中无作用,可以不传
	Province      string                        `json:"province"`        //省——其它出库单使用
	Reason        string                        `json:"reason"`          //业务原因
	Receiver      string                        `json:"receiver"`        //收件人——其它出库单使用
	Remark        string                        `json:"remark"`          //备注
	StorageCode   string                        `json:"storage_code"`    //仓库编码
}

type InRequestBillAdd struct {
	Bill InRequestBillAddBill `json:"bill"` //其他入库订单
}

func NewInRequestBillAddBillDetails(
	Index int64,
	Remark string,
	Size float64,
	SpecCode string,
) InRequestBillAddBillDetails {
	return InRequestBillAddBillDetails{
		Index:    Index,
		Remark:   Remark,
		Size:     Size,
		SpecCode: SpecCode,
	}
}

func NewInRequestBillAdd(
	Address string,
	BillCode string,
	City string,
	District string,
	Mobile string,
	OuterBillCode string,
	Province string,
	Reason string,
	Receiver string,
	Remark string,
	StorageCode string,
	Details []InRequestBillAddBillDetails,
) []byte {

	InRBillAdd = &InRequestBillAdd{
		Bill: InRequestBillAddBill{
			Address:       Address,
			BillCode:      BillCode,
			City:          City,
			Details:       Details,
			District:      District,
			Mobile:        Mobile,
			OuterBillCode: OuterBillCode,
			Province:      Province,
			Reason:        Reason,
			Receiver:      Receiver,
			Remark:        Remark,
			StorageCode:   StorageCode,
		},
	}

	b1, err := json.Marshal(InRBillAdd)
	if err != nil {
		log.DLog.Println(err)
	}

	return b1
}

type RespInRequestBillAdd struct {
	Code    int    //返回 0 表示执行成功 响应代码
	Data    bool   //响应结果数据
	Message string //仅执行出错时返回 响应异常信息
}

//erp/allocation/changebill/add
//新增调拨单,明细按照基本单位保存,对外开放

type CHangeBillAddBillDetails struct {
	Index    int64   `json:"index"`     //行号，查询单据的时候会返回，如果没填，erp给默认
	Remark   string  `json:"remark"`    //备注
	Size     float64 `json:"size"`      //数量
	SpecCode string  `json:"spec_code"` //规格编码
}

type CHangeBillAddBill struct {
	BillType       int                        `json:"bill_type"`        //调拨类型:1. 虚拟调拨 2. 实际调拨
	Details        []CHangeBillAddBillDetails `json:"details"`          //明细1111
	OperNick       string                     `json:"oper_nick"`        //创建员——erp中的登录账号
	Remark         string                     `json:"remark"`           //备注
	StorageInCode  string                     `json:"storage_in_code"`  //入库仓编码
	StorageOutCode string                     `json:"storage_out_code"` //出库仓编码
}

type CHangeBillAdd struct {
	Bill CHangeBillAddBill `json:"bill"` //调拨单
}

func NewCHangeBillAddBillDetails(
	Index int64,
	Remark string,
	Size float64,
	SpecCode string,
) CHangeBillAddBillDetails {
	return CHangeBillAddBillDetails{
		Index:    Index,
		Remark:   Remark,
		Size:     Size,
		SpecCode: SpecCode,
	}
}

func NewCHangeBillAdd(
	BillType int,
	OperNick string,
	Remark string,
	StorageInCode string,
	StorageOutCode string,
	Details []CHangeBillAddBillDetails,
) []byte {

	CHBillAdd = &CHangeBillAdd{
		Bill: CHangeBillAddBill{
			BillType:       BillType,
			OperNick:       OperNick,
			Remark:         Remark,
			StorageInCode:  StorageInCode,
			StorageOutCode: StorageOutCode,
			Details:        Details,
		},
	}

	b1, err := json.Marshal(CHBillAdd)
	if err != nil {
		log.DLog.Println(err)
	}

	return b1
}

type RespCHangeBillAdd struct {
	Code    int    //返回 0 表示执行成功 响应代码
	Data    bool   //响应结果数据
	Message string //仅执行出错时返回 响应异常信息
}

//erp/stock/in/requestbill/query
//查询其他入库订单,stock_request_code modify_time 不能同时为空,对外开放

type InRequestBillQuery struct {
	StockRequestCode string `json:"stock_request_code"` //单据编码
	ModifyTime       string `json:"modify_time"`        //修改时间，只能查近3个月
	Page             int    `json:"page"`               //当前页码，从1开始
	Limit            int    `json:"limit"`              //每页大小,最大200
}

func NewInRequestBillQuery(
	StockRequestCode string,
	ModifyTime string,
	Page int,
	Limit int,
) []byte {

	InRBillQuery = &InRequestBillQuery{
		StockRequestCode: StockRequestCode,
		ModifyTime:       ModifyTime,
		Page:             Page,
		Limit:            Limit,
	}

	b1, err := json.Marshal(InRBillQuery)
	if err != nil {
		log.DLog.Println(err)
	}

	return b1
}

type RespInRequestBillQuery struct {
	Code    int    //返回 0 表示执行成功 响应代码
	Data    bool   //响应结果数据
	Message string //仅执行出错时返回 响应异常信息
}

//erp/open/inventory/items/get/by/modifytime
//查询库存,规格编码和修改时间不能都为空,对外开放

type ItemsGetByModifyTime struct {
	Storage    string `json:"storage"`     //目标仓库编码
	ModifyTime string `json:"modify_time"` //修改时间s
	PageNo     int    `json:"page_no"`     //分页号
	PageSize   int    `json:"page_size"`   //分页大小 最大允许100
	SkuCode    string `json:"sku_code"`    //规格编码
}

func NewItemsGetByModifyTime(
	Storage string,
	ModifyTime string,
	PageNo int,
	PageSize int,
	SkuCode string,
) []byte {

	IGetByModifyTime = &ItemsGetByModifyTime{
		Storage:    Storage,
		ModifyTime: ModifyTime,
		PageNo:     PageNo,
		PageSize:   PageSize,
		SkuCode:    SkuCode,
	}

	b1, err := json.Marshal(IGetByModifyTime)
	if err != nil {
		log.DLog.Println(err)
	}

	return b1
}

type RespItemsGetByModifyTimeDataBatchs struct {
	BatchNo     string `json:"batch_no"`     //批次编码
	ExpiredDate string `json:"expired_date"` //过期日期
	Num         int    `json:"num"`          //数量
	ProduceDate string `json:"produce_date"` //成产日期
}

type RespItemsGetByModifyTimeData struct {
	Batchs    []RespItemsGetByModifyTimeDataBatchs `json:"batchs"`     //批次信息
	GoodsCode string                               `json:"goods_code"` //商品编码
	LockSize  float64                              `json:"lock_size"`  //锁定库存
	Quantity  float64                              `json:"quantity"`   //数量
	SkuCode   string                               `json:"sku_code"`   //规格编码
	Underway  float64                              `json:"underway"`   //在途库存
}

type RespItemsGetByModifyTime struct {
	Code    int                            `json:"code"`    //返回 0 表示执行成功 响应代码
	Data    []RespItemsGetByModifyTimeData `json:"data"`    //响应结果数据
	Message string                         `json:"message"` //仅执行出错时返回 响应异常信息
}

//erp/allocation/changebill/close
//关闭调拨单,对外开放

type ChangeBillClose struct {
	BillCode string `json:"bill_code"` //单据编码
}

func NewChangeBillClose(
	BillCode string,
) []byte {

	CBillClose = &ChangeBillClose{
		BillCode: BillCode,
	}

	b1, err := json.Marshal(CBillClose)
	if err != nil {
		log.DLog.Println(err)
	}

	return b1
}

type RespChangeBillClose struct {
	Code    int    //返回 0 表示执行成功 响应代码
	Data    bool   //响应结果数据
	Message string //仅执行出错时返回 响应异常信息
}

//erp/open/inventory/syn
//库存同步,对外开放

type InventorySynSkuInvsBatchInvs struct {
	BatchNo     string `json:"batch_no"`     //批次编码
	ExpiredDate string `json:"expired_date"` //过期日期
	Num         int    `json:"num"`          //数量
	ProduceDate string `json:"produce_date"` //成产日期
}

type InventorySynSkuInvs struct {
	Amount      float64                        `json:"amount"`       //数量
	BatchInvs   []InventorySynSkuInvsBatchInvs `json:"batch_invs"`   //批次集合,可为空
	SkuNo       string                         `json:"sku_no"`       //sku编码
	StorageCode string                         `json:"storage_code"` //仓库编码,如没传,则使用系统初始化仓库
}

type InventorySyn struct {
	SkuInvs []InventorySynSkuInvs `json:"sku_invs"` //库存信息
}

func NewInventorySynSkuInvsBatchInvs(
	BatchNo string,
	ExpiredDate string,
	Num int,
	ProduceDate string) InventorySynSkuInvsBatchInvs {

	return InventorySynSkuInvsBatchInvs{
		BatchNo:     BatchNo,
		ExpiredDate: ExpiredDate,
		Num:         Num,
		ProduceDate: ProduceDate,
	}
}

func NewInventorySynSkuInvs(
	Amount float64,
	BatchInvs []InventorySynSkuInvsBatchInvs,
	SkuNo string,
	StorageCode string) InventorySynSkuInvs {

	return InventorySynSkuInvs{
		Amount:      Amount,
		BatchInvs:   BatchInvs,
		SkuNo:       SkuNo,
		StorageCode: StorageCode,
	}
}

func NewInventorySyn(
	SkuInvs []InventorySynSkuInvs,
) []byte {
	ISyn = &InventorySyn{
		SkuInvs: SkuInvs,
	}

	b1, err := json.Marshal(ISyn)
	if err != nil {
		log.DLog.Println(err)
	}

	return b1
}

type RespInventorySyn struct {
	Code    int    //返回 0 表示执行成功 响应代码
	Data    bool   //响应结果数据
	Message string //仅执行出错时返回 响应异常信息
}

//erp/sn/querysnbybillcode
//查询完结单据中绑定的sn，1：订单，2：出库单：3：退货入库单

type QuerySnbyBillCode struct {
	BillCode string `json:"bill_code"` //单据编号
	BillType int    `json:"bill_type"` //单据类型
}

func NewQuerySnbyBillCode(
	BillCode string,
	BillType int,
) []byte {
	QSnbyBillCode = &QuerySnbyBillCode{
		BillCode: BillCode,
		BillType: BillType,
	}

	b1, err := json.Marshal(QSnbyBillCode)
	if err != nil {
		log.DLog.Println(err)
	}

	return b1
}

type RespQuerySnbyBillCode struct {
	Code    int    //返回 0 表示执行成功 响应代码
	Data    bool   //响应结果数据
	Message string //仅执行出错时返回 响应异常信息
}

//erp/stock/in/requestbill/close
//关闭其他入库订单,对外开放

type RequestBillClose struct {
	StockRequestCode string `json:"stock_request_code"` //单据编码
	NotCloseStocked  bool   `json:"not_close_stocked"`  //已入库不关闭
}

func NewRequestBillClose(
	StockRequestCode string,
	NotCloseStocked bool,
) []byte {
	RBillClose = &RequestBillClose{
		StockRequestCode: StockRequestCode,
		NotCloseStocked:  NotCloseStocked,
	}

	b1, err := json.Marshal(RBillClose)
	if err != nil {
		log.DLog.Println(err)
	}

	return b1
}

type RespRequestBillClose struct {
	Code    int    //返回 0 表示执行成功 响应代码
	Data    bool   //响应结果数据
	Message string //仅执行出错时返回 响应异常信息
}

//erp/stock/in/stockbill/add
//新增其他入库单,明细按照基本单位保存,对外开放

type StockBillAddBillDetails struct {
	Remark   string  `json:"remark"`    //备注
	Size     float64 `json:"size"`      //数量
	SpecCode string  `json:"spec_code"` //规格编码
	SumMoney float64 `json:"sum_money"` //明细总金额，此字段无效，其他出入库的明细成本，是以当前库存的成本来计算的
}

type StockBillAddBill struct {
	Reason      string                    `json:"reason"`       //业务原因
	Remark      string                    `json:"remark"`       //备注
	StorageCode string                    `json:"storage_code"` //仓库编码
	Details     []StockBillAddBillDetails `json:"details"`      //明细
}

type StockBillAdd struct {
	Bill StockBillAddBill `json:"bill"` //其他入库订单
}

func NewStockBillAddBillDetails(
	Remark string,
	Size float64,
	SpecCode string,
	SumMoney float64,
) StockBillAddBillDetails {

	return StockBillAddBillDetails{
		Remark:   Remark,
		Size:     Size,
		SpecCode: SpecCode,
		SumMoney: SumMoney,
	}
}

func NewStockBillAdd(
	Reason string,
	Remark string,
	StorageCode string,
	Details []StockBillAddBillDetails,
) []byte {
	SBillAdd = &StockBillAdd{
		Bill: StockBillAddBill{
			Details:     Details,
			Reason:      Reason,
			Remark:      Remark,
			StorageCode: StorageCode,
		},
	}

	b1, err := json.Marshal(SBillAdd)
	if err != nil {
		log.DLog.Println(err)
	}

	return b1
}

type RespStockBillAdd struct {
	Code    int    //返回 0 表示执行成功 响应代码
	Data    bool   //响应结果数据
	Message string //仅执行出错时返回 响应异常信息
}

//erp/allocation/changebill/query
//查询调拨单,bill_code modify_time 不能同时为空,对外开放

type ChangeBillQuery struct {
	BillCode   string `json:"bill_code"`   //单据编码
	CreateTime string `json:"create_time"` //创建时间，只能查近3个月
	Page       int    `json:"page"`        //当前页码，从1开始
	Limit      int    `json:"limit"`       //每页大小,最大200
}

func NewChangeBillQuery(
	BillCode string,
	CreateTime string,
	Page int,
	Limit int,
) []byte {

	CBillQuery = &ChangeBillQuery{
		BillCode:   BillCode,
		CreateTime: CreateTime,
		Page:       Page,
		Limit:      Limit,
	}

	b1, err := json.Marshal(CBillQuery)
	if err != nil {
		log.DLog.Println(err)
	}

	return b1
}

type RespChangeBillQueryDataDetails struct {
	GoodsName    string  `json:"goods_name"`    //商品名称
	Index        int64   `json:"index"`         //行号
	Iums         int     `json:"nums"`          //数量
	Price        float64 `json:"price"`         //单价
	ReceiveNums  int     `json:"receive_nums"`  //	入库数量
	Remark       string  `json:"remark"`        //备注
	SpecCode     string  `json:"spec_code"`     //规格编码
	SpecName     string  `json:"spec_name"`     //规格名称
	StockoutNums int     `json:"stockout_nums"` //出库数量
	Uznit        string  `json:"unit"`          //单位
}

type RespChangeBillQueryData struct {
	BillCode       string                           `json:"bill_code"`        //单据编码
	BillCreater    string                           `json:"bill_creater"`     //创建者
	BillDate       string                           `json:"bill_date"`        //业务日期
	BillType       int                              `json:"bill_code"`        //调拨类型: 1. 虚拟调拨 2. 实际调拨
	CreateTime     string                           `json:"bill_type"`        //创建时间
	Details        []RespChangeBillQueryDataDetails `json:"details"`          //明细
	Remark         string                           `json:"remark"`           //备注
	StorageInCode  string                           `json:"storage_in_code"`  //入库仓库编码
	StorageInName  string                           `json:"storage_in_name"`  //入库仓库名称
	StorageOutCode string                           `json:"storage_out_code"` //出库仓库编码
	StorageOutName string                           `json:"storage_out_name"` //出库仓库名称
}

type RespChangeBillQuery struct {
	Code    int                       `json:"code"`    //返回 0 表示执行成功 响应代码
	Data    []RespChangeBillQueryData `json:"data"`    //响应结果数据
	Message string                    `json:"message"` //仅执行出错时返回 响应异常信息
}

//erp/allocation/in/changebill/query
//调拨入库单查询,bill_code,stock_code,modify_time不能同时为空,对外开放

type InChangeBillQuery struct {
	BillCode   string `json:"bill_code"`   //调拨单编码
	StockCode  string `json:"stock_code"`  //调拨入库单编码
	ModifyTime string `json:"modify_time"` //修改时间，只能查近3个月
	Page       int    `json:"page"`        //当前页码，从1开始
	Limit      int    `json:"limit"`       //每页大小,最大200
}

func NewInChangeBillQuery(
	BillCode string,
	StockCode string,
	ModifyTime string,
	Page int,
	Limit int,
) []byte {

	ICBillQuery = &InChangeBillQuery{
		BillCode:   BillCode,
		StockCode:  StockCode,
		ModifyTime: ModifyTime,
		Page:       Page,
		Limit:      Limit,
	}

	b1, err := json.Marshal(ICBillQuery)
	if err != nil {
		log.DLog.Println(err)
	}
	return b1
}

type RespInChangeBillQueryDataDetails struct {
	GoodsName string `json:"goods_name"` //商品名称
	Index     int64  `json:"index"`      //行号
	Nums      int    `json:"nums"`       //数量
	Remark    string `json:"remark"`     //备注
	SpecCode  string `json:"spec_code"`  //规格编码
	SpecName  string `json:"spec_name"`  //规格名称
	Unit      string `json:"unit"`       //单位
}

type RespInChangeBillQueryData struct {
	BillCreater             string                             `json:"bill_creater"`               //创建者
	BillDate                string                             `json:"bill_date"`                  //业务日期
	CreateTime              string                             `json:"create_time"`                //创建时间
	Details                 []RespInChangeBillQueryDataDetails `json:"details"`                    //明细
	InventoryChangeBillCode string                             `json:"inventory_change_bill_code"` //调拨单号
	ModifiedTime            string                             `json:"modified_time"`              //单据修改时间
	OperateName             string                             `json:"operate_name"`               //经手人
	Remark                  string                             `json:"remark"`                     //备注
	StockCode               string                             `json:"stock_code"`                 //出入库单编码
	StorageCode             string                             `json:"storage_code"`               //仓库编码
	StorageName             string                             `json:"storage_name"`               //仓库名称
}

type RespInChangeBillQuery struct {
	Code    int                         `json:"code"`    //返回 0 表示执行成功 响应代码
	Data    []RespInChangeBillQueryData `json:"data"`    //响应结果数据
	Message string                      `json:"message"` //仅执行出错时返回 响应异常信息
}

//erp/stock/in/stockbill/query
//其他入库单查询,stock_request_code,stock_bill_code,modify_time不能同时为空,对外开放

type InStockBillQuery struct {
	StockRequestCode string `json:"stock_request_code"` //其他入库订单单据编码
	StockCode        string `json:"stock_code"`         //其他入库单单据编码
	ModifyTime       string `json:"modify_time"`        //修改时间，只能查近3个月
	Page             int    `json:"page"`               //当前页码，从1开始
	Limit            int    `json:"limit"`              //每页大小,最大200
}

func NewInStockBillQuery(
	StockRequestCode string,
	StockCode string,
	ModifyTime string,
	Page int,
	Limit int,

) []byte {

	ISBillQuery = &InStockBillQuery{
		StockRequestCode: StockRequestCode,
		StockCode:        StockCode,
		ModifyTime:       ModifyTime,
		Page:             Page,
		Limit:            Limit,
	}

	b1, err := json.Marshal(ISBillQuery)
	if err != nil {
		log.DLog.Println(err)
	}
	return b1
}

type RespInStockBillQueryDataDetails struct {
	GoodsName   string  `json:"goods_name"`  //	商品名称
	Index       int64   `json:"index"`       //行号
	Nums        int     `json:"nums"`        //数量
	Price       float64 `json:"price"`       //单价
	Remark      string  `json:"remark"`      //备注
	Spec_code   string  `json:"spec_code"`   //规格编码
	Spec_name   string  `json:"spec_name"`   //规格名称
	Total_money float64 `json:"total_money"` //明细总价
	Unit        string  `json:"unit"`        //单位
}

type RespInStockBillQueryData struct {
	BillCreater      string                            `json:"bill_creater"`        //创建员
	BillDate         string                            `json:"bill_date"`           //业务日期
	BillType         int                               `json:"bill_type"`           //单据类型
	CreateTime       string                            `json:"create_time"`         //创建时间
	Details          []RespInStockBillQueryDataDetails `json:"details"`             //明细
	LogisticName     string                            `json:"logistic_name"`       //快递名称
	ModifiedTime     string                            `json:"modified_time"`       //单据修改时间
	OperateName      string                            `json:"operate_name"`        //经手人
	Reason           string                            `json:"reason"`              //原因
	Remark           string                            `json:"remark"`              //备注
	StockCode        string                            `json:"stock_code"`          //出入库单编码
	StockReqBillCode string                            `json:"stock_req_bill_code"` //其它出入库预约单号
	StorageCode      string                            `json:"storage_code"`        //仓库编码
	StorageName      string                            `json:"storage_name"`        //仓库名称
}

type RespInStockBillQuery struct {
	Code    int                        `json:"code"`    //返回 0 表示执行成功 响应代码
	Data    []RespInStockBillQueryData `json:"data"`    //响应结果数据
	Message string                     `json:"message"` //仅执行出错时返回响应异常信息
}

//erp/stock/out/stockbill/query
//其他出库单查询,stock_request_code,stock_code,modify_time不能同时为空,对外开放

type OutStockBillQuery struct {
	StockRequestCode string `json:"stock_request_code"` //其他出库订单单据编码
	StockCode        string `json:"stock_code"`         //其他出库单单据编码
	ModifyTime       string `json:"modify_time"`        //修改时间，只能查近3个月
	Page             int    `json:"page"`               //当前页码，从1开始
	Limit            int    `json:"limit"`              //每页大小,最大200
}

func NewOutStockBillQuery(
	StockRequestCode string,
	StockCode string,
	ModifyTime string,
	Page int,
	Limit int,
) []byte {

	OSBillQuery = &OutStockBillQuery{
		StockRequestCode: StockRequestCode,
		StockCode:        StockCode,
		ModifyTime:       ModifyTime,
		Page:             Page,
		Limit:            Limit,
	}

	b1, err := json.Marshal(OSBillQuery)
	if err != nil {
		log.DLog.Println(err)
	}
	return b1
}

type RespOutStockBillQueryDataDetails struct {
	GoodsName  string  `json:"goods_name"`  //商品名称
	Index      int64   `json:"index"`       //行号
	Nums       int     `json:"nums"`        //数量
	Price      float64 `json:"price"`       //单价
	Remark     string  `json:"remark"`      //备注
	SpecCode   string  `json:"spec_code"`   //规格编码
	SpecName   string  `json:"spec_name"`   //规格名称
	TotalMoney float64 `json:"total_money"` //明细总价
	Unit       string  `json:"unit"`        //单位
}

type RespOutStockBillQueryData struct {
	BillCreater      string                             `json:"bill_creater"`        //创建员
	BillDate         string                             `json:"bill_date"`           //业务日期
	BillType         int                                `json:"bill_type"`           //单据类型
	CreateTime       string                             `json:"create_time"`         //创建时间
	Details          []RespOutStockBillQueryDataDetails `json:"details"`             //明细
	LogisticName     string                             `json:"logistic_name"`       //快递名称
	ModifiedTime     string                             `json:"modified_time"`       //单据修改时间
	OperateName      string                             `json:"operate_name"`        //经手人
	Reason           string                             `json:"reason"`              //原因
	Remark           string                             `json:"remark"`              //备注
	StockCode        string                             `json:"stock_code"`          //出入库单编码
	StockReqBillCode string                             `json:"stock_req_bill_code"` //其它出入库预约单号
	StorageCode      string                             `json:"storage_code"`        //仓库编码
	StorageName      string                             `json:"storage_name"`        //仓库名称
}

type RespOutStockBillQuery struct {
	Code    int                         `json:"code"`    //返回 0 表示执行成功响应代码
	Data    []RespOutStockBillQueryData `json:"data"`    //响应结果数据
	Message string                      `json:"message"` //仅执行出错时返回 响应异常信息
}

//erp/stock/out/requestbill/close
//关闭其他出库订单,对外开放

type OutRequestBillClose struct {
	StockRequestCode string `json:"stock_request_code"` //单据编码
	NotCloseStocked  bool   `json:"not_close_stocked"`  //已出库不关闭
}

func NewOutRequestBillClose(
	StockRequestCode string,
	NotCloseStocked bool,
) []byte {

	ORBillClose = &OutRequestBillClose{
		StockRequestCode: StockRequestCode,
		NotCloseStocked:  NotCloseStocked,
	}

	b1, err := json.Marshal(ORBillClose)
	if err != nil {
		log.DLog.Println(err)
	}
	return b1
}

type RespOutRequestBillClose struct {
	Code    int    `json:"code"`    //返回 0 表示执行成功响应代码
	Data    bool   `json:"data"`    //响应结果数据
	Message string `json:"message"` //仅执行出错时返回 响应异常信息
}

//erp/stock/out/requestbill/query
//查询其他出库订单,stock_request_code modify_time 不能同时为空,对外开放
//2020-01-03 18:32:22

type OutRequestBillQuery struct {
	StockRequestCode string `json:"stock_request_code"` //单据编码
	ModifyTime       string `json:"modify_time"`        //修改时间，只能查近3个月
	Page             int    `json:"page"`               //当前页码，从1开始
	Limit            int    `json:"limit"`              //每页大小,最大200
}

func NewOutRequestBillQuery(
	StockRequestCode string,
	ModifyTime string,
	Page int,
	Limit int,
) []byte {

	ORequestBillQuery = &OutRequestBillQuery{
		StockRequestCode: StockRequestCode,
		ModifyTime:       ModifyTime,
		Page:             Page,
		Limit:            Limit,
	}

	b1, err := json.Marshal(ORequestBillQuery)
	if err != nil {
		log.DLog.Println(err)
	}
	return b1
}

type RespOutRequestBillQueryDataDetails struct {
	GoodsName string `json:"goods_name"` //商品名称
	Index     int64  `json:"index"`      //行号
	Remark    string `json:"remark"`     //备注
	Size      int    `json:"size"`       //数量
	SpecCode  string `json:"spec_code"`  //规格编码
	SpecName  string `json:"spec_name"`  //规格名称
	Status    int    `json:"status"`     //状态 1:未到货 2:部分到货 3:全部到货 4:已关闭
	Unit      string `json:"unit"`       //单位，默认erp的基本单位
}

type RespOutRequestBillQueryData struct {
	Address      string                               `json:"address"`       //地址
	BillCreater  string                               `json:"bill_creater"`  //创建人
	BillDate     string                               `json:"bill_date"`     //业务日期
	City         string                               `json:"city"`          //市
	CreateTime   string                               `json:"create_time"`   //创建时间
	Details      []RespOutRequestBillQueryDataDetails `json:"details"`       //明细
	LogisticName string                               `json:"logistic_name"` ///快递名称
	Mobile       string                               `json:"mobile"`        //手机
	ModifiedTime string                               `json:"modified_time"` //单据修改时间
	OperateName  string                               `json:"operate_name"`  //经手人
	Province     string                               `json:"province"`      ////省
	Reason       string                               `json:"reason"`        //原因
	Receiver     string                               `json:"receiver"`      //收货人
	Remark       string                               `json:"remark"`        //备注
	Status       int                                  `json:"status"`        //状态
	StockCode    string                               `json:"stock_code"`    //出入库预约单编码
	StorageCode  string                               `json:"storage_code"`  //仓库编码
	StorageName  string                               `json:"storage_name"`  //仓库名称
	Tel          string                               `json:"tel"`           //电话
}

type RespOutRequestBillQuery struct {
	Code    int                           `json:"code"`    //返回 0 表示执行成功响应代码
	Data    []RespOutRequestBillQueryData `json:"data"`    //响应结果数据
	Message string                        `json:"message"` //仅执行出错时返回 响应异常信息
}

//erp/open/inventory/items/get/by/modifytime //fuck
//查询库存,规格编码和修改时间不能都为空,对外开放

//erp/allocation/out/changebill/query
//调拨出库单查询,bill_code,stock_code,modify_time不能同时为空,对外开放

type OutChangeBillQuery struct {
	BillCode   string `json:"bill_code"`   //调拨单编码码
	StockCode  string `json:"stock_code"`  //调拨出库单编码
	ModifyTime string `json:"modify_time"` //修改时间，只能查近3个月
	Page       int    `json:"page"`        //当前页码，从1开始
	Limit      int    `json:"limit"`       //每页大小, 最大200
}

func NewOutChangeBillQuery(
	BillCode string,
	StockCode string,
	ModifyTime string,
	Page int,
	Limit int,
) []byte {

	OChangeBillQuery = &OutChangeBillQuery{
		BillCode:   BillCode,
		StockCode:  StockCode,
		ModifyTime: ModifyTime,
		Page:       Page,
		Limit:      Limit,
	}

	b1, err := json.Marshal(OChangeBillQuery)
	if err != nil {
		log.DLog.Println(err)
	}
	return b1
}

type RespOutChangeBillQueryDataDetails struct {
	GoodsName string `json:"goods_name"` //商品名称
	Index     int64  `json:"index"`      //行号
	Remark    string `json:"remark"`     //备注
	Nums      int    `json:"nums"`       //数量
	SpecCode  string `json:"spec_code"`  //规格编码
	SpecName  string `json:"spec_name"`  //规格名称
	Unit      string `json:"unit"`       //单位
}

type RespOutChangeBillQueryData struct {
	BillCreater             string                              `json:"bill_creater"`               //创建人
	BillDate                string                              `json:"bill_date"`                  //业务日期
	CreateTime              string                              `json:"create_time"`                //创建时间
	Details                 []RespOutChangeBillQueryDataDetails `json:"details"`                    //明细
	InventoryChangeBillCode string                              `json:"inventory_change_bill_code"` ///快递名称
	ModifiedTime            string                              `json:"modified_time"`              //单据修改时间
	OperateName             string                              `json:"operate_name"`               //经手人
	Remark                  string                              `json:"remark"`                     //备注
	StockCode               string                              `json:"stock_code"`                 //出入库预约单编码
	StorageCode             string                              `json:"storage_code"`               //仓库编码
	StorageName             string                              `json:"storage_name"`               //仓库名称
}
type RespOutChangeBillQuery struct {
	CodeResp int                          `json:"codeResp"` //返回 0 表示执行成功响应代码
	Data     []RespOutChangeBillQueryData `json:"data"`     //响应结果数据
	Message  string                       `json:"message"`  //仅执行出错时返回 响应异常信息
}

//erp/stock/out/stockbill/add
//新增其他出库单,明细按照基本单位保存,对外开放

type OutStockBillAddBillDetails struct {
	Remark   string  `json:"remark"`    //备注
	Size     float64 `json:"size"`      //数量
	SpecCode string  `json:"spec_code"` //规格编码
	SumMoney float64 `json:"sum_money"` //明细总金额，此字段无效，其他出入库的明细成本，是以当前库存的成本来计算的
}

type OutStockBillAddBill struct {
	Details     []OutStockBillAddBillDetails `json:"details"`      //明细
	Reason      string                       `json:"reason"`       //业务原因
	Remark      string                       `json:"remark"`       //备注
	StorageCode string                       `json:"storage_code"` //仓库编码
}

type OutStockBillAdd struct {
	Bill OutStockBillAddBill `json:"bill"` //其他出库订单
}

func NewOutStockBillAddBillDetails(
	Remark string,
	Size float64,
	SpecCode string,
	SumMoney float64,
) OutStockBillAddBillDetails {

	return OutStockBillAddBillDetails{
		Remark:   Remark,
		Size:     Size,
		SpecCode: SpecCode,
		SumMoney: SumMoney,
	}
}

func NewOutStockBillAdd(
	Details []OutStockBillAddBillDetails,
	Reason string,
	Remark string,
	StorageCode string,
) []byte {

	OStockBillAdd = &OutStockBillAdd{
		Bill: OutStockBillAddBill{
			Details:     Details,
			Reason:      Reason,
			Remark:      Remark,
			StorageCode: StorageCode,
		},
	}

	b1, err := json.Marshal(OStockBillAdd)
	if err != nil {
		log.DLog.Println(err)
	}
	return b1
}

type RespOutStockBillAdd struct {
	CodeResp int    `json:"codeResp"` //返回 0 表示执行成功响应代码
	Data     bool   `json:"data"`     //响应结果数据
	Message  string `json:"message"`  //仅执行出错时返回 响应异常信息
}

var RBillAdd *RequestBillAdd
var SNQTrace *SNQueryTrace
var InRBillAdd *InRequestBillAdd
var CHBillAdd *CHangeBillAdd
var InRBillQuery *InRequestBillQuery
var IGetByModifyTime *ItemsGetByModifyTime
var CBillClose *ChangeBillClose
var ISyn *InventorySyn
var QSnbyBillCode *QuerySnbyBillCode
var RBillClose *RequestBillClose
var SBillAdd *StockBillAdd
var CBillQuery *ChangeBillQuery
var ICBillQuery *InChangeBillQuery
var ISBillQuery *InStockBillQuery
var OSBillQuery *OutStockBillQuery
var ORBillClose *OutRequestBillClose
var ORequestBillQuery *OutRequestBillQuery
var OChangeBillQuery *OutChangeBillQuery
var OStockBillAdd *OutStockBillAdd
