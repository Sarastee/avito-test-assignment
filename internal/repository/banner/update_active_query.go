package banner

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/sarastee/avito-test-assignment/internal/repository"
	"github.com/sarastee/platform_common/pkg/db"
)

// UpdateActiveQuery function which updates is_active Query in database
func (r *Repo) UpdateActiveQuery(ctx context.Context, bannerID int64, IsActive bool) (int64, error) {
	queryFormat := `
		UPDATE
			%s
		SET 
		    %s = @%s
		WHERE 
		    %s = @%s
		RETURNING 
			%s
	`

	query := fmt.Sprintf(
		queryFormat,
		bannersTable,
		isActiveColumn, isActiveColumn,
		idColumn, idColumn,
		idColumn,
	)

	q := db.Query{
		Name:     "banner_repository.UpdateActiveQuery",
		QueryRaw: query,
	}

	args := pgx.NamedArgs{
		isActiveColumn: IsActive,
		idColumn:       bannerID,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	bannerID, err = pgx.CollectOneRow(rows, pgx.RowTo[int64])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, repository.ErrBannerNotFound
		}

		return 0, err
	}

	return bannerID, err
}
