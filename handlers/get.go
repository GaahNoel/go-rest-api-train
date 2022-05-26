package handlers

import (
	"encoding/json"
	"net/http"
	"rest-api/factories"
	"rest-api/structs"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	var people []structs.Person
	people = append(people, factories.MakePerson("1", "John"))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}
