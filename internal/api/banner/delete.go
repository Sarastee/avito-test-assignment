package banner

import (
	"net/http"
	"strconv"

	"github.com/pkg/errors"
	"github.com/sarastee/avito-test-assignment/internal/api"
	"github.com/sarastee/avito-test-assignment/internal/repository"
	"github.com/sarastee/avito-test-assignment/internal/utils/response"
	"github.com/sarastee/avito-test-assignment/internal/utils/validator"
)

// DeleteBanner
//
// @Summary Deletes banner by id
// @Security AdminToken
// @Description API layer function which process the request and deletes banner
// @Tags Banners
// @Param id path integer true "Banner ID"
// @Produce json
// @Success 204 "Banner successfully deleted"
// @Failure 400 {object} model.Error "Incorrect provided data"
// @Failure 401 {object} model.Error "User not authorized"
// @Failure 403 {object} model.Error "User insufficient rights"
// @Failure 404 {object} model.Error "Banner not found"
// @Failure 500 {object} model.Error "Internal server error"
// @Router /banner/{id} [delete]
func (i *Implementation) DeleteBanner(w http.ResponseWriter, r *http.Request) {
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

	if err := i.bannerService.DeleteBanner(r.Context(), bannerID); err != nil {
		if errors.Is(err, repository.ErrBannerNotFoundDelete) {
			i.logger.Info().Msg(err.Error())
			response.SendError(w, http.StatusNotFound, err, i.logger)
			return
		}

		i.logger.Error().Msg(errors.Wrap(err, "can't delete banner").Error())
		response.SendError(w, http.StatusInternalServerError, api.ErrInternalError, i.logger)

		return
	}

	response.SendStatus(w, http.StatusNoContent, nil, i.logger)
}
