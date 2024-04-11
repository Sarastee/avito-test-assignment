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

// JWTService interface for service layer
type JWTService interface {
	GenerateAccessToken(user model.User) (string, error)
	VerifyAccessToken(tokenStr string) (bool, error)
}

// AuthService interface for service layer
type AuthService interface {
	CreateUser(ctx context.Context, user model.User) error
}
