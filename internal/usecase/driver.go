package usecase

import (
	"github.com/Risminator/gog-taxi-golang/internal/domain/model"
)

type Driver interface {
	GetAllDrivers() ([]model.Driver, error)
	GetDriverByID(ID int) (*model.Driver, error)
	CreateDriver(fName string, lName string, vesselId int, status model.DriverStatus, balance float64, certFA int, certDr int) (*model.Driver, error)
	UpdateDriver(ID int, fName string, lName string, vesselId int, status model.DriverStatus, balance float64, certFA int, certD int) (*model.Driver, error)
	DeleteDriver(ID int) (*model.Driver, error)
}

type DriverRepository interface {
	GetAllDrivers() ([]model.Driver, error)
	GetDriverByID(ID int) (*model.Driver, error)
	CreateDriver(*model.Driver) error
	UpdateDriver(*model.Driver) error
	DeleteDriver(ID int) (*model.Driver, error)
}

type driverUsecase struct {
	driverRepository DriverRepository
}

func NewDriverUsecase(d DriverRepository) Driver {
	return &driverUsecase{d}
}

func (d *driverUsecase) GetAllDrivers() ([]model.Driver, error) {
	drivers, err := d.driverRepository.GetAllDrivers()
	if err != nil {
		return nil, err
	}
	return drivers, nil
}

func (d *driverUsecase) GetDriverByID(ID int) (*model.Driver, error) {
	driver, err := d.driverRepository.GetDriverByID(ID)
	if err != nil {
		return nil, err
	}
	return driver, nil
}

func (d *driverUsecase) CreateDriver(fName string, lName string, vesselId int, status model.DriverStatus, balance float64, certFA int, certDr int) (*model.Driver, error) {
	driver := model.CreateDriver(0, fName, lName, vesselId, status, balance, certFA, certDr)
	err := d.driverRepository.CreateDriver(&driver)
	if err != nil {
		return nil, err
	}
	return &driver, nil
}

func (d *driverUsecase) UpdateDriver(ID int, fName string, lName string, vesselId int, status model.DriverStatus, balance float64, certFA int, certD int) (*model.Driver, error) {
	driver, err := d.GetDriverByID(ID)
	if err != nil {
		return nil, err
	}

	driver.SetFirstName(fName)
	driver.SetLastName(lName)
	driver.SetVesselId(vesselId)
	driver.SetStatus(status)
	driver.SetBalance(balance)
	driver.SetCertFirstAid(certFA)
	driver.SetCertDriving(certD)

	err = d.driverRepository.UpdateDriver(driver)
	if err != nil {
		return nil, err
	}
	return driver, nil
}

func (d *driverUsecase) DeleteDriver(ID int) (*model.Driver, error) {
	driver, err := d.driverRepository.DeleteDriver(ID)
	if err != nil {
		return nil, err
	}
	return driver, nil
}
