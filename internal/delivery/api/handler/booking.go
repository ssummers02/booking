package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ssummers02/booking/internal/delivery/api/restmodel"
	"github.com/ssummers02/booking/internal/domain/dto"
	"github.com/ssummers02/booking/internal/domain/entity"

	"github.com/gorilla/mux"
)

// Возвращает бронирование по ID.
func (s *Server) getBookingByID(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
		id  = mux.Vars(r)["id"]
	)

	parseID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	resort, err := s.services.BookingService.GetBookingByID(ctx, parseID)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	SendOK(w, http.StatusOK, dto.BookingToRest(resort))
}

// Возвращает бронирования по ID пользователя.
func (s *Server) getBookingsByUserID(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = r.Context()
		user = r.Context().Value("user").(entity.User)
	)

	resorts, err := s.services.BookingService.GetBookingsByUserID(ctx, user.ID)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	SendOK(w, http.StatusOK, dto.BookingsToRest(resorts))
}

// Создает бронирование.
func (s *Server) createBooking(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = r.Context()
		user = r.Context().Value("user").(entity.User)
		data restmodel.Booking
	)

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		SendErr(w, http.StatusBadRequest, "invalid json")

		return
	}

	booking, err := s.services.BookingService.CreateBooking(ctx, dto.BookingCreateEntity(data, user.ID))
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	SendOK(w, http.StatusOK, dto.BookingToRest(booking))
}

func (s *Server) getBookingByResortID(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
		id  = mux.Vars(r)["id"]
	)

	parseID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	resort, err := s.services.BookingService.GetBookingByResortID(ctx, parseID)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	SendOK(w, http.StatusOK, dto.BookingsToRest(resort))
}

func (s *Server) getBookingByOwner(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
	)

	resort, err := s.services.BookingService.GetBookingByOwner(ctx)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	SendOK(w, http.StatusOK, dto.BookingsToRest(resort))
}
