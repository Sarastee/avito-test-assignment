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

// DeleteBanner is API layer function which process the request and deletes banner
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
		response.SendError(w, http.StatusNotFound, err, i.logger)
		return
	}

	if err := i.bannerService.DeleteBanner(r.Context(), bannerID); err != nil {
		if errors.Is(err, repository.ErrBannerNotFoundDelete) {
			response.SendError(w, http.StatusNotFound, err, i.logger)
			return
		}

		i.logger.Error().Msg(errors.Wrap(err, "can't delete banner").Error())
		response.SendError(w, http.StatusInternalServerError, api.ErrInternalError, i.logger)

		return
	}

	response.SendStatus(w, http.StatusNoContent, nil, i.logger)
}
