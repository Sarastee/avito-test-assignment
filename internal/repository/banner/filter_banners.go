package banner

import (
	"context"

	"github.com/sarastee/avito-test-assignment/internal/model"
	"github.com/sarastee/platform_common/pkg/db"
)

const (
	filterNullQuery = `
		SELECT banners.id, is_active, created_at, updated_at 
		FROM banners 
		LIMIT $1 OFFSET $2
	`

	filterNotNullQuery = `
		SELECT DISTINCT banners.id, is_active, created_at, updated_at
		FROM banners
		INNER JOIN banner_feature_tags as brt ON (brt.banner_id = banners.id)
		WHERE
		    (CASE WHEN $1::bigint IS NOT NULL THEN feature_id = $1 ELSE TRUE END)
		AND (CASE WHEN $2::bigint IS NOT NULL THEN tag_id = $2 ELSE TRUE END)
		LIMIT $3 OFFSET $4
	`
)

// FilterBanners function which get banners from database by provided filter
func (r *Repo) FilterBanners(ctx context.Context, bnrEntity *model.BannerInfo, offset int64, limit int64) ([]model.Banner, error) {
	args := []any{bnrEntity.FeatureID, bnrEntity.TagID, limit, offset}
	query := filterNotNullQuery

	if !bnrEntity.FeatureID.Valid && !bnrEntity.TagID.Valid {
		args = []any{limit, offset}
		query = filterNullQuery
	}

	q := db.Query{
		Name:     "banner_repository.FilterBanners",
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	banners := make([]model.Banner, 0)

	for rows.Next() {
		var filteredBanner model.Banner

		err = rows.Scan(
			&filteredBanner.ID,
			&filteredBanner.IsActive,
			&filteredBanner.CreatedAt,
			&filteredBanner.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		filteredBanner.TagIDs = make([]int64, 0)
		filteredBanner.Revisions = make([]model.Content, 0)

		banners = append(banners, filteredBanner)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return banners, nil
}
