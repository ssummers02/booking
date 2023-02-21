package handler

import (
	"booking/internal/delivery/api/restmodel"
	"booking/internal/domain/dto"
	"booking/internal/domain/entity"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Возвращает массив городов
func (s *Server) getCities(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	cities, err := s.services.ResortsService.GetCities(ctx)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	SendOK(w, http.StatusOK, dto.CitiesToRest(cities))
}

// Возвращает массив всех курортов
func (s *Server) getResorts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	resorts, err := s.services.ResortsService.GetResorts(ctx)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	SendOK(w, http.StatusOK, dto.ResortsToRest(resorts))
}

// Возвращает курорт по ID
func (s *Server) getResortByID(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
		id  = mux.Vars(r)["id"]
	)

	parseID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	resort, err := s.services.ResortsService.GetResortByID(ctx, parseID)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	SendOK(w, http.StatusOK, dto.ResortToRest(resort))
}

// Возвращает массив всех курортов по городу
func (s *Server) getResortsByFilter(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = r.Context()
		data entity.Filter
	)

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		SendErr(w, http.StatusBadRequest, "invalid json")

		return
	}

	err = s.v.Struct(data)
	if err != nil {
		SendErr(w, http.StatusBadRequest, err.Error())

		return
	}
	resort, err := s.services.ResortsService.GetResortsByFilter(ctx, data)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	SendOK(w, http.StatusOK, dto.ResortsToRest(resort))
}

// Создает новый курорт
func (s *Server) createResort(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var data restmodel.Resort

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		SendErr(w, http.StatusBadRequest, "invalid json")
		return
	}

	err = s.v.Struct(data)
	if err != nil {
		SendErr(w, http.StatusBadRequest, err.Error())

		return
	}

	resort, err := s.services.ResortsService.CreateResort(ctx, dto.ResortFromRest(data))
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	SendOK(w, http.StatusOK, resort)
}

// Обновляет курорт
func (s *Server) updateResort(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var data restmodel.Resort

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		SendErr(w, http.StatusBadRequest, "invalid json")

		return
	}

	err = s.v.Struct(data)
	if err != nil {
		SendErr(w, http.StatusBadRequest, err.Error())

		return
	}

	resort, err := s.services.ResortsService.UpdateResort(ctx, dto.ResortFromRest(data))
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	SendOK(w, http.StatusOK, resort)
}

// Удаляет курорт
func (s *Server) deleteResort(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
		id  = mux.Vars(r)["id"]
	)

	parseID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	err = s.services.ResortsService.DeleteResort(ctx, parseID)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	SendOK(w, http.StatusOK, nil)
}
