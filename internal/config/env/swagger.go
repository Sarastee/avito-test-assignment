package env

import (
	"os"

	"github.com/pkg/errors"
	"github.com/sarastee/avito-test-assignment/internal/config"
)

const (
	swaggerHostName    = "SWAGGER_HOST"
	swaggerPortEnvName = "SWAGGER_PORT"
)

// SwaggerCfgSearcher searcher for Swagger config.
type SwaggerCfgSearcher struct{}

// NewSwaggerCfgSearcher get instance for Swagger config searcher.
func NewSwaggerCfgSearcher() *SwaggerCfgSearcher {
	return &SwaggerCfgSearcher{}
}

// Get searcher for Swagger config.
func (s *SwaggerCfgSearcher) Get() (*config.SwaggerConfig, error) {
	host := os.Getenv(swaggerHostName)
	if len(host) == 0 {
		return nil, errors.New("swagger host not found")
	}

	port := os.Getenv(swaggerPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("swagger port not found")
	}

	return &config.SwaggerConfig{
		Host: host,
		Port: port,
	}, nil
}
