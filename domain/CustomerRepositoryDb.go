package domain

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/j4real2208/banking/errs"
	"github.com/j4real2208/banking/logger"
)


type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb ) ByID(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	row:= d.client.QueryRow(customerSql,id)
	var c Customer
	err:= row.Scan(&c.Id,&c.Name,&c.City,&c.Zipcode,&c.DateofBirth,&c.Status)
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

var rows *sql.Rows
var err error
if status == ""{
	findAll := "select * from customers"
	rows,err = d.client.Query(findAll)
}else{
	findAll := "select * from customers where status = ?"
	rows,err = d.client.Query(findAll,status)
}
// Query from Sql 
if err != nil {
	logger.Error("Hit an error in querying the rows map in sql possible db outage"+err.Error())
	return nil,errs.NewUnexpectedError("Unexpected error from db side")
}

// Iterating through the table 
customers := make([]Customer , 0 )

for rows.Next() {
	var c Customer
	err := rows.Scan(&c.Id,&c.Name,&c.City,&c.Zipcode,&c.DateofBirth,&c.Status)
	if err != nil {
		logger.Error("Hit an error in querying a record of the customers "+err.Error())
		return nil,errs.NewUnexpectedError("Unexpected error from db side")
	}
	customers = append(customers, c)
}

return customers,nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb{
	client, err := sql.Open("mysql", "root:codecamp@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client}
}