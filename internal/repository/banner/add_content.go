package banner

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/sarastee/platform_common/pkg/db"
)

// AddContent function which adds revision in database
func (r Repo) AddContent(ctx context.Context, bannerID int64, content json.RawMessage) error {
	queryFormat := `
	INSERT INTO 
	    %s (%s, %s) 
	VALUES 
		(@%s, @%s)
	`

	query := fmt.Sprintf(
		queryFormat,
		bannerRevisionsTable, bannerIDColumn, contentColumn,
		bannerIDColumn, contentColumn,
	)

	q := db.Query{
		Name:     "banner_repository.AddContent",
		QueryRaw: query,
	}

	args := pgx.NamedArgs{
		bannerIDColumn: bannerID,
		contentColumn:  content,
	}

	_, err := r.db.DB().ExecContext(ctx, q, args)
	if err != nil {
		return err
	}

	return nil
}
