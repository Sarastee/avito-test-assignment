package banner

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"github.com/sarastee/avito-test-assignment/internal/api"
	"github.com/sarastee/avito-test-assignment/internal/model"
	"github.com/sarastee/avito-test-assignment/internal/repository"
	"github.com/sarastee/avito-test-assignment/internal/utils/error_handler"
	"github.com/sarastee/avito-test-assignment/internal/utils/validator"
)

// CreateBanner is API layer function which process the request and creates banner
func (i *Implementation) CreateBanner(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := r.Body.Close()
		if err != nil {
			i.logger.Warn().Msg(err.Error())
		}
	}()

	err := validator.JSONValidate(r)
	if err != nil {
		i.logger.Info().Msg(err.Error())
		error_handler.HandleError(w, i.logger, err, http.StatusBadRequest)
		return
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var banner model.Banner
	if err = dec.Decode(&banner); err != nil {
		i.logger.Info().Msg(err.Error())
		error_handler.HandleError(w, i.logger, err, http.StatusBadRequest)
		return
	}

	switch {
	case banner.Content == nil:
		i.logger.Info().Msg(errors.Wrapf(api.ErrFieldNotFound, "content").Error())
		error_handler.HandleError(w, i.logger, errors.Wrapf(api.ErrFieldNotFound, "content"), http.StatusBadRequest)
		return
	case banner.FeatureID == nil:
		i.logger.Info().Msg(errors.Wrapf(api.ErrFieldNotFound, "feature_id").Error())
		error_handler.HandleError(w, i.logger, errors.Wrapf(api.ErrFieldNotFound, "feature_id"), http.StatusBadRequest)
		return
	case banner.IsActive == nil:
		i.logger.Info().Msg(errors.Wrapf(api.ErrFieldNotFound, "is_active").Error())
		error_handler.HandleError(w, i.logger, errors.Wrapf(api.ErrFieldNotFound, "is_active"), http.StatusBadRequest)
		return
	case banner.TagIDs == nil:
		i.logger.Info().Msg(errors.Wrapf(api.ErrFieldNotFound, "tag_ids").Error())
		error_handler.HandleError(w, i.logger, errors.Wrapf(api.ErrFieldNotFound, "tag_ids"), http.StatusBadRequest)
		return
	}

	bannerID, err := i.bannerService.CreateBanner(r.Context(), &banner)
	if err != nil {
		if errors.Is(err, repository.ErrTagsUniqueViolation) {
			error_handler.HandleError(w, i.logger, repository.ErrTagsUniqueViolation, http.StatusBadRequest)
			return
		}

		i.logger.Error().Msg(err.Error())
		error_handler.HandleError(w, i.logger, api.ErrInternalError, http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err = json.NewEncoder(w).Encode(model.BannerID{ID: bannerID}); err != nil {
		return
	}
}
