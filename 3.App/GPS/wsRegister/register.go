package wsregister

import (
	"encoding/json"
	socket "gps/Socket"
	"gps/location"
	"io"
	"log"
)

type WSRegister struct {
	location *location.Location
}

func NewWSRegister() socket.ClientEvents {
	return &WSRegister{
		location: location.GetInstance(),
	}
}

func (ws *WSRegister) SendMessage()                  {}
func (ws *WSRegister) OnConnect(string, chan []byte) {}

func (ws *WSRegister) OnDisconnect(id string, outbound chan []byte) {
	ws.location.Remove(id)
}

func (ws *WSRegister) OnMessageRecieve(id string, r io.Reader) {
	data := location.Vehicle{}
	err := json.NewDecoder(r).Decode(&data)
	if err == io.EOF {
		// One value is expected in the message.
		err = io.ErrUnexpectedEOF
		log.Printf("Error: %s\n", err.Error())
		return
	}

	ws.location.Update(id, data)
}
