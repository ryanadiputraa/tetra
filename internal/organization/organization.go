package organization

import (
	"context"
	"time"

	"github.com/ryanadiputraa/inventra/internal/user"
)

type Organization struct {
	ID                int       `json:"id" gorm:"primaryKey;autoIncrement"`
	OwnerID           int       `json:"owner_id" gorm:"notNull"`
	Owner             user.User `json:"-"  gorm:"foreignKey:OwnerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name              string    `json:"name" gorm:"type:varchar(100);notNull"`
	CreatedAt         time.Time `json:"created_at" gorm:"notNull"`
	SubscriptionEndAt time.Time `json:"subscription_end_at" gorm:"notNull"`
	Members           []Member  `json:"-" `
}

type Member struct {
	ID             int       `json:"id" gorm:"primaryKey;autoIncrement"`
	OrganizationID int       `json:"organization_id" gorm:"notNull;constraint:OnDelete:CASCADE"`
	UserID         int       `json:"user_id" gorm:"notNull"`
	User           user.User `json:"-" gorm:"constraint:OnDelete:CASCADE"`
	Role           string    `json:"role" gorm:"type:varchar(10);notNull"`
	CreatedAt      time.Time `json:"created_at" gorm:"notNull"`
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

func NewMember(organizationID, userID int, role string) Member {
	return Member{
		OrganizationID: organizationID,
		UserID:         userID,
		Role:           role,
		CreatedAt:      time.Now().UTC(),
	}
}

type OrganizationService interface {
	Create(ctx context.Context, Name string, userID int) (Organization, error)
	IsSubscriptionValid(ctx context.Context, organizationID int) (bool, error)
}

type OrganizationRepository interface {
	Save(ctx context.Context, organization Organization) (Organization, error)
	FindByID(ctx context.Context, organizationID int) (Organization, error)
}
