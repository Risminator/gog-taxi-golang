package usecase

import (
	"github.com/Risminator/gog-taxi-golang/internal/domain/model"
)

type Vessel interface {
	GetAllVessels() ([]model.Vessel, error)
	GetVesselByID(ID int) (*model.Vessel, error)
	GetVesselByRequestId(reqID int) (*model.Vessel, error)
	CreateVessel(vModel string, seats int, isAppr bool, lat float64, lon float64) (*model.Vessel, error)
	UpdateVessel(ID int, vModel string, seats int, isAppr bool, lat float64, lon float64) (*model.Vessel, error)
	UpdateLocation(ID int, lat float64, lon float64) (*model.Vessel, error)
	DeleteVessel(ID int) (*model.Vessel, error)
}

type VesselRepository interface {
	GetAllVessels() ([]model.Vessel, error)
	GetVesselByID(ID int) (*model.Vessel, error)
	GetVesselByRequestId(reqID int) (*model.Vessel, error)
	CreateVessel(*model.Vessel) (int, error)
	UpdateVessel(*model.Vessel) error
	DeleteVessel(ID int) (*model.Vessel, error)
}

type vesselUsecase struct {
	vesselRepository VesselRepository
}

func NewVesselUsecase(d VesselRepository) Vessel {
	return &vesselUsecase{d}
}

func (d *vesselUsecase) GetAllVessels() ([]model.Vessel, error) {
	vessels, err := d.vesselRepository.GetAllVessels()
	if err != nil {
		return nil, err
	}
	return vessels, nil
}

func (d *vesselUsecase) GetVesselByID(ID int) (*model.Vessel, error) {
	vessel, err := d.vesselRepository.GetVesselByID(ID)
	if err != nil {
		return nil, err
	}
	return vessel, nil
}

func (d *vesselUsecase) GetVesselByRequestId(reqID int) (*model.Vessel, error) {
	vessel, err := d.vesselRepository.GetVesselByRequestId(reqID)
	if err != nil {
		return nil, err
	}
	return vessel, nil
}

func (d *vesselUsecase) CreateVessel(vModel string, seats int, isAppr bool, lat float64, lon float64) (*model.Vessel, error) {
	vessel := model.CreateVessel(0, vModel, seats, isAppr, lat, lon)
	id, err := d.vesselRepository.CreateVessel(&vessel)
	if err != nil {
		return nil, err
	}

	vessel.SetVesselId(id)
	return &vessel, nil
}

func (d *vesselUsecase) UpdateVessel(ID int, vModel string, seats int, isAppr bool, lat float64, lon float64) (*model.Vessel, error) {
	vessel, err := d.GetVesselByID(ID)
	if err != nil {
		return nil, err
	}

	vessel.SetModel(vModel)
	vessel.SetSeats(seats)
	vessel.SetIsApproved(isAppr)
	vessel.SetLatitude(lat)
	vessel.SetLongitude(lon)

	err = d.vesselRepository.UpdateVessel(vessel)
	if err != nil {
		return nil, err
	}
	return vessel, nil
}

func (d *vesselUsecase) UpdateLocation(ID int, lat float64, lon float64) (*model.Vessel, error) {
	vessel, err := d.GetVesselByID(ID)
	if err != nil {
		return nil, err
	}

	vessel.SetLatitude(lat)
	vessel.SetLongitude(lon)

	err = d.vesselRepository.UpdateVessel(vessel)
	if err != nil {
		return nil, err
	}
	return vessel, nil
}

func (d *vesselUsecase) DeleteVessel(ID int) (*model.Vessel, error) {
	vessel, err := d.vesselRepository.DeleteVessel(ID)
	if err != nil {
		return nil, err
	}
	return vessel, nil
}
