package server

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/ryanadiputraa/inventra/config"
	"github.com/ryanadiputraa/inventra/internal/middleware"
	"gorm.io/gorm"
)

func New(c config.Config, logger *slog.Logger, db *gorm.DB, rdb *redis.Client) *http.Server {
	handler := setupHandler(c, logger, db, rdb)
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
