package password

import (
	"github.com/sarastee/avito-test-assignment/internal/config"
	"golang.org/x/crypto/bcrypt"
)

// Manager struct which working with passwords.
type Manager struct {
	passwordConfig *config.PasswordConfig
}

// NewManager gets manager instance.
func NewManager(passwordConfig *config.PasswordConfig) *Manager {
	return &Manager{passwordConfig: passwordConfig}
}

// HashPassword hashes password with salt.
func (m *Manager) HashPassword(password string) (string, error) {
	passBytes, err := bcrypt.GenerateFromPassword([]byte(m.passWithSalt(password)), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(passBytes), nil
}

// CheckPasswordHash checks password with salt.
func (m *Manager) CheckPasswordHash(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(m.passWithSalt(password)))
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) passWithSalt(password string) string {
	return password + m.passwordConfig.PasswordSalt
}
