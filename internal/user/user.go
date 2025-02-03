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

func New(fullname, email, password string) (user User, err error) {
	user = User{
		Email:     email,
		Fullname:  fullname,
		CreatedAt: time.Now().UTC(),
	}

	if len(password) > 0 {
		var hashed []byte
		hashed, err = bcrypt.GenerateFromPassword([]byte(password), 10)
		if err != nil {
			return
		}

		user.Password = sql.NullString{
			Valid:  true,
			String: string(hashed),
		}
	}
	return
}

type UserService interface {
	CreateOrUpdate(ctx context.Context, fullname, email, password string) (User, error)
}

type UserRepository interface {
	Save(ctx context.Context, user User) (User, error)
	SaveOrUpdate(ctx context.Context, user User) (User, error)
	FindByEmail(ctx context.Context, email string) (User, error)
}
