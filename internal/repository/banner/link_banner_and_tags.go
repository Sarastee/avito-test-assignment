package banner

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// LinkBannerAndTags functions which links banner and tags and insert them in database
func (r Repo) LinkBannerAndTags(ctx context.Context, revisionID int64, tagIDs []int64) error {
	rows := make([][]interface{}, 0)
	for _, tagID := range tagIDs {
		rows = append(rows, []interface{}{revisionID, tagID})
	}

	_, err := r.db.DB().CopyFromContext(
		ctx,
		pgx.Identifier{bannerRevisionTagsTable},
		[]string{revisionIDColumn, tagIDColumn},
		pgx.CopyFromRows(rows))

	if err != nil {
		return err
	}

	return nil
}
