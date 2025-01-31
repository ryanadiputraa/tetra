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
	writeHeader(w, code)
	if data == nil {
		data = map[string]string{"status": "ok"}
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		handleEncodingError(w)
	}
}

func (wr *httpWriter) WriteErrorResponse(w http.ResponseWriter, code int, message string) {
	writeHeader(w, code)
	err := json.NewEncoder(w).Encode(ErrorMessage{Message: message})
	if err != nil {
		handleEncodingError(w)
	}
}

func (wr *httpWriter) WriteErrorResponseWithDetail(w http.ResponseWriter, code int, message string, errors map[string]string) {
	writeHeader(w, code)
	err := json.NewEncoder(w).Encode(ErrorDetail{
		Message: message,
		Errors:  errors,
	})
	if err != nil {
		handleEncodingError(w)
	}
}

func writeHeader(w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
}

func handleEncodingError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(errors.ServerError))
}
