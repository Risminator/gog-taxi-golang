package model

type Request struct {
	RequestId     int           `json:"request_id"`
	ClientId      int           `json:"client_id"`
	DriverId      int           `json:"driver_id"`
	DepartureId   int           `json:"departure_id"`
	DestinationId int           `json:"destination_id"`
	Price         float64       `json:"price"`
	Status        RequestStatus `json:"status"`
}

type RequestStatus uint

const (
	FindingDriver RequestStatus = iota
	WaitingForDriver
	InProgress
	Completed
	Canceled
)

// Automatically set to FindingDriver when creating an order
func (r *Request) CreateRequest(reqId, clId, drId, depId, destId int, price float64) Request {
	return Request{reqId, clId, drId, depId, destId, price, FindingDriver}
}

func (r *Request) SetClientId(c int) {
	r.ClientId = c
}

func (r *Request) SetDriverId(d int) {
	r.DriverId = d
}

func (r *Request) SetDepartureId(d int) {
	r.DepartureId = d
}

func (r *Request) SetDestinationId(d int) {
	r.DestinationId = d
}

func (r *Request) SetPrice(p float64) {
	r.Price = p
}

// Add business rule and tests, where you can only:
// Change statuses successively (where Completed is the last one)
// Go from FindingDriver or WaitingForDriver to Canceled
// Go from WaitingForDriver to FindingDriver
func (r *Request) SetStatus(s RequestStatus) {
	r.Status = s
}
