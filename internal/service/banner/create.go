package banner

import (
	"context"
	"encoding/json"
	"fmt"
)

// CreateBanner is Service layer function which process request
func (s Service) CreateBanner(ctx context.Context, isActive bool, content json.RawMessage, featureID int64, tagIDs []int64) (int64, error) {
	s.logger.Debug().Msg("attempt to create a banner")

	var bannerID int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var txErr error

		bannerID, txErr = s.bannerRepo.CreateBanner(ctx, isActive)
		if txErr != nil {
			s.logger.Err(txErr).Msg("failed attempt to create banner")
			return fmt.Errorf("failed attempt to create banner: %w", txErr)
		}

		txErr = s.bannerRepo.AddContent(ctx, bannerID, content)
		if txErr != nil {
			s.logger.Err(txErr).Msg("failed attempt to add content")
			return fmt.Errorf("failed attempt to add content: %w", txErr)
		}

		txErr = s.bannerRepo.LinkFeatureAndTags(ctx, bannerID, featureID, tagIDs)
		if txErr != nil {
			s.logger.Err(txErr).Msg("failed attempt to link feature and tags to banner")
			return fmt.Errorf("failed attempt to to link feature and tags to banner: %w", txErr)
		}

		return nil
	})
	if err != nil {
		return 0, err
	}

	return bannerID, nil
}
