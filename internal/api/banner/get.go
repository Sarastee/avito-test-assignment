package banner

//// GetUserBanner is API layer function which process the request and gets banner
//// Fields: tag_id: required, feature_id: required, version: not_required, use_last_revision: not_required
//// if use_last_revision == false -> cache, if cache == nil -> db -> getBanner -> cacheSet
//// if use_last_revision == true -> db -> getBanner -> cacheSet
//func (i *Implementation) GetUserBanner() http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		defer func() {
//			err := r.Body.Close()
//			if err != nil {
//				i.logger.Warn().Msg(err.Error())
//			}
//		}()
//
//		tagIDString := r.URL.Query().Get("tag_id")
//		featureIDString := r.URL.Query().Get("feature_id")
//
//		if tagIDString == "" || featureIDString == "" {
//			request.HandleError(w, i.logger, api.ErrTagOrFeatureNotProvided, http.StatusBadRequest)
//			return
//		}
//
//		tagID, err := strconv.Atoi(tagIDString)
//		if err != nil {
//			request.HandleError(w, i.logger, api.ErrTagIsNotANumber, http.StatusBadRequest)
//			return
//		}
//
//		featureID, err := strconv.Atoi(featureIDString)
//		if err != nil {
//			request.HandleError(w, i.logger, api.ErrFeatureIDIsNotANumber, http.StatusBadRequest)
//			return
//		}
//
//		banner, err := i.bannerService.GetBanner(r.Context(), int64(tagID), int64(featureID), isRoleAdmin)
//		if err != nil {
//			if errors.Is(err, repository.ErrBannerNotFound) {
//				request.HandleError(w, i.logger, repository.ErrBannerNotFound, http.StatusNotFound)
//				return
//			}
//
//			i.logger.Error().Msg(err.Error())
//			request.HandleError(w, i.logger, api.ErrInternalError, http.StatusInternalServerError)
//			return
//		}
//
//		w.Header().Add("Content-Type", "application/json")
//		w.WriteHeader(http.StatusOK)
//
//		_, err = w.Write([]byte(banner))
//		if err != nil {
//			i.logger.Error().Msg(err.Error())
//		}
//	}
//}

// TODO: имплементировать кэш редиса только в гет запросах с версиями. не забыть про use_last_revision
