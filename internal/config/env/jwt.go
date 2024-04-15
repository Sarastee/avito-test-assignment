package env

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/sarastee/avito-test-assignment/internal/config"
)

const (
	jwtKeyEnvName          = "JWT_SECRET_KEY"
	jwtAccessExpireEnvName = "JWT_ACCESS_TOKEN_EXPIRES_MIN"
)

// JWTCfgSearcher JWT config searcher
type JWTCfgSearcher struct{}

// NewJWTCfgSearcher get instance for JWT config searcher
func NewJWTCfgSearcher() *JWTCfgSearcher {
	return &JWTCfgSearcher{}
}

// Get config for JWT.
func (j *JWTCfgSearcher) Get() (*config.JWTConfig, error) {
	jwtSecret := os.Getenv(jwtKeyEnvName)
	if len(jwtSecret) == 0 {
		return nil, errors.New("secret for JWT not found")
	}

	accessExpireStr := os.Getenv(jwtAccessExpireEnvName)
	if len(accessExpireStr) == 0 {
		return nil, errors.New("jwt access token lifetime not found")
	}

	accessExpireInt, err := strconv.Atoi(accessExpireStr)
	if err != nil {
		return nil, errors.New("jwt access token incorrect format")
	}

	return &config.JWTConfig{
		JWTSecretKey:                jwtSecret,
		JWTAccessTokenExpireThrough: time.Duration(accessExpireInt) * time.Minute,
	}, nil
}
