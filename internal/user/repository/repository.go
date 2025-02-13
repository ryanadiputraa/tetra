package repository

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/ryanadiputraa/inventra/internal/errors"
	"github.com/ryanadiputraa/inventra/internal/user"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	redisKeyUserData = "users:"
)

type repository struct {
	db  *gorm.DB
	rdb *redis.Client
}

func New(db *gorm.DB, rdb *redis.Client) user.UserRepository {
	return &repository{
		db:  db,
		rdb: rdb,
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
	id := strconv.Itoa(userID)
	cache, err := r.rdb.Get(ctx, redisKeyUserData+id).Result()
	if err == redis.Nil {
		err = r.db.Table("users").
			Select("users.id", "users.email", "users.password", "users.fullname", "users.created_at, members.organization_id, members.role").
			Joins("LEFT JOIN members ON members.user_id = users.id").
			Where("users.id = ?", userID).
			Scan(&result).Error
		if err == gorm.ErrRecordNotFound {
			err = errors.NewServiceErr(errors.BadRequest, errors.RecordNotFound)
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
		err = r.rdb.Set(ctx, redisKeyUserData+id, val, time.Hour*6).Err()
		return
	} else if err != nil {
		return
	}

	err = json.Unmarshal([]byte(cache), &result)
	return
}

func (r *repository) FindByEmail(ctx context.Context, email string) (result user.User, err error) {
	err = r.db.Model(&user.User{}).
		Select("users.id", "users.email", "users.password", "users.fullname", "users.created_at").
		Where("users.email = ?", email).
		Scan(&result).Error
	if err != gorm.ErrRecordNotFound {
		err = errors.NewServiceErr(errors.BadRequest, errors.RecordNotFound)
		return
	}
	return
}

func (r *repository) UpdatePassword(ctx context.Context, userID int, password string) error {
	return r.db.Model(&user.User{}).Where("id = ?", userID).Update("password", password).Error
}
