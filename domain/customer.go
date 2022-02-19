package domain

import "github.com/j4real2208/banking/errs"
type Customer struct {
	Id string			`db:"customer_id"`
	Name string	
	City string
	Zipcode string
	DateofBirth string	`db:"date_of_birth"`
	Status string
}

type CustomerRepository interface {
	FindAll(string) ([]Customer , *errs.AppError)
	ByID(string) (*Customer , *errs.AppError )
}
