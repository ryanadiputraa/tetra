package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/ryanadiputraa/inventra/config"
	"github.com/ryanadiputraa/inventra/domain"
	serviceError "github.com/ryanadiputraa/inventra/internal/errors"
	"github.com/ryanadiputraa/inventra/internal/organization"
	"github.com/ryanadiputraa/inventra/internal/user"
	"github.com/ryanadiputraa/inventra/pkg/jwt"
	"github.com/ryanadiputraa/inventra/pkg/mail"
	"github.com/ryanadiputraa/inventra/pkg/secure"
)

type service struct {
	config         config.Config
	logger         *slog.Logger
	jwt            jwt.JWTService
	smtpMail       mail.SMTPMail
	secure         secure.Secure
	repository     organization.OrganizationRepository
	userRepository user.UserRepository
}

func New(
	config config.Config,
	logger *slog.Logger,
	jwt jwt.JWTService,
	smtpMail mail.SMTPMail,
	secure secure.Secure,
	repository organization.OrganizationRepository,
	userRepository user.UserRepository,
) organization.OrganizationService {
	return &service{
		config:         config,
		logger:         logger,
		jwt:            jwt,
		smtpMail:       smtpMail,
		secure:         secure,
		repository:     repository,
		userRepository: userRepository,
	}
}

func (s *service) Create(ctx context.Context, Name string, userID int) (result domain.Organization, err error) {
	o := domain.NewOrganization(Name, userID)
	result, err = s.repository.Save(ctx, o)
	if err != nil {
		if !errors.As(err, new(*serviceError.Error)) {
			s.logger.Error(
				"Fail to register new organization",
				"error", err.Error(),
				"user_id", userID,
				"organization_name", Name,
			)
		}
		return
	}

	s.logger.Info(
		"New organization registered",
		"id", result.ID,
		"name", result.Name,
		"owner", result.Owner.ID,
		"created_at", result.CreatedAt,
	)
	return
}

func (s *service) GetByID(ctx context.Context, organizationID int) (result domain.OrganizationData, err error) {
	org, err := s.repository.FindByID(ctx, organizationID)
	if err != nil {
		if !errors.As(err, new(*serviceError.Error)) {
			s.logger.Error(
				"Fail to fetch organization data",
				"error", err.Error(),
				"organiaztion_id", organizationID,
			)
		}
		return
	}

	result.ID = org.ID
	result.Owner = org.Owner
	result.Name = org.Name
	result.CreatedAt = org.CreatedAt
	result.SubscriptionEndAt = org.SubscriptionEndAt
	result.Features.Dashboard = false

	if org.OdooURL != nil && org.OdooDB != nil && org.OdooUsername != nil && org.OdooPassword != nil &&
		org.IntellitrackUsername != nil && org.IntellitrackPassword != nil {
		result.Features.Dashboard = true
	}
	return
}

func (s *service) IsSubscriptionValid(ctx context.Context, organizationID int) (isValid bool, err error) {
	organization, err := s.GetByID(ctx, organizationID)
	if err != nil {
		return
	}
	isValid = time.Now().UTC().Before(organization.SubscriptionEndAt)
	return
}

func (s *service) Delete(ctx context.Context, organizationID, userID int) (err error) {
	err = s.repository.Delete(ctx, organizationID, userID)
	if err != nil {
		if !errors.As(err, new(*serviceError.Error)) {
			s.logger.Error(
				"Fail to delete organization",
				"error", err.Error(),
				"organization_id", organizationID,
			)
		}
		return
	}
	return
}

func (s *service) ListMember(ctx context.Context, organizationID int) (result []domain.MemberData, err error) {
	result, err = s.repository.FetchMembers(ctx, organizationID)
	if err != nil {
		if !errors.As(err, new(*serviceError.Error)) {
			s.logger.Error(
				"Fail to fetch organization members",
				"error", err.Error(),
				"organization_id", organizationID,
			)
		}
		return
	}
	return
}

