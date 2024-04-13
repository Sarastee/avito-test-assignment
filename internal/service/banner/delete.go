package banner

import (
	"context"
	"fmt"
)

// DeleteBanner is Service layer function which process request
func (s *Service) DeleteBanner(ctx context.Context, id int64) error {
	err := s.bannerRepo.DeleteBanner(ctx, id)
	if err != nil {
		s.logger.Err(err).Msg("failed attempt to delete from database")
		return fmt.Errorf("failed attempt to delete banner: %w", err)
	}

	return nil
}
