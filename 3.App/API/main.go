package main

import (
	"api/data"
	"api/server"
	"api/service"
	"log"
	"net/http"
	"time"
)

func main() {
	//mysql, err := data.NewMysqlRepository()
	repo, err := data.NewPostgresRepository()
	if err != nil {
		log.Fatal(err)
	}

	service := service.NewService(repo)
	server := server.New(service)
	// log.Fatal(http.ListenAndServe(":8080", s.Router()))

	srv := &http.Server{
		Handler: server.Router,
		Addr:    "127.0.0.1:80",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
