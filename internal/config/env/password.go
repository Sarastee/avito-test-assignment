package env

import (
	"errors"
	"os"

	"github.com/sarastee/avito-test-assignment/internal/config"
)

const (
	passwordSalt = "PASSWORD_SALT"
)

// PasswordCfgSearcher password config searcher
type PasswordCfgSearcher struct{}

// NewPasswordCfgSearcher get instance for password config searcher.
func NewPasswordCfgSearcher() *PasswordCfgSearcher {
	return &PasswordCfgSearcher{}
}

// Get config for password
func (p *PasswordCfgSearcher) Get() (*config.PasswordConfig, error) {
	salt := os.Getenv(passwordSalt)
	if len(salt) == 0 {
		return nil, errors.New("salt for password not found")
	}

	return &config.PasswordConfig{
		PasswordSalt: salt,
	}, nil
}
