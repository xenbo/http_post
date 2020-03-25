package sale

import (
	"encoding/json"
	"github.com/xenbo/http_post/log"
)

//erp/open/return/order/query
//查询售后单,bill_code,create_time,modify_time 不能全为空,对外开放
type ReturnOrderQuery struct {
	BillCode    string `json:"bill_code"`    //单据编码
	StartTime   string `json:"start_time"`   //开始时间
	EndTime     string `json:"end_time"`     //结束时间
	TimeType    int    `json:"time_type"`    //时间类型 1:创建时间 2:处理时间 3:完成时间 默认为1
	StorageCode string `json:"storage_code"` //仓库编码
	Page        int    `json:"page"`         //当前页码，从1开始
	Limit       int    `json:"limit"`        //每页大小, 最大200
	ShopNick    string `json:"shop_nick"`    //店铺昵称
}

func NewReturnOrderQuery(
	BillCode string,
	StartTime string,
	EndTime string,
	TimeType int,
	StorageCode string,
	Page int,
	Limit int,
	ShopNick string,
) []byte {

	ROrderQuery = &ReturnOrderQuery{
		BillCode:    BillCode,
		StartTime:   StartTime,
		EndTime:     EndTime,
		TimeType:    TimeType,
		StorageCode: StorageCode,
		Page:        Page,
		Limit:       Limit,
		ShopNick:    ShopNick,
	}

	b1, err := json.Marshal(ROrderQuery)
	if err != nil {
		log.DLog.Println(err)
	}
	return b1

}

type RespReturnOrderQueryDataItems struct {
	IsPackage    bool    `json:"is_package"`     //是否组合商品
	OlnTradeCode string  `json:"oln_trade_code"` //线上单号
	Payment      float64 `json:"payment"`        //售后明细总金额
	Size         int     `json:"size"`           //数量
	SkuCode      string  `json:"sku_code"`       //规格编码
}

type RespReturnOrderQueryData struct {
	Address       string                          `json:"address"`         //买家地址
	BillCode      string                          `json:"bill_code"`       //售后单号
	Buyer         string                          `json:"buyer"`           //买家昵称
	BuyerName     string                          `json:"buyer_name"`      //买家姓名
	CreateTime    string                          `json:"create_time"`     //开始时间
	CreateUser    string                          `json:"create_user"`     //申请人
	Describe      string                          `json:"describe"`        //问题详情
	EndTime       string                          `json:"end_time"`        //完成时间
	ExpressCode   string                          `json:"express_code"`    //快递单号
	Items         []RespReturnOrderQueryDataItems `json:"items"`           //商品明细集
	LogisticCode  string                          `json:"logistic_code"`   //退回快递
	OlnReturnCode string                          `json:"oln_return_code"` //线上售后单号
	OlnTradeCode  string                          `json:"oln_trade_code"`  //线上订单号
	Phone         string                          `json:"phone"`           //买家手机号
	Reason        string                          `json:"reason"`          //退回原因
	ShopNick      string                          `json:"shop_nick"`       //店铺昵称
	Status        int                             `json:"status"`          //状态: -2:待审核 -1：待审核 0：处理中 1：完成 2：关闭 3：拒绝 4：取消
	StorageCode   string                          `json:"storage_code"`    //仓库编码
	StorageName   string                          `json:"storage_name"`    //仓库名称
	TotalPay      float64                         `json:"total_pay"`       //退款金额
	TradeCode     string                          `json:"trade_code"`      //系统订单号
	Type1         int                             `json:"type"`            //类型 0:退货 1:补发 2:换货(万里牛的换货分为2步，先退货，再发货，所以换货单中没有需要发的明细，会生成销售订单) 3:其他 4:仅退款
}

type RespReturnOrderQuery struct {
	Code    int                        `json:"code"`    //返回 0 表示执行成功响应代码
	Data    []RespReturnOrderQueryData `json:"data"`    //响应结果数据
	Message string                     `json:"message"` //仅执行出错时返回 响应异常信息
}

//erp/sale/stock/in/add
//添加退货入库单单（不绑定单号）,对外开放

type StockInAddBillDetails struct {
	Nums    int     `json:"nums"`     //数量
	SkuNo   string  `json:"sku_no"`   //sku编码（万里牛编码）
	SumSale float64 `json:"sum_sale"` //总售价,包含优惠
}

