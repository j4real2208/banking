package service

import (
	"github.com/j4real2208/banking/domain"
	"github.com/j4real2208/banking/dto"
	"github.com/j4real2208/banking/errs"
)
type CustomerService interface {
	GetAllCustomer(string) ([]domain.Customer , *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse , *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer ,*errs.AppError) {
	if status == "active"{
		status="1"
	}else if status == "inactive"{
		status="0"
	}else{
		status=""
	}
	return s.repo.FindAll(status)
}


func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse,*errs.AppError) {
	c, err :=  s.repo.ByID(id)
	if err != nil {
		return nil, err
	}
	response:= c.ToDto()
	return &response, nil

}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService{
	return DefaultCustomerService{repository}
}