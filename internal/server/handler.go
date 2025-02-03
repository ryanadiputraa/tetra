package server

import (
	"net/http"

	"github.com/ryanadiputraa/inventra/config"
	authHandler "github.com/ryanadiputraa/inventra/internal/auth/handler"
	authService "github.com/ryanadiputraa/inventra/internal/auth/service"
	oauthHandler "github.com/ryanadiputraa/inventra/internal/oauth/handler"
	userRepository "github.com/ryanadiputraa/inventra/internal/user/repository"
	userService "github.com/ryanadiputraa/inventra/internal/user/service"
	"github.com/ryanadiputraa/inventra/pkg/jwt"
	"github.com/ryanadiputraa/inventra/pkg/logger"
	"github.com/ryanadiputraa/inventra/pkg/oauth"
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
	oauth := oauth.NewGoogleOauth(&oauth.GoogleOauthConfig{
		State:        c.OauthState,
		RedirectURL:  c.OauthRedirectURL,
		ClientID:     c.GoogleClientID,
		ClientSecret: c.GoogleClientSecret,
	})

	userRepository := userRepository.New(db)
	userService := userService.New(log, userRepository)

	authService := authService.New(log, jwt, userRepository)
	authHandler := authHandler.New(writer, validator, jwt, authService)

	oauthHandler := oauthHandler.New(log, c, oauth, userService, jwt)

	router.Handle("POST /auth/login", authHandler.Login())
	router.Handle("POST /auth/register", authHandler.Register())

	router.Handle("GET /oauth/login/google", oauthHandler.GoogleSignin())
	router.Handle("GET /oauth/callback/google", oauthHandler.GoogleCallback())
	return router
}
