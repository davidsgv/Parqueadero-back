package socket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Hub struct {
	clients           []*Client
	clientConstructor func() ClientEvents
	register          chan *Client
	unregister        chan *Client
	mutex             *sync.Mutex
	outbound          chan []byte
}

func NewHub(clientConstructor func() ClientEvents) *Hub {
	return &Hub{
		clients:           make([]*Client, 0),
		register:          make(chan *Client),
		unregister:        make(chan *Client),
		mutex:             &sync.Mutex{},
		clientConstructor: clientConstructor,
		outbound:          make(chan []byte),
	}
}

func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.register:
			hub.onConnect(client)
		case client := <-hub.unregister:
			hub.onDisconnect(client)
		}
	}
}

func (hub *Hub) Broadcast() {
	for _, e := range hub.clients {
		e.clientEvents.SendMessage()
	}
}

func (hub *Hub) onConnect(client *Client) {
	log.Println("Client connected", client.socket.RemoteAddr())

	hub.mutex.Lock()
	defer hub.mutex.Unlock()
	client.id = client.socket.RemoteAddr().String()
	hub.clients = append(hub.clients, client)
}

func (hub *Hub) onDisconnect(client *Client) {
	log.Println("Client disconnected", client.socket.RemoteAddr())

	hub.mutex.Lock()
	defer hub.mutex.Unlock()

	for i, item := range hub.clients {
		if item.id == client.id {
			index := len(hub.clients) - 1
			hub.clients[i] = hub.clients[index]
			hub.clients = hub.clients[:index]
			break
		}
	}
}

func (hub *Hub) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error upgrading connection", http.StatusInternalServerError)
		return
	}

	client := NewClient(hub, socket, hub.clientConstructor())

	hub.register <- client

	go client.Listen()
	go client.WriteMessage()
}
