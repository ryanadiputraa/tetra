package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	JWTExpiresTime = time.Hour * 24
)

type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

type JWT struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   string `json:"expires_at"`
}

type JWTService interface {
	GenerateJWTWithClaims(userID int) (JWT, error)
	ParseJWTClaims(accessToken string) (*Claims, error)
}

type service struct {
	secretKey string
}

func NewJWT(secretKey string) JWTService {
	return &service{
		secretKey: secretKey,
	}
}

func (s *service) GenerateJWTWithClaims(userID int) (tokens JWT, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		userID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(JWTExpiresTime)),
		},
	})
	accessToken, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return
	}

	tokens = JWT{
		AccessToken: accessToken,
		ExpiresAt:   time.Now().Add(JWTExpiresTime).Format(time.RFC3339Nano),
	}
	return
}

func (s *service) ParseJWTClaims(accessToken string) (claims *Claims, err error) {
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signin method: %v", t.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})
	if err != nil {
		return
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, fmt.Errorf("fail to cast jwt claims")
	}
	return
}
