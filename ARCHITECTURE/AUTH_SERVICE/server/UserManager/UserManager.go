package UserManager

import (
	"fmt"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/server/UserManager/buffer"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/server/UserManager/database"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/server/UserManager/logger"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/server/UserManager/models"
)

//UserManager manages operations with buffer and database
type UserManager struct {
	buffer *buffer.Buffer
	db     *database.Helper
	logger *logger.MEMOLogger
}

// New - creates new instance of UserManager
func New() *UserManager {
	var manager UserManager
	manager.db = database.New()
	manager.buffer = buffer.New()
	manager.logger = logger.New()
	return &manager
}

// StartSyncProcess - starts a synchronization process, when UserManager send a buffer data to DATABASE_SERVICE
func (u *UserManager) StartSyncProcess() {
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
