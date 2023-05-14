package restmodel

import "time"

type Stats struct {
	Count       int64     `json:"count"`
	InventoryID int64     `json:"inventory_id"`
	Date        time.Time `json:"date"`
}
