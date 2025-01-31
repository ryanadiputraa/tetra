package service

import (
	"context"

	"github.com/ryanadiputraa/inventra/internal/errors"
	"github.com/ryanadiputraa/inventra/internal/user"
	"github.com/ryanadiputraa/inventra/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	log        logger.Logger
	repository user.UserRepository
}

func New(log logger.Logger, repository user.UserRepository) user.UserService {
	return &service{
		log:        log,
		repository: repository,
	}
}

func (s *service) Login(ctx context.Context, email, password string) (user user.User, err error) {
	user, err = s.repository.FindByEmail(ctx, email)
	if err != nil {
		s.log.Error("Fail to find user by email. Err: ", err.Error())
		return
	}

	// Handle user only sign with social (password is still empty)
	if user.Password.String == "" {
		err = errors.NewServiceErr(errors.Unauthorized, errors.Unauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(password))
	if err != nil {
		err = errors.NewServiceErr(errors.Unauthorized, errors.Unauthorized)
		return
	}
	return
}
