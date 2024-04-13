package banner_cache

import (
	"github.com/sarastee/avito-test-assignment/internal/repository"
	"github.com/sarastee/avito-test-assignment/internal/service"
)

var _ service.BannerCacheService = (*BannerCacheService)(nil)

type BannerCacheService struct {
	bannerCacheRepo repository.BannerCacheRepository
}

func NewService(bannerCacheRepository repository.BannerCacheRepository) *BannerCacheService {
	return &BannerCacheService{
		bannerCacheRepo: bannerCacheRepository,
	}
}
