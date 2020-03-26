package sale

import (
	"encoding/json"
	"github.com/xenbo/http_post/log"
)

//erp/opentrade/query/trades
//查询订单,bill_code,create_time,modify_time,send_goods_time,finish_time 不能全为空,对外开放

type QueryTrades struct {
	BillCode      string `json:"bill_code"`       //单据编码
	CreateTime    string `json:"create_time"`     //单据创建时间,修改时间为空是使用,仅支持近3个月的日期查询
	ModifyTime    string `json:"modify_time"`     //单据修改时间,仅支持近3个月的日期查询
	SendGoodsTime string `json:"send_goods_time"` //发货时间，创建时间为空时使用,仅支持近3个月的日期查询
	FinishTime    string `json:"finish_time"`     //完成时间，发货时间为空时使用,仅支持近3个月的日期查询
	EndTime       string `json:"end_time"`        //查询结束时间,默认是当前时间
	PayTime       string `json:"pay_time"`        //付款时间，仅支持近3个月的日期查询
	StorageCode   string `json:"storage_code"`    //仓库编码
	Page          int    `json:"page"`            //当前页码，从1开始
	Limit         int    `json:"limit"`           //每页大小,最大200
	UsSplit       bool   `json:"is_split"`        //是否拆分组合商品
	Mobile        string `json:"mobile"`          //手机
	TradeStatus   string `json:"trade_status"`    //订单状态0:审核 1：打单配货 2：验货3：称重4：待发货5：财审8：已发货9：完成10：关闭13：配货中15：打包 PS:分销商订单不支持状态查询
}

func NewQueryTrades(
	BillCode string,
	CreateTime string,
	ModifyTime string,
	SendGoodsTime string,
	FinishTime string,
	EndTime string,
	PayTime string,
	StorageCode string,
	Page int,
	Limit int,
	UsSplit bool,
	Mobile string,
	TradeStatus string,
) []byte {
	QTrades = &QueryTrades{
		BillCode:      BillCode,
		CreateTime:    CreateTime,
		ModifyTime:    ModifyTime,
		SendGoodsTime: SendGoodsTime,
		FinishTime:    FinishTime,
		EndTime:       EndTime,
		PayTime:       PayTime,
		StorageCode:   StorageCode,
		Page:          Page,
		Limit:         Limit,
		UsSplit:       UsSplit,
		Mobile:        Mobile,
		TradeStatus:   TradeStatus,
	}

	b1, err := json.Marshal(QTrades)
	if err != nil {
		log.DLog.Println(err)
	}
	return b1
}

type RespQueryTradesDataInvoice struct {
	Header        string `json:"header"`         //发票抬头
	InvoiceDetail string `json:"invoice_detail"` //发票详情
	InvoiceTaxId  string `json:"invoice_tax_id"` //税号
	Type1         int    `json:"type"`           //类型:1=普通发票，2=增值税普通发票, 3=电子增票
}

type RespQueryTradesDataOrders struct {
	CurrencyCode string  `json:"currency_code"` //原始货币种类
	CurrencySum  float64 `json:"currency_sum"`  //原始货币金额
	HasRefund    int     `json:"has_refund"`    //是否退款
	IsGift       int     `json:"is_gift"`       //明细是否赠品
	IsPackage    bool    `json:"is_package"`    //是否组合商品
	ItemName     string  `json:"item_name"`     //商品名称
	OlnItemId    string  `json:"oln_item_id"`   //线上商品id
	OlnItemName  string  `json:"oln_item_name"` //线上商品名称
	OlnSkuId     string  `json:"oln_sku_id"`    //线上规格id
	OlnStatus    int     `json:"oln_status"`    //线上状态:1:等待付款 2:等待发货 ,部分发货 3:已完成 4:已关闭 5: 等待确认 6:已签收 0: 未建交易
	OrderId      string  `json:"order_id"`      //明细id，单据级唯一
	Payment      float64 `json:"payment"`       //销售金额
	Price        float64 `json:"price"`         //单价(商品标价)
	Receivable   float64 `json:"receivable"`    //应收
	Remark       string  `json:"remark"`        //明细备注
	Size         int     `json:"size"`          //数量
	SkuCode      string  `json:"sku_code"`      //规格编码
	SnValue      string  `json:"sn_value"`      //序列号，此字段以不返回数据，请使接口——/erp/sn/querysnbybillcode查询单据中的序列号
	Status       int     `json:"status"`        //状态:1:等待付款 2:等待发货 ,部分发货 3:已完成 4:已关闭 5: 等待确认 6:已签收 0: 未建交易
	SubOrderNo   string  `json:"sub_order_no"`  //子订单号
	TpOids       string  `json:"tp_oids"`       //线上明细ID
	TpTid        string  `json:"tp_tid"`        //线上单号
}

