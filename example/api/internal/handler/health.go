package handler

import (
	"bflobox-api/internal/model"
	"bflobox-api/internal/util"
	"net/http"
)

// Health return Generic API health
func (k *API) Health(w http.ResponseWriter, r *http.Request) {
	util.JSONResponse(w, http.StatusOK, model.HealthResponse{
		Status: "healthy",
	})
}
