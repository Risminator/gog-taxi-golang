package model

import (
	"errors"
	"time"
)

type TaxiRequest struct {
	TaxiRequestId        int               `json:"taxiRequestId" gorm:"primaryKey"`
	CustomerId           int               `json:"customerId"`
	DriverId             int               `json:"driverId"`
	DepartureId          int               `json:"departureId"`
	DestinationId        int               `json:"destinationId"`
	DepartureLongitude   float64           `json:"departureLongitude"`
	DepartureLatitude    float64           `json:"departureLatitude"`
	DestinationLongitude float64           `json:"destinationLongitude"`
	DestinationLatitude  float64           `json:"destinationLatitude"`
	Price                float64           `json:"price"`
	PlannedTime          *time.Time        `json:"plannedTime"`
	Status               TaxiRequestStatus `json:"status"`
}

type TaxiRequestStatus string

const (
	UnknownTaxiRequestStatus TaxiRequestStatus = ""
	FindingDriver            TaxiRequestStatus = "findingDriver"
	WaitingForDriver         TaxiRequestStatus = "waitingForDriver"
	WaitingForCustomer       TaxiRequestStatus = "waitingForCustomer"
	InProgress               TaxiRequestStatus = "inProgress"
	Completed                TaxiRequestStatus = "completed"
	Canceled                 TaxiRequestStatus = "canceled"
)

func TaxiRequestStatusFromString(s string) (TaxiRequestStatus, error) {
	switch s {
	case "findingDriver":
		return FindingDriver, nil
	case "waitingForDriver":
		return WaitingForDriver, nil
	case "inProgress":
		return InProgress, nil
	case "completed":
		return Completed, nil
	case "canceled":
		return Canceled, nil
	}

	return UnknownTaxiRequestStatus, errors.New("Unknown TaxiRequestStatus: " + s)
}

// Automatically set to FindingDriver when creating an order
func CreateTaxiRequest(reqId, clId, drId, depId, destId int, departureLon, departureLat, destinationLon, destinationLat, price float64, plannedTime *time.Time) TaxiRequest {
	return TaxiRequest{reqId, clId, drId, depId, destId, departureLon, departureLat, destinationLon, destinationLat, price, plannedTime, FindingDriver}
}

func (r *TaxiRequest) SetTaxiRequestId(t int) {
	r.TaxiRequestId = t
}

func (r *TaxiRequest) SetCustomerId(c int) {
	r.CustomerId = c
}

func (r *TaxiRequest) SetDriverId(d int) {
	r.DriverId = d
}

func (r *TaxiRequest) SetDepartureId(d int) {
	r.DepartureId = d
}

func (r *TaxiRequest) SetDestinationId(d int) {
	r.DestinationId = d
}

func (r *TaxiRequest) SetPrice(p float64) {
	r.Price = p
}

// Add business rule and tests, where you can only:
// Change statuses successively (where Completed is the last one)
// Go from FindingDriver or WaitingForDriver to Canceled
// Go from WaitingForDriver to FindingDriver
func (r *TaxiRequest) SetStatus(s TaxiRequestStatus) {
	r.Status = s
}
