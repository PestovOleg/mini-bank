package main

import (
	"context"

	"github.com/PestovOleg/mini-bank/internal/config"
	"github.com/PestovOleg/mini-bank/pkg/handler"
	"github.com/PestovOleg/mini-bank/pkg/server"
	"github.com/PestovOleg/mini-bank/pkg/util"
	"golang.org/x/sys/unix"
)

func main() {
	logger := *util.Getlogger("server")
	cfg := config.LoadConfig()

	ctx := util.NewSignalContextHandle(unix.SIGINT, unix.SIGTERM)
	errChan := make(chan error)
	api := handler.NewRouter()
	server := server.NewServer(cfg.HTTPServerConfig, api)

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
