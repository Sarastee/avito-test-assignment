package middleware

import (
	"github.com/rs/zerolog"
	"github.com/sarastee/avito-test-assignment/internal/service"
)

// Middleware for auth Implementation
type Middleware struct {
	logger     *zerolog.Logger
	jwtService service.JWTService
}

// NewImplementation function which get new Implementation instance
func NewImplementation(logger *zerolog.Logger, jwtService service.JWTService) *Middleware {
	return &Middleware{
		logger:     logger,
		jwtService: jwtService,
	}
}
