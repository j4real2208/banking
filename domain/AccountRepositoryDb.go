package domain

import (
	"strconv"

	"github.com/j4real2208/banking/errs"
	"github.com/j4real2208/banking/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb ) Save(a Account) (*Account , *errs.AppError) {
	sqlinsert:= "INSERT INTO accounts (customer_id, opening_date ,account_type , amount , status ) values (?,?,?,?,?)"
	result , err := d.client.Exec(sqlinsert,a.CustomerId,a.OpeningDate,a.AccountType,a.Amount,a.Status)
	if err != nil {
		logger.Error("Error while creating a new account: "+err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	id , err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting the id of the new account: "+err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	a.AccountId = strconv.FormatInt(id,10)
	return &a , nil


}

// Helper func to create new instance 
func NewAccountRepositoryDb(dbClient *sqlx.DB ) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}