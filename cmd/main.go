package main

import (
	"context"
	"fmt"
	"time"

	"github.com/PestovOleg/mini-bank/pkg/handler"
	"github.com/PestovOleg/mini-bank/pkg/server"
	"github.com/PestovOleg/mini-bank/pkg/util"
	"golang.org/x/sys/unix"
)

func main() {
	config := server.Config{
		Addr:              ":3333",
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
		MaxHeadersBytes:   1000,
		ShutDownTime:      5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}

	ctx := util.NewSignalContextHandle(unix.SIGINT, unix.SIGTERM)
	errChan := make(chan error)
	api := handler.NewRouter()
	server := server.NewServer(config, api)

	go func() {
		errChan <- server.Run()
	}()
	select {
	case err := <-errChan:
		fmt.Println(err)
	case <-ctx.Done():
		server.Stop(context.Background())
	}
}
