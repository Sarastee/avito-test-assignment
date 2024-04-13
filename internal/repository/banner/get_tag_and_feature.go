package banner

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sarastee/avito-test-assignment/internal/model"
	"github.com/sarastee/platform_common/pkg/db"
)

const (
	getTagQuery = `
		SELECT banner_id, array_agg(tag_id), feature_id 
		FROM banner_revision_tags
		WHERE banner_id = ANY ($1::bigint[])
		GROUP BY banner_id, feature_id
	`
)

// GetTagAndFeature function which get tags and feature for banners from database
func (r *Repo) GetTagAndFeature(ctx context.Context, banners []model.Banner) ([]model.Banner, error) {
	q := db.Query{
		Name:     "banner_repository.GetTagAndFeature",
		QueryRaw: getTagQuery,
	}

	bannerIDs := make([]int64, len(banners))
	bannerIdxs := make(map[int64]int64)

	for index, bnr := range banners {
		bannerIDs[index] = bnr.ID
		bannerIdxs[bnr.ID] = int64(index)
	}

	rows, err := r.db.DB().QueryContext(ctx, q, bannerIDs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tags pgtype.Array[int64]

		var bannerID, featureID int64

		err = rows.Scan(
			&bannerID,
			&tags,
			&featureID,
		)
		if err != nil {
			return nil, err
		}

		banners[bannerIdxs[bannerID]].FeatureID = &featureID
		banners[bannerIdxs[bannerID]].TagIDs = []int64{}

		if tags.Valid {
			banners[bannerIdxs[bannerID]].TagIDs = tags.Elements
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return banners, nil
}
