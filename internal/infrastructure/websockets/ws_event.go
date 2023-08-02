package websockets

import (
	"encoding/json"
)

// Event is the Messages sent over the websocket
// Used to differ between different actions
type Event struct {
	// Type is the message type sent
	Type string `json:"type"`
	// Payload is the data Based on the Type
	Payload json.RawMessage `json:"payload"`
}

// EventHandler is a function signature that is used to affect messages on the socket and triggered
// depending on the type
type EventHandler func(event Event, c *WebsocketClient) error

const (
	// EventNewTaxiRequest is sent to drivers waiting for new TaxiRequests when new one is created
	EventNewTaxiRequest          = "new_taxi_request"
	EventLocationUpdate          = "location_update"
	EventTaxiRequestStatusUpdate = "taxi_request_status_update"
)
