package service

import (
	"time"

	"github.com/j4real2208/banking/domain"
	"github.com/j4real2208/banking/dto"
	"github.com/j4real2208/banking/errs"
)

const dbTSLayout = "2006-01-02 15:04:05"


type AccountService interface {
	NewAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse , *errs.AppError)
	MakeTransaction(request dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
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


func (s DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse , *errs.AppError)  {
	// Validating the trnx amount and type is valid
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	// Server side validation 
	if req.IsTransactionTypeWithdrawal() {
		account , err := s.repo.FindBy(req.AccountId)
		if err != nil {
			return nil, err
		}
		if !account.CanWithdraw(req.Amount){
			return nil , errs.NewValidationError("Insufficeint Funds to intiate the transaction")
		}
	}
	// If all condition holds create a dto to execute the tranx
	t := domain.Transaction{
		AccountId: req.AccountId,
		Amount: req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format(dbTSLayout),
	}

	transaction , appError := s.repo.SaveTransaction(t)
	if appError != nil {
		return nil , appError
	}

	response := transaction.ToDto()
	return &response , nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}