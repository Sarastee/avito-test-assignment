package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sarastee/avito-test-assignment/internal/model"
)

// GenerateAccessToken is Service layer function which process request
func (s *Service) GenerateAccessToken(user model.User) (string, error) {
	s.logger.Debug().Int64("userID", user.ID).Msg("attempt to generate access token")
	return s.generateToken(user, s.jwtConfig.JWTAccessTokenExpireThrough)
}

func (s *Service) generateToken(user model.User, duration time.Duration) (string, error) {
	claims := model.UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
		UserID:   fmt.Sprintf("%d", user.ID),
		UserName: user.Name,
		Role:     string(user.Role),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.jwtConfig.JWTSecretKey))
	if err != nil {
		return "", fmt.Errorf("unable to sign token: %w", err)
	}

	return signedToken, nil
}
