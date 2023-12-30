package repo

import (
	"encoding/json"
	"eta/model"
	"eta/utils"
	"fmt"
	"strconv"
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

func (ur *ReceiptRepo) EInvoiceListUnposted() ([]model.UnpostedReceipts, *model.ReceiptSubmitRequest, error) {
	var resp []model.UnpostedReceipts
	rows, err := ur.db.Raw("EXEC StkTrEInvoiceListUnpostedReceipts").Rows()
	if utils.CheckErr(&err) {
		return nil, nil, err
	}
	defer rows.Close()
	var body model.ReceiptSubmitRequest
	var receipts []model.Receipt
	for rows.Next() {
		var rec model.UnpostedReceipts
		var receipt model.Receipt
		var reqBody string
		err := rows.Scan(
			&rec.Serial,
			&reqBody,
		)

		if utils.CheckErr(&err) {
			return nil, &body, err
		}

		err = json.Unmarshal([]byte(reqBody), &receipt)
		if utils.CheckErr(&err) {
			return nil, &body, err
		}
		rec.RequestBody = receipt
		receipts = append(receipts, rec.RequestBody)

		resp = append(resp, rec)
	}
	body = model.ReceiptSubmitRequest{
		Receipts: receipts,
	}
	return resp, &body, nil
}
func (ur *ReceiptRepo) ReceiptUpdate(req model.ReceiptUpdateRequest) error {
	rows, err := ur.db.Raw("EXEC StkTrEInvoiceHeadUpdate @Serial = ?,@requestBody =?  , @uuid = ? , @posted = ? ", req.Serial, req.RequestBody, req.Uuid, req.Posted).Rows()
	if utils.CheckErr(&err) {
		return err
	}
	defer rows.Close()
	if utils.CheckErr(&err) {
		return err
	}
	return nil
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
func (ur *ReceiptRepo) FindReceiptData(req *model.GenerateUUIDRequest, companyInfo *model.CompanyInfo) (*model.Receipt, *model.EtaAuthunticatePOSRequest, int, error) {
	var serial int
	// var internalId string
	var clientId string
	var clientSecret string
	var totalTax float32

	// var addressSerial int
	strSerial := strconv.Itoa(req.Serial)
	rows, err := ur.db.Raw("EXEC StkTrEInvoiceFind @serials = ? , @store = ? ", strSerial, req.Store).Rows()

	if utils.CheckErr(&err) {
		return nil, nil, serial, err
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
			return nil, nil, serial, err
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
			if utils.CheckErr(&err) {
				return nil, nil, serial, err
			}
			reciept.ItemData = append(reciept.ItemData, rec)

		}
	}

	loginRequest := model.EtaAuthunticatePOSRequest{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		PosOsVersion: "os",
		PosSerial:    reciept.Seller.DeviceSerialNumber,
		Presharedkey: "",
	}
	return &reciept, &loginRequest, serial, nil
}
