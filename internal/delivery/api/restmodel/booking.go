package restmodel

import "time"

type Booking struct {
	ID          int64     `json:"id" validate:"required"`
	UserID      int64     `json:"user_id"`
	InventoryID int64     `json:"inventory_id"`
	Inventory   Inventory `json:"inventory"`
	Resort      Resort    `json:"resort"`
	TotalPrice  float64   `json:"total_price"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Duration    *int64    `json:"duration"`
}
