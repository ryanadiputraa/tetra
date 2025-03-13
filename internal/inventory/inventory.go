package inventory

import (
	"context"
	"time"

	serviceError "github.com/ryanadiputraa/inventra/internal/errors"
	"github.com/ryanadiputraa/inventra/internal/organization"
)

type ItemType string

const (
	Consumable ItemType = "consumable"  // Items that are used up (stock decreases)
	FixedAsset ItemType = "fixed_asset" // Items that are not used up (tracked but not depleted)
)

type Item struct {
	ID             int                       `json:"id" gorm:"primaryKey;autoIncrement"`
	OrganizationID int                       `json:"-" gorm:"notNull"`
	Organization   organization.Organization `json:"-"  gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ItemName       string                    `json:"item_name" gorm:"type:varchar(100);notNull"`
	ItemType       ItemType                  `json:"item_type" gorm:"type:varchar(50);notNull"`
	Stock          []ItemPrice               `json:"stock"`
	CreatedAt      time.Time                 `json:"created_at" gorm:"notNull"`
}

type ItemPrice struct {
	ID        int       `json:"-" gorm:"primaryKey;autoIncrement"`
	ItemID    int       `json:"-" gorm:"notNull"`
	Item      Item      `json:"-"  gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Price     int       `json:"price" gorm:"notNull"`
	Quantity  int       `json:"quantity" gorm:"notNull"`
	CreatedAt time.Time `json:"created_at" gorm:"notNull"`
}

type ItemPayload struct {
	ItemName string         `json:"item_name" validate:"required"`
	Type     string         `json:"type" validate:"required"`
	Prices   []PricePayload `json:"prices" validate:"required,min=1,dive"`
}

type PricePayload struct {
	Price    int `json:"price" validate:"required"`
	Quantity int `json:"quantity" validate:"required,min=1"`
}

func NewItem(organizationID int, p ItemPayload) (i Item, s []ItemPrice, err error) {
	switch p.Type {
	case string(Consumable), string(FixedAsset):
		break
	default:
		err = serviceError.NewServiceErr(serviceError.BadRequest, serviceError.InvalidItemType)
		return
	}

	now := time.Now().UTC()
	i = Item{
		OrganizationID: organizationID,
		ItemName:       p.ItemName,
		ItemType:       ItemType(p.Type),
		CreatedAt:      now,
	}

	s = make([]ItemPrice, 0)
	for _, price := range p.Prices {
		s = append(s, ItemPrice{
			Price:     price.Price,
			Quantity:  price.Quantity,
			CreatedAt: now,
		})
	}
	return
}

type InventoryService interface {
	AddItem(ctx context.Context, organizationID int, payload ItemPayload) (Item, error)
	ListItems(ctx context.Context, organizationID, page, size int) ([]Item, int64, error)
}

type InventoryRepository interface {
	SaveItem(ctx context.Context, item Item, prices []ItemPrice) (Item, error)
	FetchItems(ctx context.Context, organizationID, page, size int) ([]Item, int64, error)
}
