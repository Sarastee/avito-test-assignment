package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/sarastee/avito-test-assignment/internal/model"
	"github.com/sarastee/avito-test-assignment/internal/repository"
	"github.com/sarastee/platform_common/pkg/db"
)

// GetUser function which get user instance from database
func (r *Repo) GetUser(ctx context.Context, name string) (*model.User, error) {
	// SELECT id, password_hash, role FROM users WHERE name = 'имя_пользователя';
	queryFormat := `
	SELECT 
	    %s, %s, %s 
	FROM 
	    %s 
	WHERE 
	    %s = @%s
	`

	query := fmt.Sprintf(
		queryFormat,
		idColumn, passwordHashColumn, roleColumn,
		usersTable,
		nameColum, nameColum,
	)

	q := db.Query{
		Name:     "auth_repository.GetUser",
		QueryRaw: query,
	}

	args := pgx.NamedArgs{
		nameColum: name,
	}

	var (
		userID       int64
		passwordHash string
		userRole     string
	)

	err := r.db.DB().QueryRowContext(ctx, q, args).Scan(&userID, &passwordHash, &userRole)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repository.ErrUserNotFound
		}

		return nil, err
	}

	userModel := model.User{
		ID:       userID,
		Name:     name,
		Role:     userRole,
		Password: passwordHash,
	}

	return &userModel, nil
}
