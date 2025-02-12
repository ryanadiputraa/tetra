package service

import (
	"context"
	"errors"
	"log/slog"
	"time"

	serviceError "github.com/ryanadiputraa/inventra/internal/errors"
	"github.com/ryanadiputraa/inventra/internal/organization"
)

type service struct {
	logger     *slog.Logger
	repository organization.OrganizationRepository
}

func New(logger *slog.Logger, repository organization.OrganizationRepository) organization.OrganizationService {
	return &service{
		logger:     logger,
		repository: repository,
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
