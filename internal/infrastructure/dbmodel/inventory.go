package dbmodel

type Inventory struct {
	ID       int64  `db:"id"`
	TypeID   int64  `db:"type_id"`
	ResortID int64  `db:"resort_id"`
	Price    int64  `db:"price"`
	Photo    string `db:"photo"`
}

type InventoryType struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}
