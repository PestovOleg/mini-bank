package server

import (
	"context"
	"net/http"
	"time"

	"github.com/PestovOleg/mini-bank/pkg/util"
	"go.uber.org/zap"
)

type Config struct {
	Addr              string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	MaxHeadersBytes   int
	ShutDownTime      time.Duration
	ReadHeaderTimeout time.Duration
}

type HTTPServer struct {
	server *http.Server
	config Config
	logger *zap.Logger
}

func NewServer(config Config, handler http.Handler) *HTTPServer {
	server := &http.Server{
		Addr:              config.Addr,
		Handler:           handler,
		ReadTimeout:       config.ReadTimeout,
		WriteTimeout:      config.WriteTimeout,
		MaxHeaderBytes:    config.MaxHeadersBytes,
		IdleTimeout:       120 * time.Second,
		ReadHeaderTimeout: config.ReadHeaderTimeout,
	}

	return &HTTPServer{
		server: server,
		config: config,
		logger: util.Getlogger("server"),
	}
}

func (s *HTTPServer) Run() error {
	s.logger.Info("Start server on " + s.server.Addr)
	err := s.server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		s.logger.Error("Server error " + err.Error())
	}

	return nil
}

func (s *HTTPServer) Stop(ctx context.Context) error {
	s.logger.Info("Initiating server shutdown.")
	cancelCtx, cancel := context.WithTimeout(context.Background(), s.config.ShutDownTime)

	var err error

	defer cancel()

	stop := make(chan struct{})

	go func() {
		err = s.server.Shutdown(ctx)

		close(stop)
	}()

	select {
	case <-cancelCtx.Done():
		s.logger.Error("Shutdown failed: Timeout.")
	case <-stop:
		if err != nil {
			s.logger.Error("An error occurred while stopping server: ", zap.Error(err))
		} else {
			s.logger.Info("Shutdown finished")
		}
	}
	s.logger.Info("Server is stopped")
	return nil
}
