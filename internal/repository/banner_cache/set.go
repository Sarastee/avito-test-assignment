package banner_cache

import (
	"context"
)

func (r *BannerCacheRepo) SetCache(ctx context.Context, key string, content string) error {
	_, err := r.client.DB().DoContext(ctx, setCommand, key, content, exCommand, r.redisConfig.TTL.Seconds())
	if err != nil {
		return err
	}

	return nil
}
