package repository

import (
	"log"

	"github.com/Risminator/gog-taxi-golang/internal/domain/model"
	"github.com/Risminator/gog-taxi-golang/internal/usecase"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const vesselTableName = "gog_demo.vessel"

type vesselRepository struct {
	db *gorm.DB
}

func NewVesselRepository(db *gorm.DB) usecase.VesselRepository {
	return &vesselRepository{db}
}

func (vr *vesselRepository) GetAllVessels() ([]model.Vessel, error) {
	var vessels []model.Vessel
	err := vr.db.Table(vesselTableName).Find(&vessels).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return vessels, nil
}

func (vr *vesselRepository) GetVesselByID(ID int) (*model.Vessel, error) {
	var vessel model.Vessel
	err := vr.db.Table(vesselTableName).First(&vessel, "vessel_id = ?", ID).Error
	if err != nil {
		return nil, err
	}

	return &vessel, nil
}

func (vr *vesselRepository) GetVesselByRequestId(reqID int) (*model.Vessel, error) {
	var vessel model.Vessel
	err := vr.db.Raw(
		"SELECT v.* FROM gog_demo.vessel v JOIN gog_demo.driver dr ON v.vessel_id=dr.vessel_id JOIN gog_demo.taxi_request r ON dr.driver_id = r.driver_id WHERE r.taxi_request_id=?",
		reqID).Scan(&vessel).Error
	if err != nil {
		return nil, err
	}

	return &vessel, nil
}

func (vr *vesselRepository) CreateVessel(vessel *model.Vessel) (int, error) {
	err := vr.db.Table(vesselTableName).Select("model", "seats", "is_approved", "latitude", "longitude").Create(vessel).Error
	if err != nil {
		return 0, err
	}

	return vessel.VesselId, nil
}

func (vr *vesselRepository) UpdateVessel(vessel *model.Vessel) error {
	err := vr.db.Clauses(clause.Returning{}).Table(vesselTableName).Updates(&vessel).Error
	if err != nil {
		return err
	}

	return nil
}

func (vr *vesselRepository) DeleteVessel(ID int) (*model.Vessel, error) {
	var vessel model.Vessel
	err := vr.db.Clauses(clause.Returning{}).Table(vesselTableName).Delete(&vessel, ID).Error
	if err != nil {
		return nil, err
	}

	return &vessel, nil
}
