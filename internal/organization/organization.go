package organization

import (
	"context"
	"time"

	"github.com/ryanadiputraa/inventra/domain"
)

type OrganizationCache struct {
	ID                   int         `json:"id"`
	Owner                domain.User `json:"owner"`
	Name                 string      `json:"name"`
	CreatedAt            time.Time   `json:"created_at"`
	SubscriptionEndAt    time.Time   `json:"subscription_end_at"`
	OdooUsername         *string     `json:"odoo_username"`
	OdooPassword         *string     `json:"odoo_password"`
	IntellitrackUsername *string     `json:"intellitrack_username"`
	IntellitrackPassword *string     `json:"intellitrack_password"`
}

type OrganizationPayload struct {
	Name string `json:"name" validate:"required"`
}

type InvitePayload struct {
	Email string `json:"email" validate:"required,email"`
}

type AcceptInvitationPayload struct {
	Code string `json:"code" validate:"required"`
}

type ChangeMemberPayload struct {
	Role string `json:"role" validate:"required"`
}

type DashboardSettings struct {
	OdooUsername         *string `json:"odoo_username"`
	OdooPassword         *string `json:"odoo_password"`
	IntellitrackUsername *string `json:"intellitrack_username"`
	IntellitrackPassword *string `json:"intellitrack_password"`
}

func CacheFromOrg(org domain.Organization) OrganizationCache {
	return OrganizationCache{
		ID:                   org.ID,
		Owner:                org.Owner,
		Name:                 org.Name,
		CreatedAt:            org.CreatedAt,
		SubscriptionEndAt:    org.SubscriptionEndAt,
		OdooUsername:         org.OdooUsername,
		OdooPassword:         org.OdooPassword,
		IntellitrackUsername: org.IntellitrackUsername,
		IntellitrackPassword: org.IntellitrackPassword,
	}
}

func OrgFromCache(cache OrganizationCache) domain.Organization {
	return domain.Organization{
		ID:                   cache.ID,
		Owner:                cache.Owner,
		Name:                 cache.Name,
		CreatedAt:            cache.CreatedAt,
		SubscriptionEndAt:    cache.SubscriptionEndAt,
		OdooUsername:         cache.OdooUsername,
		OdooPassword:         cache.OdooPassword,
		IntellitrackUsername: cache.IntellitrackUsername,
		IntellitrackPassword: cache.IntellitrackPassword,
	}
}

type OrganizationService interface {
	Create(ctx context.Context, Name string, userID int) (domain.Organization, error)
	GetByID(ctx context.Context, organizationID int) (domain.OrganizationData, error)
	IsSubscriptionValid(ctx context.Context, organizationID int) (bool, error)
	Delete(ctx context.Context, organizationID, userID int) error
	ListMember(ctx context.Context, organizationID int) ([]domain.MemberData, error)
	InviteUser(ctx context.Context, organizationID int, email string) error
	Join(ctx context.Context, organizationID, userID int) (domain.Member, error)
	RemoveMember(ctx context.Context, organizationID, memberID int) error
	ChangeMemberRole(ctx context.Context, organizationID, memberID int, role string) error
	Leave(ctx context.Context, organizationID, memberID int) error
	UpdateDashboardSettings(ctx context.Context, organizationID int, settings DashboardSettings) error
}

type OrganizationRepository interface {
	Save(ctx context.Context, organization domain.Organization) (domain.Organization, error)
	FindByID(ctx context.Context, organizationID int) (domain.Organization, error)
	Delete(ctx context.Context, organizationID, userID int) error
	AddMember(ctx context.Context, member domain.Member) (domain.Member, error)
	FetchMembers(ctx context.Context, organizationID int) ([]domain.MemberData, error)
	DeleteMember(ctx context.Context, organizationID, memberID int) error
	UpdateMemberRole(ctx context.Context, organizationID, memberID int, role string) error
	UpdateDashboardSettings(ctx context.Context, organizationID int, settings DashboardSettings) error
}
