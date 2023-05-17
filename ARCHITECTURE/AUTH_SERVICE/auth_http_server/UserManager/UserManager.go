package UserManager

import (
	"errors"
	"fmt"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/auth_http_server/UserManager/buffer"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/auth_http_server/UserManager/database"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/auth_http_server/UserManager/logger"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/auth_http_server/UserManager/models"
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

func (u *UserManager) SetScheme(scheme string) {
	u.database.SetScheme(scheme)
}

// SyncWithDb - starts a synchronization process, when UserManager send a buffer data to DATABASE_SERVICE
func (u *UserManager) SyncWithDb() {
	u.logger.ManagerLog("Sync with db")
}

// InitManager - Gets users from db and create buffer
func (u *UserManager) InitManager() {
	usersFromDb, _ := u.database.ReadAll()
	count := u.buffer.Create(usersFromDb)
	//fmt.Println("COUNT:", count)
	u.logger.CacheLog(fmt.Sprintf("Total get from base: %d", count))
}

// SetInterval - sets a interval of sending buffer data to database service
func (u *UserManager) SetInterval(minutes int) {
	u.logger.ManagerLog("Set sync interval")
}

// CreateUsers - register a new users
func (u *UserManager) CreateUsers(UserData []models.User) error {
	count := u.buffer.Create(UserData)
	u.logger.CacheLog(fmt.Sprintf("Total: %d users created", count))
	if count == 0 {
		return errors.New("Not created")
	}
	return nil
}

// ReadUsers - auth user
func (u *UserManager) ReadUsers(UserData []models.User) []models.User {
	users, count := u.buffer.Read(UserData)
	if count == 0 {
		users, count = u.database.Read(UserData)
	}
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
