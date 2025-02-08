package service

import (
	"context"

	"github.com/ryanadiputraa/inventra/internal/organization"
	"github.com/ryanadiputraa/inventra/pkg/logger"
)

type service struct {
	log        logger.Logger
	repository organization.OrganizationRepository
}

func New(log logger.Logger, repository organization.OrganizationRepository) organization.OrganizationService {
	return &service{
		log:        log,
		repository: repository,
	}
}

func (s *service) Create(ctx context.Context, Name string, userID int) (res organization.Organization, err error) {
	o := organization.New(Name, userID)
	res, err = s.repository.Save(ctx, o)
	if err != nil {
		s.log.Error("Fail to create new organization. Err: ", err.Error())
		return
	}
	return
}
