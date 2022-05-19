package repo

import (
	"eta/model"
	"eta/utils"

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

func (ur *ReceiptRepo) FindReceiptData(serial *uint64, invoice *model.Receipt) error {
	// var accountSerial int
	// var addressSerial int
	// err := ur.db.Raw("EXEC StkTrEInvoiceFind @serial = ? ", serial).Row().Scan(
	// 	&invoice.DateTimeIssued,
	// 	&invoice.InternalId,
	// 	&invoice.TotalSalesAmount,
	// 	&invoice.TotalItemsDiscountAmount,
	// 	&invoice.TotalDiscountAmount,
	// 	&invoice.ExtraDiscountAmount,
	// 	&invoice.NetAmount,
	// 	&invoice.TotalAmount,
	// 	&accountSerial,
	// 	&invoice.Issuer.Address.BranchId,
	// 	&invoice.Issuer.Id,
	// 	&invoice.Issuer.Name,
	// 	&invoice.Issuer.Type,
	// 	&invoice.TaxpayerActivityCode,
	// )
	// if utils.CheckErr(&err) {
	// 	return err
	// }
	// err = ur.db.Raw("EXEC StkTrEInvoiceFindRecieverAddress  @accountSerial = ? ", accountSerial).Row().Scan(
	// 	&addressSerial,
	// 	&invoice.Reciever.Address.Country,
	// 	&invoice.Reciever.Address.Governate,
	// 	&invoice.Reciever.Address.RegionCity,
	// 	&invoice.Reciever.Address.Street,
	// 	&invoice.Reciever.Address.BuildingNumber,
	// 	&invoice.Reciever.Name,
	// 	&invoice.Reciever.Type,
	// 	&invoice.Reciever.Id,
	// )
	// if utils.CheckErr(&err) {
	// 	return err
	// }
	// err = ur.db.Raw("EXEC StkTrEInvoiceFindIssuerAddress  @storeCode = ? ", invoice.Issuer.Address.BranchId).Row().Scan(
	// 	&addressSerial,
	// 	&invoice.Issuer.Address.Country,
	// 	&invoice.Issuer.Address.Governate,
	// 	&invoice.Issuer.Address.RegionCity,
	// 	&invoice.Issuer.Address.Street,
	// 	&invoice.Issuer.Address.BuildingNumber,
	// )
	// if utils.CheckErr(&err) {
	// 	return err
	// }
	// rows, err := ur.db.Raw("EXEC StkTrEInvoiceFindItems  @serial = ? ", serial).Rows()
	// if utils.CheckErr(&err) {
	// 	return err
	// }
	// defer rows.Close()
	// for rows.Next() {
	// 	var line model.InvoiceLine
	// 	err = rows.Scan(
	// 		&line.ItemType,
	// 		&line.ItemCode,
	// 		&line.UnitType,
	// 		&line.Quantity,
	// 		&line.UnitValue.AmountEGP,
	// 		&line.TotalTaxableFees,
	// 		&line.ItemsDiscount,
	// 		&line.SalesTotal,
	// 		&line.Total,
	// 		&line.NetTotal,
	// 	)
	// 	if utils.CheckErr(&err) {
	// 		return err
	// 	}
	// 	line.UnitValue.CurrencySold = "EGP"
	// 	invoice.InvoiceLines = append(invoice.InvoiceLines, line)
	// }
	return nil
}
