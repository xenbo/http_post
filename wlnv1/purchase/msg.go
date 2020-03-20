package purchase

import (
	"encoding/json"
	"github.com/xenbo/http_post/log"
)

//erp/purchase/purchasereturnbill/close

type Close struct {
	BillCode        string `json:"bill_code"`         //单据编码
	CloseRemark     string `json:"close_remark"`      //关闭备注
	NotCloseStocked bool   `json:"not_close_stocked"` //已出库不关闭
}

func NewClose(BillCode string, CloseRemark string, NotCloseStocked bool) []byte {
	Clo = &Close{
		BillCode:        BillCode,
		CloseRemark:     CloseRemark,
		NotCloseStocked: NotCloseStocked,
	}

	b1, err := json.Marshal(Clo)
	log.DLog.Println(err)

	return b1
}

type RespClose struct {
	Code    int    //返回 0 表示执行成功 响应代码
	Data    bool   //响应结果数据
	Message string //仅执行出错时返回 响应异常信息
}

//-------------------------------------------------------------------------

type Detail struct {
	Index    int     `json:"index"`     //行号，查询单据的时候会返回，如果没填，erp给默认
	Price    float64 `json:"price"`     //采购价
	Remark   string  `json:"remark"`    //备注
	Size     float64 `json:"size"`      //数量
	SpecCode string  `json:"spec_code"` //规格编码
	Sum      float64 `json:"sum"`       //总价
}

type AddBill struct {
	BillCode     string `json:"bill_code"`     //单号,如果为空，则使用erp自己规则生成的单号,请保证单号不能重复，如重复报错
	Details      Detail `json:"details"`       //明细
	Remark       string `json:"remark"`        //备注
	StorageCode  string `json:"storage_code"`  //仓库编码
	SupplierCode string `json:"supplier_code"` //供应商编码
}

type Add struct {
	Bill AddBill `json:"bill"`
}

type RespAdd struct {
	Code    int    `json:"code"`    //返回 0 表示执行成功  响应代码
	Data    string `json:"data"`    //响应结果数据
	Message string `json:"message"` //仅执行出错时返回 响应异常信息
}

func NewAdd(BillCode string, Index int, Price float64, Remark string, Size float64, SpecCode string,
	Sum float64, Remark2 string, StorageCode string, SupplierCode string) []byte {
	ABill = &Add{
		Bill: AddBill{
			BillCode: BillCode,
			Details: Detail{
				Index:    Index,
				Price:    Price,
				Remark:   Remark,
				Size:     Size,
				SpecCode: SpecCode,
				Sum:      Sum,
			},
			Remark:       Remark2,
			StorageCode:  StorageCode,
			SupplierCode: SupplierCode,
		},
	}
	b1, err := json.Marshal(ABill)
	log.DLog.Println(err)

	return b1
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
	log.DLog.Println(err)

	return b1
}

//-------------------------------------------------------------------------

var Clo *Close
var ABill *Add
var BQuery *BillQuery
