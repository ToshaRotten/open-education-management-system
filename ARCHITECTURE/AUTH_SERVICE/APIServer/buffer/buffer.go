package buffer

import (
	"fmt"
	"github.com/ToshaRotten/open-educaton-management-system/ARCHITECTURE/AUTH_SERVICE/APIServer/APIServer/models"
)

type Buffer struct {
	Users []models.User
}

func New() *Buffer {
	return &Buffer{
		Users: nil,
	}
}

func (b *Buffer) Add(user models.User) {
	b.Users = append(b.Users, user)
}

func (b *Buffer) Remove(user models.User) {
}

func (b *Buffer) Search(user models.User) bool {
	for _, temp := range b.Users {
		if temp == user {
			return true
		}
	}
	return false
}

func (b *Buffer) Count() int {
	var counter int
	for counter, _ = range b.Users {
		counter++
	}
	return counter
}

func (b *Buffer) SearchByHash(user models.User) bool {
	for _, temp := range b.Users {
		if temp.Hash == user.Hash {
			return true
		}
	}
	return false
}

func (b *Buffer) Replace(user models.User) {
	for i, temp := range b.Users {
		if temp.Hash == user.Hash {
			b.Users[i] = user
		}
	}
}

func (b *Buffer) Print() {
	for i := 0; i < 6; i++ {
		if i == 0 {
			fmt.Print("ID \t NAME \t")
		}
		if i == 0 {
			fmt.Print("SNAME \t")
		}
		if i == 0 {
			fmt.Print("TNAME \t")
		}
		if i == 0 {
			fmt.Print("DOB \t")
		}
		if i == 0 {
			fmt.Print("ROLE \t")
		}
		if i == 0 {
			fmt.Print("HASH \t\n")
		}
	}
	for i, temp := range b.Users {
		fmt.Printf("%d \t", i)
		fmt.Printf("%s \t", temp.FirstName)
		fmt.Printf("%s \t", temp.LastName)
		fmt.Printf("%s \t", temp.ThirdName)
		fmt.Printf("%s \t", temp.DOB)
		fmt.Printf("%s \t", temp.Role)
		fmt.Printf("%s \t", temp.Hash)
		fmt.Println()
	}
}
