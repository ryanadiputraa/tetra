package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/ryanadiputraa/inventra/config"
	serviceError "github.com/ryanadiputraa/inventra/internal/errors"
	"github.com/ryanadiputraa/inventra/internal/organization"
	"github.com/ryanadiputraa/inventra/internal/user"
	"github.com/ryanadiputraa/inventra/pkg/jwt"
	"github.com/ryanadiputraa/inventra/pkg/mail"
)

type service struct {
	config         config.Config
	logger         *slog.Logger
	jwt            jwt.JWT
	smtpMail       mail.SMTPMail
	repository     organization.OrganizationRepository
	userRepository user.UserRepository
}

func New(
	config config.Config,
	logger *slog.Logger,
	jwt jwt.JWT,
	smtpMail mail.SMTPMail,
	repository organization.OrganizationRepository,
	userRepository user.UserRepository,
) organization.OrganizationService {
	return &service{
		config:         config,
		logger:         logger,
		jwt:            jwt,
		smtpMail:       smtpMail,
		repository:     repository,
		userRepository: userRepository,
	}
}

func (s *service) Create(ctx context.Context, Name string, userID int) (result organization.Organization, err error) {
	o := organization.New(Name, userID)
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
		"owner", result.OwnerID,
		"created_at", result.CreatedAt,
	)
	return
}

func (s *service) IsSubscriptionValid(ctx context.Context, organizationID int) (isValid bool, err error) {
	organization, err := s.repository.FindByID(ctx, organizationID)
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

	isValid = time.Now().UTC().Before(organization.SubscriptionEndAt)
	return
}

func (s *service) ListMember(ctx context.Context, organizationID int) (result []organization.MemberData, err error) {
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
		s.logger.Error(
			"Fail to fetch user data",
			"error", err.Error(),
			"email", email,
		)
		return
	}

	if user.OrganizationID != nil {
		err = serviceError.NewServiceErr(serviceError.BadRequest, serviceError.UserHasJoinedOrg)
		return
	}

	org, err := s.repository.FindByID(ctx, organizationID)
	if err != nil {
		s.logger.Error(
			"Fail to fetch organization data",
			"error", err.Error(),
			"organization_id", organizationID,
		)
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
		body := organization.GenrateInvitationMailBody(org.Name, s.config.AppDomain, jwt.AccessToken)
		s.smtpMail.SendMail(context.Background(), email, subject, body)
	}()
	return
}
