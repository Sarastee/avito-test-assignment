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
	return string(passBytes), err
}

// CheckPasswordHash checks password with salt.
func (m *Manager) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(m.passWithSalt(password)))
	return err == nil
}

func (m *Manager) passWithSalt(password string) string {
	return password + m.passwordConfig.PasswordSalt
}
