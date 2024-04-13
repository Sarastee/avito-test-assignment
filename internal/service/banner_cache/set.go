package banner_cache

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
)

func (b *BannerCacheService) SetCache(ctx context.Context, featureID int64, tagID int64, revisionID sql.NullInt64, content json.RawMessage) error {
	key := fmt.Sprintf("%d-%d", featureID, tagID)
	if revisionID.Valid {
		key = fmt.Sprintf("%s-%d", key, revisionID.Int64)
	}

	return b.bannerCacheRepo.SetCache(ctx, key, string(content))
}
