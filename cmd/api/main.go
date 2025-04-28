package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ryanadiputraa/tetra/config"
	"github.com/ryanadiputraa/tetra/internal/server"
	"github.com/ryanadiputraa/tetra/pkg/cache"
	"github.com/ryanadiputraa/tetra/pkg/db"
	"github.com/ryanadiputraa/tetra/pkg/secure"
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

	rdb, err := cache.NewRedis(c)
	if err != nil {
		logger.Error("Fail to open redis client", "error", err.Error())
		os.Exit(1)
		return
	}

	secure, err := secure.New(c.EncryptionKey)
	if err != nil {
		logger.Error("Fail to setup encryption", "error", err.Error())
		os.Exit(1)
		return
	}

	s := server.New(c, logger, db, rdb, secure)
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
	if err := sqlDB.Close(); err != nil {
		logger.Error("Fail to close DB connection", "error", err.Error())
	}
	if err := rdb.Close(); err != nil {
		logger.Error("Fail to close redis client", "error", err.Error())
	}

	slog.Info("Server stop successfully")
}
