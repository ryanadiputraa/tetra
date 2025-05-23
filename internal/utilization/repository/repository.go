package repository

import (
	"context"

	"github.com/ryanadiputraa/tetra/domain"
	"github.com/ryanadiputraa/tetra/internal/utilization"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) utilization.UtilizationRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) Import(ctx context.Context, data []domain.Utilization) error {
	return r.db.Create(&data).Error
}
