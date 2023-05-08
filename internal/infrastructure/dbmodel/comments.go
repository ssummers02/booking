package dbmodel

import "time"

type Comment struct {
	ID          int64     `db:"id"`
	UserID      int64     `db:"user_id"`
	InventoryID int64     `db:"inventory_id"`
	Rating      int64     `db:"rating"`
	Text        string    `db:"text"`
	CreatedAt   time.Time `db:"created_at"`
}
