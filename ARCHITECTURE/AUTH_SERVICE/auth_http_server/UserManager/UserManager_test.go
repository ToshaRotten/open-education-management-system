package UserManager

import (
	"fmt"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/auth_http_server/UserManager/models"
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

func TestCreate(t *testing.T) {
	manager := New()
	//manager.CreateUsers([]models.User{testUser})

	fmt.Println(manager.ValidateUsers([]models.User{testUser}))
}

func TestUserManager_InitManager(t *testing.T) {
	manager := New()
	manager.InitManager()

	//manager.database.DBHealth()

	fmt.Println(manager.buffer)
}
