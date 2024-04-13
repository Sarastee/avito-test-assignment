package banner

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sarastee/avito-test-assignment/internal/model"
)

// GetAdminBanners is Service layer function which get banner array from database
func (s *Service) GetAdminBanners(ctx context.Context, featureID sql.NullInt64, tagID sql.NullInt64, offset sql.NullInt64, limit sql.NullInt64) ([]model.Banner, error) {
	var entityOffset int64 = defaultOffset
	var entityLimit int64 = defaultLimit

	if offset.Valid {
		entityOffset = offset.Int64
	}

	if limit.Valid {
		entityLimit = limit.Int64
	}

	var banners []model.Banner
	bnrEntity := &model.BannerInfo{
		FeatureID: featureID,
		TagID:     tagID,
	}

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var txErr error
		banners, txErr = s.bannerRepo.FilterBanners(ctx, bnrEntity, entityOffset, entityLimit)
		if txErr != nil {
			s.logger.Err(txErr).Msg("failed attempt to filter banners")
			return fmt.Errorf("failed attempt to filter banner: %w", txErr)
		}

		if len(banners) == 0 {
			return nil
		}

		banners, txErr = s.bannerRepo.GetTagAndFeature(ctx, banners)
		if txErr != nil {
			s.logger.Err(txErr).Msg("failed attempt to filter banners")
			return fmt.Errorf("failed attempt to filter banner: %w", txErr)
		}

		banners, txErr = s.bannerRepo.GetContent(ctx, banners)
		if txErr != nil {
			s.logger.Err(txErr).Msg("failed attempt to filter banners")
			return fmt.Errorf("failed attempt to filter banner: %w", txErr)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return banners, nil
}
