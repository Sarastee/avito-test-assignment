package banner

import (
	"net/http"

	"github.com/sarastee/avito-test-assignment/internal/api"
	"github.com/sarastee/avito-test-assignment/internal/converter"
	"github.com/sarastee/avito-test-assignment/internal/utils/response"
	"github.com/sarastee/avito-test-assignment/internal/utils/validator"
)

// GetAdminBanners is API layer function which process the request and pull out banners from database
func (i *Implementation) GetAdminBanners(w http.ResponseWriter, r *http.Request) {
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

	tagID, err := validator.ParseQueryParamToInt64(r, TagIDParam, nil, api.ErrTagIsNotANumber, i.logger)
	if err != nil {
		response.SendError(w, http.StatusBadRequest, err, i.logger)
		return
	}

	featureID, err := validator.ParseQueryParamToInt64(r, FeatureIDParam, nil, api.ErrFeatureIsNotANumber, i.logger)
	if err != nil {
		response.SendError(w, http.StatusBadRequest, err, i.logger)
		return
	}

	limit, err := validator.ParseQueryParamToInt64(r, LimitParam, nil, api.ErrLimitIsNotANumber, i.logger)
	if err != nil {
		response.SendError(w, http.StatusBadRequest, err, i.logger)
		return
	}

	offset, err := validator.ParseQueryParamToInt64(r, OffsetParam, nil, api.ErrOffsetIsNotANumber, i.logger)
	if err != nil {
		response.SendError(w, http.StatusBadRequest, err, i.logger)
		return
	}

	nullTagID := converter.Int64PointerToSQLNullInt64(tagID)
	nullFeatureID := converter.Int64PointerToSQLNullInt64(featureID)
	nullLimit := converter.Int64PointerToSQLNullInt64(limit)
	nullOffset := converter.Int64PointerToSQLNullInt64(offset)

	banners, err := i.bannerService.GetAdminBanners(r.Context(), nullFeatureID, nullTagID, nullOffset, nullLimit)
	if err != nil {
		response.SendError(w, http.StatusInternalServerError, api.ErrInternalError, i.logger)
	}

	response.SendStatus(w, http.StatusOK, banners, i.logger)
}
