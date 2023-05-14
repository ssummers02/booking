package restmodel

import "time"

type Inventory struct {
	ID       int64  `json:"id"`
	TypeID   int64  `json:"type_id" validate:"required"`
	ResortID int64  `json:"resort_id" validate:"required"`
	Price    int64  `json:"price" validate:"required"`
	Photo    []byte `json:"photo"`
	Name     string `json:"inventory_name"`
}

type InventoryType struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type InventoryFilter struct {
	ResortID  int64      `json:"resort_id"`
	TypeID    *int64     `json:"type_id"`
	StartTime *time.Time `json:"start_time"`
	EndTime   *time.Time `json:"end_time"`
}

type Img struct {
	ID          int64  `json:"id"`
	InventoryID int64  `json:"inventory_id"`
	Name        string `json:"name"`
	Bytes       []byte `json:"bytes"`
}
