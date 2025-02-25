package auth

import (
	"context"
	"time"

	"github.com/ryanadiputraa/inventra/internal/user"
)

type Role string

const (
	JWTExpiresTime = time.Hour * 24

	// Role
	Admin      Role = "admin"
	Supervisor Role = "supervisor"
	Staff      Role = "staff"
)

var AccessLevel = map[Role]int{
	Staff:      1,
	Supervisor: 2,
	Admin:      3,
}

type JWT struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   string `json:"expires_at"`
}

type AppContext struct {
	UserID         int
	OrganizationID *int
	Role           Role
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

func IsValidRole(r Role) bool {
	switch r {
	case Admin, Supervisor, Staff:
		return true
	default:
		return false
	}
}

type AuthService interface {
	Login(ctx context.Context, email, password string) (user.UserData, error)
	Register(ctx context.Context, payload RegisterPayload) (user.User, error)
	GenerateJWT(ctx context.Context, userID int) (JWT, error)
}
