package service

import (
	"context"
	"errors"
	"log/slog"

	"github.com/ryanadiputraa/inventra/domain"
	serviceError "github.com/ryanadiputraa/inventra/internal/errors"
	"github.com/ryanadiputraa/inventra/internal/inventory"
)

type service struct {
	logger     *slog.Logger
	repository inventory.InventoryRepository
}

func New(logger *slog.Logger, repository inventory.InventoryRepository) inventory.InventoryService {
	return &service{
		logger:     logger,
		repository: repository,
	}
}

func (s *service) AddItem(ctx context.Context, organizationID int, payload inventory.ItemPayload) (result domain.Item, err error) {
	prices := make([]domain.ItemPrice, 0)
	for _, p := range payload.Prices {
		prices = append(prices, domain.ItemPrice{
			Price:    p.Price,
			Quantity: p.Quantity,
		})
	}

	item, prices, err := domain.NewItem(organizationID, payload.ItemName, payload.Type, prices)
	if err != nil {
		return
	}

	result, err = s.repository.SaveItem(ctx, item, prices)
	if err != nil {
		if !errors.As(err, new(*serviceError.Error)) {
			s.logger.Error(
				"Fail to add inventory item",
				"error", err.Error(),
				"organization_id", organizationID,
				"item_name", item.ItemName,
			)
		}
		return
	}
	return
}

func (s *service) ListItems(ctx context.Context, organizationID, page, size int) (result []domain.Item, total int64, err error) {
	result, total, err = s.repository.FetchItems(ctx, organizationID, page, size)
	if err != nil {
		if !errors.As(err, new(*serviceError.Error)) {
			s.logger.Error(
				"Fail to fetch inventory item",
				"error", err.Error(),
				"organization_id", organizationID,
			)
		}
		return
	}
	return
}
