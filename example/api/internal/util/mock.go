package util

import "net/http"

// MockClient mock HTTP client
type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

var (
	// GetDoFunc fetch the mock client's `Do` func
	GetDoFunc func(req *http.Request) (*http.Response, error)
)

// Do Mock client's `Do` func allowing us to dependency inject this during testing
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return GetDoFunc(req)
}
