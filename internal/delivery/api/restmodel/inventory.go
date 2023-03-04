package restmodel

type Inventory struct {
	ID       int64  `json:"id"`
	TypeID   int64  `json:"type_id" validate:"required"`
	ResortID int64  `json:"resort_id" validate:"required"`
	Price    int64  `json:"price" validate:"required"`
	Photo    string `json:"photo"`
}

type InventoryType struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type InventoryFilter struct {
	ResortID  int64   `json:"resort_id"`
	TypeID    *int64  `json:"type_id"`
	StartDate *string `json:"start_date"`
	Duration  *int64  `json:"duration"`
}
