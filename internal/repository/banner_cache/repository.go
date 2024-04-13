package banner_cache

import (
	"github.com/sarastee/avito-test-assignment/internal/repository"
	"github.com/sarastee/platform_common/pkg/memory_db"
)

const (
	setCommand = "SET"
	getCommand = "GET"
	exCommand  = "EX"
)

var _ repository.BannerCacheRepository = (*BannerCacheRepo)(nil)

type BannerCacheRepo struct {
	client memory_db.Client
}

func NewBannerCacheRepo(client memory_db.Client) *BannerCacheRepo {
	return &BannerCacheRepo{
		client: client,
	}
}
