package utilization

import (
	"context"
	"mime/multipart"

	"github.com/ryanadiputraa/inventra/domain"
)

type UtilizationService interface {
	Import(ctx context.Context, file multipart.File) error
}

type UtilizationRepository interface {
	Import(ctx context.Context, data []domain.Utilization) error
}
