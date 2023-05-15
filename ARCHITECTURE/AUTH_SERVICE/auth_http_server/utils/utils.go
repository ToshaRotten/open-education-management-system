package utils

import (
	"encoding/json"
	"fmt"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/auth_http_server/UserManager/models"
	"io"
	"net/http"
)

// ParseUsersFromJSON - returns a user
func ParseUsersFromJSON(r *http.Request) (models.Users, error) {
	var temp models.Users
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return models.Users{}, err
	}
	fmt.Println(string(data))
	err = json.Unmarshal(data, &temp)
	if err != nil {
		return models.Users{}, err
	}
	return temp, err
}
