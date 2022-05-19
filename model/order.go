package model

type Order struct {
	Serial    int     `json:"serial"`
	DocNo     string  `json:"docNo"`
	DocDate   string  `json:"docDate"`
	Discount  float64 `json:"discount"`
	TotalCash float64 `json:"totalCash"`
	TotalTax  float64 `json:"totalTax"`
}
type EInvoice struct {
	Serial              int     `json:"serial"`
	InternalID          string  `json:"internlID"`
	StoreCode           string  `json:"storeCode"`
	TotalDiscountAmount float64 `json:"totalDiscountAmount"`
	TotalAmount         float64 `json:"totalAmount"`
	TotalTax            float64 `json:"totalTax"`
	StkTr01Serial       float64 `json:"stkTr01Serial"`
}
