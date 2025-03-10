package inventory

import (
	"context"
	"time"

	// serviceError "github.com/ryanadiputraa/inventra/internal/errors"
	"github.com/ryanadiputraa/inventra/internal/organization"
)

type ItemType string

const (
	Consumable ItemType = "consumable"  // Items that are used up (stock decreases)
	FixedAsset ItemType = "fixed_asset" // Items that are not used up (tracked but not depleted)
)

type Item struct {
	ID             int                       `json:"id" gorm:"primaryKey;autoIncrement"`
	OrganizationID int                       `json:"organization_id" gorm:"notNull"`
	Organization   organization.Organization `json:"-"  gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ItemName       string                    `json:"item_name" gorm:"type:varchar(100);notNull"`
	ItemType       ItemType                  `json:"item_type" gorm:"type:varchar(50);notNull"`
	Stock          []ItemPrice               `json:"stocks"`
	CreatedAt      time.Time                 `json:"created_at" gorm:"notNull"`
}

type ItemPrice struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	ItemID    int       `json:"item_id" gorm:"notNull"`
	Item      Item      `json:"-"  gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Price     int       `json:"price" gorm:"notNull"`
	Quantity  int       `json:"quantity" gorm:"notNull"`
	CreatedAt time.Time `json:"created_at" gorm:"notNull"`
}

type ItemPayload struct {
	ItemName string         `json:"item_name" validate:"required"`
	Type     string         `json:"type" validate:"required"`
	Prices   []PricePayload `json:"prices" validate:"required,dive"`
}

type PricePayload struct {
	Price    int `json:"price" validate:"required"`
	Quantity int `json:"quantity" validate:"required,min=1"`
}

type InventoryService interface {
	AddItem(ctx context.Context, organizationID int, payload ItemPayload) (Item, error)
}
