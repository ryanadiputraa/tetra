package repository

import (
	"context"

	"github.com/ryanadiputraa/inventra/internal/errors"
	"github.com/ryanadiputraa/inventra/internal/user"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) FindByEmail(ctx context.Context, email string) (user user.User, err error) {
	err = r.db.First(&user, "email = ?", email).Error
	if err == gorm.ErrRecordNotFound {
		err = errors.NewServiceErr(errors.Unauthorized, errors.Unauthorized)
		return
	}
	return
}
