package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ryanadiputraa/inventra/internal/auth"
	"github.com/ryanadiputraa/inventra/internal/errors"
	"github.com/ryanadiputraa/inventra/internal/user"
	"github.com/ryanadiputraa/inventra/pkg/jwt"
	"github.com/ryanadiputraa/inventra/pkg/validator"
	"github.com/ryanadiputraa/inventra/pkg/writer"
)

type handler struct {
	writer      writer.HTTPWriter
	validator   validator.Validator
	jwt         jwt.JWT
	service     auth.AuthService
	userService user.UserService
}

func New(writer writer.HTTPWriter, validator validator.Validator, jwt jwt.JWT, service auth.AuthService, userService user.UserService) *handler {
	return &handler{
		writer:      writer,
		validator:   validator,
		jwt:         jwt,
		service:     service,
		userService: userService,
	}
}

func (h *handler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p auth.LoginPayload
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			h.writer.WriteErrorResponse(w, http.StatusBadRequest, errors.BadRequest)
			return
		}

		errMap, err := h.validator.Validate(p)
		if err != nil {
			h.writer.WriteErrorResponseWithDetail(w, http.StatusBadRequest, errors.BadRequest, errMap)
			return
		}

		user, err := h.userService.Login(r.Context(), p.Email, p.Password)
		if err != nil {
			if sErr, ok := err.(*errors.ServiceErr); ok {
				h.writer.WriteErrorResponse(w, errors.HttpErrMap[sErr.ErrCode], err.Error())
				return
			}
			h.writer.WriteErrorResponse(w, http.StatusInternalServerError, errors.ServerError)
			return
		}

		jwt, err := h.service.GenerateJWT(r.Context(), user.ID)
		if err != nil {
			if sErr, ok := err.(*errors.ServiceErr); ok {
				h.writer.WriteErrorResponse(w, errors.HttpErrMap[sErr.ErrCode], errors.Unauthorized)
				return
			}
			h.writer.WriteErrorResponse(w, http.StatusInternalServerError, errors.ServerError)
			return
		}

		h.writer.WriteResponseData(w, http.StatusOK, jwt)
	}
}
