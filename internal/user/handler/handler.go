package handler

import (
	"net/http"

	"github.com/ryanadiputraa/inventra/internal/auth"
	"github.com/ryanadiputraa/inventra/internal/errors"
	"github.com/ryanadiputraa/inventra/internal/user"
	"github.com/ryanadiputraa/inventra/pkg/writer"
)

type handler struct {
	writer  writer.HTTPWriter
	service user.UserService
}

func New(writer writer.HTTPWriter, service user.UserService) *handler {
	return &handler{
		writer:  writer,
		service: service,
	}
}

func (h *handler) GetUserData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := r.Context().(*auth.AuthContext)

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
