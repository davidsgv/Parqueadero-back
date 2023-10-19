package wsregister

import (
	"gps/location"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	hub      *Hub
	id       string
	socket   *websocket.Conn
	Callback func()
}

func NewClient(hub *Hub, socket *websocket.Conn, callback func()) *Client {
	return &Client{
		hub:      hub,
		socket:   socket,
		Callback: callback,
	}
}

func (c *Client) Listen() {
	instance := location.GetInstance()

	for {
		ubicacion := struct {
			Latitud  float64 `json:"latitud"`
			Longitud float64 `json:"longitud"`
			Placa    string  `json:"placa"`
		}{}

		err := c.socket.ReadJSON(&ubicacion)
		if err != nil {
			log.Println("read:", err)
			break
		}
		instance.GPS.Update(ubicacion.Placa, ubicacion.Latitud, ubicacion.Longitud)
		log.Printf("ubicacion: %+v", instance.GPS.GetBuses())
		c.Callback()
	}
}

func (c Client) Close() {
	c.socket.Close()
}
