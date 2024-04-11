package auth

import (
	"github.com/rs/zerolog"
	"github.com/sarastee/avito-test-assignment/internal/service"
)

// Implementation for auth Server
type Implementation struct {
	logger      *zerolog.Logger
	authService service.AuthService
	jwtService  service.JWTService
}

// NewImplementation function which get new Implementation instance
func NewImplementation(logger *zerolog.Logger, authService service.AuthService, jwtService service.JWTService) *Implementation {
	return &Implementation{
		logger:      logger,
		authService: authService,
		jwtService:  jwtService,
	}
}
