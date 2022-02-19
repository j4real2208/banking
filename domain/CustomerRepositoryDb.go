package domain

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/j4real2208/banking/errs"
	"github.com/j4real2208/banking/logger"
	"github.com/jmoiron/sqlx"
)


type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb ) ByID(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	
	var c Customer
	err := d.client.Get(&c,customerSql,id)
	
	if err != nil {
		if err==sql.ErrNoRows{
			logger.Error("Error while scanning customer" + err.Error())
			return nil,errs.NewNotFoundError("customer not found")
		}else{
			logger.Error("Hit an error in querying while querying from db possible db outage "+err.Error())
			return nil , errs.NewUnexpectedError("unexpected db error")
		}
	}
	return &c,nil

}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer , *errs.AppError) {
	
	
	var err error
	customers := make([]Customer , 0 )
if status == ""{
	findAll := "select * from customers"
	err = d.client.Select(&customers,findAll)
	// rows,err = d.client.Query(findAll)
}else{
	findAll := "select * from customers where status = ?"
	err = d.client.Select(&customers,findAll,status)
}
// Query from Sql 
if err != nil {
	logger.Error("Hit an error in querying the rows of customer table in sql possible db outage"+err.Error())
	return nil,errs.NewUnexpectedError("Unexpected error from db side")
}

return customers,nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb{
	return CustomerRepositoryDb{dbClient}
}