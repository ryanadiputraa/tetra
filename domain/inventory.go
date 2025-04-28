package domain

import (
	"time"

	serviceError "github.com/ryanadiputraa/tetra/internal/errors"
)

type ItemType string

const (
	Consumable ItemType = "consumable"  // Items that are used up (stock decreases)
	FixedAsset ItemType = "fixed_asset" // Items that are not used up (tracked but not depleted)
)

type Item struct {
	ID             int          `json:"id" gorm:"primaryKey;autoIncrement"`
	OrganizationID int          `json:"-" gorm:"notNull"`
	Organization   Organization `json:"-"  gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ItemName       string       `json:"item_name" gorm:"type:varchar(100);notNull"`
	ItemType       ItemType     `json:"item_type" gorm:"type:varchar(50);notNull"`
	Stock          []ItemPrice  `json:"stock"`
	CreatedAt      time.Time    `json:"created_at" gorm:"notNull"`
}

type ItemPrice struct {
	ID        int       `json:"-" gorm:"primaryKey;autoIncrement"`
	ItemID    int       `json:"-" gorm:"notNull"`
	Item      Item      `json:"-"  gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Price     int       `json:"price" gorm:"notNull"`
	Quantity  int       `json:"quantity" gorm:"notNull"`
	CreatedAt time.Time `json:"created_at" gorm:"notNull"`
}

func NewItem(organizationID int, name, itemType string, prices []ItemPrice) (i Item, s []ItemPrice, err error) {
	switch itemType {
	case string(Consumable), string(FixedAsset):
		break
	default:
		err = serviceError.NewServiceErr(serviceError.BadRequest, serviceError.InvalidItemType)
		return
	}

	now := time.Now().UTC()
	i = Item{
		OrganizationID: organizationID,
		ItemName:       name,
		ItemType:       ItemType(itemType),
		CreatedAt:      now,
	}

	s = make([]ItemPrice, 0)
	for _, price := range prices {
		s = append(s, ItemPrice{
			Price:     price.Price,
			Quantity:  price.Quantity,
			CreatedAt: now,
		})
	}
	return
}
