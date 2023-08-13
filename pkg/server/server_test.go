package server

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestHTTPServer_Run(t *testing.T) {
	config := Config{
		Addr:              ":3333",
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
		MaxHeadersBytes:   1000,
		ShutDownTime:      5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server test request")
	})

	t.Logf("Testing running HTTP server with config: %v", config)
	server := NewServer(config, handler)

	go func() {
		err := server.Run()
		if err != nil {
			t.Fatalf("Server run failed: %v", err)
		}
	}()

	defer func() {
		err := server.Stop(context.Background())
		if err != nil {
			t.Fatalf("Server stop failed: %v", err)
		}
	}()

	time.Sleep(1 * time.Second)

	resp, err := http.Get("http://localhost" + config.Addr)
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
}

func TestHTTPServerStop(t *testing.T) {
	config := Config{
		Addr:              ":3333",
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
		MaxHeadersBytes:   1000,
		ShutDownTime:      5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server test request")
	})

	server := NewServer(config, handler)
	t.Logf("HTTP server is running with config: %v", config)
	time.Sleep(1 * time.Second)
	t.Log("Testing stopping HTTP server")
	err := server.Stop(context.Background())
	if err != nil {
		t.Fatalf("Failed to stop server: %v", err)
	}

	_, err = http.Get("http://localhost" + config.Addr)
	if err == nil {
		t.Fatalf("Server should be stopped, but it is still accepting connections")
	}
}