type RespQueryTradesData struct {
	Address          string                       `json:"address"`       //地址
	Buyer            string                       `json:"buyer"`         //买家昵称
	BuyerAccount     string                       `json:"buyer_account"` //账号址
	BuyerMobile      string                       `json:"buyer_mobile"`
	BuyerMsg         string                       `json:"buyer_msg"`          //买家留言
	City             string                       `json:"city"`               //市
	Country          string                       `json:"country"`            //国家
	CreateTime       string                       `json:"create_time"`        //创建时间
	CurrencyCode     string                       `json:"currency_code"`      //原始货币种类
	CurrencySum      float64                      `json:"currency_sum"`       //原始货币金额
	DiscountFee      float64                      `json:"discount_fee"`       //优惠金额
	District         string                       `json:"district"`           //区
	EndTimes         string                       `json:"end_times"`          //完成时间：交易结束或交易成功的时间
	ExpressCode      string                       `json:"express_code"`       //快递单号
	Flag             int                          `json:"flag"`               //旗子颜色0:无1：红2：黄3：绿4：蓝5：粉
	HasRefund        int                          `json:"has_refund"`         //是否有退款
	IdentityName     string                       `json:"identity_name"`      //身份证名称
	IdentityNum      string                       `json:"identity_num"`       //身份信息
	Invoice          []RespQueryTradesDataInvoice `json:"invoice"`            //发票
	IsExceptionTrade bool                         `json:"is_exception_trade"` //是否异常订单
	IsPay            bool                         `json:"is_pay"`             //是否已付款
	IsSmallTrade     bool                         `json:"is_small_trade"`     //是否jit小单
	JzInstallCode    string                       `json:"jz_install_code"`    //安装服务商编码-- 淘系家装类订单字段
	JzInstallName    string                       `json:"jz_install_name"`    //安装服务商名称-- 淘系家装类订单字段
	JzServerCode     string                       `json:"jz_server_code"`     //物流服务商编码-- 淘系家装类订单字段
	JzServerName     string                       `json:"jz_server_name"`     //物流服务商名称-- 淘系家装类订单字段
	LogisticCode     string                       `json:"logistic_code"`      //万里牛ERP快递公司代码，用户自定义代码
	Mark             string                       `json:"mark"`               //订单标记
	ModifyTime       string                       `json:"modify_time"`        //修改时间
	OlnOrderList     []string                     `json:"oln_order_list"`     //明细线上单号集合
	OlnStatus        int                          `json:"oln_status"`         //线上状态:1:等待付款 2:等待发货 ,部分发货 3:已完成 4:已关闭 5: 等待确认 6:已签收 0: 未建交易
	Orders           []RespQueryTradesDataOrders  `json:"orders"`             //明细集
	PaidFee          float64                      `json:"paid_fee"`           //实际支付金额
	PayNo            string                       `json:"pay_no"`             //外部支付单号
	PayTime          string                       `json:"pay_time"`           //付款时间
	PayType          string                       `json:"pay_type"`           //支付类型
	Phone            string                       `json:"phone"`              //手机号，手机号为空的时候返回电话
	PostFee          float64                      `json:"post_fee"`           //邮费
	PrintRemark      string                       `json:"print_remark"`       //打印备注
	PrintTime        string                       `json:"print_time"`         //打单时间
	ProcessStatus    int                          `json:"process_status"`     //万里牛单据处理状态: -3:分销商审核 -2:到账管理 -1:未付款 0:审核 1:打单配货 2:验货 3:称重 4:待发货 5：财审 8:已发货 9:成功 10:关闭 11:异常结束 12:异常处理 13:外部系统配货中 14:预售 15:打包
	Province         string                       `json:"province"`           //省
	Receiver         string                       `json:"receiver"`           //收件人
	Remark           string                       `json:"remark"`             //备注
	SaleNans         string                       `json:"sale_mans"`          //业务员
	SellerMsg        string                       `json:"seller_msg"`         //卖家留言址
	SendTime         string                       `json:"send_time"`          //发货时间
	ServiceFee       float64                      `json:"service_fee"`        //服务费
	ShopId           string                       `json:"shop_id"`            //
	ShopName         string                       `json:"shop_name"`          //店铺名称(页面上显示)
	ShopNick         string                       `json:"shop_nick"`          //店铺昵称(店铺唯一)
	SourcePlatform   string                       `json:"source_platform"`    //订单来源平台
	Status           int                          `json:"status"`             //状态：1：处理中 2：发货 3：完成 4: 关闭 5:其他
	StorageCode      string                       `json:"storage_code"`       //仓库编码
	StorageName      string                       `json:"storage_name"`       //仓库名称
	SumSale          float64                      `json:"sum_sale"`           //总金额,包含优惠
	Tel              string                       `json:"tel"`                //电话
	TpTid            string                       `json:"tp_tid"`             //线上单号,如果是线下订单，则是万里牛的单号，合单情况下会将单号合并，使用|做分隔符
	TradeNo          string                       `json:"trade_no"`           //订单编码
	TradeType        int                          `json:"trade_type"`         //订单类型 :1:普通线上订单 5：货到付款 6：分销 7：团购类型 9：天猫国际物流订单类型 50：普通线下订单 51：售后订单（一般换货创建的订单）
	Zip              string                       `json:"zip"`                //邮编
}

