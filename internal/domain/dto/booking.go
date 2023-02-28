package dto

import (
	"booking/internal/delivery/api/restmodel"
	"booking/internal/domain/entity"
	"booking/internal/infrastructure/dbmodel"
)

func BookingFromDB(booking dbmodel.Booking) entity.Booking {
	return entity.Booking{
		ID:          booking.ID,
		UserID:      booking.UserID,
		InventoryID: booking.InventoryID,
		Inventory:   InventoryFromDB(booking.Inventory),
		Resort:      ResortFromDB(booking.Resort),
		TotalPrice:  booking.TotalPrice,
		StartDate:   booking.StartDate,
		EndDate:     booking.EndDate,
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
		InventoryID: booking.InventoryID,
		Inventory:   InventoryToRest(booking.Inventory),
		Resort:      ResortToRest(booking.Resort),
		TotalPrice:  booking.TotalPrice,

		StartDate: booking.StartDate,
		EndDate:   booking.EndDate,
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
		StartDate:   booking.StartDate,
		EndDate:     booking.EndDate,
	}
}

func BookingCreateEntity(booking restmodel.Booking, userID int64) entity.Booking {
	endDay := booking.StartDate.AddDate(0, 0, int(*booking.Duration))

	return entity.Booking{
		UserID:      userID,
		InventoryID: booking.InventoryID,
		StartDate:   booking.StartDate,
		EndDate:     endDay,
	}
}