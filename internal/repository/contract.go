package repository

import (
	"context"
	"encoding/json"
)

// BannerRepository interface for repository layer
type BannerRepository interface {
	GetBannerFromDatabase(ctx context.Context, tagID int64, featureID int64, isAdmin bool) (string, error)
	// GetAllBanners
	// GetAllRevisions

	CreateBanner(ctx context.Context, featureID int64, content json.RawMessage, isActive bool) (int64, int64, error)
	// UpdateBanner
	// DeleteBanner
	// DeleteBannerByID

	// SelectRevision
	LinkBannerAndTags(ctx context.Context, bannerID int64, tagIDs []int64) error
	InsertInSelectedVersions(ctx context.Context, bannerID int64, revisionID int64, featureID int64, tagIDs []int64) error
}
