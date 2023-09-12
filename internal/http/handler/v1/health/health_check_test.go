package health

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PestovOleg/mini-bank/pkg/util"
)

func TestHealthCheckHandler_ServeHTTP(t *testing.T) {
	err := util.InitMockLogger()

	if err != nil {
		t.Fatalf("Cannot initialize logger, %v", err)
	}

	handler := NewHealthCheckHandler()

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/health", nil)
	if err != nil {
		t.Fatalf("Cannot make request, error: %v", err)
	}

	rr := httptest.NewRecorder()

	handlerFunc := http.HandlerFunc(handler.ServeHTTP)
	handlerFunc.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status: got %v want %v", status, http.StatusOK)
	}

	expected := "Service is healthy - Hello from Health Check Handler Endpoint"
	bodyBytes, err := io.ReadAll(rr.Body)

	if err != nil {
		t.Fatalf("Can't read Body, err: %v", err)
	}

	bodyString := string(bodyBytes)
	if bodyString != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", bodyString, expected)
	}
}
