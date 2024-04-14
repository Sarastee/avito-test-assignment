package banner

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/sarastee/avito-test-assignment/internal/repository"
	"github.com/sarastee/platform_common/pkg/db"
)

// GetBannerID function which gets banner id from database
func (r *Repo) GetBannerID(ctx context.Context, providedID int64) (int64, error) {
	queryFormat := `
		SELECT 
		    %s 
		FROM 
		    %s 
		WHERE 
		    %s = @%s
	`

	query := fmt.Sprintf(
		queryFormat,
		bannerIDColumn,
		bannerFeatureTagsTable,
		bannerIDColumn, bannerIDColumn,
	)

	q := db.Query{
		Name:     "banner_repository.GetBannerID",
		QueryRaw: query,
	}

	args := pgx.NamedArgs{
		bannerIDColumn: providedID,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	bannerID, err := pgx.CollectOneRow(rows, pgx.RowTo[int64])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, repository.ErrBannerNotFound
		}

		return 0, err
	}

	return bannerID, err
}
