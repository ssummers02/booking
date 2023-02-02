package handler

import (
	"booking/internal/delivery/api/restmodel"
	"booking/internal/domain/dto"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Получает инвентарь по id инвентаря
func (s *Server) getInventoryByID(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
		id  = mux.Vars(r)["id"]
	)

	parseID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	inventory, err := s.services.InventoryService.GetInventoryByID(ctx, parseID)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	SendOK(w, http.StatusOK, dto.InventoryToRest(inventory))
}

// Массив доступных инвентарей по id отеля
func (s *Server) getInventoryByResortID(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
		id  = mux.Vars(r)["id"]
	)

	parseID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	inventory, err := s.services.InventoryService.GetInventoryByResortID(ctx, parseID)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	SendOK(w, http.StatusOK, dto.InventorysToRest(inventory))
}

// Создает инвентарь
func (s *Server) createInventory(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
	)

	var data restmodel.Inventory

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

	inventory, err := s.services.InventoryService.CreateInventory(ctx, dto.InventoryFromRest(data))
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	SendOK(w, http.StatusOK, dto.InventoryToRest(inventory))
}

// обновляет инвентарь
func (s *Server) updateInventory(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = r.Context()
		data restmodel.Inventory
	)

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		SendErr(w, http.StatusBadRequest, "invalid json")
		return
	}

	inventory, err := s.services.InventoryService.UpdateInventory(ctx, dto.InventoryFromRest(data))
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	SendOK(w, http.StatusOK, dto.InventoryToRest(inventory))
}

func (s *Server) deleteInventory(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
		id  = mux.Vars(r)["id"]
	)

	parseID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	err = s.services.InventoryService.DeleteInventory(ctx, parseID)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	SendOK(w, http.StatusOK, nil)
}

// Возвращает доступные типы инвентаря для создания
func (s *Server) getInventoryTypes(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
	)

	inventoryTypes, err := s.services.InventoryService.GetInventoriesTypes(ctx)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	SendOK(w, http.StatusOK, dto.InventoryTypesToRest(inventoryTypes))
}

// Возвращает инвертарь по курорту
func (s *Server) getInventoryByResort(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
		id  = mux.Vars(r)["id"]
	)

	parseID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	inventory, err := s.services.InventoryService.GetInventoryByResortID(ctx, parseID)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	SendOK(w, http.StatusOK, dto.InventorysToRest(inventory))
}
