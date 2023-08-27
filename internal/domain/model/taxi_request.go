package model

import (
	"errors"
)

type TaxiRequest struct {
	TaxiRequestId int               `json:"requestId" gorm:"primaryKey"`
	CustomerId    int               `json:"customerId"`
	DriverId      int               `json:"driverId"`
	DepartureId   int               `json:"departureId"`
	DestinationId int               `json:"destinationId"`
	Price         float64           `json:"price"`
	Status        TaxiRequestStatus `json:"status"`
}

type TaxiRequestStatus string

const (
	UnknownTaxiRequestStatus TaxiRequestStatus = ""
	FindingDriver            TaxiRequestStatus = "findingDriver"
	WaitingForDriver         TaxiRequestStatus = "waitingForDriver"
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
func CreateTaxiRequest(reqId, clId, drId, depId, destId int, price float64) TaxiRequest {
	return TaxiRequest{reqId, clId, drId, depId, destId, price, FindingDriver}
}

func (r *TaxiRequest) SetId(t int) {
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
