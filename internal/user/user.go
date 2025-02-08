package user

import (
	"context"
	"time"

	"github.com/ryanadiputraa/inventra/internal/organization"
	"golang.org/x/crypto/bcrypt"
)

type Status string

type User struct {
	ID             int                       `json:"id" gorm:"primaryKey;autoIncrement"`
	Email          string                    `json:"email" gorm:"type:varchar(100);unique;notNull"`
	Password       *string                   `json:"-" gorm:"type:varchar(100)"`
	Fullname       string                    `json:"fullname" gorm:"type:varchar(100);notNull"`
	Role           *string                   `json:"role" gorm:"type:varchar(10)"`
	OrganizationID *int                      `json:"organization_id"`
	Organization   organization.Organization `json:"-" gorm:"constraint:OnDelete:SET NULL;"`
	CreatedAt      time.Time                 `json:"created_at" gorm:"notNull"`
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
		pass := string(hashed)
		user.Password = &pass
	}
	return
}

type UserService interface {
	CreateOrUpdate(ctx context.Context, fullname, email, password string) (User, error)
	GetByID(ctx context.Context, userID int) (User, error)
}

type UserRepository interface {
	Save(ctx context.Context, user User) (User, error)
	SaveOrUpdate(ctx context.Context, user User) (User, error)
	FindByID(ctx context.Context, userID int) (User, error)
	FindByEmail(ctx context.Context, email string) (User, error)
}
