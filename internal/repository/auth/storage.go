package auth

import (
	"github.com/rs/zerolog"
	"github.com/sarastee/avito-test-assignment/internal/repository"
	"github.com/sarastee/platform_common/pkg/db"
)

const (
	usersTable = "users"

	nameColum          = "name"
	passwordHashColumn = "password_hash"
	roleColumn         = "role"
)

var _ repository.AuthRepository = (*Repo)(nil)

// Repo auth repository for authentication
type Repo struct {
	logger *zerolog.Logger
	db     db.Client
}

// NewRepo function which get new repo instance
func NewRepo(logger *zerolog.Logger, dbClient db.Client) *Repo {
	return &Repo{
		logger: logger,
		db:     dbClient,
	}
}