type RespQueryTrades struct {
	Code    int                   `json:"code"`    //返回 0 表示执行成功响应代码
	Data    []RespQueryTradesData `json:"data"`    //响应结果数据
	Message string                `json:"message"` //仅执行出错时返回 响应异常信息
}

//erp/sale/stock/out/query
//查询销售/线下出库单,bill_code modify_time 不能同时为空,只能查询modify_time近7天的单据,对外开放

type StockOutQuery struct {
	BillCode   string `json:"bill_code"`   //单据编码
	ModifyTime string `json:"modify_time"` //修改时间，只能查近3个月
	Page       int    `json:"page"`        //当前页码，从1开始
	Limit      int    `json:"limit"`       //每页大小, 最大200
	IsSplit    bool   `json:"is_split"`    //是否拆分组合商品
}

func NewStockOutQuery(
	BillCode string,
	ModifyTime string,
	Page int,
	Limit int,
	IsSplit bool,
) []byte {
	SOutQuery = &StockOutQuery{
		BillCode:   BillCode,
		ModifyTime: ModifyTime,
		Page:       Page,
		Limit:      Limit,
		IsSplit:    IsSplit,
	}

	b1, err := json.Marshal(SOutQuery)
	if err != nil {
		log.DLog.Println(err)
	}
	return b1
}

type RespStockOutQueryDataDetailsBatchInfos struct {
	BatchNo string `json:"batch_no"` //批次号
	Nums    int    `json:"nums"`     //数量
}

