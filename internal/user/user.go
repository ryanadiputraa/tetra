package user

import (
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Status string

type User struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	Email     string    `gorm:"type:varchar(100);unique;notNull"`
	Password  *string   `gorm:"type:varchar(100)"`
	Fullname  string    `gorm:"type:varchar(100);notNull"`
	CreatedAt time.Time `gorm:"notNull"`
}

type UserData struct {
	ID             int       `json:"id"`
	Email          string    `json:"email"`
	Password       *string   `json:"-"`
	Fullname       string    `json:"fullname"`
	CreatedAt      time.Time `json:"created_at"`
	OrganizationID *int      `json:"organization_id"`
	MemberID       *int      `json:"member_id"`
	Role           string    `json:"role"`
}

type ChangePassowrdPayload struct {
	Password string `json:"password" validate:"required,min=8"`
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
	GetByID(ctx context.Context, userID int) (UserData, error)
	ChangePassword(ctx context.Context, userID int, password string) error
}

type UserRepository interface {
	Save(ctx context.Context, user User) (User, error)
	SaveOrUpdate(ctx context.Context, user User) (User, error)
	FindByID(ctx context.Context, userID int) (UserData, error)
	FindByEmail(ctx context.Context, email string) (UserData, error)
	UpdatePassword(ctx context.Context, userID int, password string) error
}
