package auth

import (
	"context"
	"time"

	"github.com/ryanadiputraa/inventra/internal/user"
)

type Role string

const (
	Admin          Role = "admin"
	Supervisor     Role = "supervisor"
	Staff          Role = "staff"
	JWTExpiresTime      = time.Hour * 24
)

var AccessLevel = map[Role]int{
	Admin:      3,
	Supervisor: 2,
	Staff:      1,
}

type JWT struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   string `json:"expires_at"`
}

type AuthContext struct {
	UserID int
	context.Context
}

type LoginPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Fullname string `json:"fullname" validate:"required"`
}

type AuthService interface {
	Login(ctx context.Context, email, password string) (user.User, error)
	Register(ctx context.Context, payload RegisterPayload) (user.User, error)
	GenerateJWT(ctx context.Context, userID int) (JWT, error)
}
