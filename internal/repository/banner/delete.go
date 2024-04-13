package banner

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/sarastee/avito-test-assignment/internal/repository"
	"github.com/sarastee/platform_common/pkg/db"
)

// DeleteBanner function which deletes banner from database
func (r *Repo) DeleteBanner(ctx context.Context, bannerID int64) error {
	queryFormat := `
	DELETE FROM 
	    %s 
	WHERE 
		%s 
	IN (SELECT DISTINCT 
	       %s 
	    FROM 
	       %s 
	    WHERE %s = @%s)
    `

	query := fmt.Sprintf(
		queryFormat,
		bannersTable,
		idColumn,
		bannerIDColumn,
		bannerRevisionTagsTable,
		bannerIDColumn, bannerIDColumn,
	)

	q := db.Query{
		Name:     "banner_repository.DeleteBanner",
		QueryRaw: query,
	}

	args := pgx.NamedArgs{
		bannerIDColumn: bannerID,
	}

	tag, err := r.db.DB().ExecContext(ctx, q, args)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return repository.ErrBannerNotFoundDelete
	}

	return nil
}
