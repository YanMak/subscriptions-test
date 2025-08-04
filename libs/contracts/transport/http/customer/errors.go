package customerHttpContract

import (
	"errors"
	"net/http"

	customerDomainContract "project1.v0/contracts/domain/customer"
)

func HttpStatusForError(err error) int {
	switch {
	case errors.Is(err, customerDomainContract.ErrSubscriptionNotFound):
		return http.StatusNotFound // 404
	case errors.Is(err, customerDomainContract.ErrInvalidID) || errors.Is(err, customerDomainContract.ErrInvalidUserID):
		return http.StatusBadRequest // 400
	case err != nil:
		return http.StatusInternalServerError // 500 (по-умолчанию)
	default:
		return http.StatusOK
	}
}
