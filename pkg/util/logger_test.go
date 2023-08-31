package util

import (
	"testing"
)

func TestGetLogger(t *testing.T) {
	err := InitMockLogger()

	if err != nil {
		t.Fatalf("Cannot initialize logger, %v", err)
	}

	logger1 := GetLogger("server1")
	logger2 := GetLogger("server2")
	logger1Again := GetLogger("server1")

	if logger1 == logger2 {
		t.Errorf("Expected different loggers for different system names")
	}

	if logger1 != logger1Again {
		t.Errorf("Expected the same logger for the same system name")
	}
}

func TestSugaredLogger(t *testing.T) {
	err := InitMockLogger()

	if err != nil {
		t.Fatalf("Cannot initialize logger, %v", err)
	}

	logger1 := GetSugaredLogger("server")
	if logger1 == nil {
		t.Errorf("Expected Sugared Logger not nil")
	}
}
