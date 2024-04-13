package banner

import (
	"context"
	"errors"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/sarastee/avito-test-assignment/internal/repository"
)

// LinkFeatureAndTags functions which links banner, feature and tags and insert them in database
func (r *Repo) LinkFeatureAndTags(ctx context.Context, bannerID int64, featureID int64, tagIDs []int64) error {
	rows := make([][]interface{}, 0)
	for _, tagID := range tagIDs {
		rows = append(rows, []interface{}{bannerID, featureID, tagID})
	}

	_, err := r.db.DB().CopyFromContext(
		ctx,
		pgx.Identifier{bannerRevisionTagsTable},
		[]string{bannerIDColumn, featureIDColumn, tagIDColumn},
		pgx.CopyFromRows(rows))

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			r.logger.Debug().Msg(pgErr.Message)
			return repository.ErrTagsUniqueViolation
		}

		return err
	}

	return nil
}
