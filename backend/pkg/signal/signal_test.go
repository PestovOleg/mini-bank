package signal

import (
	"syscall"
	"testing"
	"time"
)

func TestNewSignalContextHandler(t *testing.T) {
	stt := syscall.SIGUSR1
	ctx := NewSignalContextHandle(stt)

	// проверяем не отменен ли контекст
	select {
	case <-ctx.Done():
		t.Fatalf("Context should not be cancelled")
	default:
	}

	if err := syscall.Kill(syscall.Getpid(), stt); err != nil {
		t.Fatalf("Failed to send signal: %v", err)
	}
	// ждем немного обработчик
	time.Sleep(time.Millisecond * 100)

	select {
	case <-ctx.Done():
		t.Logf("Context is cancelled as I wanted")
	default:
		t.Fatalf("Context should be cancelled")
	}
}
