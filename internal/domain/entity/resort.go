package entity

import "time"

type Resort struct {
	ID          int64
	Name        string
	CityID      int64
	Description string
	OwnerID     int64
	Address     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type City struct {
	ID   int64
	Name string
}

type Filter struct {
	CityID    *int64  `json:"city_id"`
	TypeID    *int64  `json:"type_id"`
	StartDate *string `json:"start_date"`
	Duration  *int64  `json:"duration"`
}
