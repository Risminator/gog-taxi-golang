package repository

import (
	"log"

	"github.com/Risminator/gog-taxi-golang/internal/domain/model"
	"github.com/Risminator/gog-taxi-golang/internal/usecase"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const driverTable = "gog_demo.driver"

type driverRepository struct {
	db *gorm.DB
}

func NewDriverRepository(db *gorm.DB) usecase.DriverRepository {
	return &driverRepository{db}
}

func (cr *driverRepository) GetAllDrivers() ([]model.Driver, error) {
	var drivers []model.Driver
	err := cr.db.Table(driverTable).Find(&drivers).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return drivers, nil
}

func (cr *driverRepository) GetDriverByID(ID int) (*model.Driver, error) {
	var driver model.Driver
	err := cr.db.Table(driverTable).First(&driver, "driver_id = ?", ID).Error
	if err != nil {
		return nil, err
	}

	return &driver, nil
}

func (cr *driverRepository) CreateDriver(driver *model.Driver) error {
	err := cr.db.Table(driverTable).Select("phone", "first_name", "last_name").Create(driver).Error
	if err != nil {
		return err
	}

	return nil
}

func (cr *driverRepository) UpdateDriver(driver *model.Driver) error {
	err := cr.db.Clauses(clause.Returning{}).Table(driverTable).Updates(&driver).Error
	if err != nil {
		return err
	}

	return nil
}

func (cr *driverRepository) DeleteDriver(ID int) (*model.Driver, error) {
	var driver model.Driver
	err := cr.db.Clauses(clause.Returning{}).Table(driverTable).Delete(&driver, ID).Error
	if err != nil {
		return nil, err
	}

	return &driver, nil
}
