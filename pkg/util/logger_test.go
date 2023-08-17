package util

import "testing"

func TestGetLogger(t *testing.T) {
	logger1 := Getlogger("server1")
	logger2 := Getlogger("server2")
	logger1Again := Getlogger("server1")

	if logger1 == logger2 {
		t.Errorf("Expected different loggers for different system names")
	}

	if logger1 != logger1Again {
		t.Errorf("Expected the same logger for the same system name")
	}
}

func TestSugaredLogger(t *testing.T) {
	logger1 := GetSugaredLogger("server")
	if logger1 == nil {
		t.Errorf("Expected Sugared Logger not nil")
	}
}
