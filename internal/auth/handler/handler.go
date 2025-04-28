package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ryanadiputraa/tetra/internal/auth"
	"github.com/ryanadiputraa/tetra/internal/errors"
	"github.com/ryanadiputraa/tetra/pkg/jwt"
	"github.com/ryanadiputraa/tetra/pkg/validator"
	"github.com/ryanadiputraa/tetra/pkg/writer"
)

type handler struct {
	writer    writer.HTTPWriter
	validator validator.Validator
	jwt       jwt.JWTService
	service   auth.AuthService
}

func New(writer writer.HTTPWriter, validator validator.Validator, jwt jwt.JWTService, service auth.AuthService) *handler {
	return &handler{
		writer:    writer,
		validator: validator,
		jwt:       jwt,
		service:   service,
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

		user, err := h.service.Login(r.Context(), p.Email, p.Password)
		if err != nil {
			if sErr, ok := err.(*errors.Error); ok {
				h.writer.WriteErrorResponse(w, errors.HttpErrMap[sErr.ErrCode], err.Error())
				return
			}
			h.writer.WriteErrorResponse(w, http.StatusInternalServerError, errors.ServerError)
			return
		}

		jwt, err := h.jwt.GenerateJWTWithClaims(user.ID)
		if err != nil {
			if sErr, ok := err.(*errors.Error); ok {
				h.writer.WriteErrorResponse(w, errors.HttpErrMap[sErr.ErrCode], err.Error())
				return
			}
			h.writer.WriteErrorResponse(w, http.StatusInternalServerError, errors.ServerError)
			return
		}

		h.writer.WriteResponseData(w, http.StatusOK, jwt)
	}
}

func (h *handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p auth.RegisterPayload
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			h.writer.WriteErrorResponse(w, http.StatusBadRequest, errors.BadRequest)
			return
		}

		errMap, err := h.validator.Validate(p)
		if err != nil {
			h.writer.WriteErrorResponseWithDetail(w, http.StatusBadRequest, errors.BadRequest, errMap)
			return
		}

		u, err := h.service.Register(r.Context(), p)
		if err != nil {
			if sErr, ok := err.(*errors.Error); ok {
				h.writer.WriteErrorResponse(w, errors.HttpErrMap[sErr.ErrCode], err.Error())
				return
			}
			h.writer.WriteErrorResponse(w, http.StatusInternalServerError, errors.ServerError)
			return
		}

		jwt, err := h.jwt.GenerateJWTWithClaims(u.ID)
		if err != nil {
			if sErr, ok := err.(*errors.Error); ok {
				h.writer.WriteErrorResponse(w, errors.HttpErrMap[sErr.ErrCode], err.Error())
				return
			}
			h.writer.WriteErrorResponse(w, http.StatusInternalServerError, errors.ServerError)
			return
		}

		h.writer.WriteResponseData(w, http.StatusOK, jwt)
	}
}
