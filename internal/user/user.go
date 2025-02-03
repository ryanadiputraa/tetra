package user

import (
	"context"
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Status string

type User struct {
	ID        int            `json:"id" gorm:"primaryKey;autoIncrement"`
	Email     string         `json:"email" gorm:"type:varchar(100);unique;notNull"`
	Password  sql.NullString `json:"-" gorm:"type:varchar(100)"`
	Fullname  string         `json:"fullname" gorm:"type:varchar(100);notNull"`
	Role      sql.NullString `json:"role" gorm:"type:varchar(10)"`
	CompanyID sql.NullInt64  `json:"status"`
	CreatedAt time.Time      `json:"created_at" gorm:"notNull"`
}

func NewUser(email, password, fullname string) (user User, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return
	}

	user = User{
		Email: email,
		Password: sql.NullString{
			Valid:  true,
			String: string(hashedPassword),
		},
		Fullname:  fullname,
		CreatedAt: time.Now().UTC(),
	}
	return
}

type UserService interface{}

type UserRepository interface {
	Save(ctx context.Context, user User) (User, error)
	FindByEmail(ctx context.Context, email string) (User, error)
}
