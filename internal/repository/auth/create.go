package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/sarastee/avito-test-assignment/internal/repository"
	"github.com/sarastee/platform_common/pkg/db"
)

// CreateUser function which creates new user and insert it in database
func (r *Repo) CreateUser(ctx context.Context, name string, passwordHash string, role string) (int64, error) {
	queryFormat := `
	INSERT INTO 
	    %s (%s, %s, %s) 
	VALUES 
		(@%s, @%s, @%s)
	RETURNING %s
	`
	query := fmt.Sprintf(
		queryFormat,
		usersTable,
		nameColum, passwordHashColumn, roleColumn,
		nameColum, passwordHashColumn, roleColumn,
		idColumn,
	)

	q := db.Query{
		Name:     "auth_repository.Register",
		QueryRaw: query,
	}

	args := pgx.NamedArgs{
		nameColum:          name,
		passwordHashColumn: passwordHash,
		roleColumn:         role,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	userID, err := pgx.CollectOneRow(rows, pgx.RowTo[int64])
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return 0, repository.ErrUserAlreadyRegistered
		}
		return 0, err
	}

	return userID, nil
}
