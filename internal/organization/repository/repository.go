package repository

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/ryanadiputraa/inventra/internal/auth"
	"github.com/ryanadiputraa/inventra/internal/errors"
	"github.com/ryanadiputraa/inventra/internal/organization"
	"gorm.io/gorm"
)

const (
	redisKeyOrganizationData = "organization:"
)

type repository struct {
	db  *gorm.DB
	rdb *redis.Client
}

func New(db *gorm.DB, rdb *redis.Client) organization.OrganizationRepository {
	return &repository{
		db:  db,
		rdb: rdb,
	}
}

func (r *repository) Save(ctx context.Context, data organization.Organization) (result organization.Organization, err error) {
	err = r.db.Transaction(func(tx *gorm.DB) error {
		err = r.db.Create(&data).Error
		if err == gorm.ErrDuplicatedKey {
			err = errors.NewServiceErr(errors.BadRequest, errors.OrganizationAlreadyExists)
			return err
		}
		if err != nil {
			return err
		}

		owner := organization.NewMember(data.ID, data.OwnerID, string(auth.Admin))
		if err = r.db.Create(&owner).Error; err != nil {
			return err
		}

		result = data
		return nil
	})
	return
}

func (r *repository) FindByID(ctx context.Context, organizationID int) (result organization.Organization, err error) {
	id := strconv.Itoa(organizationID)
	cache, err := r.rdb.Get(ctx, redisKeyOrganizationData+id).Result()
	if err == redis.Nil {
		err = r.db.Table("organizations").
			Select("organizations.id, organizations.owner_id, organizations.name, organizations.created_at, organizations.subscription_end_at").
			Where("organizations.id = ?", organizationID).
			Scan(&result).Error
		if err != nil {
			return
		}

		var val []byte
		val, err = json.Marshal(result)
		if err != nil {
			return
		}
		err = r.rdb.Set(ctx, redisKeyOrganizationData+id, val, time.Hour*6).Err()
		return
	}
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(cache), &result)
	return
}

func (r *repository) FetchMembers(ctx context.Context, organizationID int) (result []organization.MemberData, err error) {
	result = make([]organization.MemberData, 0)
	err = r.db.Table("members").
		Select("members.id, members.user_id, users.fullname, users.email, members.role").
		Joins("LEFT JOIN users ON users.id = members.user_id").
		Where("members.organization_id = ?", organizationID).
		Order("users.fullname ASC").
		Scan(&result).Error
	return
}
