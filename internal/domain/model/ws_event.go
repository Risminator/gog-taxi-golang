package model

import "encoding/json"

// Event is the Messages sent over the websocket
// Used to differ between different actions
type Event struct {
	// Type is the message type sent
	Type string `json:"type"`
	// Payload is the data Based on the Type
	Payload json.RawMessage `json:"payload"`
}

// Event types
const (
	/*
		EventSendWsClientUpdate is sent from frontend for backend to transfer new location to the other client
		Expected Payload:
			"latitude":		double
			"longitude":	double
	*/
	EventSendLocationUpdate = "send_location_update"

	/*
		EventTaxiRequestUpdate is sent from frontend for backend to transfer updated request parameters
		Expected Payload:
			"taxiRequestId":	int
			"customerId":		int
			"driverId":			int
			"departureId":		int
			"destinationId":	int
			"price":			double
			"status":			string
	*/
	EventSendTaxiRequestUpdate = "send_taxi_request_update"

	/*
		EventCancelTaxiRequest is sent from frontend for backend to transfer the command to the other client
		No expected payload
	*/
	EventCancelTaxiRequest = "cancel_taxi_request"

	/*
		EventWsClientUpdate is sent from backend to the other client with new location details
		Sent Payload matches that of EventSendLocationUpdate
	*/
	EventLocationUpdate = "location_update"

	/*
		EventTaxiRequestUpdate is sent from backend to the other client with new request details
		Sent Payload matches that of EventSendTaxiRequestUpdate
	*/
	EventTaxiRequestUpdate = "taxi_request_update"

	/*
		EventNewTaxiRequest is sent from backend to drivers waiting for new TaxiRequests when new one is created
		Sent Payload matches that of EventSendTaxiRequestUpdate
	*/
	EventNewTaxiRequest = "new_taxi_request"
)
