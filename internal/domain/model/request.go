package model

type TaxiRequest struct {
	TaxiRequestId int               `json:"taxi_request_id" gorm:"primaryKey"`
	ClientId      int               `json:"client_id"`
	DriverId      int               `json:"driver_id"`
	DepartureId   int               `json:"departure_id"`
	DestinationId int               `json:"destination_id"`
	Price         float64           `json:"price"`
	Status        TaxiRequestStatus `json:"status"`
}

type TaxiRequestStatus string

const (
	FindingDriver    TaxiRequestStatus = "FindingDriver"
	WaitingForDriver TaxiRequestStatus = "WaitingForDriver"
	InProgress       TaxiRequestStatus = "InProgress"
	Completed        TaxiRequestStatus = "Completed"
	Canceled         TaxiRequestStatus = "Canceled"
)

func ParseTaxiRequestStatus(str string) TaxiRequestStatus {
	switch str {
	case "FindingDriver":
		return FindingDriver
	case "WaitingForDriver":
		return WaitingForDriver
	case "InProgress":
		return InProgress
	case "Completed":
		return Completed
	case "Canceled":
		return Canceled
	}
	panic("Could not parse string '" + str + "' to TaxiRequestStatus")
}

// Automatically set to FindingDriver when creating an order
func CreateTaxiRequest(reqId, clId, drId, depId, destId int, price float64) TaxiRequest {
	return TaxiRequest{reqId, clId, drId, depId, destId, price, FindingDriver}
}

func (r *TaxiRequest) SetClientId(c int) {
	r.ClientId = c
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
