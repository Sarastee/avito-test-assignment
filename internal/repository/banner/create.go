package banner

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/sarastee/avito-test-assignment/internal/repository"
	"github.com/sarastee/platform_common/pkg/db"
)

// CreateBanner function which creates new banner and insert it in database
func (r Repo) CreateBanner(ctx context.Context, featureID int64, content json.RawMessage, isActive bool) (int64, int64, error) {
	queryFormat := `
	INSERT INTO
		%s
	DEFAULT VALUES RETURNING
		banner_id
    `

	query := fmt.Sprintf(queryFormat, bannersTable)
	q := db.Query{
		Name:     "banner_repository.Create",
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q)
	if err != nil {
		return 0, 0, err
	}
	defer rows.Close()

	bannerID, err := pgx.CollectOneRow(rows, pgx.RowTo[int64])
	if err != nil {
		return 0, 0, err
	}

	queryFormat = `
	INSERT INTO 
	    %s (%s, %s, %s, %s) 
	VALUES 
		(@%s, @%s, @%s, @%s) 
	RETURNING revision_id
	`

	query = fmt.Sprintf(
		queryFormat,
		bannerRevisionsTable,
		bannerIDColumn, featureIDColumn, contentColumn, isActiveColumn,
		bannerIDColumn, featureIDColumn, contentColumn, isActiveColumn,
	)

	q = db.Query{
		Name:     "banner_repository.Create",
		QueryRaw: query,
	}

	args := pgx.NamedArgs{
		bannerIDColumn:  bannerID,
		featureIDColumn: featureID,
		contentColumn:   content,
		isActiveColumn:  isActive,
	}

	rows, err = r.db.DB().QueryContext(ctx, q, args)
	if err != nil {
		return 0, 0, err
	}

	revisionID, err := pgx.CollectOneRow(rows, pgx.RowTo[int64])
	if err != nil {
		return 0, 0, err
	}

	queryFormat = `
	UPDATE 
	    %s 
	SET 
	    %s = @%s 
	WHERE 
	    %s = @%s
	`

	query = fmt.Sprintf(
		queryFormat,
		bannersTable,
		selectedRevisionIDColumn, selectedRevisionIDColumn,
		bannerIDColumn, bannerIDColumn,
	)

	q = db.Query{
		Name:     "banner_repository.Create",
		QueryRaw: query,
	}

	args = pgx.NamedArgs{
		selectedRevisionIDColumn: revisionID,
		bannerIDColumn:           bannerID,
	}

	cmdTag, err := r.db.DB().ExecContext(ctx, q, args)
	if err != nil {
		return 0, 0, err
	}

	if cmdTag.RowsAffected() == 0 {
		return 0, 0, repository.ErrNoRowsAffected
	}

	return bannerID, revisionID, nil
}
