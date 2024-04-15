package banner

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
)

// GetBannerFromDatabase is Service layer function which get banner instance from database
func (s *Service) GetBannerFromDatabase(ctx context.Context, tagID int64, featureID int64, revisionID sql.NullInt64) (json.RawMessage, error) {
	s.logger.Debug().Msg("attempt to get banner from database")

	banner, err := s.bannerRepo.GetBannerFromDatabase(ctx, tagID, featureID, revisionID)
	if err != nil {
		return nil, fmt.Errorf("failed attempt to get banner from database: %w", err)
	}

	return json.RawMessage(banner), nil
}
