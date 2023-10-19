package wsget

import (
	"gps/location"

	"github.com/gorilla/websocket"
)

type Client struct {
	hub      *Hub
	id       string
	socket   *websocket.Conn
	outbound chan bool
}

func NewClient(hub *Hub, socket *websocket.Conn) *Client {
	return &Client{
		hub:      hub,
		socket:   socket,
		outbound: make(chan bool),
	}
}

func (c *Client) SendLocation() {
	for {
		instance := location.GetInstance()
		select {
		case <-c.outbound:
			buses := instance.GetBuses()
			c.socket.WriteJSON(buses)
		}
	}
}

func (c Client) Close() {
	c.socket.Close()
	close(c.outbound)
}
