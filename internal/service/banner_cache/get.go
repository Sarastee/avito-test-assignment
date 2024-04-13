package banner_cache

import (
	"context"
	"database/sql"
	"fmt"
)

func (b *BannerCacheService) GetCache(ctx context.Context, featureID int64, tagID int64, revisionID sql.NullInt64) (string, error) {
	key := fmt.Sprintf("%d-%d", featureID, tagID)
	if revisionID.Valid {
		key = fmt.Sprintf("%s-%d", key, revisionID.Int64)
	}

	return b.bannerCacheRepo.GetCache(ctx, key)
}
