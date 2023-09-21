package server

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/PestovOleg/mini-bank/backend/pkg/logger"
)

func TestHTTPServer_Run(t *testing.T) {
	config := HTTPServerConfig{
		Addr:              ":3333",
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
		MaxHeadersBytes:   1000,
		ShutDownTime:      5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       120 * time.Second,
	}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server test request")
	})
	err := logger.InitMockLogger()

	if err != nil {
		t.Fatalf("Cannot initialize logger, %v", err)
	}

	t.Logf("Testing running HTTP server with config: %v", config)
	server := NewServer(config, handler)
	errCh := make(chan error, 1)

	go func() {
		err := server.Run()
		errCh <- err
	}()

	time.Sleep(1 * time.Second)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "http://localhost"+config.Addr, nil)
	if err != nil {
		t.Fatalf("Could not create GET request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		t.Fatalf("Could not make GET request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Could not read response body: %v", err)
	}

	expected := "Server test request"
	if string(body) != expected {
		t.Errorf("Expected %q, but got %q", expected, string(body))
	}

	err = server.Stop(context.Background())
	if err != nil {
		t.Fatalf("Server stop failed: %v", err)
	}

	err = <-errCh
	if err != nil {
		t.Fatalf("Server run failed: %v", err)
	}
}

func TestHTTPServerStop(t *testing.T) {
	config := HTTPServerConfig{
		Addr:              ":3333",
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
		MaxHeadersBytes:   1000,
		ShutDownTime:      5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       120 * time.Second,
	}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server test request")
	})

	err := logger.InitMockLogger()

	if err != nil {
		t.Fatalf("Cannot initialize logger, %v", err)
	}

	server := NewServer(config, handler)
	t.Logf("HTTP server is running with config: %v", config)
	time.Sleep(1 * time.Second)
	t.Log("Testing stopping HTTP server")

	err = server.Stop(context.Background())
	if err != nil {
		t.Fatalf("Failed to stop server: %v", err)
	}

	time.Sleep(1 * time.Second)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "http://localhost"+config.Addr, nil)
	if err != nil {
		t.Fatalf("Could not create GET request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)

	if err == nil {
		defer resp.Body.Close()
		t.Fatalf("Server should be stopped, but it is still accepting connections")
	} else {
		t.Logf("Server is stopped as expected: %v", err)
	}
}
