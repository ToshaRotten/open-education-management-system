package buffer

import "github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/server/UserManager/models"

func remove(slice []models.User, s models.User) []models.User {
	return append(slice[:s.ID], slice[s.ID+1:]...)
}

func (b *Buffer) getData(user models.User) models.User {
	for _, data := range b.data {
		if data.Email == user.Email {
			return data
		}
	}
	return models.User{}
}
