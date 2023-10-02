package app

import (
	"context"
	"net/http"

	postgresConnect "github.com/PestovOleg/mini-bank/backend/pkg/database/postgres"
	"github.com/PestovOleg/mini-bank/backend/pkg/logger"
	"github.com/PestovOleg/mini-bank/backend/pkg/middleware"
	"github.com/PestovOleg/mini-bank/backend/pkg/server"
	"github.com/PestovOleg/mini-bank/backend/pkg/signal"
	"github.com/PestovOleg/mini-bank/backend/services/auth/internal/config"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"golang.org/x/sys/unix"
)

type Services struct {
	UserService    *user.Service
	AccountService *account.Service
}

func NewServices(u *user.Service, a *account.Service) *Services {
	return &Services{
		UserService:    u,
		AccountService: a,
	}
}

type App struct {
	services *Services
	cfg      *config.AppConfig
	router   http.Handler
}

func NewRouter(s *Services) http.Handler {
	r := mux.NewRouter()
	subRouterV1 := r.PathPrefix("/api/v1").Subrouter()
	subRouterV1.Use(middleware.LoggerMiddleware)

	subRouterV1L := r.PathPrefix("/api/v1").Subrouter()
	subRouterV1L.Use(middleware.LoggerMiddleware)
	subRouterV1L.Use(middleware.BasicAuthMiddleware(s.UserService))
	SetHandler(subRouterV1, BaseRoutes(s))
	SetHandler(subRouterV1L, BaseRoutesL(s))

	// Настраиваем CORS (как минимум для swagger'а)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // TODO: ограничить
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
	})

	// Оборачиваем роутер в CORS
	return c.Handler(r)
}

func NewApp(cfg *config.AppConfig) App {
	conn := postgresConnect.NewDBCon(
		cfg.PostgresDBConfig.User,
		cfg.PostgresDBConfig.Password,
		cfg.PostgresDBConfig.Host,
		cfg.PostgresDBConfig.Port,
		cfg.PostgresDBConfig.Name,
		cfg.PostgresDBConfig.SSLMode,
	)
	logger := logger.GetLogger("APP")

	pgClient, err := postgresConnect.GetDBCon(conn)
	if err != nil {
		logger.Fatal("Unexpected error with DB connection: " + err.Error())
	}

	logger.Info("DB connection is established")

	err = InitUnleash(cfg)
	if err != nil {
		logger.Sugar().Fatalf("Couldn't establish a connection to the Unleash Server: %s", err.Error())
		panic(err.Error())
	}

	userRepo := repoUser.NewUserSQL(pgClient)
	userService := user.NewService(userRepo)

	accountRepo := postgres.NewAccountSQL(pgClient)
	accountService := account.NewService(accountRepo)
	s := NewServices(userService, accountService)

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
	logger := logger.GetLogger("server")
	ctx := signal.NewSignalContextHandle(unix.SIGINT, unix.SIGTERM)
	errChan := make(chan error)

	server := server.NewServer(srvCfg, a.router)

	go func() {
		errChan <- server.Run()
	}()

	select {
	case err := <-errChan:
		logger.Error(err.Error())

		return err
	case <-ctx.Done():
		err := server.Stop(context.Background())
		if err != nil {
			logger.Fatal("Unexpected error: " + err.Error())
		}
	}

	return nil
}