type RespStockOutQueryDataDetails struct {
	BatchInfos  []RespStockOutQueryDataDetailsBatchInfos `json:"batch_infos"`  //批次号集合
	DetailId    string                                   `json:"detail_id"`    //明细id，如果是组合商品拆分出来的明细，这个id相同
	DiscountFee float64                                  `json:"discount_fee"` //优惠
	GoodsName   string                                   `json:"goods_name"`   //商品名称
	Nums        int                                      `json:"nums"`         //数量
	SkuName     string                                   `json:"sku_name"`     //sku名称
	SkuNo       string                                   `json:"sku_no"`       //sku编码（万里牛编码）
	SkuProp1    string                                   `json:"sku_prop1"`    //规格扩展属性，商品未删除，且设置了会返回
	SkuProp2    string                                   `json:"sku_prop2"`    //规格扩展属性，商品未删除，且设置了会返回
	SnInfos     []string                                 `json:"sn_infos"`     //序列号集合，此字段以不返回数据，请使接口——/erp/sn/querysnbybillcode查询单据中的序列号
	SumCost     float64                                  `json:"sum_cost"`     //成本
	SumSale     float64                                  `json:"sum_sale"`     //总售价,包含优惠
	Unit        string                                   `json:"unit"`         //单位
}

type RespStockOutQueryData struct {
	BillDate             string                         `json:"bill_date"`               //单据日期
	BillType             int                            `json:"bill_type"`               //单据类型: 1: 销售出库单 2： 线下出库单
	Country              string                         `json:"country"`                 //国家
	CreateTime           string                         `json:"create_time"`             //创建时间
	CustomCode           string                         `json:"custom_code"`             //客户编码
	CustomName           string                         `json:"custom_name"`             //客户名称
	CustomerNick         string                         `json:"customer_nick"`           //客户昵称
	CustomerNickType     int                            `json:"customer_nick_type"`      //客户来源平台值
	CustomerNickTypeName string                         `json:"customer_nick_type_name"` //客户来源平台名称 比如:淘宝,京东
	Details              []RespStockOutQueryDataDetails `json:"details"`                 //明细集合
	DiscountFee          float64                        `json:"discount_fee"`            //优惠金额
	FromTradeNo          string                         `json:"from_trade_no"`           //万里牛中的原始交易编号
	InvNo                string                         `json:"inv_no"`                  //单据编号
	PaidFee              float64                        `json:"paid_fee"`                //实际支付金额
	PayType              int                            `json:"pay_type"`                //付款方式 0 线上支付 1 现金 2 刷卡 3 微信 4 支付宝 5 转账 6 储值支付 7 1688支付 52 其它转账支付
	PostFee              float64                        `json:"post_fee"`                //邮费
	Remark               string                         `json:"remark"`                  //备注
	SaleMan              string                         `json:"sale_man"`                //业务员
	ServiceFee           float64                        `json:"service_fee"`             //服务费
	ShopName             string                         `json:"shop_name"`               //店铺名称
	ShopNick             string                         `json:"shop_nick"`               //店铺昵称
	ShopSource           string                         `json:"shop_source"`             //店铺来源
	StorageCode          string                         `json:"storage_code"`            //仓库编码
	StorageName          string                         `json:"storage_name"`            //仓库名称
	SumSale              float64                        `json:"sum_sale"`                //总金额,包含优惠
	TpTid                string                         `json:"tp_tid"`                  //外部订单号(如淘宝/京东单号)
}

type RespStockOutQuery struct {
	Code    int                     `json:"code"`    //返回 0 表示执行成功响应代码
	Data    []RespStockOutQueryData `json:"data"`    //响应结果数据
	Message string                  `json:"message"` //仅执行出错时返回 响应异常信息
}

//erp/opentrade/modify/remark
//修改订单备注信息,最大长度为2048字节,超出部分忽略,对外开发

type ModifyRemarkOrders struct {
	OrderId string `json:"order_id"` //明细id，单据级唯一,接口：/query/trades 获取
	Remark  string `json:"remark"`   //备注最大长度为2048字节,超出部分忽略
}

