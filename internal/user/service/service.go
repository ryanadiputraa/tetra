package service

import (
	"context"
	"errors"
	"log/slog"

	"github.com/ryanadiputraa/inventra/domain"
	serviceError "github.com/ryanadiputraa/inventra/internal/errors"
	"github.com/ryanadiputraa/inventra/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	logger     *slog.Logger
	repository user.UserRepository
}

func New(logger *slog.Logger, repository user.UserRepository) user.UserService {
	return &service{
		logger:     logger,
		repository: repository,
	}
}

func (s *service) CreateOrUpdate(ctx context.Context, fullname, email, password string) (result domain.User, err error) {
	u, err := domain.NewUser(fullname, email, password)
	if err != nil {
		s.logger.Error(
			"Fail to save user data",
			"error", err.Error(),
			"fullname", fullname,
			"email", email,
		)
		return
	}

	result, err = s.repository.SaveOrUpdate(ctx, u)
	if err != nil {
		if !errors.As(err, new(*serviceError.Error)) {
			s.logger.Error(
				"Fail to save user data",
				"error", err.Error(),
				"fullname", fullname,
				"email", email,
			)
		}
		return
	}
	return
}

func (s *service) GetByID(ctx context.Context, userID int) (user domain.UserData, err error) {
	user, err = s.repository.FindByID(ctx, userID)
	if err != nil {
		if !errors.As(err, new(*serviceError.Error)) {
			s.logger.Error(
				"Fail to fetch user data",
				"error", err.Error(),
				"user_id", userID,
			)
		}
		return
	}
	return
}

func (s *service) ChangePassword(ctx context.Context, userID int, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		s.logger.Error("Fail to hash user password", "error", err.Error())
	}

	err = s.repository.UpdatePassword(ctx, userID, string(hashed))
	if err != nil {
		if !errors.As(err, new(*serviceError.Error)) {
			s.logger.Error(
				"Fail to change user password",
				"error", err.Error(),
				"user_id", userID,
			)
		}
		return err
	}
	return err
}
