package banner

import (
	"context"
	"fmt"
)

// DeleteBanner is Service layer function which process request
func (s *Service) DeleteBanner(ctx context.Context, id int64) error {
	s.logger.Debug().Msg("attempt to delete a banner")

	err := s.bannerRepo.DeleteBanner(ctx, id)
	if err != nil {
		return fmt.Errorf("failed attempt to delete banner: %w", err)
	}

	return nil
}
