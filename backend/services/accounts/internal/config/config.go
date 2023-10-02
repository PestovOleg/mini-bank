package config

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/PestovOleg/mini-bank/backend/pkg/logger"
	"github.com/ilyakaznacheev/cleanenv"
)

//nolint:gochecknoglobals
var once sync.Once

type AppConfig struct {
	Env                 string              `yaml:"env" env-default:"development"`
	HTTPServerAppConfig HTTPServerAppConfig `yaml:"http_server_app_config"`
	LoggerCfgs          []LoggerConfig      `yaml:"logger_cfgs"`
	PostgresDBConfig    PostgresDBConfig    `yaml:"postgres_db_config"`
	UnleashServerConfig UnleashServerConfig `yaml:"unleash_server_config"`
}

type HTTPServerAppConfig struct {
	Addr              string        `yaml:"addr" env-default:":8080"`
	ReadTimeout       time.Duration `yaml:"read_timeout" env-default:"5s"`
	WriteTimeout      time.Duration `yaml:"write_timeout" env-default:"5s"`
	MaxHeadersBytes   int           `yaml:"max_headers_bytes" env-default:"1000"`
	ShutDownTime      time.Duration `yaml:"shut_down_time" env-default:"5s"`
	ReadHeaderTimeout time.Duration `yaml:"read_header_timeout" env-default:"5s"`
	IdleTimeout       time.Duration `yaml:"idle_timeout" env-default:"120s"`
}

type LoggerConfig struct {
	Encoding string `yaml:"encoding" env-default:"console"`
	Output   string `yaml:"output" env-default:"stdout"`
	Level    string `yaml:"level" env-default:"debug"`
}

type PostgresDBConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host" env-default:"0.0.0.0"`
	Port     string `yaml:"port" env-default:"5432"`
	Name     string `yaml:"name"`
	SSLMode  string `yaml:"ssl_mode"`
}

type UnleashServerConfig struct {
	AppName  string `yaml:"app_name"`
	URL      string `yaml:"URL"`
	APIToken string `yaml:"api_token"`
}

func LoadConfig() AppConfig {
	var cfg AppConfig

	once.Do(func() {
		configPath := os.Getenv("CONFIG_PATH")
		if configPath == "" {
			log.Fatal("CONFIG_PATH environment variable is not set")
		}

		if _, err := os.Stat(configPath); err != nil {
			log.Fatalf("error while opening config file: %s", err)
		}

		err := cleanenv.ReadConfig(configPath, &cfg)
		if err != nil {
			log.Fatalf("error while reading config file: %s", err)
		}

		log.Println("Config is loaded")
	})

	return cfg
}

func (c *AppConfig) GetAllConfig() []logger.LogPathCfg {
	a := make([]logger.LogPathCfg, 0, len(c.LoggerCfgs))
	for _, loggerCfg := range c.LoggerCfgs {
		a = append(a, logger.NewLogPathCfg(loggerCfg.Encoding, loggerCfg.Output, loggerCfg.Level))
	}

	return a
}
