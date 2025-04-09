package organization

import (
	"context"

	"github.com/ryanadiputraa/inventra/domain"
)

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

type OrganizationService interface {
	Create(ctx context.Context, Name string, userID int) (domain.Organization, error)
	GetByID(ctx context.Context, organizationID int) (domain.Organization, error)
	IsSubscriptionValid(ctx context.Context, organizationID int) (bool, error)
	Delete(ctx context.Context, organizationID, userID int) error
	ListMember(ctx context.Context, organizationID int) ([]domain.MemberData, error)
	InviteUser(ctx context.Context, organizationID int, email string) error
	Join(ctx context.Context, organizationID, userID int) (domain.Member, error)
	RemoveMember(ctx context.Context, organizationID, memberID int) error
	ChangeMemberRole(ctx context.Context, organizationID, memberID int, role string) error
	Leave(ctx context.Context, organizationID, memberID int) error
}

type OrganizationRepository interface {
	Save(ctx context.Context, organization domain.Organization) (domain.Organization, error)
	FindByID(ctx context.Context, organizationID int) (domain.Organization, error)
	Delete(ctx context.Context, organizationID, userID int) error
	AddMember(ctx context.Context, member domain.Member) (domain.Member, error)
	FetchMembers(ctx context.Context, organizationID int) ([]domain.MemberData, error)
	DeleteMember(ctx context.Context, organizationID, memberID int) error
	UpdateMemberRole(ctx context.Context, organizationID, memberID int, role string) error
}
