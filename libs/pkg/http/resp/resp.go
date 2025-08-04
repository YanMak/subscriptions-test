package resp

import (
	"encoding/json"
	"net/http"

	httpContractCommons "project1.v0/contracts/transport/http"
	// "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// when writing go data
func Json(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// when writing json data
func JsonDirect(w http.ResponseWriter, jsonData []byte, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonData)
}

func WriteError(w http.ResponseWriter, msg string, statusCode int) {
	errMsg := msg
	if statusCode == http.StatusInternalServerError {
		errMsg = httpContractCommons.ErrInternalServerError.Error()
	}
	Json(w,
		httpContractCommons.ErrorResponse{
			Message: errMsg,
			Code:    statusCode,
		},
		statusCode,
	)
}
