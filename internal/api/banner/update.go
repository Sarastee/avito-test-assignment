package banner

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/sarastee/avito-test-assignment/internal/api"
	"github.com/sarastee/avito-test-assignment/internal/converter"
	"github.com/sarastee/avito-test-assignment/internal/model"
	"github.com/sarastee/avito-test-assignment/internal/repository"
	"github.com/sarastee/avito-test-assignment/internal/utils/response"
	"github.com/sarastee/avito-test-assignment/internal/utils/validator"
)

// UpdateBanner is API layer function which process the request and updates banner
func (i *Implementation) UpdateBanner(w http.ResponseWriter, r *http.Request) {
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

	bannerID, err := strconv.ParseInt(r.PathValue(BannerIDField), 10, 64)
	if err != nil {
		i.logger.Info().Msg(err.Error())
		response.SendError(w, http.StatusNotFound, err, i.logger)
		return
	}

	var updateBanner model.UpdateBanner
	if code, err := validator.ParseRequestBody(r.Body, &updateBanner, model.ValidateUpdateBanner, i.logger); err != nil {
		i.logger.Info().Msg(err.Error())
		response.SendError(w, code, err, i.logger)
		return
	}

	updateBannerSQL := converter.UpdateBannerToUpdateBannerSQL(&updateBanner)

	if _, err = i.bannerService.UpdateBanner(r.Context(), bannerID, &updateBannerSQL); err != nil {
		if errors.Is(err, repository.ErrBannerNotFound) {
			i.logger.Info().Msg(err.Error())
			response.SendError(w, http.StatusNotFound, err, i.logger)
			return
		}

		if errors.Is(err, repository.ErrBannerConflict) {
			i.logger.Info().Msg(err.Error())
			response.SendError(w, http.StatusConflict, err, i.logger)
			return
		}

		i.logger.Error().Msg(err.Error())
		response.SendError(w, http.StatusInternalServerError, api.ErrInternalError, i.logger)
		return
	}

	response.SendStatus(w, http.StatusOK, nil, i.logger)
}
