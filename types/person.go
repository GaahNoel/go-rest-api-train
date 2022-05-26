package types

type Person struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"firstname,omitempty"`
	Password  string `json:"-"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
}
