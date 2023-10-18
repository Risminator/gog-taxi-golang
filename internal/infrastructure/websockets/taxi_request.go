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

type wsTaxiRequestHandler struct {
	wsManager     *WebsocketManager
	taxiUsecase   usecase.TaxiRequest
	vesselUsecase usecase.Vessel
}

func NewWsTaxiRequestHandler(wsManager *WebsocketManager, taxiUsecase usecase.TaxiRequest, vesselUsecase usecase.Vessel) v1.TaxiRequestWsGateway {
	h := wsTaxiRequestHandler{wsManager, taxiUsecase, vesselUsecase}

	h.wsManager.handlers[model.EventSendLocationUpdate] = h.SendNewLocation
	h.wsManager.handlers[model.EventSendTaxiRequestUpdate] = h.SendTaxiRequestUpdate
	h.wsManager.handlers[model.EventCancelTaxiRequest] = h.CancelTaxiRequest

	return &h
}

/*
Driver's location is sent to the driver's client.
Event should include fields (may be appended in future):
request:	TaxiRequest
origin:		UserRole???
latitude:	float64 (double precision)
longitude:	float64 (double precision)
*/
func (h *wsTaxiRequestHandler) SendNewLocation(event model.Event, c *WebsocketClient) error {
	var location model.Location
	if err := json.Unmarshal(event.Payload, &location); err != nil {
		return fmt.Errorf("bad payload in event: %v", err)
	}

	// Update DB
	vessel, err := h.vesselUsecase.GetVesselByRequestId(c.requestId)
	if err != nil {
		log.Printf("failed to get vessel info: %v", err)
		return err
	}
	h.vesselUsecase.UpdateLocation(vessel.VesselId, location.Latitude, location.Longitude)

	// Create an event for user to receive
	var outEvent model.Event
	outEvent.Type = model.EventLocationUpdate
	data, err := json.Marshal(location)
	if err != nil {
		log.Printf("failed to marshal request info: %v", err)
		return err
	}
	outEvent.Payload = data

	// Send location to relevant users (need to find a client knowing user)
	// !!!!!!!!!!! HOW TO COMPARE USERS?
	for client := range c.manager.clients {
		if client.requestId == c.requestId && c != client {
			client.egress <- outEvent
			break
		}
	}

	return nil
}

/*
Status update is sent to the customer.
Event should include fields (may be appended in future):
status: string
*/
func (h *wsTaxiRequestHandler) SendTaxiRequestUpdate(event model.Event, c *WebsocketClient) error {

	var rUpd model.TaxiRequest
	if err := json.Unmarshal(event.Payload, &rUpd); err != nil {
		return fmt.Errorf("bad payload in event: %v", err)
	}

	// Update db
	_, err := h.taxiUsecase.UpdateRequest(rUpd.TaxiRequestId, rUpd.CustomerId, rUpd.DriverId, rUpd.DepartureId, rUpd.DestinationId, rUpd.DepartureLongitude, rUpd.DepartureLatitude, rUpd.DestinationLongitude, rUpd.DestinationLatitude, rUpd.Price, rUpd.Status)
	if err != nil {
		return fmt.Errorf("failed to update request: %v", err)
	}

	switch rUpd.Status {
	// Driver accepted the taxi request
	case model.WaitingForDriver:
		c.setWsClientType(model.DriverCurrentTaxiRequestInfo)
		c.setRequestId(rUpd.TaxiRequestId)
	// Driver met Customer
	case model.InProgress:
	// Driver reached the destination with Customer
	case model.Completed:
		defer c.manager.removeClient(c)
	}

	// Create an event for user to receive
	var outEvent model.Event
	outEvent.Type = model.EventTaxiRequestUpdate
	data, err := json.Marshal(rUpd)
	if err != nil {
		log.Printf("failed to marshal request info: %v", err)
		return err
	}
	outEvent.Payload = data

	// Send updated taxi request to relevant customer
	for client := range c.manager.clients {
		if client.requestId == c.requestId && c != client {
			client.egress <- outEvent
			break
		}
	}

	return nil
}

func (h *wsTaxiRequestHandler) CancelTaxiRequest(event model.Event, c *WebsocketClient) error {
	defer c.manager.removeClient(c)

	request, err := h.taxiUsecase.GetRequestById(c.requestId)
	if err != nil {
		return err
	}

	sendMessageFlag := false

	// The cancellation is coming from customer (can be done only when FindingDriver or WaitingForDriver)
	if c.user.Role == model.CustomerRole && (request.Status == model.FindingDriver || request.Status == model.WaitingForDriver) {
		request.SetStatus(model.Canceled)
		if request.DriverId != 0 {
			sendMessageFlag = true
		}
	}

	// The cancellation is coming from driver (can be done only when WaitingForDriver)
	if c.user.Role == model.DriverRole && request.Status == model.WaitingForDriver {
		sendMessageFlag = true
		request.SetStatus(model.FindingDriver)
		request.SetDriverId(0)
	}

	// update db
	_, err = h.taxiUsecase.UpdateRequest(request.TaxiRequestId, request.CustomerId, request.DriverId, request.DepartureId, request.DestinationId, request.DepartureLongitude, request.DepartureLatitude, request.DestinationLongitude, request.DestinationLatitude, request.Price, request.Status)
	if err != nil {
		return fmt.Errorf("failed to update request: %v", err)
	}

	// Send the cancellation information to the relevant client
	if sendMessageFlag {
		var outEvent model.Event
		outEvent.Type = model.EventTaxiRequestUpdate
		data, err := json.Marshal(request)
		if err != nil {
			log.Printf("failed to marshal request info: %v", err)
			return err
		}
		outEvent.Payload = data

		for client := range c.manager.clients {
			if client.requestId == c.requestId && c != client {
				client.egress <- outEvent
				break
			}
		}
	}

	return nil
}

// Customer sends new TaxiRequest to drivers
func (h *wsTaxiRequestHandler) SendNewTaxiRequest(req model.TaxiRequest) error {
	// Create an event for drivers to receive
	var outEvent model.Event
	outEvent.Type = model.EventNewTaxiRequest
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

func (h *wsTaxiRequestHandler) ConnectWebsocket(w http.ResponseWriter, r *http.Request, userId int, role model.UserRole, t model.WebsocketClientType, reqId int, initEvent *model.Event) error {
	u := model.NewUser(userId, role)
	client, err := h.wsManager.serveWS(w, r, &u, t, reqId)
	if err != nil {
		return err
	}

	if initEvent != nil {
		client.egress <- *initEvent
	}

	return nil
}
