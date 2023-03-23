package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ssummers02/booking/internal/bootstrap"
	"github.com/ssummers02/booking/internal/delivery/api/handler"
	"github.com/ssummers02/booking/internal/delivery/api/middleware"
	"github.com/ssummers02/booking/internal/domain/service"
	"github.com/ssummers02/booking/internal/infrastructure/repository"

	"github.com/gocraft/dbr/v2"
	"github.com/sirupsen/logrus"
)

const pathToMigrations = "migrations"

func Run() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	c, err := bootstrap.NewConfig()
	if err != nil {
		log.Fatalln("Error loading config:", err)
	}

	dbPool, err := bootstrap.NewDBConn(c.DB.Username, c.DB.Password, c.DB.Name, c.DB.Host, c.DB.Port)
	if err != nil {
		log.Fatal("Error connecting to DB:", err)
	}

	defer func(dbPool *dbr.Connection) {
		err := dbPool.Close()
		if err != nil {
			log.Println("Error closing DB pool:", err)
		}
	}(dbPool)

	err = bootstrap.UpMigrations(dbPool.DB, c.DB.Name, pathToMigrations)
	if err != nil {
		log.Fatal("Error up migrations:", err)
	}

	repo := repository.NewRepository(dbPool)
	middlewares := middleware.NewMiddlewares(repo.User)

	services := service.NewServices(repo)
	srv := handler.NewServer(c.HTTPPort, services, middlewares)

	go func() {
		if err := srv.Run(); err != nil {
			logrus.Fatalf("error occurred while running http server: %s", err.Error())
		}
	}()

	logrus.Print("App Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("App Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occurred on server shutting down: %s", err.Error())
	}
}
