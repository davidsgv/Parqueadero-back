package wsget

import (
	"encoding/json"
	socket "gps/Socket"
	"gps/location"
	"io"
	"log"
)

type WSGet struct {
	location *location.Location
	outbound chan []byte
}

func NewWSGet() socket.ClientEvents {
	return &WSGet{
		location: location.GetInstance(),
	}
}

func (ws *WSGet) OnConnect(id string, outbound chan []byte) {
	ws.outbound = outbound
}

func (ws *WSGet) OnDisconnect(id string, outbound chan []byte) {}

func (ws *WSGet) SendMessage() {
	locations := ws.location.GetLocations()
	jbyte, err := json.Marshal(locations)
	if err != nil {
		log.Println("Error parsing struct")
	}
	ws.outbound <- jbyte
}

func (ws WSGet) OnMessageRecieve(string, io.Reader) {}
