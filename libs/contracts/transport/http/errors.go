package httpContractCommons

import (
	"errors"
	"net/http"
)

// ErrorResponse — commonly used struct for err
// // @Description error response body
// // @name ErrorResponse
type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`
}

var ErrLimitExceeded = errors.New("too many requests")
var ErrRequestValidationError = errors.New("request validation error")
var ErrInternalServerError = errors.New("internal server error")
var ErrInvalidPaginationParams = errors.New("invalid pagination params")

const (
	HttpStatusForValidationError int = 400
)

func HttpStatusForError(err error) int {
	switch {
	case errors.Is(err, ErrRequestValidationError):
		return http.StatusBadRequest // 400
	case errors.Is(err, ErrLimitExceeded):
		return http.StatusTooManyRequests // 429
	case err != nil:
		return http.StatusInternalServerError // 500
	default:
		return http.StatusOK
	}
}
