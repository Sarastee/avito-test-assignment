package env

import (
	"errors"
	"os"

	"github.com/sarastee/avito-test-assignment/internal/config"
)

const (
	passwordSalt = "PASSWORD_SALT"
)

// PasswordConfigSearcher password config searcher
type PasswordConfigSearcher struct{}

// NewPasswordConfigSearcher get instance for password config searcher.
func NewPasswordConfigSearcher() *PasswordConfigSearcher {
	return &PasswordConfigSearcher{}
}

// Get config for password
func (p *PasswordConfigSearcher) Get() (*config.PasswordConfig, error) {
	salt := os.Getenv(passwordSalt)
	if len(salt) == 0 {
		return nil, errors.New("salt for password not found")
	}

	return &config.PasswordConfig{
		PasswordSalt: salt,
	}, nil
}
