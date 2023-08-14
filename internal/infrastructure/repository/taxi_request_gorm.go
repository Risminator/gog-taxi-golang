package repository

import (
	// "log"

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

// CreateRequest implements usecase.TaxiRequestRepository.
func (repo *taxiRequestRepository) CreateRequest(r *model.TaxiRequest) error {
	err := repo.db.Clauses(clause.Returning{}).Table(taxiRequestTableName).Create(&r).Error
	if err != nil {
		return err
	}
	return nil
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
