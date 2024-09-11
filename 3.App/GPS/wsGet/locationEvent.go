package wsget

import (
	socket "gps/Socket"
	"gps/location"
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
