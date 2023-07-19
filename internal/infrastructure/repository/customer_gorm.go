package repository

import (
	"log"

	"github.com/Risminator/gog-taxi-golang/internal/domain/model"
	"github.com/Risminator/gog-taxi-golang/internal/usecase"
	"gorm.io/gorm"
)

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) usecase.CustomerRepository {
	return &customerRepository{db}
}

func (cr *customerRepository) GetAllCustomers() ([]model.Customer, error) {
	var customers []model.Customer
	err := cr.db.Table("gog_demo.customer").Find(&customers).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return customers, nil
}

func (cr *customerRepository) GetCustomerByID(ID int) (*model.Customer, error) {
	var customer model.Customer
	err := cr.db.Table("gog_demo.customer").First(&customer, "customer_id = ?", ID).Error

	if err != nil {
		return nil, err
	}

	return &customer, nil
}
