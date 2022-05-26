package factories

import (
	"rest-api/structs"
)

func MakePerson(id string, firstName string) structs.Person {
	return structs.Person{ID: id, FirstName: firstName}
}
