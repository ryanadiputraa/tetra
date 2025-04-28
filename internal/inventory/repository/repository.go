package repository

import (
	"context"
	"errors"

	"github.com/ryanadiputraa/tetra/domain"
	"github.com/ryanadiputraa/tetra/internal/inventory"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) inventory.InventoryRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) SaveItem(ctx context.Context, item domain.Item, prices []domain.ItemPrice) (result domain.Item, err error) {
	err = r.db.Transaction(func(tx *gorm.DB) error {
		err = tx.Create(&item).Error
		if err != nil {
			return err
		}

		if item.ID == 0 {
			return errors.New("failed to retrieve item id")
		}

		for i := range prices {
			prices[i].ItemID = item.ID
		}
		err = tx.Create(&prices).Error
		if err != nil {
			return err
		}

		result = item
		result.Stock = prices
		return nil
	}, nil)
	return
}

func (r *repository) FetchItems(ctx context.Context, organizationID, page, size int) (result []domain.Item, total int64, err error) {
	err = r.db.Model(&domain.Item{}).Count(&total).Error
	if err != nil {
		return
	}

	err = r.db.Preload("Stock").
		Where("organization_id = ?", organizationID).
		Order("created_at DESC").
		Limit(size).Offset((page - 1) * size).
		Find(&result).Error
	return
}
