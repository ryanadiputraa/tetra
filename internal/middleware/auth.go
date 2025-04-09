package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/ryanadiputraa/inventra/domain"
	"github.com/ryanadiputraa/inventra/internal/auth"
	serviceError "github.com/ryanadiputraa/inventra/internal/errors"
	"github.com/ryanadiputraa/inventra/internal/organization"
	"github.com/ryanadiputraa/inventra/internal/user"
	"github.com/ryanadiputraa/inventra/pkg/jwt"
	"github.com/ryanadiputraa/inventra/pkg/writer"
)

type Middleware struct {
	writer              writer.HTTPWriter
	jwt                 jwt.JWTService
	userService         user.UserService
	organizationService organization.OrganizationService
}

func NewAuthMiddleware(writer writer.HTTPWriter, jwt jwt.JWTService, userService user.UserService, organizationService organization.OrganizationService) *Middleware {
	return &Middleware{
		writer:              writer,
		jwt:                 jwt,
		userService:         userService,
		organizationService: organizationService,
	}
}

func (m *Middleware) AuthorizeUser(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		claims, err := m.parseJWT(authorization)
		if err != nil {
			m.writer.WriteErrorResponse(w, http.StatusUnauthorized, err.Error())
			return
		}

		ac := &auth.AppContext{
			UserID:  claims.UserID,
			Context: r.Context(),
		}
		rc := r.WithContext(ac)
		h.ServeHTTP(w, rc)
	})
}

func (m *Middleware) AuthorizeUserRole(h http.Handler, level int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		claims, err := m.parseJWT(authorization)
		if err != nil {
			m.writer.WriteErrorResponse(w, http.StatusUnauthorized, err.Error())
			return
		}

		user, err := m.userService.GetByID(r.Context(), claims.UserID)
		if err != nil {
			if sErr, ok := err.(*serviceError.Error); ok {
				m.writer.WriteErrorResponse(w, serviceError.HttpErrMap[sErr.ErrCode], sErr.Error())
				return
			} else {
				m.writer.WriteErrorResponse(w, http.StatusInternalServerError, serviceError.ServerError)
				return
			}
		}

		if user.OrganizationID == nil {
			m.writer.WriteErrorResponse(w, http.StatusUnauthorized, serviceError.Unauthorized)
			return
		}

		isValid, err := m.organizationService.IsSubscriptionValid(r.Context(), *user.OrganizationID)
		if err != nil {
			if sErr, ok := err.(*serviceError.Error); ok {
				m.writer.WriteErrorResponse(w, serviceError.HttpErrMap[sErr.ErrCode], sErr.Error())
				return
			} else {
				m.writer.WriteErrorResponse(w, http.StatusInternalServerError, serviceError.ServerError)
				return
			}
		}

		if !isValid {
			m.writer.WriteErrorResponse(w, http.StatusUnauthorized, serviceError.SubscriptionEnd)
			return
		}

		if domain.AccessLevel[domain.Role(user.Role)] < level {
			m.writer.WriteErrorResponse(w, http.StatusUnauthorized, serviceError.Unauthorized)
			return
		}

		ac := &auth.AppContext{
			UserID:         user.ID,
			OrganizationID: user.OrganizationID,
			MemberID:       user.MemberID,
			Context:        r.Context(),
		}
		rc := r.WithContext(ac)
		h.ServeHTTP(w, rc)
	})
}

func (m *Middleware) parseJWT(authorization string) (claims *jwt.Claims, err error) {
	if len(authorization) == 0 {
		err = errors.New(serviceError.MissingAuthHeader)
		return
	}

	tokens := strings.Split(authorization, " ")
	if len(tokens) < 2 || tokens[0] != "Bearer" {
		err = errors.New(serviceError.InvalidAuthHeader)
		return
	}

	claims, err = m.jwt.ParseJWTClaims(tokens[1])
	if err != nil {
		return
	}
	return
}
