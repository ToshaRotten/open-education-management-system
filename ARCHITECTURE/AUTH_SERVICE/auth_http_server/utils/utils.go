package utils

import (
	"encoding/json"
	"fmt"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/auth_http_server/UserManager/models"
	"io"
	"net/http"
	"regexp"
)

// ParseUsersFromJSON - returns a user
func ParseUsersFromJSON(r *http.Request) (models.Users, error) {
	var temp models.Users
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return models.Users{}, err
	}
	//fmt.Println(string(data))
	err = json.Unmarshal(data, &temp)
	if err != nil {
		return models.Users{}, err
	}
	return temp, err
}

func ValidateUsers(UserData []models.User) bool {
	email, _ := regexp.Compile("(?i)[-A-Za-z0-9!#$%&'*+/=?^_`{|}~]+(?:\\.[-A-Za-z0-9!#$%&'*+/=?^_`{|}~]+)*@(?:[A-Za-z0-9](?:[-A-Za-z0-9]*[A-Za-z0-9])?\\.)+[A-Za-z0-9](?:[-A-Za-z0-9]*[A-Za-z0-9])?")
	phone, _ := regexp.Compile("\\+7\\s+\\([0-9]+\\)-[0-9]+-[0-9]+")
	RUS_word, _ := regexp.Compile("[ЁёА-я]")
	DOB, _ := regexp.Compile("[0-9]+-(0?[1-9]|[1][0-2])-[0-9]+")
	for _, user := range UserData {
		flag := true
		if !email.MatchString(user.Email) {
			flag = false
			fmt.Println("Bad Email")
		}
		if !phone.MatchString(user.Phone) {
			flag = false
			fmt.Println("Bad Phone")
		}
		if !RUS_word.MatchString(user.FirstName) {
			flag = false
			fmt.Println("Bad FirstName")
		}
		if !RUS_word.MatchString(user.LastName) {
			flag = false
			fmt.Println("Bad LastName")
		}
		if !RUS_word.MatchString(user.ThirdName) {
			flag = false
			fmt.Println("Bad ThirdName")
		}
		if !DOB.MatchString(user.DOB) {
			flag = false
			fmt.Println("Bad DOB")
		}
		if flag == false {
			return false
		}
	}
	return true
}
