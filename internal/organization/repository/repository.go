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

func (r *repository) Save(ctx context.Context, data organization.Organization) (res organization.Organization, err error) {
	err = r.db.Transaction(func(tx *gorm.DB) error {
		if err = r.db.Create(&data).Error; err != nil {
			if err == gorm.ErrDuplicatedKey {
				err = errors.NewServiceErr(errors.BadRequest, errors.OrganizationAlreadyExists)
				return err
			}
			return err
		}

		owner := organization.NewMember(data.ID, data.OwnerID, string(auth.Admin))
		if err = r.db.Create(&owner).Error; err != nil {
			return err
		}

		res = data
		return nil
	})
	return
}
