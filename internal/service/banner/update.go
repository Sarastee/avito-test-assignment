package banner

import (
	"context"
	"fmt"

	"github.com/sarastee/avito-test-assignment/internal/model"
)

// UpdateBanner is Service layer function which process request
func (s *Service) UpdateBanner(ctx context.Context, bannerID int64, bnr *model.UpdateBannerSQL) (int64, error) {
	s.logger.Debug().Msg("attempt to update a banner")

	var updatedID int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var txErr error

		updatedID, txErr = s.bannerRepo.GetBannerID(ctx, bannerID)
		if txErr != nil {
			return fmt.Errorf("failed attempt to update banner: %w", txErr)
		}

		if bnr.IsActive.Valid {
			updatedID, txErr = s.bannerRepo.UpdateActiveQuery(ctx, bannerID, bnr.IsActive.Bool)
			if txErr != nil {
				return fmt.Errorf("failed attempt to update banner: %w", txErr)
			}
		}

		if bnr.Content.Valid {
			txErr = s.bannerRepo.AddContent(ctx, bannerID, bnr.Content.V)
			if txErr != nil {
				return fmt.Errorf("failed attempt to update banner: %w", txErr)
			}
		}

		txErr = s.bannerRepo.UpdateBannerInfo(ctx, bannerID, bnr)
		if txErr != nil {
			return fmt.Errorf("failed attempt to update banner: %w", txErr)
		}

		return nil
	})
	if err != nil {
		return 0, err
	}

	return updatedID, nil
}
