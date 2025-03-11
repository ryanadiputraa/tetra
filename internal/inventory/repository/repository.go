package repository

import (
	"context"
	"errors"

	"github.com/ryanadiputraa/inventra/internal/inventory"
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

func (r *repository) SaveItem(ctx context.Context, item inventory.Item, prices []inventory.ItemPrice) (result inventory.Item, err error) {
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
