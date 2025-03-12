package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ryanadiputraa/inventra/internal/auth"
	"github.com/ryanadiputraa/inventra/internal/errors"
	"github.com/ryanadiputraa/inventra/internal/inventory"
	"github.com/ryanadiputraa/inventra/pkg/validator"
	"github.com/ryanadiputraa/inventra/pkg/writer"
)

type handler struct {
	writer    writer.HTTPWriter
	validator validator.Validator
	service   inventory.InventoryService
}

func New(writer writer.HTTPWriter, validator validator.Validator, service inventory.InventoryService) *handler {
	return &handler{
		writer:    writer,
		validator: validator,
		service:   service,
	}
}

func (h *handler) FetchItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := r.Context().(*auth.AppContext)
		q := r.URL.Query()
		p := q.Get("page")
		s := q.Get("size")

		page, err := strconv.Atoi(p)
		if err != nil {
			h.writer.WriteErrorResponse(w, http.StatusBadRequest, errors.BadRequest)
			return
		}
		size, err := strconv.Atoi(s)
		if err != nil {
			h.writer.WriteErrorResponse(w, http.StatusBadRequest, errors.BadRequest)
			return
		}

		items, total, err := h.service.ListItems(c, *c.OrganizationID, page, size)
		if err != nil {
			if sErr, ok := err.(*errors.Error); ok {
				h.writer.WriteErrorResponse(w, errors.HttpErrMap[sErr.ErrCode], sErr.Error())
				return
			} else {
				h.writer.WriteErrorResponse(w, http.StatusInternalServerError, errors.ServerError)
				return
			}
		}

		h.writer.WriteResponseDataWithPagination(w, http.StatusOK, items, "items", page, size, total)
	}
}

func (h *handler) AddInventoryItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := r.Context().(*auth.AppContext)
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

		item, err := h.service.AddItem(c, *c.OrganizationID, p)
		if err != nil {
			if sErr, ok := err.(*errors.Error); ok {
				h.writer.WriteErrorResponse(w, errors.HttpErrMap[sErr.ErrCode], sErr.Error())
				return
			} else {
				h.writer.WriteErrorResponse(w, http.StatusInternalServerError, errors.ServerError)
				return
			}
		}

		h.writer.WriteResponseData(w, http.StatusCreated, item)
	}
}
