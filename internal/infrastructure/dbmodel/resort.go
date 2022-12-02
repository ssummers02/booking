package dbmodel

import "time"

type Resort struct {
	ID          int64     `db:"id"`
	Name        string    `db:"name"`
	CityID      int64     `db:"city_id"`
	OwnerID     int64     `db:"user_id"`
	Description string    `db:"description"`
	Address     string    `db:"address"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type City struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}
