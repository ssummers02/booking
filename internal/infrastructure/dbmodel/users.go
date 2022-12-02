package dbmodel

import (
	"time"
)

type User struct {
	ID         int64     `db:"id"`
	FirstName  string    `db:"first_name"`
	Surname    string    `db:"surname"`
	MiddleName string    `db:"middle_name"`
	Email      string    `db:"email"`
	Password   string    `db:"password"`
	Phone      string    `db:"phone"`
	RoleID     int64     `db:"role_id"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
