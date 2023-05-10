package entity

import "time"

type Comment struct {
	ID          int64
	UserID      int64
	UserName    string
	InventoryID int64
	Rating      int64
	Text        string
	CreatedAt   time.Time
}
