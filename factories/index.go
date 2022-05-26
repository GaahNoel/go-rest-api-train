package factories

import (
	"rest-api/types"
)

type PersonFactory struct {
	Id        string
	FirstName string
	Password  string
	Email     string
	Phone     string
}

func MakePerson(params PersonFactory) (person types.Person) {
	person = types.Person{
		ID:        params.Id,
		FirstName: params.FirstName,
		Password:  params.Password,
		Email:     params.Email,
		Phone:     params.Phone,
	}
	return
}
