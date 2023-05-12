package models

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	ThirdName string `json:"thirdName"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	DOB       string `json:"DOB"`
	Role      string `json:"role"`
	Hash      string `json:"hash"`
}

type Users struct {
	Users []User `json:"users"`
}
