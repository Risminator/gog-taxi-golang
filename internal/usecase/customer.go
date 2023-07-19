package usecase

import (
	"github.com/Risminator/gog-taxi-golang/internal/domain/model"
)

type Customer interface {
	GetAllCustomers() ([]model.Customer, error)
	GetCustomerByID(ID int) (*model.Customer, error)
}

type CustomerRepository interface {
	GetAllCustomers() ([]model.Customer, error)
	GetCustomerByID(ID int) (*model.Customer, error)
}

type customerUsecase struct {
	customerRepository CustomerRepository
}

func (c *customerUsecase) GetAllCustomers() ([]model.Customer, error) {
	return c.customerRepository.GetAllCustomers()
}

func (c *customerUsecase) GetCustomerByID(ID int) (*model.Customer, error) {
	customer, err := c.customerRepository.GetCustomerByID(ID)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func NewCustomerUsecase(c CustomerRepository) Customer {
	return &customerUsecase{c}
}
