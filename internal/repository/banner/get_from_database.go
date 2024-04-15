package banner

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/sarastee/avito-test-assignment/internal/model"
	"github.com/sarastee/avito-test-assignment/internal/repository"
	"github.com/sarastee/platform_common/pkg/db"
)

const (
	queryGet = `
	SELECT rb.revision_id, rb.created_at, rb.content
	FROM banners
	INNER JOIN banner_feature_tags ON (banner_feature_tags.banner_id = banners.id)
	LEFT JOIN banner_revisions AS rb ON (rb.banner_id = banners.id)
	WHERE is_active 
		AND rb.revision_id = COALESCE($3::bigint, banners.selected_revision)
		AND feature_id = $1 
	  	AND tag_id = $2
		LIMIT 1
	`
)

// GetBannerFromDatabase function which get banner instance from database
func (r *Repo) GetBannerFromDatabase(ctx context.Context, tagID int64, featureID int64, revisionID sql.NullInt64) (string, error) {
	q := db.Query{
		Name:     "banner_repository.GetBannerFromDatabase",
		QueryRaw: queryGet,
	}

	var revisionIDParam sql.NullInt64
	if revisionID.Valid {
		revisionIDParam = revisionID
	} else {
		revisionIDParam.Valid = false
	}

	rows, err := r.db.DB().QueryContext(ctx, q, featureID, tagID, revisionIDParam)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var bannerContent model.Content
	var hasRows bool
	for rows.Next() {
		hasRows = true
		err = rows.Scan(
			&bannerContent.Revision,
			&bannerContent.CreatedAt,
			&bannerContent.Content,
		)
		if err != nil {
			return "", err
		}
	}
	if !hasRows {
		return "", repository.ErrBannerNotFound
	}

	jsonContent, err := json.Marshal(bannerContent)
	if err != nil {
		return "", err
	}

	return string(jsonContent), nil
}
