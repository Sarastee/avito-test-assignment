package app

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/rs/zerolog"
	"github.com/sarastee/avito-test-assignment/internal/api/auth"
	"github.com/sarastee/avito-test-assignment/internal/api/banner"
	"github.com/sarastee/avito-test-assignment/internal/api/middleware"
	"github.com/sarastee/avito-test-assignment/internal/config"
	"github.com/sarastee/avito-test-assignment/internal/config/env"
	"github.com/sarastee/avito-test-assignment/internal/repository"
	authRepository "github.com/sarastee/avito-test-assignment/internal/repository/auth"
	bannerRepository "github.com/sarastee/avito-test-assignment/internal/repository/banner"
	bannerCacheRepository "github.com/sarastee/avito-test-assignment/internal/repository/banner_cache"
	"github.com/sarastee/avito-test-assignment/internal/service"
	authService "github.com/sarastee/avito-test-assignment/internal/service/auth"
	bannerService "github.com/sarastee/avito-test-assignment/internal/service/banner"
	bannerCacheService "github.com/sarastee/avito-test-assignment/internal/service/banner_cache"
	jwtService "github.com/sarastee/avito-test-assignment/internal/service/jwt"
	"github.com/sarastee/avito-test-assignment/internal/utils/password"
	"github.com/sarastee/platform_common/pkg/memory_db"
	"github.com/sarastee/platform_common/pkg/memory_db/rs"

	"github.com/sarastee/platform_common/pkg/closer"
	"github.com/sarastee/platform_common/pkg/db"
	"github.com/sarastee/platform_common/pkg/db/pg"
)

type serviceProvider struct {
	logger         *zerolog.Logger
	passManager    *password.Manager
	pgConfig       *config.PgConfig
	redisConfig    *config.RedisConfig
	httpConfig     *config.HTTPConfig
	swaggerConfig  *config.SwaggerConfig
	passwordConfig *config.PasswordConfig
	jwtConfig      *config.JWTConfig

	dbClient      db.Client
	txManager     db.TxManager
	redisDbClient memory_db.Client

	bannerCacheRepo repository.BannerCacheRepository
	bannerRepo      repository.BannerRepository
	authRepo        repository.AuthRepository

	bannerCacheService service.BannerCacheService
	bannerService      service.BannerService
	authService        service.AuthService
	jwtService         service.JWTService

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
		cfgSearcher := env.NewPasswordCfgSearcher()
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
		cfgSearcher := env.NewJWTCfgSearcher()
		cfg, err := cfgSearcher.Get()
		if err != nil {
			log.Fatalf("unable to get JWT config: %s", err.Error())
		}

		s.jwtConfig = cfg
	}

	return s.jwtConfig
}

func (s *serviceProvider) RedisConfig() *config.RedisConfig {
	if s.redisConfig == nil {
		cfgSearcher := env.NewRedisCfgSearcher()
		cfg, err := cfgSearcher.Get()
		if err != nil {
			log.Fatalf("unable to get Redis config:%s", err.Error())
		}

		s.redisConfig = cfg
	}

	return s.redisConfig
}

// SwaggerConfig ...
func (s *serviceProvider) SwaggerConfig() *config.SwaggerConfig {
	if s.swaggerConfig == nil {
		cfgSearcher := env.NewSwaggerCfgSearcher()
		cfg, err := cfgSearcher.Get()
		if err != nil {
			log.Fatalf("unable to get Swagger config: %s", err.Error())
		}

		s.swaggerConfig = cfg
	}

	return s.swaggerConfig
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

func (s *serviceProvider) RedisDBClient(_ context.Context) memory_db.Client {
	if s.redisDbClient == nil {
		redisConfig := s.RedisConfig()
		redisPool := &redis.Pool{
			MaxIdle:     redisConfig.MaxIdle,
			IdleTimeout: redisConfig.IdleTimeout,
			DialContext: func(ctx context.Context) (redis.Conn, error) {
				return redis.DialContext(ctx, "tcp", redisConfig.Address())
			},
			TestOnBorrowContext: func(_ context.Context, conn redis.Conn, lastUsed time.Time) error {
				if time.Since(lastUsed) < time.Minute {
					return nil
				}
				_, err := conn.Do("PING")
				return err
			},
		}
		s.redisDbClient = rs.New(redisPool)

		log.Printf("Redis connected at %s", redisConfig.Address())

		closer.Add(s.redisDbClient.Close)
	}

	return s.redisDbClient
}

// TxManager ...
func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = pg.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) PassManager() *password.Manager {
	if s.passManager == nil {
		s.passManager = password.NewManager(s.PasswordConfig())
	}

	return s.passManager
}

func (s *serviceProvider) BannerCacheRepository(ctx context.Context) repository.BannerCacheRepository {
	if s.bannerCacheRepo == nil {
		s.bannerCacheRepo = bannerCacheRepository.NewBannerCacheRepo(
			s.RedisDBClient(ctx),
			s.RedisConfig())
	}

	return s.bannerCacheRepo
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

func (s *serviceProvider) BannerCacheService(ctx context.Context) service.BannerCacheService {
	if s.bannerCacheService == nil {
		s.bannerCacheService = bannerCacheService.NewService(
			s.BannerCacheRepository(ctx))
	}

	return s.bannerCacheService
}

func (s *serviceProvider) BannerImpl(ctx context.Context) *banner.Implementation {
	if s.bannerImpl == nil {
		s.bannerImpl = banner.NewImplementation(
			s.Logger(),
			s.BannerService(ctx),
			s.BannerCacheService(ctx))
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
			s.PassManager())
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
