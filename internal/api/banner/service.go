package banner

import (
	"github.com/rs/zerolog"
	"github.com/sarastee/avito-test-assignment/internal/service"
)

const (
	TagIDParam      = "tag_id"      // TagIDParam database param
	FeatureIDParam  = "feature_id"  // FeatureIDParam database param
	RevisionIDParam = "revision_id" // RevisionIDParam database param
	LimitParam      = "limit"       // LimitParam database param
	OffsetParam     = "offset"      // OffsetParam database param
)

// Implementation for banner Server
type Implementation struct {
	logger             *zerolog.Logger
	bannerService      service.BannerService
	bannerCacheService service.BannerCacheService
}

// NewImplementation function which get new Implementation instance
func NewImplementation(logger *zerolog.Logger, bannerService service.BannerService, bannerCacheService service.BannerCacheService) *Implementation {
	return &Implementation{
		logger:             logger,
		bannerService:      bannerService,
		bannerCacheService: bannerCacheService,
	}
}
