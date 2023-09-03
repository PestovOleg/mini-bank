package main

import (
	"context"

	"github.com/PestovOleg/mini-bank/internal/config"
	"github.com/PestovOleg/mini-bank/internal/handler"
	"github.com/PestovOleg/mini-bank/pkg/server"
	"github.com/PestovOleg/mini-bank/pkg/util"
	"golang.org/x/sys/unix"
)

func main() {
	cfg := config.LoadConfig()
	err := util.InitLogger(&cfg)

	if err != nil {
		panic("Logger cannot be initialized")
	}

	logger := util.GetLogger("server")
	srvCfg := server.HTTPServerConfig{
		Addr:              cfg.HTTPServerAppConfig.Addr,
		ReadTimeout:       cfg.HTTPServerAppConfig.ReadTimeout,
		WriteTimeout:      cfg.HTTPServerAppConfig.WriteTimeout,
		MaxHeadersBytes:   cfg.HTTPServerAppConfig.MaxHeadersBytes,
		ShutDownTime:      cfg.HTTPServerAppConfig.ShutDownTime,
		ReadHeaderTimeout: cfg.HTTPServerAppConfig.ReadHeaderTimeout,
		IdleTimeout:       cfg.HTTPServerAppConfig.IdleTimeout,
	}
	ctx := util.NewSignalContextHandle(unix.SIGINT, unix.SIGTERM)
	errChan := make(chan error)
	api := handler.NewRouter()
	server := server.NewServer(srvCfg, api)

	go func() {
		errChan <- server.Run()
	}()

	select {
	case err := <-errChan:
		logger.Error("There is error on server's side")
		logger.Error(err.Error())
	case <-ctx.Done():
		err := server.Stop(context.Background())
		if err != nil {
			logger.Fatal("Unexpected error: " + err.Error())
		}
	}
}
