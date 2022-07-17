package repo

import (
	"eta/model"
	"eta/utils"

	"github.com/jinzhu/gorm"
)

type LogRepo struct {
	db *gorm.DB
}

func NewLogRepo(db *gorm.DB) LogRepo {
	return LogRepo{
		db: db,
	}
}

func (ur *LogRepo) ELogInsert(req *model.Log) (*int, error) {
	var resp int
	err := ur.db.Raw("EXEC EtaLogInsert @internalID = ?,@submissionID = ?, @storeCode = ?, @serials = ?, @logText = ?, @errText = ?, @posted = ?",
		req.InternalID,
		req.SubmissionID,
		req.StoreCode,
		req.Serials,
		req.LogText,
		req.ErrText,
		req.Posted,
	).Row().Scan(&resp)
	if utils.CheckErr(&err) {
		return nil, err
	}
	return &resp, nil
}
