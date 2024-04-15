package banner

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sarastee/avito-test-assignment/internal/api"
	"github.com/sarastee/avito-test-assignment/internal/converter"
	"github.com/sarastee/avito-test-assignment/internal/repository"
	"github.com/sarastee/avito-test-assignment/internal/utils/response"
	"github.com/sarastee/avito-test-assignment/internal/utils/validator"
)

// GetUserBanner
//
// @Summary Get banner by id
// @Security AdminToken
// @Security UserToken
// @Description API layer function which process the request and pull out banner from database
// @Tags Get User Banner
//
// @Param Content-Type header string true "Content Type" default(application/json)
// @Param tag_id query integer true "Tag ID"
// @Param feature_id query integer true "Feature ID"
// @Param revision_id query integer false "Revision ID"
// @Param use_last_revision query boolean false "Use last revision"
// @Produce json
//
// @Success 200 {object} any "Banner in JSON format"
// @Failure 400 {object} model.Error "Incorrect provided data"
// @Failure 401 {object} model.Error "User not authorized"
// @Failure 403 {object} model.Error "User insufficient rights"
// @Failure 404 {object} model.Error "Banner not found"
// @Failure 500 {object} model.Error "Internal server error"
// @Router /user_banner [get]
func (i *Implementation) GetUserBanner(w http.ResponseWriter, r *http.Request) {
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

	tagID, err := validator.ParseQueryParamToInt64(r, TagIDParam, api.ErrTagNotFound, api.ErrTagIsNotANumber, i.logger)
	if err != nil {
		response.SendError(w, http.StatusBadRequest, err, i.logger)
		return
	}

	featureID, err := validator.ParseQueryParamToInt64(r, FeatureIDParam, api.ErrFeatureNotFound, api.ErrFeatureIsNotANumber, i.logger)
	if err != nil {
		response.SendError(w, http.StatusBadRequest, err, i.logger)
		return
	}

	revisionID, err := validator.ParseQueryParamToInt64(r, RevisionIDParam, nil, api.ErrRevisionIsNotANumber, i.logger)
	if err != nil {
		response.SendError(w, http.StatusBadRequest, err, i.logger)
		return
	}

	nullRevisionID := converter.Int64PointerToSQLNullInt64(revisionID)

	var useLastRevision bool
	useLastRevisionStr := r.URL.Query().Get("use_last_revision")
	if useLastRevisionStr == "true" {
		useLastRevision = true
	} else {
		useLastRevision = false
	}

	switch {
	case !useLastRevision:
		banner, err := i.bannerCacheService.GetCache(r.Context(), *featureID, *tagID, nullRevisionID)
		if err != nil {
			if !errors.Is(err, repository.ErrCacheNotFound) {
				i.logger.Error().Msg(err.Error())
			}
			i.logger.Info().Msg(err.Error())
		} else {
			i.logger.Info().Msg("cache found")
			if err != nil {
				i.logger.Error().Msg(err.Error())
			}
			response.SendStatus(w, http.StatusOK, json.RawMessage(banner), i.logger)
			break
		}

		fallthrough

	case useLastRevision:
		banner, err := i.bannerService.GetBannerFromDatabase(r.Context(), *tagID, *featureID, nullRevisionID)
		if err != nil {
			if errors.Is(err, repository.ErrBannerNotFound) {
				response.SendError(w, http.StatusNotFound, repository.ErrBannerNotFound, i.logger)
				return
			}
			i.logger.Error().Msg(err.Error())
			response.SendError(w, http.StatusInternalServerError, api.ErrInternalError, i.logger)
			return
		}
		err = i.bannerCacheService.SetCache(r.Context(), *featureID, *tagID, nullRevisionID, banner)
		if err != nil {
			i.logger.Error().Msg(err.Error())
		}
		response.SendStatus(w, http.StatusOK, banner, i.logger)
	}

}
