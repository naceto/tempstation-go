package rest

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
)

type ErrorResponse struct {
	Message    string `json:"messge"`
	StatusCode int    `json:"statusCode"`
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

func EncodeAndReturnError(logger *slog.Logger, w http.ResponseWriter, err error, headers map[string]string) int {
	status := http.StatusInternalServerError
	var body any
	var er *ErrorResponse
	if errors.As(err, &er) {
		status = er.StatusCode
		dup := *er
		dup.Message = err.Error()
		body = &dup
	} else {
		body = &ErrorResponse{
			Message:    err.Error(),
			StatusCode: status,
		}
	}

	if logger != nil && status >= http.StatusInternalServerError {
		logger.Error("API request failed", "error", err)
	}

	return EncodeAndReturn(w, status, body, headers)
}

func EncodeAndReturn(w http.ResponseWriter, statusCode int, body any, headers map[string]string) int {
	if !BodyAllowedForStatus(statusCode) {
		return Return(w, statusCode, headers)
	}

	switch b := body.(type) {
	case []byte:
		writeHeaders(w, statusCode, headers)
		_, _ = w.Write(b)
	case string:
		writeHeaders(w, statusCode, headers)
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write([]byte(b))
	default:
		w.Header().Set("Content-Type", "application/json")
		writeHeaders(w, statusCode, headers)
		_ = json.NewEncoder(w).Encode(body)
	}

	return statusCode
}

func BodyAllowedForStatus(statusCode int) bool {
	switch {
	case statusCode >= 100 && statusCode <= 199:
		return false
	case statusCode == http.StatusNoContent:
		return false
	case statusCode == http.StatusNotModified:
		return false
	}

	return true
}

func Return(w http.ResponseWriter, statusCode int, headers map[string]string) int {
	writeHeaders(w, statusCode, headers)
	return statusCode
}

func writeHeaders(w http.ResponseWriter, statusCode int, headers map[string]string) {
	for h, v := range headers {
		w.Header().Set(h, v)
	}
	w.WriteHeader(statusCode)
}
