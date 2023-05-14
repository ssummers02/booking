package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ssummers02/booking/internal/delivery/api/restmodel"
	"github.com/ssummers02/booking/internal/domain/dto"
	"github.com/ssummers02/booking/internal/domain/entity"
)

// Возвращает комментарий по ID.
func (s *Server) getCommentByID(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
		id  = mux.Vars(r)["id"]
	)

	parseID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	resort, err := s.services.CommentService.GetCommentByID(ctx, parseID)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	SendOK(w, http.StatusOK, dto.CommentToRest(resort))
}

// Возвращает комментарии по ID инвентаря.
func (s *Server) getCommentsByInventory(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
		id  = mux.Vars(r)["id"]
	)

	parseID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	resorts, err := s.services.CommentService.GetCommentsByInventory(ctx, parseID)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	SendOK(w, http.StatusOK, dto.CommentsToRest(resorts))
}

// Создает бронирование.
func (s *Server) createComment(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = r.Context()
		data restmodel.Comment
	)

	user, ok := ctx.Value("user").(entity.User)
	if !ok {
		SendErr(w, http.StatusForbidden, "not authorized")
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		SendErr(w, http.StatusBadRequest, fmt.Errorf("invalid json: %w", err).Error())

		return
	}

	Comment, err := s.services.CommentService.CreateComment(ctx, dto.CommentCreateEntity(data, user.ID))
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	SendOK(w, http.StatusOK, dto.CommentToRest(Comment))
}

func (s *Server) getCommentByResortID(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
		id  = mux.Vars(r)["id"]
	)

	parseID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	resort, err := s.services.CommentService.GetCommentByResortID(ctx, parseID)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	SendOK(w, http.StatusOK, dto.CommentsToRest(resort))
}

func (s *Server) deleteComment(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
		id  = mux.Vars(r)["id"]
	)

	parseID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	err = s.services.CommentService.DeleteComment(ctx, parseID)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	SendOK(w, http.StatusOK, nil)
}
