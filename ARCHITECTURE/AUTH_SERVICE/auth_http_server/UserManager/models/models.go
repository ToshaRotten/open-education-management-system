package models

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	ThirdName string `json:"thirdName"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	DOB       string `json:"DOB"`
	Role      string `json:"role"`
	Group     string `json:"group"`
	Hash      string `json:"hash"`
	Job       string `json:"job"`
}

type Users struct {
	Users []User `json:"users"`
}

func New() *Users {
	return &Users{nil}
}

func Validate(UserData User) {

}
