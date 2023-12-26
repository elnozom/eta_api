package model

type ReceiptSubmitRequest struct {
	Receipts []Receipt `json:"receipts"`
}
type Receipt struct {
	Header                  Header       `json:"header"`
	Seller                  Seller       `json:"seller"`
	Buyer                   Buyer        `json:"buyer"`
	DocumentType            DocumentType `json:"documentType"`
	ItemData                []ItemData   `json:"itemData"`
	TotalSales              float64      `json:"totalSales"`
	TotalCommersialDiscount float64      `json:"totalCommercialDiscount"`
	TotalItemsDiscount      float64      `json:"totalItemsDiscount"`
	NetAmount               float64      `json:"netAmount"`
	FeesAmount              float64      `json:"feesAmount"`
	TotalAmount             float64      `json:"totalAmount"`
	TaxTotals               []TaxTotal   `json:"taxTotals"`
	PaymentMethod           string       `json:"paymentMethod"`
}

type DocumentType struct {
	ReceiptType string `json:"receiptType"`
	TypeVersion string `json:"typeVersion"`
}
type TaxTotal struct {
	TaxType string  `json:"taxType"`
	Amount  float64 `json:"amount"`
}
type Header struct {
	DateTimeIssued string `json:"dateTimeIssued"`
	ReceiptNumber  string `json:"receiptNumber"`
	Uuid           string `json:"uuid"`
	PreviousUUID   string `json:"previousUUID"`
	Currency       string `json:"currency"`
}
type Seller struct {
	Rin                string          `json:"rin"`
	CompanyTradeName   string          `json:"companyTradeName"`
	BranchCode         string          `json:"branchCode"`
	BranchAddress      ReceiverAddress `json:"branchAddress"`
	DeviceSerialNumber string          `json:"deviceSerialNumber"`
	ActivityCode       string          `json:"activityCode"`
}
type Buyer struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}
type ItemData struct {
	InternalCode string         `json:"internalCode"`
	Description  string         `json:"description"`
	ItemType     string         `json:"itemType"`
	ItemCode     string         `json:"itemCode"`
	UnitType     string         `json:"unitType"`
	TaxableItems []TaxableItems `json:"taxableItems"`
	Quantity     float64        `json:"quantity"`
	UnitPrice    float64        `json:"unitPrice"`
	NetSale      float64        `json:"netSale"`
	TotalSale    float64        `json:"totalSale"`
	Total        float64        `json:"total"`
}
