package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sarastee/avito-test-assignment/internal/api"
	"github.com/sarastee/avito-test-assignment/internal/model"
	"github.com/sarastee/avito-test-assignment/internal/repository"
	"github.com/sarastee/avito-test-assignment/internal/utils/error_handler"
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

	err := validator.JSONValidate(r)
	if err != nil {
		i.logger.Info().Msg(err.Error())
		error_handler.HandleError(w, i.logger, err, http.StatusBadRequest)
		return
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var user model.User
	if err = dec.Decode(&user); err != nil {
		i.logger.Info().Msg(err.Error())
		error_handler.HandleError(w, i.logger, err, http.StatusBadRequest)
		return
	}

	if !(string(user.Role) == "ADMIN" || string(user.Role) == "USER") {
		i.logger.Info().Msg(api.ErrIncorrectRole.Error())
		error_handler.HandleError(w, i.logger, api.ErrIncorrectRole, http.StatusBadRequest)
		return
	}

	err = i.authService.CreateUser(r.Context(), user)
	if err != nil {
		if errors.Is(err, repository.ErrUserAlreadyRegistered) {
			i.logger.Info().Msg(err.Error())
			error_handler.HandleError(w, i.logger, repository.ErrUserAlreadyRegistered, http.StatusBadRequest)
			return
		}

		i.logger.Error().Msg(err.Error())
		error_handler.HandleError(w, i.logger, api.ErrInternalError, http.StatusInternalServerError)
		return
	}

	tokenStr, err := i.jwtService.GenerateAccessToken(user)
	if err != nil {
		i.logger.Error().Msg(err.Error())
		error_handler.HandleError(w, i.logger, api.ErrInternalError, http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(model.Token{Token: tokenStr})
	if err != nil {
		i.logger.Error().Msg(err.Error())
	}
}
