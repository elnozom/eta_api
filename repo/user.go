package repo

import (
	"database/sql"
	"eta/model"
	"eta/utils"

	"github.com/jinzhu/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return UserRepo{
		db: db,
	}
}

func (ur *UserRepo) GetByCode(code *uint) (*model.User, error) {
	row := ur.db.Raw("EXEC GetEmp @EmpCode = ?", code).Row()
	user, err := scanUserResult(row)
	utils.CheckErr(&err)
	return user, nil
}

func scanUserResult(row *sql.Row) (*model.User, error) {
	var user model.User
	row.Scan(&user.EmpName, &user.EmpPassword, &user.EmpCode, &user.SecLevel, &user.FixEmpStore)
	return &user, nil
}