type StockInAddBill struct {
	BillDate     string                  `json:"bill_date"`     //单据日期
	CreateTime   string                  `json:"create_time"`   //创建时间
	CustomerNick string                  `json:"customer_nick"` //客户昵称
	Details      []StockInAddBillDetails `json:"details"`       //明细集合,以万里牛商品的基本单位保存
	DiscountFee  float64                 `json:"discount_fee"`  //优惠金额
	PaidFee      float64                 `json:"paid_fee"`      //实际支付金额
	PostFee      float64                 `json:"post_fee"`      //邮费
	Remark       string                  `json:"remark"`        //备注
	ServiceFee   float64                 `json:"service_fee"`   //服务费
	ShopNick     string                  `json:"shop_nick"`     //店铺昵称,对应万里牛erp中的店铺账号ID
	StorageCode  string                  `json:"storage_code"`  //仓库编码
	SumSale      float64                 `json:"sum_sale"`      //总金额,包含优惠
}

type StockInAdd struct {
	Bill StockInAddBill `json:"bill"` //销售出库单
}

func NewStockInAddBillDetails(
	Nums int,
	SkuNo string,
	SumSale float64,
) StockInAddBillDetails {

	return StockInAddBillDetails{
		Nums:    Nums,
		SkuNo:   SkuNo,
		SumSale: SumSale,
	}
}

func NewStockInAdd(
	BillDate string,
	CreateTime string,
	CustomerNick string,
	Details []StockInAddBillDetails,
	DiscountFee float64,
	PaidFee float64,
	PostFee float64,
	Remark string,
	ServiceFee float64,
	ShopNick string,
	StorageCode string,
	SumSale float64,
) []byte {

	SInAdd = &StockInAdd{
		Bill: StockInAddBill{
			BillDate:     BillDate,
			CreateTime:   CreateTime,
			CustomerNick: CustomerNick,
			Details:      Details,
			DiscountFee:  DiscountFee,
			PaidFee:      PaidFee,
			PostFee:      PostFee,
			Remark:       Remark,
			ServiceFee:   ServiceFee,
			ShopNick:     ShopNick,
			StorageCode:  StorageCode,
			SumSale:      SumSale,
		},
	}

	b1, err := json.Marshal(SInAdd)
	if err != nil {
		log.DLog.Println(err)
	}
	return b1

}

type RespStockInAdd struct {
	Code    int    `json:"code"`    //返回 0 表示执行成功响应代码
	Data    string `json:"data"`    //响应结果数据
	Message string `json:"message"` //仅执行出错时返回 响应异常信息
}

//erp/sale/stock/in/query
//查询退后入库单,bill_code modify_time 不能同时为空,只能查询modify_time近7天的单据,对外开放

type StockInQuery struct {
	BillCode   string `json:"bill_code"`   //单据编码
	ModifyTime string `json:"modify_time"` //修改时间，只能查近3个月
	Page       int    `json:"page"`        //当前页码，从1开始
	Limit      int    `json:"limit"`       //每页大小, 最大200
	IsSplit    bool   `json:"is_split"`    //是否拆分组合商品
}

func NewStockInQuery(
	BillCode string,
	ModifyTime string,
	Page int,
	Limit int,
	IsSplit bool,
) []byte {
	SInQuery = &StockInQuery{
		BillCode:   BillCode,
		ModifyTime: ModifyTime,
		Page:       Page,
		Limit:      Limit,
		IsSplit:    IsSplit,
	}

	b1, err := json.Marshal(SInQuery)
	if err != nil {
		log.DLog.Println(err)
	}
	return b1
}

type RespStockInQueryDataDetailsBatchInfos struct {
	BatchNo string `json:"batch_no"` //批次号
	Nums    int    `json:"nums"`     //数量
}

type RespStockInQueryDataDetails struct {
	BatchInfos  []RespStockInQueryDataDetailsBatchInfos `json:"batch_infos"`  //批次号集合
	DetailId    string                                  `json:"detail_id"`    //明细id，如果是组合商品拆分出来的明细，这个id相同
	DiscountFee float64                                 `json:"discount_fee"` //优惠
	GoodsName   string                                  `json:"goods_name"`   //商品名称
	Nums        int                                     `json:"nums"`         //数量
	SkuName     string                                  `json:"sku_name"`     //sku名称
	SkuProp1    string                                  `json:"sku_prop1"`    //规格扩展属性，商品未删除，且设置了会返回
	SkuProp2    string                                  `json:"sku_prop2"`    //规格扩展属性，商品未删除，且设置了会返回
	SnInfos     []string                                `json:"sn_infos"`     //序列号集合，此字段以不返回数据，请使接口——/erp/sn/querysnbybillcode查询单据中的序列号
	SumCost     float64                                 `json:"sum_cost"`     //成本
	SumSale     float64                                 `json:"sum_sale"`     //总售价,包含优惠
	unit        string                                  `json:"unit"`         //单位
}

