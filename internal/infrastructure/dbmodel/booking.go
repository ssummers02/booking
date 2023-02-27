package dbmodel

import "time"

type Booking struct {
	ID          int64     `db:"id"`
	UserID      int64     `db:"user_id"`
	InventoryID int64     `db:"inventory_id"`
	Inventory   Inventory `db:"inventory"`
	Resort      Resort    `db:"resort"`
	TotalPrice  float64   `db:"total_price"`
	StartDate   time.Time `db:"start_date"`
	EndDate     time.Time `db:"end_date"`
}
