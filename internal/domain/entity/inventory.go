package entity

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
	StartDate *string
	Duration  *int64
}
