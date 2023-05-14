package dto

import (
	"github.com/ssummers02/booking/internal/delivery/api/restmodel"
	"github.com/ssummers02/booking/internal/domain/entity"
	"github.com/ssummers02/booking/internal/infrastructure/dbmodel"
)

func BookingFromDB(booking dbmodel.Booking) entity.Booking {
	return entity.Booking{
		ID:          booking.ID,
		UserID:      booking.UserID,
		InventoryID: booking.InventoryID,
		Inventory:   InventoryFromDB(booking.Inventory),
		Resort:      ResortFromDB(booking.Resort),
		TotalPrice:  booking.TotalPrice,
		StartTime:   booking.StartTime,
		EndTime:     booking.EndTime,
	}
}

func BookingsFromDB(bookings []dbmodel.Booking) []entity.Booking {
	result := make([]entity.Booking, 0, len(bookings))

	for i := range bookings {
		result = append(result, BookingFromDB(bookings[i]))
	}

	return result
}

func BookingToRest(booking entity.Booking) restmodel.Booking {
	return restmodel.Booking{
		ID:          booking.ID,
		UserID:      booking.UserID,
		User:        UserToRestWithoutToken(booking.User),
		InventoryID: booking.InventoryID,
		Inventory:   InventoryToRest(booking.Inventory),
		Resort:      ResortToRest(booking.Resort),
		TotalPrice:  booking.TotalPrice,

		StartTime: booking.StartTime,
		EndTime:   booking.EndTime,
	}
}

func BookingsToRest(bookings []entity.Booking) []restmodel.Booking {
	result := make([]restmodel.Booking, 0, len(bookings))

	for i := range bookings {
		result = append(result, BookingToRest(bookings[i]))
	}

	return result
}

func BookingToDB(booking entity.Booking) dbmodel.Booking {
	return dbmodel.Booking{
		UserID:      booking.UserID,
		InventoryID: booking.InventoryID,
		TotalPrice:  booking.TotalPrice,
		StartTime:   booking.StartTime,
		EndTime:     booking.EndTime,
	}
}

func BookingCreateEntity(booking restmodel.Booking, userID int64) entity.Booking {
	return entity.Booking{
		UserID:      userID,
		InventoryID: booking.InventoryID,
		StartTime:   booking.StartTime,
		EndTime:     booking.EndTime,
	}
}

func StatFromDB(stats dbmodel.Stats) entity.Stats {
	return entity.Stats{
		Count:       stats.Count,
		InventoryID: stats.InventoryID,
		Date:        stats.Date,
	}
}

func StatsFromDB(stats []dbmodel.Stats) []entity.Stats {
	result := make([]entity.Stats, 0, len(stats))

	for i := range stats {
		result = append(result, StatFromDB(stats[i]))
	}

	return result
}
func StatToRest(stats entity.Stats) restmodel.Stats {
	return restmodel.Stats{
		Count:       stats.Count,
		InventoryID: stats.InventoryID,
		Date:        stats.Date,
	}
}

func StatsToRest(stats []entity.Stats) []restmodel.Stats {
	result := make([]restmodel.Stats, 0, len(stats))

	for i := range stats {
		result = append(result, StatToRest(stats[i]))
	}

	return result
}
