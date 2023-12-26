package repo

import (
	"eta/model"
	"eta/utils"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

type ReceiptRepo struct {
	db *gorm.DB
}

func NewReceiptRepo(db *gorm.DB) ReceiptRepo {
	return ReceiptRepo{
		db: db,
	}
}

func (ur *ReceiptRepo) ListReceiptsByPosted(posted *bool) (*[]model.EInvoice, error) {
	rows, err := ur.db.Raw("EXEC StkTrEInvoiceHeadList @posted = ? ", posted).Rows()
	utils.CheckErr(&err)
	defer rows.Close()
	if utils.CheckErr(&err) {
		return nil, err
	}
	result, err := scanEInvoiceResult(rows)
	return result, nil
}

func _prepareReceipt(info *model.CompanyInfo, serial int, receipt *model.Receipt, storeCode *int) {
	receipt.Seller.Rin = info.EtaRegistrationId
	internalID := fmt.Sprintf("%d-%d", *storeCode, serial)
	receipt.Header.ReceiptNumber = internalID
	// receipt.Seller.Type = info.EtaType
	// receipt.TaxpayerActivityCode = info.EtaActivityCode
	receipt.Seller.CompanyTradeName = info.ComName
	receipt.Seller.ActivityCode = info.EtaActivityCode
	receipt.Buyer.Type = "P"
	receipt.DocumentType.ReceiptType = "S"
	receipt.DocumentType.TypeVersion = "1.2"
	receipt.PaymentMethod = "C"
	// receipt.DocumentTypeVersion = "1.0"

	receipt.NetAmount = _removeTax(&receipt.TotalAmount)
	taxAmount := receipt.TotalAmount - receipt.NetAmount
	tax := model.TaxTotal{TaxType: "T1", Amount: _roundFloat(&taxAmount, 5)}
	receipt.TaxTotals = append(receipt.TaxTotals, tax)
	receipt.TotalSales = receipt.NetAmount

}
func (ur *ReceiptRepo) FindReceiptData(req *model.PostInvoicessRequest, companyInfo *model.CompanyInfo) (*model.Receipt, error) {
	var serial int
	// var internalId string
	var clientId string
	var clientSecret string
	var totalTax float32

	// var addressSerial int
	rows, err := ur.db.Raw("EXEC StkTrEInvoiceFind @serials = ? , @store = ? ", req.Serilas, req.Store).Rows()

	if utils.CheckErr(&err) {
		return nil, err
	}
	var reciept model.Receipt
	for rows.Next() {
		err := rows.Scan(
			&serial,
			&reciept.Header.DateTimeIssued,
			&reciept.Header.ReceiptNumber,
			&reciept.Header.Currency,
			&reciept.TotalCommersialDiscount,
			&reciept.TotalAmount,
			&reciept.Seller.DeviceSerialNumber,
			&reciept.Seller.BranchCode,
			&totalTax,
			&reciept.NetAmount,
			&reciept.Seller.BranchAddress.Country,
			&reciept.Seller.BranchAddress.Governate,
			&reciept.Seller.BranchAddress.RegionCity,
			&reciept.Seller.BranchAddress.Street,
			&reciept.Seller.BranchAddress.BuildingNumber,
			&clientId,
			&clientSecret,
		)
		if utils.CheckErr(&err) {
			return nil, err
		}

		reciept.Header.DateTimeIssued = fixDate(reciept.Header.DateTimeIssued)

	}
	_prepareReceipt(companyInfo, serial, &reciept, &req.Store)
	if rows.NextResultSet() {
		var headSerial int
		var discount float32
		// var counter int

		// currentInvoice := invoices[counter]
		for rows.Next() {
			var rec model.ItemData
			var rate float64
			var totalVat float64
			// var head int
			err := rows.Scan(
				&headSerial,
				&rec.Description,
				&rec.ItemType,
				&rec.ItemCode,
				&rate,
				&rec.UnitType,
				&rec.Quantity,
				&rec.UnitPrice,
				&discount,
				&rec.TotalSale,
				&rec.Total,
				&totalVat,
			)
			rec.NetSale = rec.TotalSale
			taxItem := model.TaxableItems{
				TaxType: "T1",
				Amount:  totalVat,
				SubType: "V009",
				Rate:    rate,
			}
			taxItems := []model.TaxableItems{taxItem}
			rec.ItemType = strings.ReplaceAll(rec.ItemType, " ", "")
			rec.TaxableItems = taxItems
			rec.InternalCode = rec.ItemCode
			// _prepareInvoiceItem(&rec)
			if utils.CheckErr(&err) {
				return nil, err
			}
			reciept.ItemData = append(reciept.ItemData, rec)
			// if invoices[counter].Serial != headSerial {
			// 	counter++
			// }
			// invoices[counter].InvoiceLines = append(invoices[counter].InvoiceLines, rec)
			// invoices[counter].TaxTotals[0].Amount += rec.TaxableItems[0].Amount
			// invoices[counter].TaxTotals[0].Amount = _roundFloat(&invoices[counter].TaxTotals[0].Amount, 5)
		}
	}
	return &reciept, nil
}

// func (ur *InvoiceRepo) FindInvoiceData(req *model.PostInvoicessRequest, companyInfo *model.CompanyInfo) ([]model.Invoice, error) {
// 	var invoices []model.Invoice
// 	// var invoicesLines []model.InvoiceItem
// 	rows, err := ur.db.Raw("EXEC StkTrEInvoiceFind @serials = ? , @store = ? ", req.Serilas, req.Store).Rows()
// 	if utils.CheckErr(&err) {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var rec model.Invoice
// 		err := rows.Scan(
// 			&rec.Serial,
// 			&rec.DateTimeIssued,
// 			&rec.InternalID,
// 			&rec.TotalDiscountAmount,
// 			&rec.TotalAmount,
// 			&rec.Issuer.Address.BranchId,
// 			&rec.Issuer.Address.Country,
// 			&rec.Issuer.Address.Governate,
// 			&rec.Issuer.Address.RegionCity,
// 			&rec.Issuer.Address.Street,
// 			&rec.Issuer.Address.BuildingNumber,
// 		)
// 		if utils.CheckErr(&err) {
// 			return nil, err
// 		}
// 		_prepareInvoice(companyInfo, &rec, &req.Store)
// 		invoices = append(invoices, rec)
// 	}
// 	err = ur.db.ScanRows(rows, &invoices)
// 	if rows.NextResultSet() {
// 		var headSerial int
// 		var counter int

// 		// currentInvoice := invoices[counter]
// 		for rows.Next() {
// 			var rec model.InvoiceLine
// 			// var head int
// 			err := rows.Scan(
// 				&headSerial,
// 				&rec.Description,
// 				&rec.ItemType,
// 				&rec.ItemCode,
// 				&rec.UnitType,
// 				&rec.Quantity,
// 				&rec.UnitValue.AmountEGP,
// 				&rec.ItemsDiscount,
// 				&rec.SalesTotal,
// 				&rec.Total,
// 			)
// 			_prepareInvoiceItem(&rec)
// 			if utils.CheckErr(&err) {
// 				return nil, err
// 			}

// 			if invoices[counter].Serial != headSerial {
// 				counter++
// 			}
// 			invoices[counter].InvoiceLines = append(invoices[counter].InvoiceLines, rec)
// 			invoices[counter].TaxTotals[0].Amount += rec.TaxableItems[0].Amount
// 			invoices[counter].TaxTotals[0].Amount = _roundFloat(&invoices[counter].TaxTotals[0].Amount, 5)
// 		}
// 	}
// 	return invoices, nil
// }
