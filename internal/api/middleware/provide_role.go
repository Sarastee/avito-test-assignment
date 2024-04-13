package middleware

//
//func (m *Middleware) ProvideRole(next func(bool) http.HandlerFunc) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		isRoleAdmin, err := m.checkIsRoleAdmin(r)
//		if err != nil {
//			m.logger.Info().Msg(err.Error())
//			response.SendError(w, http.StatusUnauthorized, err, m.logger)
//			return
//		}
//
//		next(isRoleAdmin).ServeHTTP(w, r)
//	})
//}

// TODO: delete if not needed
