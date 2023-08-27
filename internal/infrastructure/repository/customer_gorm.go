package repository

import (
	"log"

	"github.com/Risminator/gog-taxi-golang/internal/domain/model"
	"github.com/Risminator/gog-taxi-golang/internal/usecase"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const customerTableName = "gog_demo.customer"

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) usecase.CustomerRepository {
	return &customerRepository{db}
}

func (cr *customerRepository) GetAllCustomers() ([]model.Customer, error) {
	var customers []model.Customer
	err := cr.db.Table(customerTableName).Find(&customers).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return customers, nil
}

func (cr *customerRepository) GetCustomerByID(ID int) (*model.Customer, error) {
	var customer model.Customer
	err := cr.db.Table(customerTableName).First(&customer, "customer_id = ?", ID).Error
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (cr *customerRepository) CreateCustomer(customer *model.Customer) (int, error) {
	err := cr.db.Table(customerTableName).Select("phone", "first_name", "last_name").Create(customer).Error
	if err != nil {
		return 0, err
	}

	return customer.CustomerId, nil
}

func (cr *customerRepository) UpdateCustomer(customer *model.Customer) error {
	err := cr.db.Clauses(clause.Returning{}).Table(customerTableName).Updates(&customer).Error
	if err != nil {
		return err
	}

	return nil
}

func (cr *customerRepository) DeleteCustomer(ID int) (*model.Customer, error) {
	var customer model.Customer
	err := cr.db.Clauses(clause.Returning{}).Table(customerTableName).Delete(&customer, ID).Error
	if err != nil {
		return nil, err
	}

	return &customer, nil
}
