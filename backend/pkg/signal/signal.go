package signal

import (
	"context"
	"os"
	"os/signal"
)

func NewSignalContextHandle(signals ...os.Signal) context.Context {
	var (
		stop        = make(chan os.Signal, 1)
		ctx, cancel = context.WithCancel(context.Background())
	)

	go func() {
		<-stop
		cancel()
	}()

	signal.Notify(stop, signals...)

	return ctx
}
