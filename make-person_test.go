package main

import (
	"reflect"
	"rest-api/factories"
	"rest-api/types"
	"testing"
)

func TestMakePerson(t *testing.T) {
	got := factories.MakePerson(factories.PersonFactory{Id: "1", FirstName: "John", Password: "123", Email: "teste@teste.com", Phone: "99999999"})
	want := types.Person{ID: "1", FirstName: "John"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}

}
