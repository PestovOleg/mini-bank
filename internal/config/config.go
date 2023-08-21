package config

import (
	"os"
	"time"

	"github.com/PestovOleg/mini-bank/pkg/util"
	"github.com/ilyakaznacheev/cleanenv"
)

type AppConfig struct {
	Env              string           `yaml:"env" env-default:"development"`
	HTTPServerConfig HTTPServerConfig `yaml:"http_server"`
}

type HTTPServerConfig struct {
	Addr              string        `yaml:"address" env-default:":8080"`
	ReadTimeout       time.Duration `yaml:"read_timeout" env-default:"5s"`
	WriteTimeout      time.Duration `yaml:"write_timeout" env-default:"5s"`
	MaxHeadersBytes   int           `yaml:"max_header_bytes" env-default:"1000"`
	ShutDownTime      time.Duration `yaml:"shutdown_time" env-default:"5s"`
	ReadHeaderTimeout time.Duration `yaml:"read_header_timeout" env-default:"5s"`
	IdleTimeout       time.Duration `yaml:"idle_timeout" env-default:"120s"`
}

func LoadConfig() AppConfig {
	logger := util.Getlogger("config")

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		logger.Fatal("CONFIG_PATH environment variable is not set")
	}

	if _, err := os.Stat(configPath); err != nil {
		logger.Sugar().Fatalf("error opening config file (loading HTTPserver): %s", err)
	}

	var cfg AppConfig

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		logger.Sugar().Fatalf("error reading config file (loading HTTPserver): %s", err)
	}

	logger.Info("Config is loaded")

	return cfg
}
