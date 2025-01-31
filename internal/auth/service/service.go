package service

import (
	"context"

	"github.com/ryanadiputraa/inventra/internal/auth"
	"github.com/ryanadiputraa/inventra/pkg/jwt"
	"github.com/ryanadiputraa/inventra/pkg/logger"
)

type service struct {
	log logger.Logger
	jwt jwt.JWT
}

func New(log logger.Logger, jwt jwt.JWT) auth.AuthService {
	return &service{
		log: log,
		jwt: jwt,
	}
}

func (s *service) GenerateJWT(ctx context.Context, userID int) (tokens auth.JWT, err error) {
	tokens, err = s.jwt.GenereateJWTWithClaims(userID)
	if err != nil {
		s.log.Error(err, "fail to generate jwt. err: ")
	}
	return
}
