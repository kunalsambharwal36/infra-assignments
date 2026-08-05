package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/kunalsambharwal36/infra-assignments/internal/handler"
	"github.com/kunalsambharwal36/infra-assignments/internal/service"
)

func main() {

	configService := service.NewConfigService()

	h := handler.NewHandler(configService)

	router := mux.NewRouter()

	router.HandleFunc("/ping", h.Ping).Methods("GET")
	router.HandleFunc("/configs", h.CreateConfig).Methods("POST")
	router.HandleFunc("/configs/{id}", h.GetConfig).Methods("GET")

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	log.Println("Config Service started on :8080")

	log.Fatal(server.ListenAndServe())
}
