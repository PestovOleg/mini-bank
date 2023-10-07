// @title           minibank
// @version         1.0
// @description     minibank.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
package main

import (
	"github.com/PestovOleg/mini-bank/backend/pkg/config"
	"github.com/PestovOleg/mini-bank/backend/pkg/logger"
	"github.com/PestovOleg/mini-bank/backend/services/account/cmd/app"
)

func main() {
	cfg := config.LoadConfig()
	err := logger.InitLogger(&cfg)

	if err != nil {
		panic("Logger cannot be initialized")
	}

	logger := logger.GetLogger("APP")
	app := app.NewApp(&cfg)

	err = app.Run()
	if err != nil {
		logger.Error("There is error on server's side:" + err.Error())
	}
}
