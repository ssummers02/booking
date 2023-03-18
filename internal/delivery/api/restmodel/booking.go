package restmodel

import "time"

type Booking struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	InventoryID int64     `json:"inventory_id"`
	Inventory   Inventory `json:"inventory"`
	Resort      Resort    `json:"resort"`
	TotalPrice  float64   `json:"total_price"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
}
