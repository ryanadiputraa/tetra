package repository

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/ryanadiputraa/inventra/domain"
	serviceError "github.com/ryanadiputraa/inventra/internal/errors"
	"github.com/ryanadiputraa/inventra/internal/user"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type repository struct {
	db    *gorm.DB
	cache *redis.Client
}

func New(db *gorm.DB, rdb *redis.Client) user.UserRepository {
	return &repository{
		db:    db,
		cache: rdb,
	}
}

func (r *repository) Save(ctx context.Context, user domain.User) (result domain.User, err error) {
	err = r.db.Create(&user).Error
	if err == gorm.ErrDuplicatedKey {
		err = serviceError.NewServiceErr(serviceError.BadRequest, serviceError.EmailTaken)
		return
	}
	result = user
	return
}

func (r *repository) SaveOrUpdate(ctx context.Context, user domain.User) (result domain.User, err error) {
	err = r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "email"}},
		DoUpdates: clause.AssignmentColumns([]string{"fullname"}),
	}).Create(&user).Error
	result = user
	return
}

func (r *repository) FindByID(ctx context.Context, userID int) (result domain.UserData, err error) {
	id := strconv.Itoa(userID)
	cache, err := r.cache.Get(ctx, "users:"+id).Result()
	if err == redis.Nil {
		err = r.db.Table("users").
			Select("users.id", "users.email", "users.password", "users.fullname", "users.created_at, members.organization_id, members.id AS member_id, members.role").
			Joins("LEFT JOIN members ON members.user_id = users.id").
			Where("users.id = ?", userID).
			First(&result).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = serviceError.NewServiceErr(serviceError.BadRequest, serviceError.RecordNotFound)
			return
		}
		if err != nil {
			return
		}

		var val []byte
		val, err = json.Marshal(result)
		if err != nil {
			return
		}
		err = r.cache.Set(ctx, "users:"+id, val, time.Hour*6).Err()
		return
	} else if err != nil {
		return
	}

	err = json.Unmarshal([]byte(cache), &result)
	return
}

func (r *repository) FindByEmail(ctx context.Context, email string) (result domain.UserData, err error) {
	err = r.db.Table("users").
		Select("users.id, users.email, users.password, users.fullname, users.created_at, members.organization_id, members.id AS member_id, members.role").
		Joins("LEFT JOIN members ON members.user_id = users.id").
		Where("users.email = ?", email).
		First(&result).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = serviceError.NewServiceErr(serviceError.BadRequest, serviceError.RecordNotFound)
		return
	}
	return
}

func (r *repository) UpdatePassword(ctx context.Context, userID int, password string) error {
	return r.db.Table("users").Where("id = ?", userID).Update("password", password).Error
}
