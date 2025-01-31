package server

import (
	"net/http"

	"github.com/ryanadiputraa/inventra/config"
	authHandler "github.com/ryanadiputraa/inventra/internal/auth/handler"
	authService "github.com/ryanadiputraa/inventra/internal/auth/service"
	userRepository "github.com/ryanadiputraa/inventra/internal/user/repository"
	userService "github.com/ryanadiputraa/inventra/internal/user/service"
	"github.com/ryanadiputraa/inventra/pkg/jwt"
	"github.com/ryanadiputraa/inventra/pkg/logger"
	"github.com/ryanadiputraa/inventra/pkg/validator"
	"github.com/ryanadiputraa/inventra/pkg/writer"
	"gorm.io/gorm"
)

func setupHandler(c config.Config, db *gorm.DB) http.Handler {
	router := http.NewServeMux()

	log := logger.New()
	writer := writer.NewHTTPWriter()
	validator := validator.NewValidator()
	jwt := jwt.NewJWT(c.JWTSecret)

	userRepository := userRepository.New(db)
	userService := userService.New(log, userRepository)

	authService := authService.New(log, jwt)
	authHandler := authHandler.New(writer, validator, jwt, authService, userService)

	router.Handle("POST /auth/login", authHandler.Login())
	return router
}
