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

func NewClose(b string, c string, n bool) []byte {
	Clo = &Close{
		BillCode:        b,
		CloseRemark:     c,
		NotCloseStocked: n,
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

type Add struct {
	BillCode     string `json:"bill_code"`     //单号,如果为空，则使用erp自己规则生成的单号,请保证单号不能重复，如重复报错
	Details      Detail `json:"details"`       //明细
	Remark       string `json:"remark"`        //备注
	StorageCode  string `json:"storage_code"`  //仓库编码
	SupplierCode string `json:"supplier_code"` //供应商编码
}

type RespAdd struct {
	Code    int    `json:"code"`    //返回 0 表示执行成功  响应代码
	Data    string `json:"data"`    //响应结果数据
	Message string `json:"message"` //仅执行出错时返回 响应异常信息
}

func NewAdd(b string, i int, p float64, r2 string, s3 float64, s4 string, s5 float64, r string, s1 string, s2 string) []byte {
	AddBill = &Add{
		BillCode: b,
		Details: Detail{
			Index:    i,
			Price:    p,
			Remark:   r2,
			Size:     s3,
			SpecCode: s4,
			Sum:      s5,
		},
		Remark:       r,
		StorageCode:  s1,
		SupplierCode: s2,
	}
	b1, err := json.Marshal(AddBill)
	log.DLog.Println(err)

	return b1
}

//-------------------------------------------------------------------------

var Clo *Close
var AddBill *Add
