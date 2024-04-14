package banner

import (
	"context"
	"errors"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/sarastee/avito-test-assignment/internal/model"
	"github.com/sarastee/avito-test-assignment/internal/repository"
	"github.com/sarastee/platform_common/pkg/db"
)

const (
	updateFeatureQuery = `
		UPDATE banner_feature_tags 
		SET feature_id = $2 
		WHERE banner_id = $1
	`

	deleteFeatureTagsQuery = `
	WITH deleted_features AS (
		DELETE FROM banner_feature_tags 
		WHERE banner_id = $1 
		RETURNING feature_id
	)
	SELECT DISTINCT feature_id
	FROM deleted_features
	LIMIT 1
	`
)

// UpdateBannerInfo function which updates banner info in database
func (r *Repo) UpdateBannerInfo(ctx context.Context, bannerID int64, bnr *model.UpdateBannerSQL) error {
	switch {
	case !bnr.TagsIDs.Valid && bnr.FeatureID.Valid:
		q := db.Query{
			Name:     "banner_repository.UpdateBannerInfo",
			QueryRaw: updateFeatureQuery,
		}

		_, err := r.db.DB().ExecContext(ctx, q, bannerID, bnr.FeatureID.Int64)
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			r.logger.Debug().Msg(pgErr.Message)
			return repository.ErrBannerConflict
		}
	case bnr.TagsIDs.Valid:
		q := db.Query{
			Name:     "banner_repository.UpdateBannerInfo",
			QueryRaw: deleteFeatureTagsQuery,
		}

		var featureID int64
		featureIDRows, err := r.db.DB().QueryContext(ctx, q, bannerID)
		if err != nil {
			return err
		}
		defer featureIDRows.Close()

		featureID, err = pgx.CollectOneRow(featureIDRows, pgx.RowTo[int64])
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return repository.ErrBannerNotFound
			}

			return err
		}

		if bnr.FeatureID.Valid {
			featureID = bnr.FeatureID.Int64
		}

		err = r.LinkFeatureAndTags(ctx, bannerID, featureID, bnr.TagsIDs.V)
		if err != nil {
			return err
		}
	}

	return nil
}
