package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"github.com/ryanadiputraa/tetra/config"
	"github.com/ryanadiputraa/tetra/domain"
	"github.com/ryanadiputraa/tetra/internal/errors"
	"github.com/ryanadiputraa/tetra/internal/utilization"
	"github.com/ryanadiputraa/tetra/pkg/writer"
)

type handler struct {
	logger  *slog.Logger
	config  config.Config
	writer  writer.HTTPWriter
	service utilization.UtilizationService
}

func New(logger *slog.Logger, c config.Config, writer writer.HTTPWriter, service utilization.UtilizationService) *handler {
	return &handler{
		logger:  logger,
		config:  c,
		writer:  writer,
		service: service,
	}
}

func (h *handler) Import() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(10 << 20) // 10MB max file size
		if err != nil {
			h.writer.WriteErrorResponse(w, http.StatusBadRequest, errors.Max10MBFileSize)
			return
		}

		file, _, err := r.FormFile("file")
		if err != nil {
			h.writer.WriteErrorResponse(w, http.StatusBadRequest, errors.BadRequest)
			return
		}
		defer file.Close()

		err = h.service.Import(r.Context(), file)
		if err != nil {
			if sErr, ok := err.(*errors.Error); ok {
				h.writer.WriteErrorResponse(w, errors.HttpErrMap[sErr.ErrCode], err.Error())
				return
			}
			h.writer.WriteErrorResponse(w, http.StatusInternalServerError, errors.ServerError)
			return
		}

		h.writer.WriteResponseData(w, http.StatusNoContent, nil)
	}
}

func (h *handler) GetUtilizations() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		date := q.Get("date")
		startTime := q.Get("start_time")
		endTime := q.Get("end_time")

		_, err := time.Parse("2006-01-02", date)
		if err != nil {
			h.writer.WriteErrorResponse(w, http.StatusBadRequest, errors.BadRequest)
			return
		}
		_, err = time.Parse("15:04:05", startTime)
		if err != nil {
			h.writer.WriteErrorResponse(w, http.StatusBadRequest, errors.BadRequest)
			return
		}
		_, err = time.Parse("15:04:05", endTime)
		if err != nil {
			h.writer.WriteErrorResponse(w, http.StatusBadRequest, errors.BadRequest)
			return
		}

		params := url.Values{}
		params.Add("date", date)
		params.Add("start_time", startTime)
		params.Add("end_time", endTime)

		url := h.config.DashboardServiceURI + "/api/v1/utilizations?" + params.Encode()
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			h.writer.WriteErrorResponse(w, http.StatusInternalServerError, errors.ServerError)
			return

		}

		req.Header.Add("Authorization", r.Header.Get("Authorization"))

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			h.writer.WriteErrorResponse(w, http.StatusBadRequest, errors.BadRequest)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			var serviceErr utilization.UtilizationServiceError
			if err = json.NewDecoder(resp.Body).Decode(&serviceErr); err != nil {
				h.writer.WriteErrorResponse(w, http.StatusInternalServerError, errors.ServerError)
				return
			}

			if resp.StatusCode == http.StatusInternalServerError {
				h.writer.WriteErrorResponse(w, http.StatusInternalServerError, errors.ServerError)
				return
			}
			h.writer.WriteErrorResponse(w, resp.StatusCode, serviceErr.Message)
			return
		}

		var data domain.Utilizations
		if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
			h.logger.Info("DEBUG", "err", err)
			h.writer.WriteErrorResponse(w, http.StatusInternalServerError, errors.ServerError)
			return
		}

		h.writer.WriteResponseData(w, http.StatusOK, data)
	}
}
