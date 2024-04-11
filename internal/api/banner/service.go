package banner

import (
	"github.com/rs/zerolog"
	"github.com/sarastee/avito-test-assignment/internal/service"
)

// Implementation for banner Server
type Implementation struct {
	logger        *zerolog.Logger
	bannerService service.BannerService
}

// NewImplementation function which get new Implementation instance
func NewImplementation(logger *zerolog.Logger, bannerService service.BannerService) *Implementation {
	return &Implementation{
		logger:        logger,
		bannerService: bannerService,
	}
}
