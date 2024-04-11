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
func (r Repo) CreateUser(ctx context.Context, name string, passwordHash string, role string) error {
	queryFormat := `
	INSERT INTO 
	    %s (%s, %s, %s) 
	VALUES 
		(@%s, @%s, @%s)
	`
	query := fmt.Sprintf(
		queryFormat,
		usersTable,
		nameColum, passwordHashColumn, roleColumn,
		nameColum, passwordHashColumn, roleColumn,
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

	_, err := r.db.DB().ExecContext(ctx, q, args)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return repository.ErrUserAlreadyRegistered
		}
		return err
	}

	return nil
}
