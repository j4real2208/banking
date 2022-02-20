package dto

import (
	"strings"

	"github.com/j4real2208/banking/errs"
)

type NewAccountRequest struct {
	CustomerId string 	`json:"customer_id"`
	AccountType string	`json:"account_type"`	
	Amount float64		`json:"amount"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000.00 {
		return errs.NewValidationError("to open a new account you need atleast 5000$ ")
	}
	if strings.ToLower(r.AccountType) != "saving" && strings.ToLower(r.AccountType) != "checking" {
		return errs.NewValidationError("you require a savings account type to start the account")
	}

	return nil
}