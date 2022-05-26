package main

import (
	"log"
	"net/http"
	"rest-api/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.GetHandler).Methods("GET")
	log.Default().Println("Listening on port 3333")
	log.Fatal(http.ListenAndServe(":3333", router))
}
