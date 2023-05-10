package entity

import "time"

type Resort struct {
	ID          int64
	Name        string
	CityID      int64
	Description string
	OwnerID     int64
	Address     string
	AvgRating   float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type City struct {
	ID   int64
	Name string
}

type Filter struct {
	CityID    *int64     `json:"city_id"`
	TypeID    *int64     `json:"type_id"`
	StartTime *time.Time `json:"start_time"`
	EndTime   *time.Time `json:"end_time"`
}
