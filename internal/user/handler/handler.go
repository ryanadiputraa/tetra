package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ryanadiputraa/inventra/internal/auth"
	"github.com/ryanadiputraa/inventra/internal/errors"
	"github.com/ryanadiputraa/inventra/internal/user"
	"github.com/ryanadiputraa/inventra/pkg/validator"
	"github.com/ryanadiputraa/inventra/pkg/writer"
)

type handler struct {
	writer    writer.HTTPWriter
	validator validator.Validator
	service   user.UserService
}

func New(writer writer.HTTPWriter, validator validator.Validator, service user.UserService) *handler {
	return &handler{
		writer:    writer,
		validator: validator,
		service:   service,
	}
}

func (h *handler) GetUserData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := r.Context().(*auth.AppContext)

		user, err := h.service.GetByID(c, c.UserID)
		if err != nil {
			if sErr, ok := err.(*errors.Error); ok {
				h.writer.WriteErrorResponse(w, errors.HttpErrMap[sErr.ErrCode], sErr.Error())
				return
			} else {
				h.writer.WriteErrorResponse(w, http.StatusInternalServerError, errors.ServerError)
				return
			}
		}

		h.writer.WriteResponseData(w, http.StatusOK, user)
	}
}

func (h *handler) ChangePassword() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := r.Context().(*auth.AppContext)
		var p user.ChangePassowrdPayload
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

		err = h.service.ChangePassword(c, c.UserID, p.Password)
		if err != nil {
			if sErr, ok := err.(*errors.Error); ok {
				h.writer.WriteErrorResponse(w, errors.HttpErrMap[sErr.ErrCode], sErr.Error())
				return
			}
			h.writer.WriteErrorResponse(w, http.StatusInternalServerError, errors.ServerError)
			return
		}

		h.writer.WriteResponseData(w, http.StatusNoContent, nil)
	}
}
