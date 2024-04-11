package banner

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/sarastee/avito-test-assignment/internal/api"
	"github.com/sarastee/avito-test-assignment/internal/repository"
	"github.com/sarastee/avito-test-assignment/internal/utils/error_handler"
)

// GetBanner is API layer function which process the request and gets banner
func (i *Implementation) GetBanner() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := r.Body.Close()
			if err != nil {
				i.logger.Warn().Msg(err.Error())
			}
		}()

		tagIDString := r.URL.Query().Get("tag_id")
		featureIDString := r.URL.Query().Get("feature_id")

		if tagIDString == "" || featureIDString == "" {
			error_handler.HandleError(w, i.logger, api.ErrTagOrFeatureNotProvided, http.StatusBadRequest)
			return
		}

		tagID, err := strconv.Atoi(tagIDString)
		if err != nil {
			error_handler.HandleError(w, i.logger, api.ErrTagIsNotANumber, http.StatusBadRequest)
			return
		}

		featureID, err := strconv.Atoi(featureIDString)
		if err != nil {
			error_handler.HandleError(w, i.logger, api.ErrFeatureIDIsNotANumber, http.StatusBadRequest)
			return
		}

		banner, err := i.bannerService.GetBanner(r.Context(), int64(tagID), int64(featureID))
		if err != nil {
			if errors.Is(err, repository.ErrBannerNotFound) {
				error_handler.HandleError(w, i.logger, repository.ErrBannerNotFound, http.StatusNotFound)
				return
			}

			i.logger.Error().Msg(err.Error())
			error_handler.HandleError(w, i.logger, api.ErrInternalError, http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, err = w.Write([]byte(banner))
		if err != nil {
			i.logger.Error().Msg(err.Error())
		}
	}
}
