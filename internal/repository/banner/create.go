package banner

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/sarastee/platform_common/pkg/db"
)

// CreateBanner function which creates new banner and insert it in database
func (r *Repo) CreateBanner(ctx context.Context, isActive bool) (int64, error) {
	queryFormat := `
	INSERT INTO
		%s (%s)
	VALUES 
	    (@%s)
	RETURNING %s
    `

	query := fmt.Sprintf(
		queryFormat,
		bannersTable, isActiveColumn,
		isActiveColumn,
		idColumn,
	)

	q := db.Query{
		Name:     "banner_repository.CreateBanner",
		QueryRaw: query,
	}

	args := pgx.NamedArgs{
		isActiveColumn: isActive,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	bannerID, err := pgx.CollectOneRow(rows, pgx.RowTo[int64])
	if err != nil {
		return 0, err
	}

	return bannerID, nil
}
