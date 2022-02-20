package service

import (
	"time"

	"github.com/j4real2208/banking/domain"
	"github.com/j4real2208/banking/dto"
	"github.com/j4real2208/banking/errs"
)

type AccountService interface {
	NewAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse , *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse , *errs.AppError) {
	//logger.Info("--------Inside account account service and going for validation-----")
	err := req.Validate()
	if err != nil{
		return nil , err
	}
	//logger.Info("--------Inside account account service and validated entries making domain object-----")
	a:= domain.Account{
		AccountId: "",
		CustomerId: req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount: req.Amount,
		Status: "1",
	}
	//logger.Info("Printing the customer_id again"+req.CustomerId+" Solutions")
	newAccount , err:= s.repo.Save(a)
	if err != nil {
		return nil, err
	}
	//logger.Info("--------Inside account account service and saved account and came back-----")
	responseDto:= newAccount.ToNewAccountResponseDto()

	return responseDto, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}