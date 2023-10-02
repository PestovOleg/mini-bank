package app

import (
	"net/http"

	"github.com/PestovOleg/mini-bank/backend/pkg/config"
	"github.com/PestovOleg/mini-bank/backend/pkg/logger"
	"github.com/Unleash/unleash-client-go/v3"
)

func InitUnleash(cfg *config.AppConfig, service string) error {
	logger := logger.GetLogger(service)
	err := unleash.Initialize(
		unleash.WithListener(&unleash.DebugListener{}),
		unleash.WithAppName(cfg.UnleashServerConfig.AppName),
		unleash.WithUrl(cfg.UnleashServerConfig.URL),
		unleash.WithCustomHeaders(http.Header{"Authorization": {cfg.UnleashServerConfig.APIToken}}),
	)

	if err != nil {
		return err
	}

	logger.Info("Unleash connections is established")

	return nil
}
