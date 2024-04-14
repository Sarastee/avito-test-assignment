package auth

import (
	"context"

	"github.com/sarastee/avito-test-assignment/internal/model"
	"github.com/sarastee/avito-test-assignment/internal/service"
)

// VerifyUser is Service layer function which process request and verifies user
func (s Service) VerifyUser(ctx context.Context, userAuth model.AuthUser) (*model.User, error) {
	userModel, err := s.authRepo.GetUser(ctx, userAuth.Name)
	if err != nil {
		return nil, err
	}

	ok := s.passManager.CheckPasswordHash(userAuth.Password, userModel.Password)
	if !ok {
		s.logger.Info().Msg("failed to hash password")
		return nil, service.ErrWrongPassword
	}
	return userModel, nil
}