type ModifyRemark struct {
	BillCode string               `json:"bill_code"` //单据编码
	Orders   []ModifyRemarkOrders `json:"orders"`    //明细集合
	Remark   string               `json:"remark"`    //备注
}

func NewModifyRemarkOrders(
	OrderId string,
	Remark string,
) ModifyRemarkOrders {
	return ModifyRemarkOrders{
		OrderId: OrderId,
		Remark:  Remark,
	}
}

func NewModifyRemark(
	BillCode string,
	Orders []ModifyRemarkOrders,
	Remark string,
) []byte {
	MRemark = &ModifyRemark{
		BillCode: BillCode,
		Orders:   Orders,
		Remark:   Remark,
	}

	b1, err := json.Marshal(MRemark)
	if err != nil {
		log.DLog.Println(err)
	}
	return b1
}

type RespModifyRemark struct {
	Code    int    `json:"code"`    //返回 0 表示执行成功响应代码
	Data    bool   `json:"data"`    //响应结果数据
	Message string `json:"message"` //仅执行出错时返回 响应异常信息
}

//erp/opentrade/modify/mark
//修改订单标记,对外开发——因部分订单标记后，不会生成出库单，请谨慎使用,具体要看用户的设置

type ModifyMark struct {
	BillCode string `json:"bill_code"` //单据编码
	MarkName string `json:"mark_name"` //标记名称
}

func NewModifyMark(
	BillCode string,
	MarkName string,
) []byte {
	MMark = &ModifyMark{
		BillCode: BillCode,
		MarkName: MarkName,
	}

	b1, err := json.Marshal(MMark)
	if err != nil {
		log.DLog.Println(err)
	}
	return b1
}

type RespModifyMark struct {
	Code    int    `json:"code"`    //返回 0 表示执行成功响应代码
	Data    bool   `json:"data"`    //响应结果数据
	Message string `json:"message"` //仅执行出错时返回 响应异常信息
}

//erp/opentrade/reply/exception/trades
//打回异常订单,返回失败的订单编码,对外开放

type ReplyExceptionTrades struct {
	BillCodes []string `json:"bill_codes"` //单据编码
	Remark    string   `json:"remark"`     //备注
}

func NewReplyExceptionTrades(
	BillCodes []string,
	Remark string,
) []byte {
	RExceptionTrades = &ReplyExceptionTrades{
		BillCodes: BillCodes,
		Remark:    Remark,
	}

	b1, err := json.Marshal(RExceptionTrades)
	if err != nil {
		log.DLog.Println(err)
	}
	return b1
}

type RespReplyExceptionTrades struct {
	Code    int      `json:"code"`    //返回 0 表示执行成功响应代码
	Data    []string `json:"data"`    //响应结果数据
	Message string   `json:"message"` //仅执行出错时返回 响应异常信息
}

//erp/opentrade/send/trades
//发货,对外开放

type SendTradesSendTradesPackagesDetails struct {
	Nums    int    `json:"nums"`     //发货数量
	SkuCode string `json:"sku_code"` //sku编码
}

type SendTradesSendTradesPackages struct {
	Details        []SendTradesSendTradesPackagesDetails `json:"details"`          //明细
	ErpExpressCode string                                `json:"erp_express_code"` //万里牛ERP快递公司代码，具体值请登录ERP，快递公司页面查看。
	ExpressCompany string                                `json:"express_company"`  //快递公司代码, 建议不要传，使用erp_express_code
	PostFee        float64                               `json:"post_fee"`         //快递成本
	WayBill        string                                `json:"way_bill"`         //快递单号
	Weight         float64                               `json:"weight"`           //重量
}

type SendTradesSendTrades struct {
	BillCode      string                         `json:"bill_code"`      //订单编码
	Packages      []SendTradesSendTradesPackages `json:"packages"`       //包裹信息集合
	PartDelivered bool                           `json:"part_delivered"` //是否部分发货, 默认全部发货
}

type SendTrades struct {
	SendTrade SendTradesSendTrades `json:"send_trade"` //发货单
	Remark    string               `json:"remark"`     ///备注
}

