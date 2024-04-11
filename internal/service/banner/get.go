package banner

import (
	"context"
	"fmt"
)

// GetBanner function which get banner instance
func (s Service) GetBanner(ctx context.Context, tagID int64, featureID int64) (string, error) {
	banner, err := s.bannerRepo.GetBannerFromDatabase(ctx, tagID, featureID, true)
	if err != nil {
		s.logger.Err(err).Msg("failed attempt to get banner from database")
		return "", fmt.Errorf("failed attempt to get banner from database: %w", err)
	}

	return banner, nil
}
