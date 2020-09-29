package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"io/ioutil"
	"encoding/json"

	"github.com/stretchr/testify/assert"
	"bflobox-api/internal/model"
)

// TestHealth validate /health returns "healthy"
func TestHealth(t *testing.T) {

	api := API{}

	request, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		panic(err.Error())
	}

	requestRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(api.Health)
	handler.ServeHTTP(requestRecorder, request)

	body, err := ioutil.ReadAll(requestRecorder.Body)

	if err != nil {
		panic(err.Error())
	}

	var healthResponse model.HealthResponse
	json.Unmarshal(body, &healthResponse)

	assert.Equal(t, http.StatusOK, requestRecorder.Code)
	assert.Equal(t, "healthy", healthResponse.Status)
}
