package middleware

import (
	"net/http"

	"github.com/sarastee/avito-test-assignment/internal/utils/response"
)

// AuthRequired is middleware which check user authorization
func (m *Middleware) AuthRequired(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := m.checkIsRoleAdmin(r)
		if err != nil {
			m.logger.Info().Msg(err.Error())
			response.SendError(w, http.StatusUnauthorized, err, m.logger)
			return
		}

		next.ServeHTTP(w, r)
	})
}
