package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)


type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb ) ByID(id string) (*Customer,error) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	row:= d.client.QueryRow(customerSql,id)
	var c Customer
	err:= row.Scan(&c.Id,&c.Name,&c.City,&c.Zipcode,&c.DateofBirth,&c.Status)
	if err != nil {
		log.Println("Hit an error in querying a record of the customers "+err.Error())
		return nil,err
	}
	return &c,nil

}

func (d CustomerRepositoryDb) FindAll() ([]Customer , error) {


// Query from Sql 
findAll := "select * from customers"
rows,err := d.client.Query(findAll)
if err != nil {
	log.Println("Hit an error in querying the rows map in sql"+err.Error())
	return nil,err
}

// Iterating through the table 
customers := make([]Customer , 0 )

for rows.Next() {
	var c Customer
	err := rows.Scan(&c.Id,&c.Name,&c.City,&c.Zipcode,&c.DateofBirth,&c.Status)
	if err != nil {
		log.Println("Hit an error in querying a record of the customers "+err.Error())
		return nil,err
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