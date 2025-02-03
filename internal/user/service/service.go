package service

import (
	"context"

	"github.com/ryanadiputraa/inventra/internal/user"
	"github.com/ryanadiputraa/inventra/pkg/logger"
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

func (s *service) CreateOrUpdate(ctx context.Context, fullname, email, password string) (res user.User, err error) {
	u, err := user.New(fullname, email, password)
	if err != nil {
		s.log.Error("Fail to create new user. Err: ", err.Error())
		return
	}

	res, err = s.repository.SaveOrUpdate(ctx, u)
	if err != nil {
		s.log.Error("Fail to save or update. Err: ", err.Error())
		return
	}
	return
}
