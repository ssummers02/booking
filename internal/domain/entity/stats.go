package entity

import "time"

type Stats struct {
	Count       int64
	InventoryID int64
	Date        time.Time
}

type StatisticFilter struct {
	StartDate time.Time        `json:"start_time" validate:"required"`
	EndDate   time.Time        `json:"end_time" validate:"required"`
	GroupBy   StatisticGroupBy `json:"group_by" validate:"required"`
}

type StatisticGroupBy string

const (
	StatisticGroupByDay   StatisticGroupBy = "DAY"
	StatisticGroupByMonth StatisticGroupBy = "MONTH"
	StatisticGroupByYear  StatisticGroupBy = "YEAR"
)