func NewSendTradesSendTradesPackagesDetails(
	Nums int,
	SkuCode string,
) SendTradesSendTradesPackagesDetails {
	return SendTradesSendTradesPackagesDetails{
		Nums:    Nums,
		SkuCode: SkuCode,
	}
}

func NewSendTradesSendTradesPackages(
	Details []SendTradesSendTradesPackagesDetails,
	ErpExpressCode string,
	ExpressCompany string,
	PostFee float64,
	WayBill string,
	Weight float64,
) SendTradesSendTradesPackages {
	return SendTradesSendTradesPackages{
		Details:        Details,
		ErpExpressCode: ErpExpressCode,
		ExpressCompany: ExpressCompany,
		PostFee:        PostFee,
		WayBill:        WayBill,
		Weight:         Weight,
	}
}

func NewSendTrades(
	BillCode string,
	Packages []SendTradesSendTradesPackages,
	PartDelivered bool,

	Remark string,
) []byte {
	STrades = &SendTrades{
		SendTrade: SendTradesSendTrades{
			BillCode:      BillCode,
			Packages:      Packages,
			PartDelivered: PartDelivered,
		},
		Remark: Remark,
	}

	b1, err := json.Marshal(STrades)
	if err != nil {
		log.DLog.Println(err)
	}
	return b1
}

type RespSendTrades struct {
	Code    int    `json:"code"`    //返回 0 表示执行成功响应代码
	Data    bool   `json:"data"`    //响应结果数据
	Message string `json:"message"` //仅执行出错时返回 响应异常信息
}

//erp/sale/stock/out/add
//添加销售出库单（线下出库单）,对外开放

type StockOutAddBillDetails struct {
	Nums    int     `json:"nums"`     //数量
	SkuNo   string  `json:"sku_no"`   //sku编码（万里牛编码）
	SumSale float64 `json:"sum_sale"` //总售价,包含优惠
}

type StockOutAddBill struct {
	BillDate     string                   `json:"bill_date"`     //单据日期
	CreateTime   string                   `json:"create_time"`   //创建时间
	CustomerNick string                   `json:"customer_nick"` //客户昵称
	Details      []StockOutAddBillDetails `json:"details"`       //明细集合, 以万里牛商品的基本单位保存
	DiscountFee  float64                  `json:"discount_fee"`  //优惠金额
	PaidFee      float64                  `json:"paid_fee"`      //实际支付金额
	PostFee      float64                  `json:"post_fee"`      //邮费
	Remark       string                   `json:"remark"`        //备注
	ServiceFee   float64                  `json:"service_fee"`   //	服务费
	ShopNick     string                   `json:"shop_nick"`     //店铺昵称, 对应万里牛erp中的店铺账号ID
	StorageCode  string                   `json:"storage_code"`  //仓库编码
	SumSale      float64                  `json:"sum_sale"`      //总金额, 包含优惠
}
type StockOutAdd struct {
	Bill StockOutAddBill `json:"bill"` //销售出库单
}

func NewStockOutAddBillDetails(
	Nums int,
	SkuNo string,
	SumSale float64,
) StockOutAddBillDetails {
	return StockOutAddBillDetails{
		Nums:    Nums,
		SkuNo:   SkuNo,
		SumSale: SumSale,
	}
}

