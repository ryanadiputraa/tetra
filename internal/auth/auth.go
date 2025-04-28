package auth

import (
	"context"

	"github.com/ryanadiputraa/tetra/domain"
)

type AppContext struct {
	UserID         int
	OrganizationID *int
	MemberID       *int
	Role           domain.Role
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
	Login(ctx context.Context, email, password string) (domain.UserData, error)
	Register(ctx context.Context, payload RegisterPayload) (domain.User, error)
}
