package repository

import (
	"github.com/Risminator/gog-taxi-golang/internal/domain/model"
	"github.com/Risminator/gog-taxi-golang/internal/usecase"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const taxiRequestTableName = "gog_demo.taxi_request"

type taxiRequestRepository struct {
	db *gorm.DB
}

func NewTaxiRequestRepository(db *gorm.DB) usecase.TaxiRequestRepository {
	return &taxiRequestRepository{db}
}

// GetRequestByUserId implements usecase.TaxiRequestRepository.
func (repo *taxiRequestRepository) GetRequestByUserId(id int, role model.UserRole) (*model.TaxiRequest, error) {
	request := &model.TaxiRequest{}
	var err error

	if role == model.CustomerRole {
		err = repo.db.Table(taxiRequestTableName).Find(&request, "status != ? and status != ? and customer_id = ?", model.Canceled, model.Completed, id).Error
	} else {
		err = repo.db.Table(taxiRequestTableName).Find(&request, "status != ? and status != ? and driver_id = ?", model.Canceled, model.Completed, id).Error
	}

	if err != nil {
		return nil, err
	}
	return request, nil
}

// CreateRequest implements usecase.TaxiRequestRepository.
func (repo *taxiRequestRepository) CreateRequest(r *model.TaxiRequest) (int, error) {
	request := map[string]interface{}{
		"customer_id":           r.CustomerId,
		"driver_id":             r.DriverId,
		"departure_id":          r.DepartureId,
		"destination_id":        r.DestinationId,
		"departure_longitude":   r.DepartureLongitude,
		"departure_latitude":    r.DepartureLatitude,
		"destination_longitude": r.DestinationLongitude,
		"destination_latitude":  r.DestinationLatitude,
		"price":                 r.Price,
		"status":                r.Status,
		"planned_time":          r.PlannedTime,
	}

	if r.DriverId == 0 {
		request["driver_id"] = nil
	}

	err := repo.db.Clauses(clause.Returning{}).Table(taxiRequestTableName).Create(&request).Error
	if err != nil {
		return 0, err
	}

	return int(request["taxi_request_id"].(int32)), nil
}

// GetRequestById implements usecase.TaxiRequestRepository.
func (repo *taxiRequestRepository) GetRequestById(id int) (*model.TaxiRequest, error) {
	var req *model.TaxiRequest
	err := repo.db.Table(taxiRequestTableName).First(&req, "taxi_request_id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return req, nil
}

// GetRequestsByStatus implements usecase.TaxiRequestRepository.
func (repo *taxiRequestRepository) GetRequestsByStatus(status model.TaxiRequestStatus) ([]model.TaxiRequest, error) {
	var reqs []model.TaxiRequest
	err := repo.db.Table(taxiRequestTableName).Find(&reqs, "status = ?", status).Error
	if err != nil {
		return nil, err
	}
	return reqs, nil
}

func (repo *taxiRequestRepository) UpdateRequest(r *model.TaxiRequest) error {
	request := map[string]interface{}{
		"customer_id":           r.CustomerId,
		"driver_id":             r.DriverId,
		"departure_id":          r.DepartureId,
		"destination_id":        r.DestinationId,
		"departure_longitude":   r.DepartureLongitude,
		"departure_latitude":    r.DepartureLatitude,
		"destination_longitude": r.DestinationLongitude,
		"destination_latitude":  r.DestinationLatitude,
		"price":                 r.Price,
		"status":                r.Status,
	}

	if r.DriverId == 0 {
		request["driver_id"] = nil
	}

	err := repo.db.Clauses(clause.Returning{}).Table(taxiRequestTableName).Where("taxi_request_id = ?", r.TaxiRequestId).Updates(&request).Error
	if err != nil {
		return err
	}

	return nil
}
