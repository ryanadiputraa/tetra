package service

import (
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
