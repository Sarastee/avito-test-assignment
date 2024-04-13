package repository

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/sarastee/avito-test-assignment/internal/model"
)

// BannerRepository interface for repository layer
type BannerRepository interface {
	GetBannerFromDatabase(ctx context.Context, tagID int64, featureID int64, revisionID sql.NullInt64) (string, error)
	// GetAllBanners
	// GetAllRevisions

	CreateBanner(ctx context.Context, isActive bool) (int64, error)
	AddContent(ctx context.Context, bannerID int64, content json.RawMessage) error
	LinkFeatureAndTags(ctx context.Context, bannerID int64, featureID int64, tagIDs []int64) error

	// UpdateBanner
	// DeleteBanner
	// DeleteBannerByID

	// SelectRevision
}

// AuthRepository interface for repository layer
type AuthRepository interface {
	CreateUser(ctx context.Context, name string, passwordHash string, role string) (int64, error)
	GetUser(ctx context.Context, name string) (*model.User, error)
}

// BannerCacheRepository interface for repository layer.
type BannerCacheRepository interface {
	SetCache(ctx context.Context, key string, content string) error
	GetCache(ctx context.Context, key string) (string, error)
}
