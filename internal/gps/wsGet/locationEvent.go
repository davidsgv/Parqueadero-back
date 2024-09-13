package wsget

import (
	"parqueadero-back/internal/gps/location"
	"parqueadero-back/internal/gps/socket"
)

type LocationEvent struct {
	hub *socket.Hub
}

func NewLocationEvent(hub *socket.Hub) location.Events {
	return &LocationEvent{
		hub: hub,
	}
}

func (le *LocationEvent) OnLocationUpdate() {
	le.hub.Broadcast()
}
