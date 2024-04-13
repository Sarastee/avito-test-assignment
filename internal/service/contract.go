package service

import (
	"context"
	"encoding/json"

	"github.com/sarastee/avito-test-assignment/internal/model"
)

// BannerService interface for service layer
type BannerService interface {
	// GetBanner(ctx context.Context, tagID int64, featureID int64, isRoleAdmin bool) (string, error)
	// GetAllBanners
	// GetAllRevisions

	CreateBanner(ctx context.Context, isActive bool, content json.RawMessage, featureID int64, tagIDs []int64) (int64, error)
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
	CreateUser(ctx context.Context, user model.CreateUser) (int64, error)
	VerifyUser(ctx context.Context, userAuth model.AuthUser) (*model.User, error)
}
