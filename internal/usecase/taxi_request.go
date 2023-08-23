package usecase

import "github.com/Risminator/gog-taxi-golang/internal/domain/model"

type TaxiRequest interface {
	GetRequestsByStatus(status model.TaxiRequestStatus) ([]model.TaxiRequest, error)
	GetRequestById(id int) (*model.TaxiRequest, error)
	CreateRequest(reqId, clId, drId, depId, destId int, price float64) (*model.TaxiRequest, error)
	UpdateRequest(reqId int, clId int, drId int, depId int, destId int, price float64, status model.TaxiRequestStatus) (*model.TaxiRequest, error)
}
type TaxiRequestRepository interface {
	GetRequestsByStatus(status model.TaxiRequestStatus) ([]model.TaxiRequest, error)
	GetRequestById(id int) (*model.TaxiRequest, error)
	CreateRequest(r *model.TaxiRequest) error
	UpdateRequest(request *model.TaxiRequest) error
}

type taxiRequestUsecase struct {
	requestRepo TaxiRequestRepository
}

func NewTaxiRequestUsecase(requestRepo TaxiRequestRepository) TaxiRequest {
	return &taxiRequestUsecase{requestRepo}
}

// CreateRequest implements TaxiRequest.
func (use *taxiRequestUsecase) CreateRequest(reqId int, clId int, drId int, depId int, destId int, price float64) (*model.TaxiRequest, error) {
	req := model.CreateTaxiRequest(reqId, clId, drId, depId, destId, price)
	err := use.requestRepo.CreateRequest(&req)
	if err != nil {
		return nil, err
	}
	return &req, nil
}

// GetRequestById implements TaxiRequest.
func (use *taxiRequestUsecase) GetRequestById(id int) (*model.TaxiRequest, error) {
	req, err := use.requestRepo.GetRequestById(id)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// GetRequestsByStatus implements TaxiRequest.
func (use *taxiRequestUsecase) GetRequestsByStatus(status model.TaxiRequestStatus) ([]model.TaxiRequest, error) {
	reqs, err := use.requestRepo.GetRequestsByStatus(status)
	if err != nil {
		return nil, err
	}
	return reqs, nil
}

func (use *taxiRequestUsecase) UpdateRequest(reqId int, clId int, drId int, depId int, destId int, price float64, status model.TaxiRequestStatus) (*model.TaxiRequest, error) {
	request, err := use.GetRequestById(reqId)
	if err != nil {
		return nil, err
	}

	request.SetCustomerId(clId)
	request.SetDriverId(drId)
	request.SetDepartureId(depId)
	request.SetDestinationId(destId)
	request.SetPrice(price)
	request.SetStatus(status)

	err = use.requestRepo.UpdateRequest(request)
	if err != nil {
		return nil, err
	}
	return request, nil
}
