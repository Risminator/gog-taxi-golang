package websockets

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	v1 "github.com/Risminator/gog-taxi-golang/internal/controllers/httpgin/v1"
	"github.com/Risminator/gog-taxi-golang/internal/domain/model"
	"github.com/Risminator/gog-taxi-golang/internal/usecase"
)

type wsClientOptions struct {
	ClientType *string `json:"clientType,omitempty"`
	RequestId  *int    `json:"requestId,omitempty"`
}

type wsTaxiRequestHandler struct {
	wsManager   *WebsocketManager
	taxiUsecase usecase.TaxiRequest
}

func NewWsTaxiRequestHandler(wsManager *WebsocketManager, taxiUsecase usecase.TaxiRequest) v1.TaxiRequestWsGateway {
	wsManager.handlers[EventWsClientUpdate] = UpdateWsClient
	wsManager.handlers[EventSendLocationUpdate] = SendNewLocation
	wsManager.handlers[EventSendTaxiRequestUpdate] = SendTaxiRequestUpdate

	return &wsTaxiRequestHandler{wsManager, taxiUsecase}
}

func UpdateWsClient(event Event, c *WebsocketClient) error {
	var options wsClientOptions
	if err := json.Unmarshal(event.Payload, &options); err != nil {
		return fmt.Errorf("bad payload in event: %v", err)
	}

	if options.RequestId != nil {
		c.setRequestId(*options.RequestId)
	}

	if options.ClientType != nil {
		wsType, err := model.WsClientTypeFromString(*options.ClientType)
		if err != nil {
			fmt.Printf("bad payload in event: %v", err)
		}
		c.setWsClientType(wsType)
	}

	return nil
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
		return fmt.Errorf("bad payload in event: %v", err)
	}

	// Create an event for user to receive
	var outEvent Event
	outEvent.Type = EventLocationUpdate
	data, err := json.Marshal(location)
	if err != nil {
		log.Printf("failed to marshal request info: %v", err)
		return err
	}
	outEvent.Payload = data

	// Send location to relevant users (need to find a client knowing user)
	// !!!!!!!!!!! HOW TO COMPARE USERS?
	for client := range c.manager.clients {
		if client.requestId == c.requestId && !(client.user.Role == c.user.Role && client.user.UserId == c.user.UserId) {
			client.egress <- outEvent
		}
	}

	return nil
}

/*
Status update is sent to the customer.
Event should include fields (may be appended in future):
status: string
*/
func SendTaxiRequestUpdate(event Event, c *WebsocketClient) error {
	var reqUpdate model.TaxiRequest
	if err := json.Unmarshal(event.Payload, &reqUpdate); err != nil {
		return fmt.Errorf("bad payload in event: %v", err)
	}

	// Create an event for user to receive
	var outEvent Event
	outEvent.Type = EventTaxiRequestUpdate
	data, err := json.Marshal(outEvent)
	if err != nil {
		log.Printf("failed to marshal request info: %v", err)
		return err
	}
	outEvent.Payload = data

	// Send updated taxi request to relevant customer
	for client := range c.manager.clients {
		if client.requestId == c.requestId && !(client.user.Role == c.user.Role && client.user.UserId == c.user.UserId) {
			client.egress <- outEvent
		}
	}

	return nil
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
		if client.user.Role == model.DriverRole && client.clientType == model.DriverGetOffers {
			client.egress <- outEvent
		}
	}

	return nil
}

func (h *wsTaxiRequestHandler) ConnectWebsocket(w http.ResponseWriter, r *http.Request, userId int, role model.UserRole, t model.WebsocketClientType, reqId int) {
	u := model.NewUser(userId, role)
	h.wsManager.serveWS(w, r, &u, t, reqId)
}
