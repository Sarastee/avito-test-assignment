package service

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/sarastee/avito-test-assignment/internal/model"
)

// AuthService interface for service layer
type AuthService interface {
	CreateUser(ctx context.Context, user model.CreateUser) (int64, error)
	VerifyUser(ctx context.Context, userAuth model.AuthUser) (*model.User, error)
}

// BannerService interface for service layer
type BannerService interface {
	CreateBanner(ctx context.Context, isActive bool, content json.RawMessage, featureID int64, tagIDs []int64) (int64, error)
	// UpdateBanner
	DeleteBanner(ctx context.Context, id int64) error
	// DeleteBanner

	// SelectRevision
	GetBannerFromDatabase(ctx context.Context, tagID int64, featureID int64, revisionID sql.NullInt64) (string, error)
	// GetAllBanners
	// GetAllRevisions

}

// JWTService interface for service layer
type JWTService interface {
	GenerateAccessToken(user model.User) (string, error)
	VerifyAccessToken(tokenStr string) (bool, error)
}

// BannerCacheService interface for service layer.
type BannerCacheService interface {
	SetCache(ctx context.Context, featureID int64, tagID int64, revisionID sql.NullInt64, content json.RawMessage) error
	GetCache(ctx context.Context, featureID int64, tagID int64, revisionID sql.NullInt64) (string, error)
}
