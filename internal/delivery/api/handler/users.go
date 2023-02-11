package handler

import (
	"booking/internal/delivery/api/restmodel"
	"booking/internal/domain/dto"
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) signIn(w http.ResponseWriter, r *http.Request) {
	var (
		data restmodel.User
		ctx  = r.Context()
	)

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		SendErr(w, http.StatusBadRequest, "invalid json")

		return
	}

	user, err := s.services.UsersService.GetUsersByEmail(ctx, data.Email)
	if err != nil {
		SendErr(w, http.StatusNotFound, err.Error())

		return
	}

	credentialError := user.CheckPassword(data.Password)
	if credentialError != nil {
		SendErr(w, http.StatusUnauthorized, "invalid credentials")

		return
	}

	jwt, err := s.m.Auth.GenerateJWT(user.Email)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	SendOK(w, http.StatusOK, dto.UserToRest(user, jwt))
}

func (s *Server) register(w http.ResponseWriter, r *http.Request) {
	var (
		data restmodel.User
		ctx  = r.Context()
	)

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("error: %v", err)
		SendErr(w, http.StatusBadRequest, "invalid json")

		return
	}

	err = s.v.Struct(data)
	if err != nil {
		SendErr(w, http.StatusBadRequest, err.Error())

		return
	}

	user, err := s.services.UsersService.CreateUser(ctx, dto.UserFromRest(data))
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())

		return
	}

	jwt, err := s.m.Auth.GenerateJWT(user.Email)
	if err != nil {
		return
	}

	SendOK(w, http.StatusOK, dto.UserToRest(user, jwt))
}
