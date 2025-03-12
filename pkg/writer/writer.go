package writer

import (
	"encoding/json"
	"net/http"

	"github.com/ryanadiputraa/inventra/internal/errors"
)

type HTTPWriter interface {
	WriteResponseData(w http.ResponseWriter, code int, data any)
	WriteResponseDataWithPagination(w http.ResponseWriter, code int, data any, dataKey string, page, size int, count int64)
	WriteErrorResponse(w http.ResponseWriter, code int, message string)
	WriteErrorResponseWithDetail(w http.ResponseWriter, code int, message string, errors map[string]string)
}

type httpWriter struct{}

type Pagination struct {
	CurrentPage int   `json:"current_page"`
	TotalPages  int   `json:"total_pages"`
	Size        int   `json:"size"`
	TotalData   int64 `json:"total_data"`
}

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

func (wr *httpWriter) WriteResponseDataWithPagination(w http.ResponseWriter, code int, data any, dataKey string, page, size int, total int64) {
	w.Header().Set("Content-Type", "application/json")
	m := Pagination{
		CurrentPage: page,
		TotalPages:  int((total + int64(size) - 1) / int64(size)),
		Size:        size,
		TotalData:   total,
	}

	resp, err := json.Marshal(map[string]any{
		dataKey: data,
		"meta":  m,
	})
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
