package repo

import (
	"database/sql"
	"eta/model"
	"eta/utils"

	"github.com/jinzhu/gorm"
)

type OrderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) OrderRepo {
	return OrderRepo{
		db: db,
	}
}

func (ur *OrderRepo) ListUnConverted() (*[]model.Order, error) {
	rows, err := ur.db.Raw("EXEC StkTr01ListUnConverted").Rows()
	if utils.CheckErr(&err) {
		return nil, err
	}
	defer rows.Close()
	result, err := scanOrderResult(rows)
	if utils.CheckErr(&err) {
		return nil, err
	}
	return result, nil
}

func (ur *OrderRepo) ListByTransSerialAndConverted(req *model.ListOrdersRequest) (*[]model.Order, error) {
	rows, err := ur.db.Raw("EXEC StkTr01ListByTransSerialAndConverted @TransSerial = ? , @Converted = ? , @StoreCode = ?", req.TransSerial, req.Status, req.Store).Rows()
	if utils.CheckErr(&err) {
		return nil, err
	}
	defer rows.Close()
	result, err := scanOrderResult(rows)
	if utils.CheckErr(&err) {
		return nil, err
	}
	return result, nil
}

func (ur *OrderRepo) ConvertToEta(serial *int64) (*int, error) {
	var resp int
	err := ur.db.Raw("EXEC StkTr01ConvertInvoice @Serial = ? ", serial).Row().Scan(&resp)
	if utils.CheckErr(&err) {
		return nil, err
	}
	return &resp, nil
}

func scanOrderResult(rows *sql.Rows) (*[]model.Order, error) {
	var resp []model.Order
	for rows.Next() {
		var rec model.Order
		err := rows.Scan(&rec.Serial, &rec.DocNo, &rec.DocDate, &rec.Discount, &rec.TotalCash, &rec.TotalTax)
		if utils.CheckErr(&err) {
			return nil, err
		}
		resp = append(resp, rec)

	}

	return &resp, nil
}
