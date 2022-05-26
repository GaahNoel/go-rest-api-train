package handlers

import (
	"encoding/json"
	"net/http"
	"rest-api/factories"
	"rest-api/types"

	uuid "github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	var people []types.Person

	people = append(people, factories.MakePerson(factories.PersonFactory{Id: "1", FirstName: "John", Password: "123", Email: "teste@teste.com", Phone: "99999999"}))
	people = append(people, factories.MakePerson(factories.PersonFactory{Id: "2", FirstName: "Mary", Password: "123", Email: "teste@teste.com", Phone: ""}))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

func PostHandler(writer http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	person := types.Person{
		ID:        uuid.New().String(),
		FirstName: params["firstname"],
		Password:  params["password"],
		Email:     params["email"],
		Phone:     params["phone"],
	}
	json.NewEncoder(writer).Encode(PostReturn{Id: person.ID})
}

type PostReturn struct {
	Id string `json:"id"`
}
