package service

import (
	"context"

	"github.com/ryanadiputraa/inventra/internal/auth"
	"github.com/ryanadiputraa/inventra/internal/errors"
	"github.com/ryanadiputraa/inventra/internal/user"
	"github.com/ryanadiputraa/inventra/pkg/jwt"
	"github.com/ryanadiputraa/inventra/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	log            logger.Logger
	jwt            jwt.JWT
	userRepository user.UserRepository
}

func New(log logger.Logger, jwt jwt.JWT, userRepository user.UserRepository) auth.AuthService {
	return &service{
		log:            log,
		jwt:            jwt,
		userRepository: userRepository,
	}
}

func (s *service) Login(ctx context.Context, email, password string) (user user.User, err error) {
	user, err = s.userRepository.FindByEmail(ctx, email)
	if err != nil {
		s.log.Error("Fail to find user by email. Err: ", err.Error())
		return
	}

	// Handle user only signin with social (password is still empty)
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

func (s *service) Register(ctx context.Context, payload auth.RegisterPayload) (res user.User, err error) {
	u, err := user.New(payload.Fullname, payload.Email, payload.Password)
	if err != nil {
		s.log.Error("Fail to create new user. Err: ", err.Error())
		return
	}

	res, err = s.userRepository.Save(ctx, u)
	if err != nil {
		s.log.Error("Fail to save user. Err: ", err.Error())
		return
	}
	return
}

func (s *service) GenerateJWT(ctx context.Context, userID int) (tokens auth.JWT, err error) {
	tokens, err = s.jwt.GenereateJWTWithClaims(userID)
	if err != nil {
		s.log.Error(err, "fail to generate jwt. err: ")
	}
	return
}
