package service

import (
	"context"
	"errors"
	"log/slog"

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

func (s *service) AddItem(ctx context.Context, organizationID int, payload inventory.ItemPayload) (result inventory.Item, err error) {
	item, prices, err := inventory.NewItem(organizationID, payload)
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
	}
	return
}
