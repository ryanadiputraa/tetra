package handler

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ryanadiputraa/inventra/config"
	"github.com/ryanadiputraa/inventra/internal/auth"
	"github.com/ryanadiputraa/inventra/internal/errors"
	"github.com/ryanadiputraa/inventra/internal/user"
	"github.com/ryanadiputraa/inventra/pkg/oauth"
)

type handler struct {
	logger      *slog.Logger
	config      config.Config
	googleOauth oauth.GoogleOauth
	userService user.UserService
	authService auth.AuthService
}

func New(logger *slog.Logger, config config.Config, googleOauth oauth.GoogleOauth, userService user.UserService, authService auth.AuthService) *handler {
	return &handler{
		logger:      logger,
		config:      config,
		googleOauth: googleOauth,
		userService: userService,
		authService: authService,
	}
}

func (h *handler) GoogleSignin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := h.googleOauth.GetSigninURL()
		http.Redirect(w, r, url, http.StatusFound)
	}
}

func (h *handler) GoogleCallback() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		state := r.URL.Query().Get("state")

		if state != h.config.OauthState || code == "" {
			h.logger.Warn("Invalid oauth callback state", "provider", "Google")
			h.redirectWithError(w, r, errors.Unauthorized)
			return
		}

		u, err := h.googleOauth.ExchangeCodeWithUserInfo(r.Context(), code)
		if err != nil {
			h.logger.Warn(
				"Fail to exchange user info",
				"error", err.Error(),
				"provider", "Google",
			)
			h.redirectWithError(w, r, errors.Unauthorized)
			return
		}

		newUser, err := h.userService.CreateOrUpdate(r.Context(), u.FirstName+" "+u.LastName, u.Email, "")
		if err != nil {
			h.redirectWithError(w, r, errors.ServerError)
			return
		}

		jwt, err := h.authService.GenerateJWT(r.Context(), newUser.ID, newUser.OrganizationID)
		if err != nil {
			h.redirectWithError(w, r, errors.ServerError)
			return
		}

		http.Redirect(w, r, h.config.FrontendURL+fmt.Sprintf("/auth?access_token=%v&expires_at=%v", jwt.AccessToken, jwt.ExpiresAt), http.StatusFound)
	}
}

func (h *handler) redirectWithError(w http.ResponseWriter, r *http.Request, err string) {
	http.Redirect(w, r, h.config.FrontendURL+"/auth?err="+err, http.StatusFound)
}
