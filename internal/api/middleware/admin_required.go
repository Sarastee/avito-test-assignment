package middleware

import (
	"net/http"

	"github.com/sarastee/avito-test-assignment/internal/api"
	"github.com/sarastee/avito-test-assignment/internal/utils/response"
)

func (m *Middleware) AdminRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isRoleAdmin, err := m.checkIsRoleAdmin(r)
		if err != nil {
			m.logger.Info().Msg(err.Error())
			response.SendError(w, http.StatusUnauthorized, err, m.logger)
			return
		}

		if !isRoleAdmin {
			response.SendError(w, http.StatusForbidden, api.ErrInsufficientRights, m.logger)
			return
		}

		next.ServeHTTP(w, r)
	})
}
