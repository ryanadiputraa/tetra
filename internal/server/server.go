package server

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/ryanadiputraa/tetra/config"
	"github.com/ryanadiputraa/tetra/internal/middleware"
	"github.com/ryanadiputraa/tetra/pkg/secure"
	"gorm.io/gorm"
)

func New(c config.Config, logger *slog.Logger, db *gorm.DB, rdb *redis.Client, secure secure.Secure) *http.Server {
	handler := setupHandler(c, logger, db, rdb, secure)
	handler = registerMiddlewares(
		handler,
		middleware.CORSMiddleware,
		middleware.ThrottleMiddleware,
	)

	server := &http.Server{
		Addr:         c.Port,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return server
}

func registerMiddlewares(handler http.Handler, middlewares ...func(handler http.Handler) http.Handler) http.Handler {
	for _, m := range middlewares {
		handler = m(handler)
	}
	return handler
}
