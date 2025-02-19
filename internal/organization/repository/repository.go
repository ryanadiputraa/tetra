package repository

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/ryanadiputraa/inventra/internal/auth"
	serviceError "github.com/ryanadiputraa/inventra/internal/errors"
	"github.com/ryanadiputraa/inventra/internal/organization"
	"gorm.io/gorm"
)

type repository struct {
	db    *gorm.DB
	cache *redis.Client
}

func New(db *gorm.DB, rdb *redis.Client) organization.OrganizationRepository {
	return &repository{
		db:    db,
		cache: rdb,
	}
}

func (r *repository) Save(ctx context.Context, data organization.Organization) (result organization.Organization, err error) {
	err = r.db.Transaction(func(tx *gorm.DB) error {
		err = tx.Create(&data).Error
		if err == gorm.ErrDuplicatedKey {
			err = serviceError.NewServiceErr(serviceError.BadRequest, serviceError.OrganizationAlreadyExists)
			return err
		}
		if err != nil {
			return err
		}

		owner := organization.NewMember(data.ID, data.OwnerID, auth.Admin)
		if err = tx.Create(&owner).Error; err != nil {
			return err
		}

		result = data
		return nil
	})
	return
}

func (r *repository) FindByID(ctx context.Context, organizationID int) (result organization.Organization, err error) {
	id := strconv.Itoa(organizationID)
	cache, err := r.cache.Get(ctx, "organizations:"+id).Result()
	if err == redis.Nil {
		err = r.db.Table("organizations").
			Select("organizations.id, organizations.owner_id, organizations.name, organizations.created_at, organizations.subscription_end_at").
			Where("organizations.id = ?", organizationID).
			First(&result).Error
		if err != nil {
			return
		}

		var val []byte
		val, err = json.Marshal(result)
		if err != nil {
			return
		}
		err = r.cache.Set(ctx, "organizations:"+id, val, time.Hour*6).Err()
		return
	}
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(cache), &result)
	return
}

func (r *repository) AddMember(ctx context.Context, member organization.Member) (result organization.Member, err error) {
	var id int
	err = r.db.Table("members").
		Select("user_id").
		Where("user_id = ?", member.UserID).
		Limit(1).
		Find(&id).Error
	if err != nil {
		return
	}
	if id != 0 {
		err = serviceError.NewServiceErr(serviceError.BadRequest, serviceError.UserHasJoinedOrg)
		return
	}

	err = r.db.Create(&member).Error
	if err != nil {
		return
	}
	result = member

	userID := strconv.Itoa(result.UserID)
	err = r.cache.Del(ctx, "users:"+userID).Err()
	return
}

func (r *repository) FetchMembers(ctx context.Context, organizationID int) (result []organization.MemberData, err error) {
	result = make([]organization.MemberData, 0)
	err = r.db.Table("members").
		Select("members.id, members.user_id, users.fullname, users.email, members.role").
		Joins("LEFT JOIN users ON users.id = members.user_id").
		Where("members.organization_id = ?", organizationID).
		Order("users.fullname ASC").
		Find(&result).Error
	return
}

func (r *repository) DeleteMember(ctx context.Context, organizationID int, memberID []int) (err error) {
	err = r.db.Transaction(func(tx *gorm.DB) error {
		err = tx.Where("organization_id = ? AND id IN (?)", organizationID, memberID).Delete(&organization.Member{}).Error
		if err != nil {
			return err
		}

		for _, v := range memberID {
			id := strconv.Itoa(v)
			if err = r.cache.Del(ctx, "users:"+id).Err(); err != nil {
				return err
			}
		}

		return nil
	}, nil)
	return
}
