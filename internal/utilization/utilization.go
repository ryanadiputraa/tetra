package utilization

import (
	"context"
	"mime/multipart"

	"github.com/ryanadiputraa/tetra/domain"
)

type UtilizationServiceError struct {
	Message string `json:"message"`
}

type UtilizationService interface {
	Import(ctx context.Context, file multipart.File) error
}

type UtilizationRepository interface {
	Import(ctx context.Context, data []domain.Utilization) error
}
