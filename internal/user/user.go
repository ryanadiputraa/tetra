package user

import (
	"context"
	"database/sql"
	"time"
)

type Status string

type User struct {
	ID        int            `json:"id" gorm:"primaryKey;autoIncrement"`
	Email     string         `json:"email" gorm:"type:varchar(100);unique;notNull"`
	Password  sql.NullString `json:"-" gorm:"type:varchar(50)"`
	Fullname  string         `json:"fullname" gorm:"type:varchar(100);notNull"`
	Role      sql.NullString `json:"role" gorm:"type:varchar(10)"`
	CompanyID sql.NullInt64  `json:"status"`
	CreatedAt time.Time      `json:"created_at" gorm:"notNull"`
}

type UserService interface {
	Login(ctx context.Context, email, password string) (User, error)
}

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (User, error)
}
