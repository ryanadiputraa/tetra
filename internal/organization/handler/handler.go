package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ryanadiputraa/inventra/internal/auth"
	"github.com/ryanadiputraa/inventra/internal/errors"
	"github.com/ryanadiputraa/inventra/internal/organization"
	"github.com/ryanadiputraa/inventra/pkg/writer"
)

type handler struct {
	writer  writer.HTTPWriter
	service organization.OrganizationService
}

func New(writer writer.HTTPWriter, service organization.OrganizationService) *handler {
	return &handler{
		writer:  writer,
		service: service,
	}
}

func (h *handler) CreateOrganization() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := r.Context().(*auth.AppContext)
		var p organization.OrganizationPayload
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			h.writer.WriteErrorResponse(w, http.StatusBadRequest, errors.BadRequest)
			return
		}

		res, err := h.service.Create(c, p.Name, c.UserID)
		if err != nil {
			if sErr, ok := err.(*errors.Error); ok {
				h.writer.WriteErrorResponse(w, errors.HttpErrMap[sErr.ErrCode], sErr.Error())
				return
			}
			h.writer.WriteErrorResponse(w, http.StatusInternalServerError, errors.ServerError)
			return
		}

		h.writer.WriteResponseData(w, http.StatusCreated, res)
	}
}
