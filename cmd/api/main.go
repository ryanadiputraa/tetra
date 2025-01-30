package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ryanadiputraa/inventra/config"
	"github.com/ryanadiputraa/inventra/internal/server"
	"github.com/ryanadiputraa/inventra/pkg/db"
	"github.com/ryanadiputraa/inventra/pkg/logger"
)

func main() {
	log := logger.New()

	c, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Fail to load config. Err: ", err)
		return
	}

	db, sqlDB, err := db.NewPostgres(c)
	if err != nil {
		log.Fatal("Fail to open db connection. Err: ", err)
		return
	}
	defer func() {
		if err := sqlDB.Close(); err != nil {
			log.Error("Fail to close DB connection. Err: ", err)
		}
	}()

	s := server.New(c, db)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Info("Starting server on port", c.Port)
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Fail to start server. Err: ", err)
		}
	}()

	<-done
	log.Info("Shutting down gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	if err := s.Shutdown(ctx); err == context.DeadlineExceeded {
		log.Error("Error shutting down server. Err: ", err.Error())
	}
	log.Info("Server stop successfully")
}
