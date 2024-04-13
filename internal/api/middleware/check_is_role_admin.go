package middleware

import (
	"net/http"

	"github.com/sarastee/avito-test-assignment/internal/api"
)

func (m *Middleware) checkIsRoleAdmin(r *http.Request) (bool, error) {
	authToken := r.Header.Get("token")
	if authToken == "" {
		return false, api.ErrNoTokenProvided
	}

	isRoleAdmin, err := m.jwtService.VerifyAccessToken(authToken)
	if err != nil {
		return false, api.ErrInvalidToken
	}

	return isRoleAdmin, nil
}
