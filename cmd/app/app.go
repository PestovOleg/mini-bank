package app

import (
	"context"
	"net/http"

	"github.com/PestovOleg/mini-bank/domain/user"
	"github.com/PestovOleg/mini-bank/domain/user/postgres"
	"github.com/PestovOleg/mini-bank/internal/common"
	"github.com/PestovOleg/mini-bank/internal/config"
	v1 "github.com/PestovOleg/mini-bank/internal/http/handler/v1"
	"github.com/PestovOleg/mini-bank/internal/http/middleware"
	"github.com/PestovOleg/mini-bank/pkg/server"
	"github.com/PestovOleg/mini-bank/pkg/util"
	"github.com/gorilla/mux"
	"golang.org/x/sys/unix"
)

type App struct {
	services *common.Services
	cfg      *config.AppConfig
	router   http.Handler
}

func NewRouter(s *common.Services) http.Handler {

	r := mux.NewRouter()
	subRouterV1 := r.PathPrefix("/v1").Subrouter()
	subRouterV1.Use(middleware.LoggingMiddleware)
	subRouterV1L := r.PathPrefix("/v1").Subrouter()
	subRouterV1L.Use(middleware.LoggingMiddleware)
	v1.SetHandler(subRouterV1, v1.BaseRoutes(s))
	v1.SetHandler(subRouterV1L, v1.BaseRoutesL(s))

	return r
}

func NewApp(cfg *config.AppConfig) App {
	conn := postgres.NewDBCon(
		cfg.PostgresDBConfig.User,
		cfg.PostgresDBConfig.Password,
		cfg.PostgresDBConfig.Host,
		cfg.PostgresDBConfig.Port,
		cfg.PostgresDBConfig.Name,
		cfg.PostgresDBConfig.SSLMode,
	)
	logger := util.GetLogger("APP")
	pgClient, err := postgres.GetDBCon(conn)
	if err != nil {
		logger.Fatal("Unexpected error with DB connection: " + err.Error())
	}
	userRepo := postgres.NewUserSQL(pgClient)
	userService := user.NewService(userRepo)

	s := common.NewServices(userService)

	api := NewRouter(s)

	return App{
		cfg:      cfg,
		services: s,
		router:   api,
	}
}

func (a *App) Run() error {
	srvCfg := server.HTTPServerConfig{
		Addr:              a.cfg.HTTPServerAppConfig.Addr,
		ReadTimeout:       a.cfg.HTTPServerAppConfig.ReadTimeout,
		WriteTimeout:      a.cfg.HTTPServerAppConfig.WriteTimeout,
		MaxHeadersBytes:   a.cfg.HTTPServerAppConfig.MaxHeadersBytes,
		ShutDownTime:      a.cfg.HTTPServerAppConfig.ShutDownTime,
		ReadHeaderTimeout: a.cfg.HTTPServerAppConfig.ReadHeaderTimeout,
		IdleTimeout:       a.cfg.HTTPServerAppConfig.IdleTimeout,
	}
	logger := util.GetLogger("server")
	ctx := util.NewSignalContextHandle(unix.SIGINT, unix.SIGTERM)
	errChan := make(chan error)

	server := server.NewServer(srvCfg, a.router)

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
	return nil
}
