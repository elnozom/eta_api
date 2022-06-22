package model

type Issuer struct {
	Type    string        `json:"type"`
	Id      string        `json:"id"`
	Name    string        `json:"name"`
	Address IssuerAddress `json:"address"`
}

type RecieverAddress struct {
	Country        string `json:"country"`
	Governate      string `json:"governate"`
	RegionCity     string `json:"regionCity"`
	Street         string `json:"street"`
	BuildingNumber string `json:"buildingNumber"`
}
type IssuerAddress struct {
	BranchId       int    `json:"branchId"`
	Country        string `json:"country"`
	Governate      string `json:"governate"`
	RegionCity     string `json:"regionCity"`
	Street         string `json:"street"`
	BuildingNumber string `json:"buildingNumber"`
}
type Reciever struct {
	Type    string          `json:"type"`
	Id      string          `json:"id"`
	Name    string          `json:"name"`
	Address RecieverAddress `json:"address"`
}

type InvoiceLine struct {
	Description      string  `json:"description"`
	ItemType         string  `json:"itemType"`
	ItemCode         string  `json:"itemCode"`
	UnitType         string  `json:"unitType"`
	UnitValue        Value   `json:"unitValue"`
	Quantity         float64 `json:"quantity"`
	SalesTotal       float64 `json:"salesTotal"`
	Total            float64 `json:"total"`
	ValueDifference  float64 `json:"valueDifference"`
	TotalTaxableFees float64 `json:"totalTaxableFees"`
	NetTotal         float64 `json:"netTotal"`
	ItemsDiscount    float64 `json:"itemsDiscount"`
}
type Value struct {
	CurrencySold string  `json:"currencySold"`
	AmountEGP    float64 `json:"amountEGP"`
}
type Signature struct {
	Issuer       Issuer   `json:"=issuer"`
	Reciever     Reciever `json:"reciever"`
	DocumentType string   `json:"documentType"`
}

type Invoice struct {
	Issuer                   Issuer        `json:"issuer"`
	TaxTotals                []Value       `json:"taxTotals"`
	Reciever                 Reciever      `json:"reciever"`
	InvoiceLines             []InvoiceLine `json:"invoiceLines"`
	Signatures               []Signature   `json:"signatures"`
	DocumentType             string        `json:"documentType"`
	DocumentTypeVersion      string        `json:"documentTypeVersion"`
	TaxpayerActivityCode     string        `json:"taxpayerActivityCode"`
	DateTimeIssued           string        `json:"dateTimeIssued"`
	InternalId               string        `json:"internalId"`
	TotalSalesAmount         float64       `json:"totalSalesAmount"`
	TotalDiscountAmount      float64       `json:"totalDiscountAmount"`
	TotalItemsDiscountAmount float64       `json:"totalItemsDiscountAmount"`
	ExtraDiscountAmount      float64       `json:"extraDiscountAmount"`
	TotalAmount              float64       `json:"totalAmount"`
	NetAmount                float64       `json:"netAmount"`
}

type InvoiceSubmitRequest struct {
	Documents []Invoice `json:"documents"`
}

type InvoiceFindResp struct {
	Serial                   int
	DateTimeIssued           string
	InternalId               string
	TotalSalesAmount         float64
	TotalItemsDiscountAmount float64
	TotalDiscountAmount      float64
	ExtraDiscountAmount      float64
	NetAmount                float64
	TotalAmount              float64
}
