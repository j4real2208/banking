package domain

import (
	"github.com/j4real2208/banking/dto"
	"github.com/j4real2208/banking/errs"
)
type Customer struct {
	Id string			`db:"customer_id"`
	Name string	
	City string
	Zipcode string
	DateofBirth string	`db:"date_of_birth"`
	Status string
}
// Statu as Text method
func ( c Customer ) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0"{
		statusAsText="inactive"
	}
	return statusAsText
}
// Dto conversion method 
func (c Customer) ToDto() (dto.CustomerResponse)  {
	return dto.CustomerResponse{
		Id:				c.Id,
		Name:			c.Name,
		City:			c.City,
		Zipcode:		c.Zipcode,
		DateofBirth:	c.DateofBirth,
		Status:			c.statusAsText(),
	}	
}

type CustomerRepository interface {
	FindAll(string) ([]Customer , *errs.AppError)
	ByID(string) (*Customer , *errs.AppError )
}
