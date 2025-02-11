package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ryanadiputraa/inventra/config"
	"github.com/ryanadiputraa/inventra/internal/server"
	"github.com/ryanadiputraa/inventra/pkg/db"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	c, err := config.LoadConfig()
	if err != nil {
		logger.Error("Fail to load config", "error", err.Error())
		os.Exit(1)
		return
	}

	db, sqlDB, err := db.NewPostgres(c)
	if err != nil {
		logger.Error("Fail to open DB connection", "error", err.Error())
		os.Exit(1)
		return
	}
	defer func() {
		if err := sqlDB.Close(); err != nil {
			logger.Error("Fail to close DB connection", "error", err.Error())
		}
	}()

	s := server.New(c, logger, db)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

	go func() {
		slog.Info("Starting server", "port", c.Port)
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Fail to start server", "error", err.Error())
		}
	}()

	<-done
	slog.Info("Shutting down gracefully")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	if err := s.Shutdown(ctx); err == context.DeadlineExceeded {
		slog.Error("Error while shutting down server", "error", err.Error())
	}
	slog.Info("Server stop successfully")
}
