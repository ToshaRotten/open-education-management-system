package buffer

import (
	"fmt"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/APIServer/UserManager/models"
)

// Buffer struct used for store users
type Buffer struct {
	Users []models.User
}

// New - returns a pointer to new instance of Buffer
func New() *Buffer {
	return &Buffer{
		Users: nil,
	}
}

// Add - Append user to buffer
func (b *Buffer) Add(user models.User) {
	b.Users = append(b.Users, user)
}

// Search - search user in buffer, returns true if it is existed
func (b *Buffer) Search(user models.User) bool {
	for _, temp := range b.Users {
		if temp == user {
			return true
		}
	}
	return false
}

// Count - returns number of entries
func (b *Buffer) Count() int {
	var counter int
	for counter, _ = range b.Users {
		counter++
	}
	return counter
}

// SearchByHash - search user by hash, returns true if it is existed
func (b *Buffer) SearchByHash(user models.User) bool {
	for _, temp := range b.Users {
		if temp.Hash == user.Hash {
			return true
		}
	}
	return false
}

// SearchByEmail - search user by email, returns true if it is existed
func (b *Buffer) SearchByEmail(user models.User) bool {
	for _, temp := range b.Users {
		if temp.Email == user.Email {
			return true
		}
	}
	return false
}

// Replace - replace user in buffer struct
func (b *Buffer) Replace(user models.User) {
	for i, temp := range b.Users {
		if temp.Hash == user.Hash {
			b.Users[i] = user
		}
	}
}

// Print - cli representation of Buffer struct
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
