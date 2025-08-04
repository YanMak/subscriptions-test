package httpContractCommons

import (
	"fmt"
	"net/http"
	"strconv"
)

func ParsePagination(r *http.Request) (limit, offset int, err error) {
	const (
		defaultLimit = 20
		maxLimit     = 100
	)

	limitStr := r.URL.Query().Get(ParamLimit)
	if limitStr == "" {
		limit = defaultLimit
	} else {
		limit, err = strconv.Atoi(limitStr)
		if err != nil || limit < 1 {
			return 0, 0, fmt.Errorf("invalid limit")
		}
		if limit > maxLimit {
			limit = maxLimit
		}
	}

	offsetStr := r.URL.Query().Get(ParamOffset)
	if offsetStr == "" {
		offset = 0
	} else {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil || offset < 0 {
			return 0, 0, fmt.Errorf("invalid offset")
		}
	}

	return limit, offset, nil
}
