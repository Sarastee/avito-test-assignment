package banner

import (
	"context"
	"fmt"

	"github.com/sarastee/avito-test-assignment/internal/model"
)

// CreateBanner is Service layer function which process request
func (s Service) CreateBanner(ctx context.Context, banner *model.Banner) (int64, error) {
	s.logger.Debug().Msg("attempt to create a banner")

	var bannerID, revisionID int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var txErr error

		bannerID, revisionID, txErr = s.bannerRepo.CreateBanner(ctx, *banner.FeatureID, banner.Content, *banner.IsActive)
		if txErr != nil {
			s.logger.Err(txErr).Msg("failed attempt to create banner")
			return fmt.Errorf("failed attempt to create banner: %w", txErr)
		}

		txErr = s.bannerRepo.LinkBannerAndTags(ctx, revisionID, banner.TagIDs)
		if txErr != nil {
			s.logger.Err(txErr).Msg("failed attempt to link banner and tags")
			return fmt.Errorf("failed attempt to link banner and tags: %w", txErr)
		}

		txErr = s.bannerRepo.InsertInSelectedVersions(ctx, bannerID, revisionID, *banner.FeatureID, banner.TagIDs)
		if txErr != nil {
			s.logger.Err(txErr).Msg("failed attempt to insert in selected versions")
			return fmt.Errorf("failed attempt to insert in selected versions: %w", txErr)
		}

		return nil
	})
	if err != nil {
		return 0, err
	}

	return bannerID, nil
}
