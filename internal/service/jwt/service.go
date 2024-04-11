package jwt

import (
	"github.com/rs/zerolog"
	"github.com/sarastee/avito-test-assignment/internal/config"
	"github.com/sarastee/avito-test-assignment/internal/service"
)

var _ service.JWTService = (*Service)(nil)

// Service JWT Service
type Service struct {
	logger    *zerolog.Logger
	jwtConfig *config.JWTConfig
}

// NewService function which get new JWT Service instance
func NewService(logger *zerolog.Logger, jwtConfig *config.JWTConfig) *Service {
	return &Service{
		logger:    logger,
		jwtConfig: jwtConfig,
	}
}
