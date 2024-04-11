package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sarastee/avito-test-assignment/internal/model"
)

// VerifyAccessToken is Service layer function which process request
func (s *Service) VerifyAccessToken(tokenStr string) (bool, error) {
	s.logger.Debug().Str("access_token", tokenStr).Msg("verifying access token")
	return s.verifyToken(tokenStr)
}

func (s *Service) verifyToken(tokenStr string) (bool, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&model.UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpeted sign method")
			}

			return []byte(s.jwtConfig.JWTSecretKey), nil
		},
	)
	if err != nil {
		return false, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*model.UserClaims)
	if !ok {
		return false, fmt.Errorf("invalid token body")
	}

	if claims.Role != "ADMIN" {
		return false, nil
	}

	return true, nil
}
