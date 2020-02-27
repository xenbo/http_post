package wln

import (
	"gitlab.com/eosforce/vbbirdworker/event"
	"strconv"
)

type System struct {
	Appkey    string `json:"app_key"`
	TimeStamp int64  `json:"timestamp"`
	Sign      string `json:"sign"`
	Format    string `json:"format"`
}

type Args struct {
	ModifyTime string `json:"modify_time"`
	Page_no    int32  `json:"page"`  //1
	Page_size  int32  `json:"limit"` //20
}

type Args2 struct {
	ModifyTime string `json:"modify_time"`
	Page_no    int32  `json:"page"`  //1
	Page_size  int32  `json:"limit"` //20
	BillType   int32  `json:"bill_type"`
}

type Msg interface {
	name() string
	RequestData(int32, string, string, string)

	GetNodeID0(i int) string
	GetNodeRole0(i int) string
	NodeActorID0(i int) string
	NodeActorRole0(i int) string

	GetNodeID1(i int, j int) string
	GetNodeRole1(i int, j int) string
	NodeActorID1(i int, j int) string
	NodeActorRole1(i int, j int) string

	GetContainerId(int,int) string //sku
	GetExpressId() string   //物流编号

	RespN() int
	ContainerN(i int) int

	GetType()  event.EvType
}

//---------------------------------------------
type StockDetails struct {
	DetailId string  `json:"detail_id"`
	SkuNo    string  `json:"sku_no"`
	Size     float64 `json:"size"`
	SnValue  string  `json:"sn_value"`
}

type StockResponse struct {
	//销售出库单接
	ExpressCode bool           `json:"express_code"`
	Express     string         `json:"express"`
	CustomCode  string         `json:"custom_code"`
	CustomNick  string         `json:"custom_nick"`
	CustomName  string         `json:"custom_name"`
	StorageCode string         `json:"storage_code"`
	StorageName string         `json:"storage_name"`
	BillDate    string         `json:"bill_date"`
	Provice     string         `json:"provice"`
	City        string         `json:"city"`
	Company     string         `json:"company"`
	ShopNick    string         `json:"shop_nick"`
	Details     []StockDetails `json:"details"`
}

type StockMsg struct {
	Success   bool            `json:"success"`
	ErrorCode string          `json:"error_code"`
	ErrorMsg  string          `json:"error_msg"`
	Response  []StockResponse `json:"response"`
	Tp 		  event.EvType 	  `json:"type"`
}

func NewStockMsg() StockMsg {
	return StockMsg {
		Tp: event.TypeSendDealer,
	}
}


func (msg *StockMsg) name() string {
	return "stock"
}

func (msg *StockMsg) GetNodeID0(i int) string {
	return string(msg.Response[i].StorageCode)
}

func (msg *StockMsg) GetNodeRole0(i int) string {
	return string(msg.Response[i].StorageName)
}

func (msg *StockMsg) NodeActorID0(i int) string {
	return string(msg.Response[i].BillDate)
}

func (msg *StockMsg) NodeActorRole0(i int) string {
	return string(msg.Response[i].Company)
}

func (msg *StockMsg) GetNodeID1(i int, j int) string {
	return string(msg.Response[i].Details[j].DetailId)
}

func (msg *StockMsg) GetNodeRole1(i int, j int) string {
	return strconv.FormatFloat(msg.Response[i].Details[j].Size, 'f', 6, 64)
}

func (msg *StockMsg) NodeActorID1(i int, j int) string {
	return msg.Response[i].CustomCode
}

func (msg *StockMsg) NodeActorRole1(i int, j int) string {
	return ""
}

func (msg *StockMsg) GetExpressId() string {
	return ""
}

func (msg *StockMsg) GetContainerId(i int, j int) string {
	return msg.Response[i].Details[j].SkuNo
} //sku

func (msg *StockMsg) RespN() int {
	return len(msg.Response)
}

func (msg *StockMsg) ContainerN(i int) int {
	return len(msg.Response[i].Details)
}

func (msg *StockMsg)GetType()event.EvType  {
	return msg.Tp
}

//----------------------------------------------------

type PurchaseDetails struct {
	DetailId string  `json:"detail_id"`
	SkuNo    string  `json:"sku_no"`
	Size     float64 `json:"size"`
	SnValue  string  `json:"sn_value"`
}

type PurchaseResponse struct {
	StorageCode string            `json:"storage_code"`
	StorageName string            `json:"storage_name"`
	BillDate    string            `json:"bill_date"`
	Company     string            `json:"company"`
	Details     []PurchaseDetails `json:"details"`
}

