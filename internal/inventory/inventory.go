package inventory

import (
	"context"

	"github.com/ryanadiputraa/inventra/domain"
)

type ItemPayload struct {
	ItemName string         `json:"item_name" validate:"required"`
	Type     string         `json:"type" validate:"required"`
	Prices   []PricePayload `json:"prices" validate:"required,min=1,dive"`
}

type PricePayload struct {
	Price    int `json:"price" validate:"required"`
	Quantity int `json:"quantity" validate:"required,min=1"`
}

type InventoryService interface {
	AddItem(ctx context.Context, organizationID int, payload ItemPayload) (domain.Item, error)
	ListItems(ctx context.Context, organizationID, page, size int) ([]domain.Item, int64, error)
}

type InventoryRepository interface {
	SaveItem(ctx context.Context, item domain.Item, prices []domain.ItemPrice) (domain.Item, error)
	FetchItems(ctx context.Context, organizationID, page, size int) ([]domain.Item, int64, error)
}
