package inventory

import "github.com/ryanadiputraa/inventra/internal/organization"

type ItemType string

const (
	Consumable ItemType = "consumable"  // Items that are used up (stock decreases)
	FixedAsset ItemType = "fixed_asset" // Items that are not used up (tracked but not depleted)
)

type Item struct {
	ID             int                       `json:"id" gorm:"primaryKey;autoIncrement"`
	OrganizationID int                       `json:"organization_id" gorm:"notNull"`
	Organization   organization.Organization `json:"-"  gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ItemType       ItemType                  `json:"item_type" gorm:"type:varchar(50);notNull"`
	ItemName       string                    `json:"item_name" gorm:"type:varchar(100);notNull"`
	Stock          int                       `json:"stock" gorm:"notNull"`
}

type ItemPrice struct {
	ID       int  `json:"id" gorm:"primaryKey;autoIncrement"`
	ItemID   int  `json:"item_id" gorm:"notNull"`
	Item     Item `json:"-"  gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Price    int  `json:"price" gorm:"notNull"`
	Quantity int  `json:"quantity" gorm:"notNull"`
}