func (s *service) InviteUser(ctx context.Context, organizationID int, email string) (err error) {
	user, err := s.userRepository.FindByEmail(ctx, email)
	if err != nil {
		if !errors.As(err, new(*serviceError.Error)) {
			s.logger.Error(
				"Fail to fetch user data",
				"error", err.Error(),
				"email", email,
			)
		}
		return
	}

	if user.OrganizationID != nil {
		err = serviceError.NewServiceErr(serviceError.BadRequest, serviceError.UserHasJoinedOrg)
		return
	}

	org, err := s.GetByID(ctx, organizationID)
	if err != nil {
		return
	}

	// organization id stored in user_id field to be used in /api/join endpoint to join sender org
	jwt, err := s.jwt.GenerateJWTWithClaims(organizationID)
	if err != nil {
		s.logger.Error(
			"Fail to generate jwt",
			"error", err.Error(),
		)
		return
	}

	go func() {
		subject := fmt.Sprintf("Undangan bergabung dengan %s di Inventra", org.Name)
		body := domain.GenrateInvitationMailBody(org.Name, s.config.FrontendURL, jwt.AccessToken)
		if err = s.smtpMail.SendMail(context.Background(), email, subject, body); err != nil {
			s.logger.Error(
				"Fail to send invitation mail",
				"error", err.Error(),
				"organization_id", organizationID,
				"address", email,
			)
		}
	}()
	return
}

func (s *service) Join(ctx context.Context, organizationID, userID int) (result domain.Member, err error) {
	m := domain.NewMember(organizationID, userID, domain.Staff)
	result, err = s.repository.AddMember(ctx, m)
	if err != nil {
		if !errors.As(err, new(*serviceError.Error)) {
			s.logger.Error(
				"Fail to add organization member",
				"error", err.Error(),
				"organization_id", organizationID,
				"user_id", userID,
			)
		}
		return
	}

	s.logger.Info(
		"New organization member joined",
		"organization_id", result.OrganizationID,
		"user_id", result.UserID,
	)
	return
}

func (s *service) RemoveMember(ctx context.Context, organizationID, memberID int) (err error) {
	err = s.repository.DeleteMember(ctx, organizationID, memberID)
	if err != nil {
		if !errors.As(err, new(*serviceError.Error)) {
			s.logger.Error(
				"Fail to remove member",
				"error", err.Error(),
				"organization_id", organizationID,
				"memberID", memberID,
			)
		}
		return
	}
	return
}

func (s *service) ChangeMemberRole(ctx context.Context, organizationID, memberID int, role string) (err error) {
	if !domain.IsValidRole(domain.Role(role)) {
		return serviceError.NewServiceErr(serviceError.BadRequest, serviceError.InvalidRole)
	}

	err = s.repository.UpdateMemberRole(ctx, organizationID, memberID, role)
	if err != nil {
		if !errors.As(err, new(*serviceError.Error)) {
			s.logger.Error(
				"Fail to change member role",
				"error", err.Error(),
				"organization_id", organizationID,
				"memberID", memberID,
				"role", role,
			)
		}
		return
	}
	return
}

func (s *service) Leave(ctx context.Context, organizationID, memberID int) (err error) {
	members, err := s.repository.FetchMembers(ctx, organizationID)
	if err != nil {
		return
	}

	adminCnt := 0
	for _, m := range members {
		if m.Role == string(domain.Admin) {
			adminCnt++
		}
	}
	if adminCnt < 2 {
		err = serviceError.NewServiceErr(serviceError.BadRequest, serviceError.NotEnoughAdmin)
		return
	}

	err = s.repository.DeleteMember(ctx, organizationID, memberID)
	if err != nil {
		if !errors.As(err, new(*serviceError.Error)) {
			s.logger.Error(
				"Fail to leave organization",
				"error", err.Error(),
				"organization_id", organizationID,
				"memberID", memberID,
			)
		}
		return
	}
	return
}

func (s *service) UpdateDashboardSettings(ctx context.Context, organizationID int, settings organization.DashboardSettings) (err error) {
	odooEncPass, err := s.secure.Encrypt(*settings.OdooPassword)
	if err != nil {
		s.logger.Error(
			"Fail to encrypt Odoo password",
			"error", err.Error(),
		)
	}
	intellitrackEncPass, err := s.secure.Encrypt(*settings.IntellitrackPassword)
	if err != nil {
		s.logger.Error(
			"Fail to encrypt Intellitrack password",
			"error", err.Error(),
		)
	}

	settings.OdooPassword = &odooEncPass
	settings.IntellitrackPassword = &intellitrackEncPass

	err = s.repository.UpdateDashboardSettings(ctx, organizationID, settings)
	if err != nil {
		if !errors.As(err, new(*serviceError.Error)) {
			s.logger.Error(
				"Fail to update organization dashboard settgins",
				"error", err.Error(),
				"organization_id", organizationID,
			)
		}
		return
	}
	return
}
