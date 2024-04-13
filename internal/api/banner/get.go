package banner

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sarastee/avito-test-assignment/internal/api"
	"github.com/sarastee/avito-test-assignment/internal/repository"
	"github.com/sarastee/avito-test-assignment/internal/utils/response"
	"github.com/sarastee/avito-test-assignment/internal/utils/validator"
)

// GetUserBanner is API layer function which process the request and gets banner
// Fields: tag_id: required, feature_id: required, version: not_required, use_last_revision: not_required
// if use_last_revision == false -> cache, if cache == nil -> db -> getBanner -> cacheSet
// if use_last_revision == true -> db -> getBanner -> cacheSet
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

	var nullRevisionID sql.NullInt64
	if revisionID != nil {
		nullRevisionID = sql.NullInt64{Int64: *revisionID, Valid: true}
	} else {
		nullRevisionID = sql.NullInt64{Valid: false}
	}

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
			err = i.bannerCacheService.SetCache(r.Context(), *featureID, *tagID, nullRevisionID, json.RawMessage(banner))
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
		err = i.bannerCacheService.SetCache(r.Context(), *featureID, *tagID, nullRevisionID, json.RawMessage(banner))
		if err != nil {
			i.logger.Error().Msg(err.Error())
		}
		response.SendStatus(w, http.StatusOK, json.RawMessage(banner), i.logger)
	}

}
