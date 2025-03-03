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
	"gorm.io/gorm/clause"
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
		err = r.db.InnerJoins("Owner").First(&result, "organizations.id = ?", organizationID).Error
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

func (r *repository) Delete(ctx context.Context, organizationID, userID int) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var userIDs []int
		err := tx.Model(&organization.Member{}).
			Where("organization_id = ?", organizationID).
			Pluck("user_id", &userIDs).Error
		if err != nil {
			return err
		}

		err = tx.Where("organization_id = ?", organizationID).Delete(&organization.Member{}).Error
		if err != nil {
			return err
		}

		err = tx.Where("owner_id = ? AND id = ?", userID, organizationID).Delete(&organization.Organization{}).Error
		if err != nil {
			return err
		}

		err = r.cache.Del(ctx, "organizations:"+strconv.Itoa(organizationID)).Err()
		if err != nil {
			return err
		}

		if len(userIDs) > 0 {
			pipe := r.cache.Pipeline()
			for _, id := range userIDs {
				pipe.Del(ctx, "users:"+strconv.Itoa(id))
			}
			if _, err = pipe.Exec(ctx); err != nil {
				return err
			}
		}

		return nil
	}, nil)
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

func (r *repository) DeleteMember(ctx context.Context, organizationID, memberID int) (err error) {
	var member organization.Member
	err = r.db.Clauses(clause.Returning{Columns: []clause.Column{{Name: "user_id"}}}).
		Where("organization_id = ? AND id = ?", organizationID, memberID).
		Delete(&member).Error
	if err != nil {
		return
	}

	userID := strconv.Itoa(member.UserID)
	return r.cache.Del(ctx, "users:"+userID).Err()
}

func (r *repository) UpdateMemberRole(ctx context.Context, organizationID, memberID int, role string) (err error) {
	var member organization.Member
	err = r.db.Model(&member).Clauses(clause.Returning{Columns: []clause.Column{{Name: "user_id"}}}).
		Where("organization_id = ? AND id = ? AND role <> 'admin'", organizationID, memberID).
		Update("role", role).Error
	if err != nil {
		return
	}

	userID := strconv.Itoa(member.UserID)
	return r.cache.Del(ctx, "users:"+userID).Err()
}
