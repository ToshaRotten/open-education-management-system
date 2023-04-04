package UserManager

import (
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/APIServer/UserManager/buffer"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/APIServer/UserManager/database"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/APIServer/UserManager/logger"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/APIServer/UserManager/models"
)

//UserManager manages operations with buffer and database
type UserManager struct {
	buf    *buffer.Buffer
	db     *database.Helper
	logger *logger.MEMOLogger
}

// New - creates new instance of UserManager
func New() *UserManager {
	var manager UserManager
	manager.db = database.New()
	manager.buf = buffer.New()
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

// RegisterUser - register a new user
func (u *UserManager) RegisterUser(UserData models.User) {
	u.buf.Add(UserData)
	u.logger.CacheLog("User is registered")
}

// AuthUser - auth user
func (u *UserManager) AuthUser(UserData models.User) bool {
	if u.buf.SearchByEmail(UserData) {
		u.logger.CacheLog("Auth user")
		return true
	}
	return false
}

func (u *UserManager) EditUserProps(UserData models.User) bool {
	return false
}

func (u *UserManager) GetUserList() models.Users {
	var users models.Users
	users.Users = u.buf.Users
	return users
}
