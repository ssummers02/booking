package entity

import "time"

type Inventory struct {
	ID       int64
	TypeID   int64
	ResortID int64
	Price    int64
	Photo    string
}

type InventoryType struct {
	ID   int64
	Name string
}

type InventoryFilter struct {
	ResortID  int64
	TypeID    *int64
	StartTime *time.Time
	EndTime   *time.Time
}

type Img struct {
	ID          int64
	InventoryID int64
	Name        string
	Bytes       []byte
}
