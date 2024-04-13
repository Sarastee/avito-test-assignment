package auth

import (
	"errors"
	"net/http"

	"github.com/sarastee/avito-test-assignment/internal/api"
	"github.com/sarastee/avito-test-assignment/internal/converter"
	"github.com/sarastee/avito-test-assignment/internal/model"
	"github.com/sarastee/avito-test-assignment/internal/repository"
	"github.com/sarastee/avito-test-assignment/internal/utils/response"
	"github.com/sarastee/avito-test-assignment/internal/utils/validator"
)

// CreateUser is API layer function which process the request and creates user
func (i *Implementation) CreateUser(w http.ResponseWriter, r *http.Request) {
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

	var createUser model.CreateUser
	if code, err := validator.ParseRequestBody(r.Body, &createUser, model.ValidateCreateUser, i.logger); err != nil { // nolint
		response.SendError(w, code, err, i.logger)
		return
	}

	userID, err := i.authService.CreateUser(r.Context(), createUser)
	if err != nil {
		if errors.Is(err, repository.ErrUserAlreadyRegistered) {
			i.logger.Info().Msg(err.Error())
			response.SendError(w, http.StatusBadRequest, repository.ErrUserAlreadyRegistered, i.logger)
			return
		}

		i.logger.Error().Msg(err.Error())
		response.SendError(w, http.StatusInternalServerError, api.ErrInternalError, i.logger)
		return
	}

	user := converter.CreateUserToUser(userID, &createUser)

	tokenStr, err := i.jwtService.GenerateAccessToken(*user)
	if err != nil {
		i.logger.Error().Msg(err.Error())
		response.SendError(w, http.StatusInternalServerError, api.ErrInternalError, i.logger)
		return
	}

	response.SendStatus(w, http.StatusCreated, model.Token{Token: tokenStr}, i.logger)
}
