package main

import (
	"log"
	"net/http"
	"os"
	"parqueadero-back/internal/api/data"
	"parqueadero-back/internal/api/server"
	"parqueadero-back/internal/api/service"
	"time"
)

func main() {
	conInfo := data.ConnectionInfo{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBname:   os.Getenv("DB_NAME"),
	}

	repo, err := data.NewPostgresRepository(conInfo)
	if err != nil {
		log.Fatal(err)
	}

	service := service.NewService(repo)
	server := server.New(service)
	// log.Fatal(http.ListenAndServe(":8080", s.Router()))

	srv := &http.Server{
		Handler: server.Router,
		Addr:    "0.0.0.0:80",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
