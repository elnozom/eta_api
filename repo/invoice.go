package repo

import (
	"database/sql"
	"eta/model"
	"eta/utils"

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

func (ur *InvoiceRepo) ListEInvoicesByPosted(posted *bool) (*[]model.EInvoice, error) {
	rows, err := ur.db.Raw("EXEC StkTrEInvoiceHeadList @posted = ? ", posted).Rows()
	utils.CheckErr(&err)
	defer rows.Close()
	if utils.CheckErr(&err) {
		return nil, err
	}
	result, err := scanEInvoiceResult(rows)
	return result, nil
}

func (ur *InvoiceRepo) FindInvoiceData(serial *uint64, invoice *model.Invoice) error {
	var accountSerial int
	var addressSerial int
	err := ur.db.Raw("EXEC StkTrEInvoiceFind @serial = ? ", serial).Row().Scan(
		&invoice.DateTimeIssued,
		&invoice.InternalId,
		&invoice.TotalSalesAmount,
		&invoice.TotalItemsDiscountAmount,
		&invoice.TotalDiscountAmount,
		&invoice.ExtraDiscountAmount,
		&invoice.NetAmount,
		&invoice.TotalAmount,
		&accountSerial,
		&invoice.Issuer.Address.BranchId,
		&invoice.Issuer.Id,
		&invoice.Issuer.Name,
		&invoice.Issuer.Type,
		&invoice.TaxpayerActivityCode,
	)
	if utils.CheckErr(&err) {
		return err
	}
	err = ur.db.Raw("EXEC StkTrEInvoiceFindRecieverAddress  @accountSerial = ? ", accountSerial).Row().Scan(
		&addressSerial,
		&invoice.Reciever.Address.Country,
		&invoice.Reciever.Address.Governate,
		&invoice.Reciever.Address.RegionCity,
		&invoice.Reciever.Address.Street,
		&invoice.Reciever.Address.BuildingNumber,
		&invoice.Reciever.Name,
		&invoice.Reciever.Type,
		&invoice.Reciever.Id,
	)
	if utils.CheckErr(&err) {
		return err
	}
	err = ur.db.Raw("EXEC StkTrEInvoiceFindIssuerAddress  @storeCode = ? ", invoice.Issuer.Address.BranchId).Row().Scan(
		&addressSerial,
		&invoice.Issuer.Address.Country,
		&invoice.Issuer.Address.Governate,
		&invoice.Issuer.Address.RegionCity,
		&invoice.Issuer.Address.Street,
		&invoice.Issuer.Address.BuildingNumber,
	)
	if utils.CheckErr(&err) {
		return err
	}
	rows, err := ur.db.Raw("EXEC StkTrEInvoiceFindItems  @serial = ? ", serial).Rows()
	if utils.CheckErr(&err) {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var line model.InvoiceLine
		err = rows.Scan(
			&line.ItemType,
			&line.ItemCode,
			&line.UnitType,
			&line.Quantity,
			&line.UnitValue.AmountEGP,
			&line.TotalTaxableFees,
			&line.ItemsDiscount,
			&line.SalesTotal,
			&line.Total,
			&line.NetTotal,
		)
		if utils.CheckErr(&err) {
			return err
		}
		line.UnitValue.CurrencySold = "EGP"
		invoice.InvoiceLines = append(invoice.InvoiceLines, line)
	}
	return nil
}

func scanEInvoiceResult(rows *sql.Rows) (*[]model.EInvoice, error) {
	var resp []model.EInvoice
	for rows.Next() {
		var rec model.EInvoice
		err := rows.Scan(&rec.Serial, &rec.InternalID, &rec.StoreCode, &rec.TotalDiscountAmount, &rec.TotalAmount, &rec.TotalTax, &rec.StkTr01Serial)
		if utils.CheckErr(&err) {
			return nil, err
		}
		resp = append(resp, rec)
	}
	return &resp, nil
}
