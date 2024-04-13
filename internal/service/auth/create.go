package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/sarastee/avito-test-assignment/internal/model"
	"github.com/sarastee/avito-test-assignment/internal/service"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser is Service layer function which process request and creates user
func (s Service) CreateUser(ctx context.Context, user model.CreateUser) (int64, error) {
	hashedPassword, err := s.passManager.HashPassword(user.Password)
	if err != nil {
		s.logger.Info().Err(err).Msg("failed to hash password")
		if errors.Is(err, bcrypt.ErrPasswordTooLong) {
			return 0, service.ErrPasswordToLong
		}

		return 0, err
	}

	user.Password = hashedPassword

	userID, err := s.authRepo.CreateUser(ctx, user.Name, user.Password, user.Role)
	if err != nil {
		return 0, fmt.Errorf("failed to create user: %w", err)
	}

	return userID, nil
}
