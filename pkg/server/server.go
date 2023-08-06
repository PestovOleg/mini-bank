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
	s.logger.Info("Start server")
	err := s.server.ListenAndServe()
	s.logger.Info("Server is started")

	if err != nil {
		s.logger.Error("Not error")
	}

	s.logger.Info("Server stopped")

	return nil
}

func (s *HTTPServer) Stop(ctx context.Context) {
	s.logger.Info("Start to stop")
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
		s.logger.Error("Shutdown failed- Timeout")
	case <-stop:
		if err != nil {
			s.logger.Error("one more error")
		} else {
			s.logger.Info("Shutdown finished")
		}
	}
	s.logger.Info("Finish to stop")
}
