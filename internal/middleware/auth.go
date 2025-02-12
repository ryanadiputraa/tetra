package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/ryanadiputraa/inventra/internal/auth"
	serviceError "github.com/ryanadiputraa/inventra/internal/errors"
	"github.com/ryanadiputraa/inventra/pkg/jwt"
	"github.com/ryanadiputraa/inventra/pkg/writer"
)

type Middleware struct {
	writer    writer.HTTPWriter
	jwt       jwt.JWT
	jwtSecret string
}

func NewAuthMiddleware(jwt jwt.JWT, jwtSecret string) *Middleware {
	return &Middleware{
		writer:    writer.NewHTTPWriter(),
		jwt:       jwt,
		jwtSecret: jwtSecret,
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

func (m *Middleware) AuthorizeUserRole(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		claims, err := m.parseJWT(authorization)
		if err != nil {
			m.writer.WriteErrorResponse(w, http.StatusUnauthorized, err.Error())
			return
		}

		// fetch organiztion id and role
		// check organiztion subscription date
		// validate access level

		ac := &auth.AppContext{
			UserID:  claims.UserID,
			Context: r.Context(),
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
