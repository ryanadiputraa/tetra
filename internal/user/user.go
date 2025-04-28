package user

import (
	"context"

	"github.com/ryanadiputraa/tetra/domain"
)

type ChangePassowrdPayload struct {
	Password string `json:"password" validate:"required,min=8"`
}

type UserService interface {
	CreateOrUpdate(ctx context.Context, fullname, email, password string) (domain.User, error)
	GetByID(ctx context.Context, userID int) (domain.UserData, error)
	ChangePassword(ctx context.Context, userID int, password string) error
}

type UserRepository interface {
	Save(ctx context.Context, user domain.User) (domain.User, error)
	SaveOrUpdate(ctx context.Context, user domain.User) (domain.User, error)
	FindByID(ctx context.Context, userID int) (domain.UserData, error)
	FindByEmail(ctx context.Context, email string) (domain.UserData, error)
	UpdatePassword(ctx context.Context, userID int, password string) error
}
