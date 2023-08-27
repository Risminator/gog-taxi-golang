package repository

import (
	"log"

	"github.com/Risminator/gog-taxi-golang/internal/domain/model"
	"github.com/Risminator/gog-taxi-golang/internal/usecase"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const driverTableName = "gog_demo.driver"

type driverRepository struct {
	db *gorm.DB
}

func NewDriverRepository(db *gorm.DB) usecase.DriverRepository {
	return &driverRepository{db}
}

func (dr *driverRepository) GetAllDrivers() ([]model.Driver, error) {
	var drivers []model.Driver
	err := dr.db.Table(driverTableName).Find(&drivers).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return drivers, nil
}

func (dr *driverRepository) GetDriverByID(ID int) (*model.Driver, error) {
	var driver model.Driver
	err := dr.db.Table(driverTableName).First(&driver, "driver_id = ?", ID).Error
	if err != nil {
		return nil, err
	}

	return &driver, nil
}

func (dr *driverRepository) CreateDriver(driver *model.Driver) (int, error) {
	err := dr.db.Table(driverTableName).Select("first_name", "last_name", "vessel_id", "status", "balance", "cert_first_aid", "cert_driving").Create(driver).Error
	if err != nil {
		return 0, err
	}

	return driver.DriverId, nil
}

func (dr *driverRepository) UpdateDriver(driver *model.Driver) error {
	err := dr.db.Clauses(clause.Returning{}).Table(driverTableName).Updates(&driver).Error
	if err != nil {
		return err
	}

	return nil
}

func (dr *driverRepository) DeleteDriver(ID int) (*model.Driver, error) {
	var driver model.Driver
	err := dr.db.Clauses(clause.Returning{}).Table(driverTableName).Delete(&driver, ID).Error
	if err != nil {
		return nil, err
	}

	return &driver, nil
}
