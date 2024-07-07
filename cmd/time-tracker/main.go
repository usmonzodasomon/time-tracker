package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/usmonzodasomon/time-tracker/internal/handler"
	"github.com/usmonzodasomon/time-tracker/pkg/logger"
	"github.com/usmonzodasomon/time-tracker/pkg/postgres"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// @title time-tracker API
// @version 1.0
// @description This is the API for the Time Tracker application.
// @host localhost:8080
// @BasePath /api
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	dbConn, err := postgres.GetConnection(postgres.Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DATABASE"),
	})
	defer postgres.CloseConnection(dbConn)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	logger.InitLogger(os.Getenv("GO_ENV"))

	router := gin.New()
	handler.NewRouter(router, dbConn)

	go func() {
		logger.Logger.Info(fmt.Sprintf("starting server on port %s", os.Getenv("PORT")))
		if err := http.ListenAndServe(":"+os.Getenv("PORT"), router); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Logger.Error("failed to start server", slog.String("error", err.Error()))
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGTERM, syscall.SIGINT)
	<-done

	logger.Logger.Info("service stopped")
}
