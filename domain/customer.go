package domain

import "github.com/j4real2208/banking/errs"
type Customer struct {
	Id string
	Name string
	City string
	Zipcode string
	DateofBirth string
	Status string
}

type CustomerRepository interface {
	FindAll() ([]Customer , *errs.AppError)
	ByID(string) (*Customer , *errs.AppError )
}
