package repository

import (
	"context"

	"github.com/ryanadiputraa/inventra/internal/errors"
	"github.com/ryanadiputraa/inventra/internal/user"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) Save(ctx context.Context, user user.User) (result user.User, err error) {
	err = r.db.Create(&user).Error
	if err == gorm.ErrDuplicatedKey {
		err = errors.NewServiceErr(errors.BadRequest, errors.EmailTaken)
		return
	}
	result = user
	return
}

func (r *repository) SaveOrUpdate(ctx context.Context, user user.User) (result user.User, err error) {
	err = r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "email"}},
		DoUpdates: clause.AssignmentColumns([]string{"fullname"}),
	}).Create(&user).Error
	result = user
	return
}

func (r *repository) FindByID(ctx context.Context, userID int) (result user.UserData, err error) {
	err = r.db.Model(&user.User{}).
		Select("users.id", "users.email", "users.fullname", "users.created_at, members.organization_id").
		Joins("LEFT JOIN members ON members.user_id = users.id").
		Scan(&result).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = errors.NewServiceErr(errors.BadRequest, errors.RecordNotFound)
			return
		}
	}
	return
}

func (r *repository) FindByEmail(ctx context.Context, email string) (result user.User, err error) {
	err = r.db.Model(&user.User{}).
		Select("users.id", "users.email", "users.fullname", "users.created_at").
		Scan(&result).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = errors.NewServiceErr(errors.BadRequest, errors.RecordNotFound)
			return
		}
	}
	return
}

func (r *repository) UpdatePassword(ctx context.Context, userID int, password string) error {
	return r.db.Model(&user.User{}).Where("id = ?", userID).Update("password", password).Error
}
