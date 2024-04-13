package app

import (
	"context"
	"log"
	"os"

	"github.com/rs/zerolog"
	"github.com/sarastee/avito-test-assignment/internal/api/auth"
	"github.com/sarastee/avito-test-assignment/internal/api/banner"
	"github.com/sarastee/avito-test-assignment/internal/api/middleware"
	"github.com/sarastee/avito-test-assignment/internal/config"
	"github.com/sarastee/avito-test-assignment/internal/config/env"
	"github.com/sarastee/avito-test-assignment/internal/repository"
	authRepository "github.com/sarastee/avito-test-assignment/internal/repository/auth"
	bannerRepository "github.com/sarastee/avito-test-assignment/internal/repository/banner"
	"github.com/sarastee/avito-test-assignment/internal/service"
	authService "github.com/sarastee/avito-test-assignment/internal/service/auth"
	bannerService "github.com/sarastee/avito-test-assignment/internal/service/banner"
	jwtService "github.com/sarastee/avito-test-assignment/internal/service/jwt"
	"github.com/sarastee/avito-test-assignment/internal/utils/password"

	"github.com/sarastee/platform_common/pkg/closer"
	"github.com/sarastee/platform_common/pkg/db"
	"github.com/sarastee/platform_common/pkg/db/pg"
)

type serviceProvider struct {
	logger         *zerolog.Logger
	passManager    *password.Manager
	pgConfig       *config.PgConfig
	httpConfig     *config.HTTPConfig
	passwordConfig *config.PasswordConfig
	jwtConfig      *config.JWTConfig

	dbClient  db.Client
	txManager db.TxManager

	bannerRepo repository.BannerRepository
	authRepo   repository.AuthRepository

	bannerService service.BannerService
	authService   service.AuthService
	jwtService    service.JWTService

	bannerImpl *banner.Implementation
	authImpl   *auth.Implementation

	middleware *middleware.Middleware
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// Logger ...
func (s *serviceProvider) Logger() *zerolog.Logger {
	if s.logger == nil {
		cfgSearcher := env.NewLogCfgSearcher()
		cfg, err := cfgSearcher.Get()
		if err != nil {
			log.Fatalf("unable to get Logger config: %s", err.Error())
		}

		s.logger = setupZeroLog(cfg)
	}

	return s.logger
}

func setupZeroLog(logConfig *config.LogConfig) *zerolog.Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: logConfig.TimeFormat}
	logger := zerolog.New(output).With().Timestamp().Logger()
	logger = logger.Level(logConfig.LogLevel)
	zerolog.TimeFieldFormat = logConfig.TimeFormat

	return &logger
}

// PgConfig ...
func (s *serviceProvider) PgConfig() *config.PgConfig {
	if s.pgConfig == nil {
		cfgSearcher := env.NewPgCfgSearcher()
		cfg, err := cfgSearcher.Get()
		if err != nil {
			log.Fatalf("unable to get PG config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

// HTTPConfig ...
func (s *serviceProvider) HTTPConfig() *config.HTTPConfig {
	if s.httpConfig == nil {
		cfgSearcher := env.NewHTTPCfgSearcher()
		cfg, err := cfgSearcher.Get()
		if err != nil {
			log.Fatalf("unable to get HTTP config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) PasswordConfig() *config.PasswordConfig {
	if s.passwordConfig == nil {
		cfgSearcher := env.NewPasswordConfigSearcher()
		cfg, err := cfgSearcher.Get()
		if err != nil {
			log.Fatalf("unable to get Password config: %s", err.Error())
		}

		s.passwordConfig = cfg
	}

	return s.passwordConfig
}

func (s *serviceProvider) JWTConfig() *config.JWTConfig {
	if s.jwtConfig == nil {
		cfgSearcher := env.NewJWTConfigSearcher()
		cfg, err := cfgSearcher.Get()
		if err != nil {
			log.Fatalf("unable to get JWT config: %s", err.Error())
		}

		s.jwtConfig = cfg
	}

	return s.jwtConfig
}

// DBClient ...
func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PgConfig().DSN(), s.Logger())
		if err != nil {
			log.Fatalf("failure while creating DB: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("no connection to DB: %s", err.Error())
		}
		closer.Add(cl.Close)

		log.Printf("DB connected at %s:%d/%s", s.PgConfig().Host, s.PgConfig().Port, s.PgConfig().DbName)

		s.dbClient = cl
	}

	return s.dbClient
}

// TxManager ...
func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = pg.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) PasswordManager() *password.Manager {
	if s.passManager == nil {
		s.passManager = password.NewManager(s.PasswordConfig())
	}

	return s.passManager
}

func (s *serviceProvider) BannerRepository(ctx context.Context) repository.BannerRepository {
	if s.bannerRepo == nil {
		s.bannerRepo = bannerRepository.NewRepo(s.Logger(), s.DBClient(ctx))
	}

	return s.bannerRepo
}

func (s *serviceProvider) BannerService(ctx context.Context) service.BannerService {
	if s.bannerService == nil {
		s.bannerService = bannerService.NewService(
			s.Logger(),
			s.TxManager(ctx),
			s.BannerRepository(ctx))
	}

	return s.bannerService
}

func (s *serviceProvider) BannerImpl(ctx context.Context) *banner.Implementation {
	if s.bannerImpl == nil {
		s.bannerImpl = banner.NewImplementation(s.Logger(), s.BannerService(ctx))
	}

	return s.bannerImpl
}

func (s *serviceProvider) AuthRepository(ctx context.Context) repository.AuthRepository {
	if s.authRepo == nil {
		s.authRepo = authRepository.NewRepo(s.Logger(), s.DBClient(ctx))
	}

	return s.authRepo
}

func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = authService.NewService(
			s.Logger(),
			s.TxManager(ctx),
			s.AuthRepository(ctx),
			s.PasswordManager())
	}

	return s.authService
}

func (s *serviceProvider) JWTService() service.JWTService {
	if s.jwtService == nil {
		s.jwtService = jwtService.NewService(
			s.Logger(),
			s.JWTConfig())
	}

	return s.jwtService
}

func (s *serviceProvider) AuthImpl(ctx context.Context) *auth.Implementation {
	if s.authImpl == nil {
		s.authImpl = auth.NewImplementation(
			s.Logger(),
			s.AuthService(ctx),
			s.JWTService())
	}

	return s.authImpl
}

func (s *serviceProvider) Middleware() *middleware.Middleware {
	if s.middleware == nil {
		s.middleware = middleware.NewImplementation(
			s.Logger(),
			s.JWTService())
	}

	return s.middleware
}
