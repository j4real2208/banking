package service

import (
	"github.com/j4real2208/banking/domain"
	"github.com/j4real2208/banking/errs"
)
type CustomerService interface {
	GetAllCustomer() ([]domain.Customer , *errs.AppError)
	GetCustomer(string) (*domain.Customer , *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer ,*errs.AppError) {
	return s.repo.FindAll()
}


func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer ,*errs.AppError) {
	return s.repo.ByID(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService{
	return DefaultCustomerService{repository}
}