package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type RoleType int64

const (
	admin = 1
	user  = 2
	owner = 3
)

type User struct {
	ID         int64
	FirstName  string
	Surname    string
	MiddleName string
	Email      string
	Password   string
	Phone      string
	RoleID     RoleType
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (user *User) HashPassword(password string) error {
	const cost = 14

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return err
	}

	user.Password = string(bytes)

	return nil
}
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}

	return nil
}
func (user *User) IsAuthorized() bool {
	return user.ID != 0
}

func (user *User) IsOwnerRole() bool {
	return user.RoleID == owner
}
