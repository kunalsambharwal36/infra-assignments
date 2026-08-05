package main

import (
	"log"
	"net/http"

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

	log.Println("Config Service started on :8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
