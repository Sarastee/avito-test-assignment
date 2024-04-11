package app

import (
	"context"
	"log"
	"os"

	"github.com/rs/zerolog"
	"github.com/sarastee/avito-test-assignment/internal/api/banner"
	"github.com/sarastee/avito-test-assignment/internal/config"
	"github.com/sarastee/avito-test-assignment/internal/config/env"
	"github.com/sarastee/avito-test-assignment/internal/repository"
	bannerRepository "github.com/sarastee/avito-test-assignment/internal/repository/banner"
	"github.com/sarastee/avito-test-assignment/internal/service"
	bannerService "github.com/sarastee/avito-test-assignment/internal/service/banner"
	"github.com/sarastee/platform_common/pkg/closer"
	"github.com/sarastee/platform_common/pkg/db"
	"github.com/sarastee/platform_common/pkg/db/pg"
)

type serviceProvider struct {
	logger     *zerolog.Logger
	pgConfig   *config.PgConfig
	httpConfig *config.HTTPConfig

	dbClient  db.Client
	txManager db.TxManager

	bannerRepo repository.BannerRepository
	// authRepo

	bannerService service.BannerService
	// authService

	bannerImpl *banner.Implementation
	// authImpl
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
