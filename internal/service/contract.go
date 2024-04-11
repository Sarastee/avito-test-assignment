package service

import (
	"context"

	"github.com/sarastee/avito-test-assignment/internal/model"
)

// BannerService interface for service layer
type BannerService interface {
	GetBanner(ctx context.Context, tagID int64, featureID int64) (string, error)
	// GetAllBanners
	// GetAllRevisions

	CreateBanner(ctx context.Context, banner *model.Banner) (int64, error)
	// UpdateBanner
	// DeleteBanner
	// DeleteBannerByID

	// SelectRevision
}
