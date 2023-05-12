package UserManager

import (
	"fmt"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/database"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/server/UserManager/buffer"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/server/UserManager/logger"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/server/UserManager/models"
	"regexp"
)

//UserManager manages operations with buffer and database
type UserManager struct {
	buffer   *buffer.Buffer
	logger   *logger.MEMOLogger
	database *database.DBController
}

// New - creates new instance of UserManager
func New() *UserManager {
	var manager UserManager
	manager.buffer = buffer.New()
	manager.logger = logger.New()
	manager.database = database.New()
	return &manager
}

// StartSyncProcess - starts a synchronization process, when UserManager send a buffer data to DATABASE_SERVICE
func (u *UserManager) SyncWithDb() {
	u.logger.ManagerLog("Starting sync process")

}

// SetInterval - sets a interval of sending buffer data to database service
func (u *UserManager) SetInterval(minutes int) {
	u.logger.ManagerLog("Set sync interval")
}

// CreateUsers - register a new users
func (u *UserManager) CreateUsers(UserData []models.User) {
	count := u.buffer.Create(UserData)
	u.logger.CacheLog(fmt.Sprintf("Total: %d users created", count))
}

// ReadUsers - auth user
func (u *UserManager) ReadUsers(UserData []models.User) []models.User {
	users, count := u.buffer.Read(UserData)
	u.logger.CacheLog(fmt.Sprintf("Total: %d users found", count))
	return users
}

// UpdateUsers - ..
func (u *UserManager) UpdateUsers(UserData []models.User, NewUserData []models.User) {
	count := u.buffer.Update(UserData, NewUserData)
	u.logger.CacheLog(fmt.Sprintf("Total: %d users updated", count))
}

// DeleteUsers - ..
func (u *UserManager) DeleteUsers(UserData []models.User) {
	count := u.buffer.Delete(UserData)
	u.logger.CacheLog(fmt.Sprintf("Total: %d users deleted", count))
}

func (u *UserManager) ValidateUsers(UserData []models.User) bool {
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
