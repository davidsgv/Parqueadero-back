package main

import (
	socket "gps/Socket"
	"gps/location"
	wsget "gps/wsGet"
	wsregister "gps/wsRegister"
	"log"
	"net/http"
)

func main() {
	//hub for register gps
	registerHub := socket.NewHub(wsregister.NewWSRegister)
	go registerHub.Run()

	//hub for get locations
	getHub := socket.NewHub(wsget.NewWSGet)

	//create port for locations events
	le := wsget.NewLocationEvent(getHub)
	location.GetInstance().EventHandler = le

	go getHub.Run()

	http.HandleFunc("/gps", registerHub.HandleWebSocket)
	http.HandleFunc("/getgps", getHub.HandleWebSocket)

	err := http.ListenAndServe("localhost:81", nil)
	if err != nil {
		log.Fatal(err)
	}
}
