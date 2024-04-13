package banner_cache

import (
	"context"
	"errors"

	"github.com/gomodule/redigo/redis"
	"github.com/sarastee/avito-test-assignment/internal/repository"
)

func (r *BannerCacheRepo) GetCache(ctx context.Context, key string) (string, error) {
	db := r.client.DB()
	content, err := db.String(db.DoContext(ctx, getCommand, key))
	if err != nil {
		if errors.Is(err, redis.ErrNil) {
			return "", repository.ErrCacheNotFound
		}

		return "", err
	}

	return content, nil
}
