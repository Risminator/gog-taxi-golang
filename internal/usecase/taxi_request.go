package usecase

import (
	"strconv"
	"time"

	"github.com/Risminator/gog-taxi-golang/internal/domain/model"
)

type TaxiRequest interface {
	GetRequestsByStatus(status model.TaxiRequestStatus) ([]model.TaxiRequest, error)
	GetRequestById(id int) (*model.TaxiRequest, error)
	GetRequestByUserId(id int, role model.UserRole) (*model.TaxiRequest, error)
	CreateRequest(reqId int, clId int, drId int, depId int, destId int, departureLon, departureLat, destinationLon, destinationLat, price float64, plannedTime *time.Time) (*model.TaxiRequest, error)
	UpdateRequest(reqId int, clId int, drId int, depId int, destId int, departureLon, departureLat, destinationLon, destinationLat, price float64, status model.TaxiRequestStatus) (*model.TaxiRequest, error)
}
type TaxiRequestRepository interface {
	GetRequestsByStatus(status model.TaxiRequestStatus) ([]model.TaxiRequest, error)
	GetRequestById(id int) (*model.TaxiRequest, error)
	GetRequestByUserId(id int, role model.UserRole) (*model.TaxiRequest, error)
	CreateRequest(r *model.TaxiRequest) (int, error)
	UpdateRequest(request *model.TaxiRequest) error
}

type taxiRequestUsecase struct {
	requestRepo TaxiRequestRepository
	routeWebApi RouteWebApi
}

func NewTaxiRequestUsecase(requestRepo TaxiRequestRepository, routeWebApi RouteWebApi) TaxiRequest {
	return &taxiRequestUsecase{requestRepo, routeWebApi}
}

// GetRequestByUserId implements TaxiRequest.
func (use *taxiRequestUsecase) GetRequestByUserId(id int, role model.UserRole) (*model.TaxiRequest, error) {
	req, err := use.requestRepo.GetRequestByUserId(id, role)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// CreateRequest implements TaxiRequest.
func (use *taxiRequestUsecase) CreateRequest(reqId int, clId int, drId int, depId int, destId int, departureLon, departureLat, destinationLon, destinationLat, price float64, plannedTime *time.Time) (*model.TaxiRequest, error) {
	geo, err := use.routeWebApi.GetRouteInfo(model.Location{Latitude: departureLat, Longitude: departureLon}, model.Location{Latitude: destinationLat, Longitude: destinationLon}, "river")
	if err != nil {
		return nil, err
	}
	trackLength, err := strconv.ParseFloat(geo.Features[0].Properties["track-length"].(string), 64)
	if err != nil {
		return nil, err
	}

	price = (230 + 206 + ((trackLength/1000)/30)*3000) * 1.18

	req := model.CreateTaxiRequest(reqId, clId, drId, depId, destId, departureLon, departureLat, destinationLon, destinationLat, price, plannedTime)
	id, err := use.requestRepo.CreateRequest(&req)
	if err != nil {
		return nil, err
	}

	req.SetTaxiRequestId(id)
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

func (use *taxiRequestUsecase) UpdateRequest(reqId int, clId int, drId int, depId int, destId int, departureLon, departureLat, destinationLon, destinationLat, price float64, status model.TaxiRequestStatus) (*model.TaxiRequest, error) {
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
	request.DepartureLongitude = departureLon
	request.DepartureLatitude = departureLat
	request.DestinationLongitude = destinationLon
	request.DestinationLatitude = destinationLat

	err = use.requestRepo.UpdateRequest(request)
	if err != nil {
		return nil, err
	}

	return request, nil
}
