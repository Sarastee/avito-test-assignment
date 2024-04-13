package banner

import (
	"github.com/rs/zerolog"
	"github.com/sarastee/avito-test-assignment/internal/service"
)

const (
	BannerIDField = "id" // BannerIDField request param

	TagIDParam      = "tag_id"      // TagIDParam request param
	FeatureIDParam  = "feature_id"  // FeatureIDParam request param
	RevisionIDParam = "revision_id" // RevisionIDParam request param
	LimitParam      = "limit"       // LimitParam request param
	OffsetParam     = "offset"      // OffsetParam request param
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
