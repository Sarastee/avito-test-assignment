package banner

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/sarastee/avito-test-assignment/internal/api"
	"github.com/sarastee/avito-test-assignment/internal/model"
	"github.com/sarastee/avito-test-assignment/internal/repository"
	"github.com/sarastee/avito-test-assignment/internal/utils/response"
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

	err := validator.CheckContentType(r)
	if err != nil {
		i.logger.Info().Msg(err.Error())
		response.SendError(w, http.StatusBadRequest, err, i.logger)
		return
	}

	var createBanner model.CreateBanner
	if code, err := validator.ParseRequestBody(r.Body, &createBanner, model.ValidateCreateBanner, i.logger); err != nil { // nolint
		response.SendError(w, code, err, i.logger)
		return
	}

	bannerID, err := i.bannerService.CreateBanner(r.Context(), createBanner.IsActive, createBanner.Content,
		createBanner.FeatureID, createBanner.TagsIDs)
	if err != nil {
		if errors.Is(err, repository.ErrTagsUniqueViolation) {
			response.SendError(w, http.StatusBadRequest, repository.ErrTagsUniqueViolation, i.logger)
			return
		}

		i.logger.Error().Msg(err.Error())
		response.SendError(w, http.StatusInternalServerError, api.ErrInternalError, i.logger)
		return
	}

	response.SendStatus(w, http.StatusCreated, model.BannerID{ID: bannerID}, i.logger)
}
