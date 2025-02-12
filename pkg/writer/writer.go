package writer

import (
	"encoding/json"
	"net/http"

	"github.com/ryanadiputraa/inventra/internal/errors"
)

type HTTPWriter interface {
	WriteResponseData(w http.ResponseWriter, code int, data any)
	WriteErrorResponse(w http.ResponseWriter, code int, message string)
	WriteErrorResponseWithDetail(w http.ResponseWriter, code int, message string, errors map[string]string)
}

type httpWriter struct{}

type ErrorMessage struct {
	Message string `json:"message"`
}

type ErrorDetail struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}

func NewHTTPWriter() HTTPWriter {
	return &httpWriter{}
}

func (wr *httpWriter) WriteResponseData(w http.ResponseWriter, code int, data any) {
	w.Header().Set("Content-Type", "application/json")
	if code != http.StatusNoContent && data == nil {
		data = map[string]string{"status": "ok"}
	}
	resp, err := json.Marshal(data)
	if err != nil {
		http.Error(w, errors.ServerError, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)
	w.Write(resp)
}

func (wr *httpWriter) WriteErrorResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(ErrorMessage{Message: message})
	if err != nil {
		http.Error(w, errors.ServerError, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)
	w.Write(resp)
}

func (wr *httpWriter) WriteErrorResponseWithDetail(w http.ResponseWriter, code int, message string, errors map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(ErrorDetail{
		Message: message,
		Errors:  errors,
	})
	if err != nil {
		http.Error(w, message, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)
	w.Write(resp)
}
