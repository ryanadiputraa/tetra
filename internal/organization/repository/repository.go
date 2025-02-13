package repository

import (
	"context"

	"github.com/ryanadiputraa/inventra/internal/auth"
	"github.com/ryanadiputraa/inventra/internal/errors"
	"github.com/ryanadiputraa/inventra/internal/organization"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) organization.OrganizationRepository {
	return &repository{
		db: db,
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
	// TODO: add cache
	err = r.db.Table("organizations").
		Select("organizations.id, organizations.owner_id, organizations.name, organizations.created_at, organizations.subscription_end_at").
		Where("organizations.id = ?", organizationID).
		Scan(&result).Error
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
