package dbmodel

import "time"

type Stats struct {
	Count       int64     `db:"count"`
	InventoryID int64     `db:"inventory_id"`
	Date        time.Time `db:"date"`
}
