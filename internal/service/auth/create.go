package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/sarastee/avito-test-assignment/internal/model"
	"github.com/sarastee/avito-test-assignment/internal/service"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser is Service layer function which process request
func (s Service) CreateUser(ctx context.Context, user model.User) error {
	hashedPassword, err := s.passManager.HashPassword(user.PasswordHash)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed to hash password")
		if errors.Is(err, bcrypt.ErrPasswordTooLong) {
			return service.ErrPasswordToLong
		}

		return err
	}

	user.PasswordHash = hashedPassword

	err = s.authRepo.CreateUser(ctx, user.Name, user.PasswordHash, string(user.Role))
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}
