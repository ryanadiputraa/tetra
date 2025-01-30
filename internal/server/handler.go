package server

import (
	"net/http"

	"github.com/ryanadiputraa/inventra/config"
	"gorm.io/gorm"
)

func setupHandler(c config.Config, db *gorm.DB) http.Handler {
	router := http.NewServeMux()
	return router
}
