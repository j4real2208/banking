package domain

import (
	"github.com/j4real2208/banking/dto"
	"github.com/j4real2208/banking/errs"
)

type Account struct {
	AccountId string
	CustomerId string
	OpeningDate string
	AccountType string
	Amount float64
	Status string
}

func (a  Account) ToNewAccountResponseDto() *dto.NewAccountResponse {
	return &dto.NewAccountResponse{a.AccountId}
}
type AccountRepository interface {
	Save(account Account) (*Account , *errs.AppError)
}