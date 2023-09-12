package main

import (
	"github.com/PestovOleg/mini-bank/cmd/app"
	"github.com/PestovOleg/mini-bank/internal/config"
	"github.com/PestovOleg/mini-bank/pkg/util"
)

func main() {
	cfg := config.LoadConfig()
	err := util.InitLogger(&cfg)

	if err != nil {
		panic("Logger cannot be initialized")
	}

	//logger := util.GetLogger("server")
	app := app.NewApp(&cfg)
	app.Run()

}
