package repo

import (
	"database/sql"
	"eta/model"
	"eta/utils"
	"math"

	"github.com/jinzhu/gorm"
)

type InvoiceRepo struct {
	db *gorm.DB
}

func NewInvoiceRepo(db *gorm.DB) InvoiceRepo {
	return InvoiceRepo{
		db: db,
	}
}

func (ur *InvoiceRepo) ListEInvoices(req *model.ListInvoicessRequest) (*[]model.EInvoice, error) {
	rows, err := ur.db.Raw("EXEC StkTrEInvoiceHeadList @posted = ? , @store = ? , @start_date = ? , @end_date = ? ", req.Posted, req.Store, req.StartDate, req.EndDate).Rows()
	utils.CheckErr(&err)
	defer rows.Close()
	if utils.CheckErr(&err) {
		return nil, err
	}
	result, err := scanEInvoiceResult(rows)
	return result, nil
}

func (ur *InvoiceRepo) EInvoiceHeadPost(serial *uint64, store *uint64) (*int, error) {
	var resp int
	err := ur.db.Raw("EXEC StkTrEInvoicePosted @serial = ? , @store = ?  ", serial, store).Row().Scan(&resp)

	if utils.CheckErr(&err) {
		return nil, err
	}
	return &resp, nil
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

// func (ur *InvoiceRepo) FindInvoiceData(req model.PostInvoicessRequest) error {

// 	rows, err := ur.db.Raw("EXEC StkTrEInvoiceFind @serials = ? , @store = ? ", req.Serilas, req.Store).Rows()

// 	// invoice.TotalSalesAmount = roundFloat(invoice.TotalSalesAmount, 5)
// 	// invoice.TotalItemsDiscountAmount = roundFloat(invoice.TotalItemsDiscountAmount, 5)
// 	// invoice.NetAmount = roundFloat(invoice.NetAmount, 5)
// 	if utils.CheckErr(&err) {
// 		return err
// 	}
// 	// var se string = strconv.FormatUint(*serial, 10)
// 	// var st string = strconv.FormatUint(*serial, 10)
// 	// invoice.InternalID = "IID" + se + st
// 	// invoice.Receiver.Type = "P"

// 	// err = ur.db.Raw("EXEC StkTrEInvoiceFindIssuerAddress  @storeCode = ? ", branchId).Row().Scan(
// 	// 	&addressSerial,
// 	// 	&invoice.Issuer.Address.Country,
// 	// 	&invoice.Issuer.Address.Governate,
// 	// 	&invoice.Issuer.Address.RegionCity,
// 	// 	&invoice.Issuer.Address.Street,
// 	// 	&invoice.Issuer.Address.BuildingNumber,
// 	// )
// 	// if utils.CheckErr(&err) {
// 	// 	return err
// 	// }
// 	// rows, err := ur.db.Raw("EXEC StkTrEInvoiceFindItems  @serial = ? , @store = ? ", serial, store).Rows()
// 	// if utils.CheckErr(&err) {
// 	// 	return err
// 	// }
// 	defer rows.Close()

// 	taxTotals := 0.0
// 	for rows.Next() {
// 		var line model.InvoiceLine
// 		err = rows.Scan(
// 			&line.ItemType,
// 			&line.ItemCode,
// 			&line.UnitType,
// 			&line.Quantity,
// 			&line.UnitValue.AmountEGP,
// 			&line.ItemsDiscount,
// 			&line.SalesTotal,
// 			&line.Total,
// 			&line.NetTotal,
// 		)

// 		line.Description = "description"
// 		line.Total = roundFloat(line.Total, 5)
// 		line.NetTotal = roundFloat(line.NetTotal, 5)
// 		line.UnitValue.AmountEGP = roundFloat(line.UnitValue.AmountEGP, 5)
// 		line.SalesTotal = roundFloat(line.SalesTotal, 5)
// 		var itemTax model.TaxableItems
// 		itemTax.Amount = roundFloat(line.SalesTotal*.14, 5)
// 		itemTax.TaxType = "T1"
// 		itemTax.SubType = " "
// 		itemTax.Rate = 14.00
// 		line.TaxableItems = append(line.TaxableItems, itemTax)

// 		taxTotals += roundFloat(line.TaxableItems[0].Amount, 5)
// 		if utils.CheckErr(&err) {
// 			return err
// 		}
// 		line.UnitValue.CurrencySold = "EGP"
// 		invoice.InvoiceLines = append(invoice.InvoiceLines, line)
// 	}

// 	var taxTotal model.TaxTotals

// 	taxTotal.TaxType = "T1"
// 	taxTotal.Amount = roundFloat(taxTotals, 5)
// 	invoice.TaxTotals = append(invoice.TaxTotals, taxTotal)
// 	// invoice.Issuer.Address.BranchId = "0"
// 	return nil
// }

// type Invoice struct {
// 	Issuer                   Issuer        `json:"issuer"`
// 	Receiver                 Receiver      `json:"receiver"`
// 	DocumentType             string        `json:"documentType"`
// 	DocumentTypeVersion      string        `json:"documentTypeVersion"`
// 	DateTimeIssued           string        `json:"dateTimeIssued"`
// 	TaxpayerActivityCode     string        `json:"taxpayerActivityCode"`
// 	InternalID               string        `json:"internalID"`
// 	InvoiceLines             []InvoiceLine `json:"invoiceLines"`
// 	TotalDiscountAmount      float64       `json:"totalDiscountAmount"`
// 	TotalItemsDiscountAmount float64       `json:"totalItemsDiscountAmount"`
// 	NetAmount                float64       `json:"netAmount"`
// 	TotalSalesAmount         float64       `json:"totalSalesAmount"`
// 	ExtraDiscountAmount      float64       `json:"extraDiscountAmount"`
// 	TotalAmount              float64       `json:"totalAmount"`
// 	TaxTotals                []TaxTotals   `json:"taxTotals"`
// }

func _loadCompnayInfoIntoInvoice(info *model.CompanyInfo, invoice *model.Invoice) {
	invoice.Issuer.Id = info.EtaRegistrationId
	invoice.Issuer.Type = info.EtaType
	invoice.TaxpayerActivityCode = info.EtaActivityCode
	invoice.Issuer.Name = info.ComName
	invoice.Receiver.Type = "P"
	invoice.DocumentType = "I"
	invoice.DocumentTypeVersion = "1.0"
}
func _roundFloat(val *float64, precision uint) *float64 {
	ratio := math.Pow(10, float64(precision))
	roundedValue := math.Round(*val*ratio) / ratio
	return &roundedValue
}
func _prepareInvoiceItem(item *model.InvoiceLine) {
	item.NetTotal = item.Total
	item.UnitValue.CurrencySold = "EGP"
	taxAmount := item.SalesTotal * .14
	taxAmountRounded := _roundFloat(&taxAmount, 5)
	tax := model.TaxableItems{TaxType: "T1", Amount: *taxAmountRounded, SubType: " ", Rate: 14}
	item.TaxableItems = append(item.TaxableItems, tax)
	// item.UnitValue.TaxableItems.Ta = "EGP"
}
func (ur *InvoiceRepo) FindInvoiceData(req *model.PostInvoicessRequest, companyInfo *model.CompanyInfo) (*[]model.Invoice, error) {
	var invoices []model.Invoice
	// var invoicesLines []model.InvoiceItem
	rows, err := ur.db.Raw("EXEC StkTrEInvoiceFind @serials = ? , @store = ? ", req.Serilas, req.Store).Rows()
	if utils.CheckErr(&err) {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var rec model.Invoice
		err := rows.Scan(
			&rec.Serial,
			&rec.DateTimeIssued,
			&rec.InternalID,
			&rec.TotalDiscountAmount,
			&rec.TotalAmount,
			&rec.Issuer.Address.BranchId,
			&rec.Issuer.Address.Country,
			&rec.Issuer.Address.Governate,
			&rec.Issuer.Address.RegionCity,
			&rec.Issuer.Address.Street,
			&rec.Issuer.Address.BuildingNumber,
		)
		if utils.CheckErr(&err) {
			return nil, err
		}
		_loadCompnayInfoIntoInvoice(companyInfo, &rec)
		invoices = append(invoices, rec)
	}
	err = ur.db.ScanRows(rows, &invoices)
	if rows.NextResultSet() {
		var headSerial int
		var counter int

		// currentInvoice := invoices[counter]
		for rows.Next() {
			var rec model.InvoiceLine
			// var head int
			err := rows.Scan(
				&headSerial,
				&rec.ItemType,
				&rec.ItemCode,
				&rec.UnitType,
				&rec.Quantity,
				&rec.UnitValue.AmountEGP,
				&rec.ItemsDiscount,
				&rec.SalesTotal,
				&rec.Total,
			)
			_prepareInvoiceItem(&rec)
			if utils.CheckErr(&err) {
				return nil, err
			}

			if invoices[counter].Serial != headSerial {
				counter++
			}
			invoices[counter].InvoiceLines = append(invoices[counter].InvoiceLines, rec)

		}
	}
	return &invoices, nil
}
func scanEInvoiceResult(rows *sql.Rows) (*[]model.EInvoice, error) {
	var resp []model.EInvoice
	for rows.Next() {
		var rec model.EInvoice
		err := rows.Scan(&rec.Serial, &rec.InternalID, &rec.StoreCode, &rec.TotalDiscountAmount, &rec.TotalAmount, &rec.TotalTax, &rec.StkTr01Serial, &rec.DateTimeIssued)
		if utils.CheckErr(&err) {
			return nil, err
		}
		resp = append(resp, rec)
	}
	return &resp, nil
}
