package model

// HealthResponse Generic API health
type HealthResponse struct {
	Status string `json:"status"`
}

// ErrorResponse Generic API error
type ErrorResponse struct {
	Error string `json:"error"`
}
