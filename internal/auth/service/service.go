package service

import (
	"context"
	"errors"
	"log/slog"

	"github.com/ryanadiputraa/inventra/domain"
	"github.com/ryanadiputraa/inventra/internal/auth"
	serviceErr "github.com/ryanadiputraa/inventra/internal/errors"
	"github.com/ryanadiputraa/inventra/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	logger         *slog.Logger
	userRepository user.UserRepository
}

func New(logger *slog.Logger, userRepository user.UserRepository) auth.AuthService {
	return &service{
		logger:         logger,
		userRepository: userRepository,
	}
}

func (s *service) Login(ctx context.Context, email, password string) (user domain.UserData, err error) {
	user, err = s.userRepository.FindByEmail(ctx, email)
	if err != nil {
		if !errors.As(err, new(*serviceErr.Error)) {
			s.logger.Error(
				"Fail to fetch user data",
				"error", err.Error(),
				"email", email,
			)
		}
		return
	}

	// Handle user only signin with social (password is still empty)
	if user.Password == nil {
		err = serviceErr.NewServiceErr(serviceErr.Unauthorized, serviceErr.Unauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(password))
	if err != nil {
		err = serviceErr.NewServiceErr(serviceErr.Unauthorized, serviceErr.Unauthorized)
		return
	}
	return
}

func (s *service) Register(ctx context.Context, payload auth.RegisterPayload) (result domain.User, err error) {
	u, err := domain.NewUser(payload.Fullname, payload.Email, payload.Password)
	if err != nil {
		s.logger.Error(
			"Fail to create new user",
			"error", err.Error(),
			"email", payload.Email,
		)
		return
	}

	result, err = s.userRepository.Save(ctx, u)
	if err != nil {
		if !errors.As(err, new(*serviceErr.Error)) {
			s.logger.Error(
				"Fail to register user",
				"error", err.Error(),
				"fullname", payload.Fullname,
				"email", payload.Email,
			)
		}
		return
	}

	s.logger.Info(
		"New user registered",
		"id", result.ID,
		"fullname", result.Fullname,
		"email", result.Email,
		"created_at", result.CreatedAt,
	)
	return
}
