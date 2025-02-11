package service

import (
	"context"
	"errors"
	"log/slog"

	"github.com/ryanadiputraa/inventra/internal/auth"
	serviceErr "github.com/ryanadiputraa/inventra/internal/errors"
	"github.com/ryanadiputraa/inventra/internal/user"
	"github.com/ryanadiputraa/inventra/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	logger         *slog.Logger
	jwt            jwt.JWT
	userRepository user.UserRepository
}

func New(logger *slog.Logger, jwt jwt.JWT, userRepository user.UserRepository) auth.AuthService {
	return &service{
		logger:         logger,
		jwt:            jwt,
		userRepository: userRepository,
	}
}

func (s *service) Login(ctx context.Context, email, password string) (user user.User, err error) {
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
	if user.Password != nil {
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

func (s *service) Register(ctx context.Context, payload auth.RegisterPayload) (res user.User, err error) {
	u, err := user.New(payload.Fullname, payload.Email, payload.Password)
	if err != nil {
		s.logger.Error(
			"Fail to create new user",
			"error", err.Error(),
			"email", payload.Email,
		)
		return
	}

	res, err = s.userRepository.Save(ctx, u)
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
		"id", res.ID,
		"fullname", res.Fullname,
		"email", res.Email,
		"orgnaization_id", res.OrganizationID,
		"role", res.Role,
		"created_at", res.CreatedAt,
	)
	return
}

func (s *service) GenerateJWT(ctx context.Context, userID int) (tokens auth.JWT, err error) {
	tokens, err = s.jwt.GenereateJWTWithClaims(userID)
	if err != nil {
		s.logger.Error("Fail to generate jwt", "error", err.Error())
	}
	return
}
