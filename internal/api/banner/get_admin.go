package banner

import (
	"net/http"

	"github.com/sarastee/avito-test-assignment/internal/api"
	"github.com/sarastee/avito-test-assignment/internal/converter"
	"github.com/sarastee/avito-test-assignment/internal/utils/response"
	"github.com/sarastee/avito-test-assignment/internal/utils/validator"
)

// GetAdminBanners ...
//
// @Summary Updates banner by id
// @Security AdminToken
// @Description API layer function which process the request and pull out banners from database
// @Tags Banners
//
// @Param tag_id query integer false "Tag ID"
// @Param feature_id query integer false "Feature ID"
// @Param revision_id query integer false "Revision ID"
// @Param limit query integer false "Limit"
// @Param offset query integer false "Offset"
// @Produce json
//
// @Success 200 {array} model.Banner "Banner array in JSON format"
// @Failure 400 {object} model.Error "Incorrect provided data"
// @Failure 401 {object} model.Error "User not authorized"
// @Failure 403 {object} model.Error "User insufficient rights"
// @Failure 500 {object} model.Error "Internal server error"
// @Router /banner [get]
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
