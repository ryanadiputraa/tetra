package repository

import (
	"context"
	"os/user"

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

func (r *repository) Save(ctx context.Context, organization organization.Organization) (res organization.Organization, err error) {
	err = r.db.Transaction(func(tx *gorm.DB) error {
		if err = r.db.Create(&organization).Error; err != nil {
			if err == gorm.ErrDuplicatedKey {
				err = errors.NewServiceErr(errors.BadRequest, errors.OrganizationAlreadyExists)
				return err
			}
			return err
		}
		res = organization
		if err = r.db.Model(&user.User{}).
			Where("id = ?", res.OwnerID).
			Update("organization_id", res.ID).
			Update("role", auth.Admin).Error; err != nil {
			return err
		}
		return nil
	})
	return
}
