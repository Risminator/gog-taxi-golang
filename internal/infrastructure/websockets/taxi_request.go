package websockets

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	v1 "github.com/Risminator/gog-taxi-golang/internal/controllers/httpgin/v1"
	"github.com/Risminator/gog-taxi-golang/internal/domain/model"
)

type wsTaxiRequestHandler struct {
	wsManager *WebsocketManager
}

func NewWsTaxiRequestHandler(wsManager *WebsocketManager) v1.TaxiRequestWsGateway {
	wsManager.handlers[EventLocationUpdate] = SendNewLocation

	return &wsTaxiRequestHandler{wsManager}
}

// Customer sends new TaxiRequest to drivers
func (h *wsTaxiRequestHandler) SendNewTaxiRequest(req model.TaxiRequest) error {
	// Create an event for drivers to receive
	var outEvent Event
	outEvent.Type = EventNewTaxiRequest
	data, err := json.Marshal(req)
	if err != nil {
		log.Printf("failed to marshal request info: %v", err)
		return err
	}
	outEvent.Payload = data

	// Go through every active websocket client
	for client := range h.wsManager.clients {
		// Send request info to drivers
		if client.user.Role == model.DriverRole {
			client.egress <- outEvent
		}
	}

	return nil
}

func (h *wsTaxiRequestHandler) ConnectWebsocket(w http.ResponseWriter, r *http.Request, userId int, role model.UserRole) {
	u := model.NewUser(userId, role)
	h.wsManager.serveWS(w, r, &u)
}

/*
Driver's location is sent to the driver's client.
Event should include fields (may be appended in future):
request:	TaxiRequest
origin:		UserRole???
latitude:	float64 (double precision)
longitude:	float64 (double precision)
*/
func SendNewLocation(event Event, c *WebsocketClient) error {
	var location model.Location
	if err := json.Unmarshal(event.Payload, &location); err != nil {
		return fmt.Errorf("bad payload in request: %v", err)
	}

	// TODO: Send location to relevant users (need to find a client knowing user)
	//for client := range c.manager.clients {
	//}

	return nil
}

/*
Status update is sent to the customer.
Event should include fields (may be appended in future):
status: string
*/
func SendNewTaxiRequestStatus(event Event, c *WebsocketClient) error {
	var status model.TaxiRequestStatus
	if err := json.Unmarshal(event.Payload, &status); err != nil {
		return fmt.Errorf("bad payload in request: %v", err)
	}

	// TODO: Send status to relevant customer (need to find a client knowing user)

	return nil
}
