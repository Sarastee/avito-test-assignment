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

// UpdateBanner
//
// @Summary Updates banner by id
// @Security AdminToken
// @Description API layer function which process the request and updates banner
// @Tags Banners
// @Param id path integer true "Banner ID"
// @Param request body model.UpdateBanner true "Banner update data"
// @Accept json
// @Produce json
// @Success 200 "Banner successfully updated"
// @Failure 400 {object} model.Error "Incorrect provided data"
// @Failure 401 {object} model.Error "User not authorized"
// @Failure 403 {object} model.Error "User insufficient rights"
// @Failure 404 {object} model.Error "Banner not found"
// @Failure 409 {object} model.Error "Banner with provided params already exists"
// @Failure 500 {object} model.Error "Internal server error"
// @Router /banner/{id} [patch]
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
