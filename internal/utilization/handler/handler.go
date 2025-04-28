package handler

import (
	"net/http"

	"github.com/ryanadiputraa/tetra/internal/errors"
	"github.com/ryanadiputraa/tetra/internal/utilization"
	"github.com/ryanadiputraa/tetra/pkg/writer"
)

type handler struct {
	writer  writer.HTTPWriter
	service utilization.UtilizationService
}

func New(writer writer.HTTPWriter, service utilization.UtilizationService) *handler {
	return &handler{
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
