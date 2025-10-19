package httputils

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

type HTTPTestConfig struct {
	Method      string
	Payload     *bytes.Buffer
	Path        string
	HandlerFunc http.HandlerFunc
}

func HTTPTestRequest(t *testing.T, config HTTPTestConfig) *httptest.ResponseRecorder {
	var req *http.Request
	if config.HandlerFunc == nil {
		t.Error(`handlerFunc not set yet`)
	}

	switch {
	case config.Payload == nil:
		req = httptest.NewRequest(config.Method, config.Path, nil)
	default:
		req = httptest.NewRequest(config.Method, config.Path, config.Payload)
	}

	response := httptest.NewRecorder()
	handler := config.HandlerFunc
	handler.ServeHTTP(response, req)
	return response
}