func NewStockOutAdd(
	BillDate string,
	CreateTime string,
	CustomerNick string,
	Details []StockOutAddBillDetails,
	DiscountFee float64,
	PaidFee float64,
	PostFee float64,
	Remark string,
	ServiceFee float64,
	ShopNick string,
	StorageCode string,
	SumSale float64,
) []byte {
	SOutAdd = &StockOutAdd{
		Bill: StockOutAddBill{
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

	b1, err := json.Marshal(SOutAdd)
	if err != nil {
		log.DLog.Println(err)
	}
	return b1
}

type RespStockOutAdd struct {
	Code    int    `json:"code"`    //返回 0 表示执行成功响应代码
	Data    string `json:"data"`    //响应结果数据
	Message string `json:"message"` //仅执行出错时返回 响应异常信息
}

//erp/opentrade/reply/approve/trades
//打回打单配货订单到审核,返回是否打回成功,对外开放

type ApproveTrades struct {
	BillCode string `json:"bill_code"` //单据系统编码
}

func NewApproveTrades(BillCode string) []byte {
	ATrades = &ApproveTrades{
		BillCode: BillCode,
	}
	b1, err := json.Marshal(ATrades)
	if err != nil {
		log.DLog.Println(err)
	}
	return b1

}

type RespApproveTrades struct {
	Code    int    `json:"code"`    //返回 0 表示执行成功响应代码
	Data    bool   `json:"data"`    //响应结果数据
	Message string `json:"message"` //仅执行出错时返回 响应异常信息
}

//erp/opentrade/trade/commit
//订单进行状态的流转,进行操作后会自动流转到下个环节，正向环节支持审核-财审-配货-验货

type TradeCommit struct {
	BillCode    string `json:"bill_code"`    //单据编码
	TradeStatus int    `json:"trade_status"` //订单当前状态
	CommitType  int    `json:"commit_type"`  //操作方式 0 正向提交 1 打回
	NextStatus  int    `json:"next_status"`  //打回的订单状态 commit_type = 1 必传
	Force       bool   `json:"force"`        //是否强制提交-库存不足的时候
}

func NewTradeCommit(
	BillCode string,
	TradeStatus int,
	CommitType int,
	NextStatus int,
	Force bool,
) []byte {
	TCommit = &TradeCommit{
		BillCode:    BillCode,
		TradeStatus: TradeStatus,
		CommitType:  CommitType,
		NextStatus:  NextStatus,
		Force:       Force,
	}
	b1, err := json.Marshal(TCommit)
	if err != nil {
		log.DLog.Println(err)
	}
	return b1

}

type RespTradeCommit struct {
	Code    int    `json:"code"`    //返回 0 表示执行成功响应代码
	Data    bool   `json:"data"`    //响应结果数据
	Message string `json:"message"` //仅执行出错时返回 响应异常信息
}

//erp/opentrade/query/mark
//查询订单标记列表
type RespQueryMarkData struct {
	Color    string `json:"color"`     //颜色
	MarkName string `json:"mark_name"` //标记名称
}

type RespQueryMark struct {
	Code    int                 `json:"code"`    //返回 0 表示执行成功响应代码
	Data    []RespQueryMarkData `json:"data"`    //响应结果数据
	Message string              `json:"message"` //仅执行出错时返回 响应异常信息
}

//erp/sale/stock/count/query
//查询出库单/退货入库单的总量

type StockCountQuery struct {
	BillType int    `json:"bill_type"` //单据类型，1代表销售出库单，2代表退货入库单，非此类型按1处理
	BillDate string `json:"bill_date"` //单据时间，只能查近3个月内某日的单据量
}

func NewStockCountQuery(
	BillType int,
	BillDate string,
) []byte {

	SCountQuery = &StockCountQuery{
		BillType: BillType,
		BillDate: BillDate,
	}
	b1, err := json.Marshal(SCountQuery)
	if err != nil {
		log.DLog.Println(err)
	}
	return b1

}

type RespQStockCountQuery struct {
	Code    int    `json:"code"`    //返回 0 表示执行成功响应代码
	Data    int    `json:"data"`    //响应结果数据
	Message string `json:"message"` //仅执行出错时返回 响应异常信息
}

var QTrades *QueryTrades
var SOutQuery *StockOutQuery
var MRemark *ModifyRemark
var MMark *ModifyMark
var RExceptionTrades *ReplyExceptionTrades
var STrades *SendTrades
var SOutAdd *StockOutAdd
var ATrades *ApproveTrades
var TCommit *TradeCommit
var SCountQuery *StockCountQuery
