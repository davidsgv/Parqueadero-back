package socket

import (
	"io"
	"log"

	"github.com/gorilla/websocket"
)

type ClientEvents interface {
	OnMessageRecieve(string, io.Reader)
	OnConnect(string, chan []byte)
	OnDisconnect(string, chan []byte)
	SendMessage()
}

type Client struct {
	hub          *Hub
	id           string
	socket       *websocket.Conn
	clientEvents ClientEvents
}

func NewClient(hub *Hub, socket *websocket.Conn, clientEvents ClientEvents) *Client {
	client := &Client{
		hub:          hub,
		socket:       socket,
		clientEvents: clientEvents,
	}

	client.clientEvents.OnConnect(client.id, client.hub.outbound)
	return client
}

func (c *Client) Listen() {
	for {
		_, r, err := c.socket.NextReader()
		if err != nil {
			if !websocket.IsCloseError(err, 1000) {
				log.Println("No Json data find or unexpected error:", err)
			}
			c.Close()
			break
		}

		c.clientEvents.OnMessageRecieve(c.id, r)
	}
}

func (c *Client) WriteMessage() {
	for message := range c.hub.outbound {
		err := c.socket.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("Error Writing Message:", err)
			c.Close()
			return
		}
	}

	c.socket.WriteMessage(websocket.CloseMessage, []byte{})
	c.Close()
}

func (c Client) Close() {
	c.clientEvents.OnDisconnect(c.id, c.hub.outbound)
	c.hub.unregister <- &c
}
