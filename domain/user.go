package domain

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Status string

type User struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Email     string    `json:"email" gorm:"type:varchar(100);unique;notNull"`
	Password  *string   `json:"-" gorm:"type:varchar(100)"`
	Fullname  string    `json:"fullname" gorm:"type:varchar(100);notNull"`
	CreatedAt time.Time `json:"created_at" gorm:"notNull"`
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

func NewUser(fullname, email, password string) (user User, err error) {
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
