package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/PestovOleg/mini-bank/internal/config"
	"github.com/PestovOleg/mini-bank/pkg/util"
	"go.uber.org/zap"
)

type HTTPServer struct {
	server *http.Server
	config config.HTTPServerConfig
	logger *zap.Logger
}

func NewServer(config config.HTTPServerConfig, handler http.Handler) *HTTPServer {
	server := &http.Server{
		Addr:              config.Addr,
		Handler:           handler,
		ReadTimeout:       config.ReadTimeout,
		WriteTimeout:      config.WriteTimeout,
		MaxHeaderBytes:    config.MaxHeadersBytes,
		IdleTimeout:       config.IdleTimeout,
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

	if err != nil && !errors.Is(err, http.ErrServerClosed) {
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
