package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initilizeRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/", GetAllStudent).Methods("GET")
	r.HandleFunc("/", PostStudent).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func main() {
	InitialMigration()
	initilizeRouter()
}
