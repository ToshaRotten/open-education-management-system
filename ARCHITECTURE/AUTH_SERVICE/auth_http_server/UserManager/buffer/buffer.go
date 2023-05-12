package buffer

import (
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/server/UserManager/models"
)

type Buffer struct {
	data  []models.User
	Limit int
	count int
}

func New() *Buffer {
	var b Buffer
	return &b
}

func (b *Buffer) Create(users []models.User) int {
	var entries int
	for _, user := range users {
		user.ID = b.count
		b.data = append(b.data, user)
		b.count++
		entries++
	}
	return entries
}

func (b *Buffer) Read(users []models.User) ([]models.User, int) {
	var entries int
	var result []models.User
	var temp []models.User
	//get ids
	for _, user := range users {
		temp = append(temp, b.getData(user))
	}
	for _, user := range temp {
		for _, data := range b.data {
			if data.ID == user.ID {
				result = append(result, data)
				entries++
			}
		}
	}

	return result, entries
}

func (b *Buffer) Update(oldData []models.User, newData []models.User) int {
	var entries int
	//get ids
	for i, user := range oldData {
		oldData[i] = b.getData(user)
	}
	for i := 0; i < len(newData); i++ {
		newData[i].ID = oldData[i].ID
	}

	for j, buf := range b.data {
		for i, new := range newData {
			if buf.ID == new.ID {
				b.data[j] = newData[i]
			}
		}
	}

	return entries
}

func (b *Buffer) Delete(users []models.User) int {
	var entries int
	//get ids
	for i, user := range users {
		users[i] = b.getData(user)
	}

	for _, user := range users {
		for _, data := range b.data {
			if data.ID == user.ID {
				b.data = remove(b.data, user)
				b.count--
			}
		}
	}
	return entries
}
