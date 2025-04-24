package service

import (
	"context"
	"encoding/csv"
	"errors"
	"io"
	"log/slog"
	"mime/multipart"
	"time"

	"github.com/ryanadiputraa/inventra/domain"
	serviceErr "github.com/ryanadiputraa/inventra/internal/errors"
	"github.com/ryanadiputraa/inventra/internal/utilization"
)

type service struct {
	logger     *slog.Logger
	repository utilization.UtilizationRepository
}

func New(logger *slog.Logger, repository utilization.UtilizationRepository) utilization.UtilizationService {
	return &service{
		logger:     logger,
		repository: repository,
	}
}

func (s *service) Import(ctx context.Context, file multipart.File) error {
	reader := csv.NewReader(file)
	var record []string
	var data []domain.Utilization

	isOnReadHeader := true
	for {
		if isOnReadHeader {
			if _, err := reader.Read(); err != nil {
				s.logger.Error("Fail to read CSV import", "error", err.Error())
				return err
			}
			isOnReadHeader = false
			continue
		}

		var err error
		record, err = reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			s.logger.Error("Fail to read CSV import", "error", err.Error())
			return err
		}

		// TODO: validate column
		date, err := time.Parse("02/01/2006", record[0])
		if err != nil {
			return err
		}

		u := domain.Utilization{
			Date:         date,
			Contract:     record[1],
			MoveType:     record[2],
			UnitCategory: record[3],
			UnitName:     record[4],
			Unit:         record[5],
			Condition:    record[6],
			CreatedAt:    time.Now().UTC(),
		}
		data = append(data, u)
	}

	err := s.repository.Import(ctx, data)
	if err != nil {
		if !errors.As(err, new(*serviceErr.Error)) {
			s.logger.Error(
				"Fail to import utilization data",
				"error", err.Error(),
			)
		}
		return err
	}

	return nil
}
