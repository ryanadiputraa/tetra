package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ryanadiputraa/inventra/internal/auth"
)

type Claims struct {
	UserID         int  `json:"user_id"`
	OrganizationID *int `json:"organization_id"`
	jwt.RegisteredClaims
}

type JWT interface {
	GenereateJWTWithClaims(userID int, organizationID *int) (auth.JWT, error)
	ParseJWTClaims(accessToken string) (*Claims, error)
}

type service struct {
	secretKey string
}

func NewJWT(secretKey string) JWT {
	return &service{
		secretKey: secretKey,
	}
}

func (s *service) GenereateJWTWithClaims(userID int, organizationID *int) (tokens auth.JWT, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		userID,
		organizationID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(auth.JWTExpiresTime)),
		},
	})
	accessToken, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return
	}

	tokens = auth.JWT{
		AccessToken: accessToken,
		ExpiresAt:   time.Now().Add(auth.JWTExpiresTime).Format(time.RFC3339Nano),
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
