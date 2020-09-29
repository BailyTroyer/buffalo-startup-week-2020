package util

import (
	"encoding/json"
	"net/http"
	"bflobox-api/internal/model"
)

// JSONResponse construct JSON response setting status with ResponseWriter and code, marshaling into output interface
func JSONResponse(w http.ResponseWriter, status int, output interface{}) {
	response, err := json.Marshal(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

// RespondError write error response given status code and message
func RespondError(w http.ResponseWriter, status int, message string) {
	JSONResponse(w, status, model.ErrorResponse{Error: message})
}
