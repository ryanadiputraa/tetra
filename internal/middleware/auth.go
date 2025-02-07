package middleware

import (
	"net/http"
	"strings"

	"github.com/ryanadiputraa/inventra/internal/auth"
	"github.com/ryanadiputraa/inventra/internal/errors"
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

func (m *Middleware) ParseJWTToken(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if len(authorization) == 0 {
			m.writer.WriteErrorResponse(w, http.StatusUnauthorized, errors.MissingAuthHeader)
			return
		}

		tokens := strings.Split(authorization, " ")
		if len(tokens) < 2 || tokens[0] != "Bearer" {
			m.writer.WriteErrorResponse(w, http.StatusUnauthorized, errors.InvalidAuthHeader)
			return
		}

		claims, err := m.jwt.ParseJWTClaims(tokens[1])
		if err != nil {
			m.writer.WriteErrorResponse(w, http.StatusForbidden, err.Error())
			return
		}

		ac := &auth.AuthContext{
			UserID:  claims.UserID,
			Context: r.Context(),
		}
		rc := r.WithContext(ac)
		h.ServeHTTP(w, rc)
	})
}
