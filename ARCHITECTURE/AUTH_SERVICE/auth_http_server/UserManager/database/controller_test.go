package database

import (
	"fmt"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/server/UserManager/models"
	"testing"
)

var (
	testUser = models.User{
		ID:        0,
		FirstName: "Вован",
		LastName:  "Свиридонов",
		ThirdName: "Александрович",
		Phone:     "+7 (000)-123-12-12",
		Email:     "notmail@mail.ru",
		DOB:       "2000-01-01",
		Role:      "developer",
		Hash:      "Есть",
		Group:     "",
	}
	testUser1 = models.User{
		ID:        0,
		FirstName: "Гадя",
		LastName:  "Хренова",
		ThirdName: "Петрович",
		Phone:     "7 (000)-000-00-02",
		Email:     "hrenova@mail.ru",
		DOB:       "1997-01-01",
		Role:      "user",
		Hash:      "Нету",
		Group:     "",
	}
	testUser2 = models.User{
		ID:        0,
		FirstName: "Гена",
		LastName:  "Козлов",
		ThirdName: "Дмитриевич",
		Phone:     "+7 (000)-123-12-12",
		Email:     "mail@mail.ru",
		DOB:       "2000-01-01",
		Role:      "developer",
		Hash:      "Есть",
		Group:     "",
	}
)

func TestDBController_Create(t *testing.T) {
	d := New()
	d.Create(testUser2)
}

func TestDBController_Read(t *testing.T) {
	d := New()
	var users models.Users
	users.Users = append(users.Users, testUser1)
	users.Users = append(users.Users, testUser)
	result, _ := d.Read(users.Users)
	fmt.Println(result)
}

func TestDBController_Update(t *testing.T) {
	d := New()
	d.Update(testUser1, testUser2)
}

func TestDBController_Delete(t *testing.T) {
	var usr models.User
	usr.ID = 1
	d := New()
	d.Delete(usr)
}

func TestDBController_DBHealth(t *testing.T) {
	d := New()
	d.DBHealth()
}