type RespStockInQueryData struct {
	BillDate             string                        `json:"bill_date"`               //单据日期
	BillType             int                           `json:"bill_type"`               //单据类型: 1: 销售出库单 2： 线下出库单
	Country              string                        `json:"country"`                 //国家
	CreateTime           string                        `json:"create_time"`             //创建时间
	CustomCode           string                        `json:"custom_code"`             //客户编码
	CustomName           string                        `json:"custom_name"`             //客户名称
	CustomerNick         string                        `json:"customer_nick"`           //客户昵称
	CustomerNickType     int                           `json:"customer_nick_type"`      //客户来源平台值
	CustomerNickTypeName string                        `json:"customer_nick_type_name"` //客户来源平台名称 比如:淘宝, 京东
	Details              []RespStockInQueryDataDetails `json:"details"`                 //明细集合
	DiscountFee          float64                       `json:"discount_fee"`            //优惠金额
	FromTradeNo          string                        `json:"from_trade_no"`           //万里牛中的原始交易编号
	InvNo                string                        `json:"inv_no"`                  //单据编号
	PaidFee              float64                       `json:"paid_fee"`                //实际支付金额
	PayType              int                           `json:"pay_type"`                //付款方式 0 线上支付 1 现金 2 刷卡 3 微信 4 支付宝 5 转账 6 储值支付 7 1688支付 52 其它转账支付
	PostFee              float64                       `json:"post_fee"`                //邮费
	Remark               string                        `json:"remark"`                  //备注
	SaleMan              string                        `json:"sale_man"`                //业务员
	ServiceFee           float64                       `json:"service_fee"`             //服务费
	ShopName             string                        `json:"shop_name"`               //店铺名称
	ShopNick             string                        `json:"shop_nick"`               //店铺昵称,对应万里牛erp中的店铺账号ID
	ShopSource           string                        `json:"shop_source"`             //店铺来源
	StorageCode          string                        `json:"storage_code"`            //仓库编码
	StorageName          string                        `json:"storage_name"`            //仓库名称
	SumSale              float64                       `json:"sum_sale"`                //总金额,包含优惠
	TpTid                string                        `json:"tp_tid"`                  //外部订单号(如淘宝/京东单号)
}

type RespStockInQuery struct {
	Code    int                    `json:"code"`    //返回 0 表示执行成功响应代码
	Data    []RespStockInQueryData `json:"data"`    //响应结果数据
	Message string                 `json:"message"` //仅执行出错时返回 响应异常信息
}

//erp/open/return/order/stock/in
//通过售后单号退货入库,支持

type OrderStockInItems struct {
	Size    int    `json:"size"`     //本次入库数量
	SkuCode string `json:"sku_code"` //规格编码
}
type OrderStockIn struct {
	BillCode string              `json:"bill_code"` //售后单号
	Remark   string              `json:"remark"`    //备注
	Items    []OrderStockInItems `json:"items"`     //商品明细
}

func NewOrderStockInItems(
	Size int,
	SkuCode string,
) OrderStockInItems {
	return OrderStockInItems{
		Size:    Size,
		SkuCode: SkuCode,
	}
}

func NewOrderStockIn(
	BillCode string,
	Remark string,
	Items []OrderStockInItems,
) []byte {
	OStockIn = &OrderStockIn{
		BillCode: BillCode,
		Remark:   Remark,
		Items:    Items,
	}

	b1, err := json.Marshal(OStockIn)
	if err != nil {
		log.DLog.Println(err)
	}
	return b1
}

type RespOrderStockIn struct {
	Code    int    `json:"code"`    //返回 0 表示执行成功响应代码
	Data    bool   `json:"data"`    //响应结果数据
	Message string `json:"message"` //仅执行出错时返回 响应异常信息
}

var ROrderQuery *ReturnOrderQuery
var SInAdd *StockInAdd
var SInQuery *StockInQuery
var OStockIn *OrderStockIn
