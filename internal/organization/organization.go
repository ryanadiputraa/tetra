package organization

import (
	"context"
	"time"
)

type Organization struct {
	ID                int       `json:"id" gorm:"primaryKey"`
	OwnerID           int       `json:"owner_id" gorm:"unique;notNull"`
	Name              string    `json:"name" gorm:"type:varchar(100);notNull"`
	CreatedAt         time.Time `json:"created_at" gorm:"notNull"`
	SubscriptionEndAt time.Time `json:"subscription_end_at" gorm:"notNull"`
}

type OrganizationPayload struct {
	Name string `json:"name"`
}

func New(Name string, userID int) Organization {
	return Organization{
		OwnerID:           userID,
		Name:              Name,
		CreatedAt:         time.Now().UTC(),
		SubscriptionEndAt: time.Now().AddDate(0, 1, 0).UTC(),
	}
}

type OrganizationService interface {
	Create(ctx context.Context, Name string, userID int) (Organization, error)
}

type OrganizationRepository interface {
	Save(ctx context.Context, organization Organization) (Organization, error)
}
