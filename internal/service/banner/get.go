package banner

import (
	"context"
	"database/sql"
	"fmt"
)

// GetBannerFromDatabase is Service layer function which get banner instance from database
func (s Service) GetBannerFromDatabase(ctx context.Context, tagID int64, featureID int64, revisionID sql.NullInt64) (string, error) {
	banner, err := s.bannerRepo.GetBannerFromDatabase(ctx, tagID, featureID, revisionID)
	if err != nil {
		s.logger.Err(err).Msg("failed attempt to get banner from database")
		return "", fmt.Errorf("failed attempt to get banner from database: %w", err)
	}

	return banner, nil
}
