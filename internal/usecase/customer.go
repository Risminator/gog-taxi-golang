package usecase

import (
	"github.com/Risminator/gog-taxi-golang/internal/domain/model"
)

type Customer interface {
	GetAllCustomers() ([]model.Customer, error)
	GetCustomerByID(ID int) (*model.Customer, error)
	CreateCustomer(phone string, fName string, lName string) (*model.Customer, error)
	UpdateCustomer(ID int, phone string, fName string, lName string) (*model.Customer, error)
	DeleteCustomer(ID int) (*model.Customer, error)
}

type CustomerRepository interface {
	GetAllCustomers() ([]model.Customer, error)
	GetCustomerByID(ID int) (*model.Customer, error)
	CreateCustomer(*model.Customer) (int, error)
	UpdateCustomer(*model.Customer) error
	DeleteCustomer(ID int) (*model.Customer, error)
}

type customerUsecase struct {
	customerRepository CustomerRepository
}

func NewCustomerUsecase(c CustomerRepository) Customer {
	return &customerUsecase{c}
}

func (c *customerUsecase) GetAllCustomers() ([]model.Customer, error) {
	customers, err := c.customerRepository.GetAllCustomers()
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (c *customerUsecase) GetCustomerByID(ID int) (*model.Customer, error) {
	customer, err := c.customerRepository.GetCustomerByID(ID)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (c *customerUsecase) CreateCustomer(phone string, fName string, lName string) (*model.Customer, error) {
	customer := model.CreateCustomer(0, phone, fName, lName)
	id, err := c.customerRepository.CreateCustomer(&customer)
	if err != nil {
		return nil, err
	}

	customer.SetCustomerId(id)
	return &customer, nil
}

func (c *customerUsecase) UpdateCustomer(ID int, phone string, fName string, lName string) (*model.Customer, error) {
	customer, err := c.GetCustomerByID(ID)
	if err != nil {
		return nil, err
	}

	customer.SetPhone(phone)
	customer.SetFirstName(fName)
	customer.SetLastName(lName)

	err = c.customerRepository.UpdateCustomer(customer)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (c *customerUsecase) DeleteCustomer(ID int) (*model.Customer, error) {
	customer, err := c.customerRepository.DeleteCustomer(ID)
	if err != nil {
		return nil, err
	}
	return customer, nil
}
