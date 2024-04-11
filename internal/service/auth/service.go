package auth

import (
	"github.com/rs/zerolog"
	"github.com/sarastee/avito-test-assignment/internal/repository"
	"github.com/sarastee/avito-test-assignment/internal/service"
	"github.com/sarastee/avito-test-assignment/internal/utils/password"
	"github.com/sarastee/platform_common/pkg/db"
)

var _ service.AuthService = (*Service)(nil)

// Service auth Service
type Service struct {
	logger      *zerolog.Logger
	txManager   db.TxManager
	authRepo    repository.AuthRepository
	passManager *password.Manager
}

// NewService function which get new Auth Service instance
func NewService(logger *zerolog.Logger, txManager db.TxManager, authRepository repository.AuthRepository, passManager *password.Manager) *Service {
	return &Service{
		logger:      logger,
		txManager:   txManager,
		authRepo:    authRepository,
		passManager: passManager,
	}
}
