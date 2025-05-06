package server

import (
	"log/slog"
	"net/http"

	"github.com/redis/go-redis/v9"
	"github.com/ryanadiputraa/tetra/config"
	"github.com/ryanadiputraa/tetra/domain"
	authHandler "github.com/ryanadiputraa/tetra/internal/auth/handler"
	authService "github.com/ryanadiputraa/tetra/internal/auth/service"
	inventoryHandler "github.com/ryanadiputraa/tetra/internal/inventory/handler"
	inventoryRepository "github.com/ryanadiputraa/tetra/internal/inventory/repository"
	inventoryService "github.com/ryanadiputraa/tetra/internal/inventory/service"
	"github.com/ryanadiputraa/tetra/internal/middleware"
	oauthHandler "github.com/ryanadiputraa/tetra/internal/oauth/handler"
	organizationHandler "github.com/ryanadiputraa/tetra/internal/organization/handler"
	organizationRepository "github.com/ryanadiputraa/tetra/internal/organization/repository"
	organizationService "github.com/ryanadiputraa/tetra/internal/organization/service"
	userHandler "github.com/ryanadiputraa/tetra/internal/user/handler"
	userRepository "github.com/ryanadiputraa/tetra/internal/user/repository"
	userService "github.com/ryanadiputraa/tetra/internal/user/service"
	utilizationHandler "github.com/ryanadiputraa/tetra/internal/utilization/handler"
	utilizationRepository "github.com/ryanadiputraa/tetra/internal/utilization/repository"
	utilizationService "github.com/ryanadiputraa/tetra/internal/utilization/service"
	"github.com/ryanadiputraa/tetra/pkg/jwt"
	"github.com/ryanadiputraa/tetra/pkg/mail"
	"github.com/ryanadiputraa/tetra/pkg/oauth"
	"github.com/ryanadiputraa/tetra/pkg/pagination"
	"github.com/ryanadiputraa/tetra/pkg/secure"
	"github.com/ryanadiputraa/tetra/pkg/validator"
	"github.com/ryanadiputraa/tetra/pkg/writer"
	"gorm.io/gorm"
)

func setupHandler(c config.Config, logger *slog.Logger, db *gorm.DB, rdb *redis.Client, secure secure.Secure) http.Handler {
	router := http.NewServeMux()

	writer := writer.NewHTTPWriter()
	validator := validator.NewValidator()
	pagination := pagination.New()
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
	utilizationRepository := utilizationRepository.New(db)

	userService := userService.New(logger, userRepository)
	authService := authService.New(logger, userRepository)
	organizationService := organizationService.New(c, logger, jwt, smtpMail, secure, organizationRepository, userRepository)
	utilizationService := utilizationService.New(logger, utilizationRepository)

	authHandler := authHandler.New(writer, validator, jwt, authService)
	oauthHandler := oauthHandler.New(logger, c, oauth, jwt, userService)
	userHandler := userHandler.New(writer, validator, userService)
	organizationHandler := organizationHandler.New(c, writer, organizationService, validator, jwt)
	utilizationHandler := utilizationHandler.New(logger, c, writer, utilizationService)

	inventoryRepository := inventoryRepository.New(db)
	inventoryService := inventoryService.New(logger, inventoryRepository)
	inventoryHandler := inventoryHandler.New(writer, validator, pagination, inventoryService)
	authMiddleware := middleware.NewAuthMiddleware(writer, jwt, userService, organizationService)

	staffAccessLv := domain.AccessLevel[domain.Staff]
	supervisorAccessLv := domain.AccessLevel[domain.Supervisor]
	adminAccessLv := domain.AccessLevel[domain.Admin]

	router.Handle("POST /auth/login", authHandler.Login())
	router.Handle("POST /auth/register", authHandler.Register())
	router.Handle("GET /oauth/login/google", oauthHandler.GoogleSignin())
	router.Handle("GET /oauth/callback/google", oauthHandler.GoogleCallback())

	router.Handle("GET /api/users/profile", authMiddleware.AuthorizeUser(userHandler.GetUserData()))
	router.Handle("POST /api/users/password", authMiddleware.AuthorizeUser(userHandler.ChangePassword()))

	router.Handle("GET /api/utilizations/dashboard", authMiddleware.AuthorizeUserRole(utilizationHandler.GetUtilizations(), staffAccessLv))
	router.Handle("POST /api/utilizations/import", authMiddleware.AuthorizeUserRole(utilizationHandler.Import(), staffAccessLv))

	router.Handle("GET /api/organizations", authMiddleware.AuthorizeUserRole(organizationHandler.FetchOrganizationData(), staffAccessLv))
	router.Handle("POST /api/organizations", authMiddleware.AuthorizeUser(organizationHandler.CreateOrganization()))
	router.Handle("DELETE /api/organizations", authMiddleware.AuthorizeUserRole(organizationHandler.DeleteOrganization(), adminAccessLv))
	router.Handle("POST /api/organizations/invite", authMiddleware.AuthorizeUserRole(organizationHandler.Invite(), adminAccessLv))
	router.Handle("POST /api/organizations/join", authMiddleware.AuthorizeUser(organizationHandler.AcceptInvitation()))
	router.Handle("DELETE /api/organizations/leave", authMiddleware.AuthorizeUserRole(organizationHandler.Leave(), staffAccessLv))
	router.Handle("GET /api/organizations/members", authMiddleware.AuthorizeUserRole(organizationHandler.FetchMembers(), staffAccessLv))
	router.Handle("DELETE /api/organizations/members/{id}", authMiddleware.AuthorizeUserRole(organizationHandler.RemoveMember(), adminAccessLv))
	router.Handle("PUT /api/organizations/members/{id}", authMiddleware.AuthorizeUserRole(organizationHandler.ChangeMemberRole(), adminAccessLv))
	router.Handle("PUT /api/organizations/dashboard", authMiddleware.AuthorizeUserRole(organizationHandler.UpdateDashboardSettings(), supervisorAccessLv))

	router.Handle("GET /api/inventory", authMiddleware.AuthorizeUserRole(inventoryHandler.FetchItems(), staffAccessLv))
	router.Handle("POST /api/inventory", authMiddleware.AuthorizeUserRole(inventoryHandler.AddInventoryItem(), staffAccessLv))
	return router
}
