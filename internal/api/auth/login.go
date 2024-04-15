package auth

import (
	"errors"
	"net/http"

	"github.com/sarastee/avito-test-assignment/internal/api"
	"github.com/sarastee/avito-test-assignment/internal/model"
	"github.com/sarastee/avito-test-assignment/internal/repository"
	"github.com/sarastee/avito-test-assignment/internal/service"
	"github.com/sarastee/avito-test-assignment/internal/utils/response"
	"github.com/sarastee/avito-test-assignment/internal/utils/validator"
)

// LogIn ...
//
// @Summary User login
// @Description API layer function which process the request and login user
// @Tags Auth
// @Param request body model.AuthUser true "Login user params"
// @Accept json
// @Produce json
// @Success 200 {object} model.Token "User has successfully logged in"
// @Failure 400 {object} model.Error "Incorrect provided data"
// @Failure 401 {object} model.Error "Incorrect password"
// @Failure 404 {object} model.Error "User not found"
// @Failure 500 {object} model.Error "Internal server error"
// @Router /login [post]
func (i *Implementation) LogIn(w http.ResponseWriter, r *http.Request) {
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

	var authUser model.AuthUser
	if code, err := validator.ParseRequestBody(r.Body, &authUser, model.ValidateAuthUser, i.logger); err != nil { // nolint
		response.SendError(w, code, err, i.logger)
		return
	}

	userModel, err := i.authService.VerifyUser(r.Context(), authUser)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			i.logger.Info().Msg(err.Error())
			response.SendError(w, http.StatusNotFound, repository.ErrUserNotFound, i.logger)
			return
		}

		if errors.Is(err, service.ErrWrongPassword) {
			i.logger.Info().Msg(err.Error())
			response.SendError(w, http.StatusUnauthorized, service.ErrWrongPassword, i.logger)
			return
		}

		i.logger.Error().Msg(err.Error())
		response.SendError(w, http.StatusInternalServerError, api.ErrInternalError, i.logger)
		return
	}

	tokenStr, err := i.jwtService.GenerateAccessToken(*userModel)
	if err != nil {
		i.logger.Error().Msg(err.Error())
		response.SendError(w, http.StatusInternalServerError, api.ErrInternalError, i.logger)
		return
	}

	response.SendStatus(w, http.StatusOK, model.Token{Token: tokenStr}, i.logger)
}
