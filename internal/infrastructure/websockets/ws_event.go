package websockets

import "github.com/Risminator/gog-taxi-golang/internal/domain/model"

// EventHandler is a function signature that is used to affect messages on the socket and triggered
// depending on the type
type EventHandler func(event model.Event, c *WebsocketClient) error
