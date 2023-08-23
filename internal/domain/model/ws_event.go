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

const (
	// EventWsClientUpdate is sent from frontend to change websocket client options
	EventWsClientUpdate = "update_ws_client"

	// EventSendWsClientUpdate is sent from frontend for backend to transfer new location to the other client
	EventSendLocationUpdate = "send_location_update"

	// EventTaxiRequestUpdate is sent from frontend for backend to transfer updated request parameters
	EventSendTaxiRequestUpdate = "send_taxi_request_update"

	// EventWsClientUpdate is sent from backend to the other client with new location details
	EventLocationUpdate = "location_update"

	// EventTaxiRequestUpdate is sent from backend to the other client with new request details
	EventTaxiRequestUpdate = "taxi_request_update"

	// EventNewTaxiRequest is sent from backend to drivers waiting for new TaxiRequests when new one is created
	EventNewTaxiRequest = "new_taxi_request"
)
