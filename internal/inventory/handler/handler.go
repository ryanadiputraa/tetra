package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ryanadiputraa/inventra/internal/errors"
	"github.com/ryanadiputraa/inventra/internal/inventory"
	"github.com/ryanadiputraa/inventra/pkg/validator"
	"github.com/ryanadiputraa/inventra/pkg/writer"
)

type handler struct {
	writer    writer.HTTPWriter
	validator validator.Validator
}

func New(writer writer.HTTPWriter, validator validator.Validator) *handler {
	return &handler{
		writer:    writer,
		validator: validator,
	}
}

func (h *handler) AddInventoryItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p inventory.ItemPayload
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			h.writer.WriteErrorResponse(w, http.StatusBadRequest, errors.BadRequest)
			return
		}

		errMap, err := h.validator.Validate(p)
		if err != nil {
			h.writer.WriteErrorResponseWithDetail(w, http.StatusBadRequest, err.Error(), errMap)
			return
		}

		h.writer.WriteResponseData(w, http.StatusCreated, nil)
	}
}
