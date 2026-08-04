package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kunalsambharwal36/infra-assignments/internal/handler"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/ping", handler.Ping).Methods("GET")

	log.Println("Config Service started on :8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
