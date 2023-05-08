package restmodel

import "time"

type Comment struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	InventoryID int64     `json:"inventory_id"`
	Rating      int64     `json:"rating"`
	Text        string    `json:"text"`
	CreatedAt   time.Time `json:"created_at"`
}
