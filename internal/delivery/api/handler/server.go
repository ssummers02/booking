package handler

import (
	"booking/internal/delivery/api/middleware"
	"booking/internal/domain/service"
	"context"

	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Server struct {
	httpServer *http.Server
	r          *mux.Router
	v          *validator.Validate
	services   *service.Service
	m          *middleware.M
}

//nolint:gomnd
func NewServer(port string, services *service.Service, m *middleware.M) *Server {
	r := mux.NewRouter()

	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + port,
			Handler:        r,
			MaxHeaderBytes: 1 << 20, // 1 MB
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		},
		r:        r,
		v:        validator.New(),
		services: services,
		m:        m,
	}
}

func (s *Server) Run() error {
	// s.r.Use(s.m.Recovery.Handler, s.m.Cors.Handler, s.m.Auth.Handler)
	s.r.Use(s.m.Cors.Handler, s.m.Auth.Handler)
	s.initRoutes()

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) initRoutes() {
	router := s.r.PathPrefix("/api").
		Subrouter()
	router.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write([]byte("ok"))
	})

	router.HandleFunc("/user/register", s.register).Methods(http.MethodPost)
	router.HandleFunc("/user/login", s.signIn).Methods(http.MethodPost)

	router.HandleFunc("/cities", s.getCities).Methods(http.MethodGet)

	router.HandleFunc("/resorts/{id:[0-9]+}", s.getResortByID).Methods(http.MethodGet)
	router.HandleFunc("/resorts/filter", s.getResortsByFilter).Methods(http.MethodPost)
	router.HandleFunc("/resorts", s.getResorts).Methods(http.MethodGet)
	router.HandleFunc("/myresorts", s.getMyResorts).Methods(http.MethodGet)
	router.HandleFunc("/resorts", s.createResort).Methods(http.MethodPost)
	router.HandleFunc("/resorts", s.updateResort).Methods(http.MethodPut)
	router.HandleFunc("/resorts/{id:[0-9]+}", s.deleteResort).Methods(http.MethodDelete)
	router.HandleFunc("/resorts/inventories/{id:[0-9]+}", s.getInventoryByResort).Methods(http.MethodGet)

	router.HandleFunc("/inventories/{id:[0-9]+}", s.getInventoryByID).Methods(http.MethodGet)
	router.HandleFunc("/inventories/filter", s.getInventoriesByFilters).Methods(http.MethodPost)

	router.HandleFunc("/inventories", s.createInventory).Methods(http.MethodPost)
	router.HandleFunc("/inventories", s.updateInventory).Methods(http.MethodPut)
	router.HandleFunc("/inventories/types", s.getInventoryTypes).Methods(http.MethodGet)
	router.HandleFunc("/inventories/{id:[0-9]+}", s.deleteInventory).Methods(http.MethodDelete)

	router.HandleFunc("/booking/{id:[0-9]+}", s.getBookingByID).Methods(http.MethodGet)
	router.HandleFunc("/user/bookings", s.getBookingsByUserID).Methods(http.MethodGet)
	router.HandleFunc("/resorts/bookings/{id:[0-9]+}", s.getBookingByResortID).Methods(http.MethodGet)
	router.HandleFunc("/resorts/bookings", s.getBookingByOwner).Methods(http.MethodGet)

	router.HandleFunc("/booking", s.createBooking).Methods(http.MethodPost)
	/*	router.HandleFunc("/user", s.getUser).
			Methods(http.MethodGet)
		router.HandleFunc("/user", s.createUser).
			Methods(http.MethodPost)
		router.HandleFunc("/user", s.updateUser).
			Methods(http.MethodPut)
		router.HandleFunc("/user", s.deleteUser).
			Methods(http.MethodDelete)*/
}
