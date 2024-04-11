package banner

import (
	"github.com/rs/zerolog"
	"github.com/sarastee/avito-test-assignment/internal/repository"
	"github.com/sarastee/avito-test-assignment/internal/service"
	"github.com/sarastee/platform_common/pkg/db"
)

var _ service.BannerService = (*Service)(nil)

// Service banner Service
type Service struct {
	logger     *zerolog.Logger
	txManager  db.TxManager
	bannerRepo repository.BannerRepository
}

// NewService function which get new Service instance
func NewService(logger *zerolog.Logger, txManager db.TxManager, bannerRepository repository.BannerRepository) *Service {
	return &Service{
		logger:     logger,
		txManager:  txManager,
		bannerRepo: bannerRepository,
	}
}
