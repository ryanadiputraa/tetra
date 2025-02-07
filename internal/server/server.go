package server

import (
	"net/http"
	"time"

	"github.com/ryanadiputraa/inventra/config"
	"github.com/ryanadiputraa/inventra/internal/middleware"
	"gorm.io/gorm"
)

func New(c config.Config, db *gorm.DB) *http.Server {
	handler := setupHandler(c, db)
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
