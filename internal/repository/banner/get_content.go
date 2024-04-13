package banner

import (
	"context"

	"github.com/sarastee/avito-test-assignment/internal/model"
	"github.com/sarastee/platform_common/pkg/db"
)

const (
	getRevisionQuery = `
		SELECT banner_id, content, revision_id, created_at 
		FROM banner_revisions 
		WHERE banner_id = ANY ($1::bigint[])
	`
)

// GetContent function which get content for banners from database
func (r *Repo) GetContent(ctx context.Context, banners []model.Banner) ([]model.Banner, error) {
	q := db.Query{
		Name:     "banner_repository.GetContent",
		QueryRaw: getRevisionQuery,
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
		var bannerContent model.Content
		var bannerID int64

		err = rows.Scan(
			&bannerID,
			&bannerContent.Content,
			&bannerContent.Revision,
			&bannerContent.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		banners[bannerIdxs[bannerID]].Revisions = append(banners[bannerIdxs[bannerID]].Revisions, bannerContent)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return banners, nil
}