type PurchaseMsg struct {
	Success   bool               `json:"success"`
	ErrorCode string             `json:"error_code"`
	ErrorMsg  string             `json:"error_msg"`
	Response  []PurchaseResponse `json:"response"`
	Tp event.EvType				 `json:"type"`
}

func NewPurchaseMsg() PurchaseMsg {
	return PurchaseMsg {
		Tp: event.TypeRecvDealer,
	}
}

func (msg *PurchaseMsg) name() string {
	return "purchase"
}

func (msg *PurchaseMsg) GetNodeID0(i int) string {
	return string(msg.Response[i].StorageCode)
}

func (msg *PurchaseMsg) GetNodeRole0(i int) string {
	return string(msg.Response[i].StorageName)
}

func (msg *PurchaseMsg) NodeActorID0(i int) string {
	return string(msg.Response[i].BillDate)
}

func (msg *PurchaseMsg) NodeActorRole0(i int) string {
	return string(msg.Response[i].Company)
}

func (msg *PurchaseMsg) GetNodeID1(i int, j int) string {

	return ""
}

func (msg *PurchaseMsg) GetNodeRole1(i int, j int) string {
	return ""
}

func (msg *PurchaseMsg) NodeActorID1(i int, j int) string {
	return ""
}

func (msg *PurchaseMsg) NodeActorRole1(i int, j int) string {
	return ""
}

func (msg *PurchaseMsg) GetExpressId() string {
	return ""
}

func (msg *PurchaseMsg) GetContainerId(i int, j int) string {
	return msg.Response[i].Details[j].SkuNo
} //sku

func (msg *PurchaseMsg) RespN() int {
	return len(msg.Response)
}

func (msg *PurchaseMsg) ContainerN(i int) int {
	return len(msg.Response[i].Details)
}

func (msg *PurchaseMsg)GetType()event.EvType  {
	return msg.Tp
}


//---------------------------------------------

type InventoryDetails struct {
	DetailId string  `json:"detail_id"`
	SkuNo    string  `json:"sku_no"`
	Size     float64 `json:"size"`
	SnValue  string  `json:"sn_value"`
}

type InventoryResponse struct {
	StorageCode string            `json:"storage_code"`
	StorageName string            `json:"storage_name"`
	BillDate    string            `json:"bill_date"`
	Company     string            `json:"company"`
	Details     []PurchaseDetails `json:"details"`
}

type InventoryMsg struct {
	Success   bool               `json:"success"`
	ErrorCode string             `json:"error_code"`
	ErrorMsg  string             `json:"error_msg"`
	Response  []PurchaseResponse `json:"response"`
	Tp  event.EvType			 `json:"type"`
}

func NewInventoryMsg0() InventoryMsg {
	return InventoryMsg{
		Tp: event.TypeSendDepot,
	}
}

func NewInventoryMsg1() InventoryMsg {
	return InventoryMsg{
		Tp: event.TypeRecvDepot,
	}
}

func (msg *InventoryMsg) name() string {
	return "inventory"
}

func (msg *InventoryMsg) GetNodeID0(i int) string {
	return string(msg.Response[i].StorageCode)
}

func (msg *InventoryMsg) GetNodeRole0(i int) string {
	return string(msg.Response[i].StorageName)
}

func (msg *InventoryMsg) NodeActorID0(i int) string {
	return string(msg.Response[i].BillDate)
}

func (msg *InventoryMsg) NodeActorRole0(i int) string {
	return string(msg.Response[i].Company)
}

func (msg *InventoryMsg) GetNodeID1(i int, j int) string {

	return ""
}

func (msg *InventoryMsg) GetNodeRole1(i int, j int) string {
	return ""
}

func (msg *InventoryMsg) NodeActorID1(i int, j int) string {
	return ""
}

func (msg *InventoryMsg) NodeActorRole1(i int, j int) string {
	return ""
}

func (msg *InventoryMsg) GetExpressId() string {
	return ""
}

func (msg *InventoryMsg) GetContainerId(i int, j int) string {
	return msg.Response[i].Details[j].SkuNo
} //sku

func (msg *InventoryMsg) RespN() int {
	return len(msg.Response)
}

func (msg *InventoryMsg) ContainerN(i int) int {
	return len(msg.Response[i].Details)
}

func (msg *InventoryMsg)GetType()event.EvType  {
	return msg.Tp
}


//---------------------------------------------
