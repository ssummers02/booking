package restmodel

type Resort struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name" validate:"required"`
	CityID      int64   `json:"city_id" validate:"required"`
	OwnerID     int64   `json:"owner_id"`
	Description string  `json:"description"`
	Address     string  `json:"address" validate:"required"`
	AvgRating   float64 `json:"avg_rating"`
}
