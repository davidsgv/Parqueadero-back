package main

import (
	wsget "gps/wsGet"
	wsregister "gps/wsRegister"
	"log"
	"net/http"
)

func main() {
	getHub := wsget.NewHub()
	go getHub.Run()

	registerHub := wsregister.NewHub(func() {
		getHub.Broadcast()
	})
	go registerHub.Run()

	http.HandleFunc("/gps", registerHub.HandleWebSocket)
	http.HandleFunc("/getgps", getHub.HandleWebSocket)
	err := http.ListenAndServe("localhost:81", nil)
	if err != nil {
		log.Fatal(err)
	}
}
