package main

import (
	"reflect"
	"rest-api/factories"
	"rest-api/structs"
	"testing"
)

func TestMakePerson(t *testing.T) {
	got := factories.MakePerson("1", "John")
	want := structs.Person{ID: "1", FirstName: "John"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}

}
