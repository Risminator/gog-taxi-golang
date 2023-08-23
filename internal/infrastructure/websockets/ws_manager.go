package websockets

import (
	"context"
	"errors"
	"log"
	"net/http"
	"sync"

	"github.com/Risminator/gog-taxi-golang/internal/domain/model"
	"github.com/gorilla/websocket"
)

var (
	/**
	websocketUpgrader is used to upgrade incomming HTTP requests into a persitent websocket connection
	*/
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

var (
	ErrEventNotSupported = errors.New("this event type is not supported")
)

// WebsocketManager is used to hold references to all Clients Registered, and Broadcasting etc
type WebsocketManager struct {
	clients WebsocketClientList

	// Using a syncMutex here to be able to lock state before editing clients
	// Could also use Channels to block
	sync.RWMutex
	// handlers are functions that are used to handle Events
	handlers map[string]EventHandler
}

// NewManager is used to initalize all the values inside the manager
func NewManager(ctx context.Context) *WebsocketManager {
	m := &WebsocketManager{
		clients:  make(WebsocketClientList),
		handlers: make(map[string]EventHandler),
	}
	m.setupEventHandlers()
	return m
}

// setupEventHandlers configures and adds all handlers
func (m *WebsocketManager) setupEventHandlers() {
	// m.handlers[EventSendMessage] = SendMessageHandler
}

// routeEvent is used to make sure the correct event goes into the correct handler
func (m *WebsocketManager) routeEvent(event model.Event, c *WebsocketClient) error {
	// Check if Handler is present in Map
	if handler, ok := m.handlers[event.Type]; ok {
		// Execute the handler and return any err
		if err := handler(event, c); err != nil {
			return err
		}
		return nil
	} else {
		return ErrEventNotSupported
	}
}

// serveWS is a HTTP Handler that the has the Manager that allows connections
func (m *WebsocketManager) serveWS(w http.ResponseWriter, r *http.Request, u *model.User, t model.WebsocketClientType, reqId int) (*WebsocketClient, error) {
	log.Println("New connection")
	// Begin by upgrading the HTTP request
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Create New Client
	client := NewWebsocketClient(conn, m, u, t, reqId)
	// Add the newly created client to the manager
	m.addClient(client)

	go client.readMessages(m)
	go client.writeMessages(m)

	return client, nil
}

// addClient will add clients to our clientList
func (m *WebsocketManager) addClient(client *WebsocketClient) {
	// Lock so we can manipulate
	m.Lock()
	defer m.Unlock()

	// Add Client
	m.clients[client] = true
}

// removeClient will remove the client and clean up
func (m *WebsocketManager) removeClient(client *WebsocketClient) {
	m.Lock()
	defer m.Unlock()

	// Check if Client exists, then delete it
	if _, ok := m.clients[client]; ok {
		// close connection
		client.connection.Close()
		// remove
		delete(m.clients, client)
	}
}
