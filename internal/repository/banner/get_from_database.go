package banner

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/sarastee/avito-test-assignment/internal/repository"
	"github.com/sarastee/platform_common/pkg/db"
)

// GetBannerFromDatabase function which get banner instance from database
func (r *Repo) GetBannerFromDatabase(ctx context.Context, tagID int64, featureID int64, revisionID sql.NullInt64) (string, error) {
	queryFormat := `
	SELECT rb.%s
	FROM %s
	INNER JOIN %s ON (%s.%s = %s.%s)
	LEFT JOIN %s AS rb ON (rb.%s = %s.%s)
	WHERE %s 
		AND rb.%s = COALESCE(@%s::bigint, %s.%s)
		AND %s = @%s and %s = @%s
		LIMIT 1
	`

	query := fmt.Sprintf(
		queryFormat,
		contentColumn,
		bannersTable,
		bannerRevisionTagsTable, bannerRevisionTagsTable, bannerIDColumn, bannersTable, idColumn,
		bannerRevisionsTable, bannerIDColumn, bannersTable, idColumn,
		isActiveColumn,
		revisionIDColumn, revisionIDColumn, bannersTable, selectedRevisionColumn,
		featureIDColumn, featureIDColumn, tagIDColumn, tagIDColumn,
	)
	q := db.Query{
		Name:     "banner_repository.GetBannerFromDatabase",
		QueryRaw: query,
	}
	var revisionIDParam sql.NullInt64
	if revisionID.Valid {
		revisionIDParam = revisionID
	} else {
		revisionIDParam.Valid = false
	}

	args := pgx.NamedArgs{
		revisionIDColumn: revisionIDParam,
		featureIDColumn:  featureID,
		tagIDColumn:      tagID,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	banner, err := pgx.CollectOneRow(rows, pgx.RowTo[string])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", repository.ErrBannerNotFound
		}

		return "", err
	}

	return banner, nil
}
