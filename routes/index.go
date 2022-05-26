package routes

import (
	"rest-api/handlers"

	"github.com/gorilla/mux"
)

func Setup(router *mux.Router) {
	router.HandleFunc("/", handlers.GetHandler).Methods("GET")
	router.HandleFunc("/", handlers.PostHandler).Methods("POST")
}
