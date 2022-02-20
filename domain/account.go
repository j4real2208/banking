package domain

import (
	"github.com/j4real2208/banking/dto"
	"github.com/j4real2208/banking/errs"
)

type Account struct {
	AccountId   string  `db:"account_id"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

func (a  Account) ToNewAccountResponseDto() *dto.NewAccountResponse {
	return &dto.NewAccountResponse{a.AccountId}
}

func (a Account) CanWithdraw(amount float64) bool {
	return  a.Amount >= amount 		
}

type AccountRepository interface {
	Save(account Account) (*Account , *errs.AppError)
	SaveTransaction(transaction Transaction) (*Transaction, *errs.AppError)
	FindBy(accountId string) (*Account, *errs.AppError)
}