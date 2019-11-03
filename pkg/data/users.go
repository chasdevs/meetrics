package data

import (
	"log"
	"time"
)

type User struct {
	ID         uint
	Email      string `gorm:"not null;unique"`
	Name       string
	Department string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (mgr *manager) AddAllUsers(users []*User) {
	for _, user := range users {
		errors := mgr.db.Create(user).GetErrors()
		if len(errors) > 0 {
			log.Fatalf("Error adding user: %v", errors[0])
		}
	}
}

func (mgr *manager) GetAllUsers() []User {
	var users []User
	mgr.db.Find(&users)
	return users
}

func (mgr *manager) GetUserByEmail(email string) User {
	var user User
	mgr.db.First(&user, "email = ?", email)
	return user
}

func (mgr *manager) GetUserById(id int) User {
	var user User
	mgr.db.First(&user, "id = ?", id)
	return user
}
