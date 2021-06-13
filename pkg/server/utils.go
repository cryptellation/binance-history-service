package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ErrorResponse will acknowledge informations on error
// swagger:response ErrorResponse
type ErrorResponse struct {
	// in:body
	Error string `json:"error"`
}

func wrapAsServerError(err error) string {
	return fmt.Sprint("[Server] ", err.Error())
}

func httpError(w http.ResponseWriter, httpCode int, sentErr error) {
	strError := wrapAsServerError(sentErr)

	// Try to write as JSON, if not send plain text
	errResp := ErrorResponse{Error: strError}
	jsonBytes, _ := json.Marshal(errResp)
	w.WriteHeader(httpCode)
	fmt.Fprintf(w, string(jsonBytes))
}

func httpResponse(w http.ResponseWriter, httpCode int, v interface{}) {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		httpError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(jsonBytes))
}
