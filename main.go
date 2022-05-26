package main

import (
	"log"
	"net/http"
	"rest-api/routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.Setup(router)
	log.Default().Println("Listening on port 3333")
	log.Fatal(http.ListenAndServe(":3333", router))
}
