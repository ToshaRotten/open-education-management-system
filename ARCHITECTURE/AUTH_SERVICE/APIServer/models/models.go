package models

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	ThirdName string `json:"thirdName"`
	DOB       string `json:"DOB"`
	Role      string `json:"role"`
	Hash      string `json:"hash"`
}

type Users struct {
	Users []User `json:"users"`
}
