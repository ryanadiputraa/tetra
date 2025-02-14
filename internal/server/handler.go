package server

import (
	"log/slog"
	"net/http"

	"github.com/redis/go-redis/v9"
	"github.com/ryanadiputraa/inventra/config"
	"github.com/ryanadiputraa/inventra/internal/auth"
	authHandler "github.com/ryanadiputraa/inventra/internal/auth/handler"
	authService "github.com/ryanadiputraa/inventra/internal/auth/service"
	"github.com/ryanadiputraa/inventra/internal/middleware"
	oauthHandler "github.com/ryanadiputraa/inventra/internal/oauth/handler"
	organizationHandler "github.com/ryanadiputraa/inventra/internal/organization/handler"
	organizationRepository "github.com/ryanadiputraa/inventra/internal/organization/repository"
	organizationService "github.com/ryanadiputraa/inventra/internal/organization/service"
	userHandler "github.com/ryanadiputraa/inventra/internal/user/handler"
	userRepository "github.com/ryanadiputraa/inventra/internal/user/repository"
	userService "github.com/ryanadiputraa/inventra/internal/user/service"
	"github.com/ryanadiputraa/inventra/pkg/jwt"
	"github.com/ryanadiputraa/inventra/pkg/mail"
	"github.com/ryanadiputraa/inventra/pkg/oauth"
	"github.com/ryanadiputraa/inventra/pkg/validator"
	"github.com/ryanadiputraa/inventra/pkg/writer"
	"gorm.io/gorm"
)

func setupHandler(c config.Config, logger *slog.Logger, db *gorm.DB, rdb *redis.Client) http.Handler {
	router := http.NewServeMux()

	writer := writer.NewHTTPWriter()
	validator := validator.NewValidator()
	jwt := jwt.NewJWT(c.JWTSecret)
	oauth := oauth.NewGoogleOauth(&oauth.GoogleOauthConfig{
		State:        c.OauthState,
		RedirectURL:  c.OauthRedirectURL,
		ClientID:     c.GoogleClientID,
		ClientSecret: c.GoogleClientSecret,
	})
	smtpMail := mail.NewSMTPMail(c.SMTPEmail, c.SMTPPassword)

	userRepository := userRepository.New(db, rdb)
	organizationRepository := organizationRepository.New(db, rdb)

	userService := userService.New(logger, userRepository)
	authService := authService.New(logger, jwt, userRepository)
	organizationService := organizationService.New(c, logger, jwt, smtpMail, organizationRepository, userRepository)

	authHandler := authHandler.New(writer, validator, jwt, authService)
	oauthHandler := oauthHandler.New(logger, c, oauth, userService, authService)
	userHandler := userHandler.New(writer, validator, userService)
	organizationHandler := organizationHandler.New(writer, organizationService, validator)

	authMiddleware := middleware.NewAuthMiddleware(writer, jwt, userService, organizationService)

	staffAccessLv := auth.AccessLevel[auth.Staff]
	// supervisorAccessLv := auth.AccessLevel[auth.Supervisor]
	adminAccessLv := auth.AccessLevel[auth.Admin]

	router.Handle("POST /auth/login", authHandler.Login())
	router.Handle("POST /auth/register", authHandler.Register())
	router.Handle("GET /oauth/login/google", oauthHandler.GoogleSignin())
	router.Handle("GET /oauth/callback/google", oauthHandler.GoogleCallback())

	router.Handle("GET /api/users/profile", authMiddleware.AuthorizeUser(userHandler.GetUserData()))
	router.Handle("POST /api/users/password", authMiddleware.AuthorizeUser(userHandler.ChangePassword()))

	router.Handle("POST /api/organizations", authMiddleware.AuthorizeUser(organizationHandler.CreateOrganization()))
	router.Handle("GET /api/organizations/members", authMiddleware.AuthorizeUserRole(organizationHandler.FetchMembers(), staffAccessLv))
	router.Handle("POST /api/organizations/invite", authMiddleware.AuthorizeUserRole(organizationHandler.Invite(), adminAccessLv))
	return router
}
