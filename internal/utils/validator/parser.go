package validator

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

const (
	errMsgCannotReadBody       = "can't read body"
	errMsgIncorrectBodyContent = "incorrect body content"
)

var (
	ErrCannotReadBody       = errors.New(errMsgCannotReadBody)       // ErrCannotReadBody is Can't Read Body Error object
	ErrIncorrectBodyContent = errors.New(errMsgIncorrectBodyContent) // ErrIncorrectBodyContent is Incorrect Body content Error object
)

const (
	base10 = 10
	size64 = 64
)

func ParseRequestBody(reqBody io.ReadCloser, output any, validation func([]byte) error, logger *zerolog.Logger) (int, error) {
	body, err := io.ReadAll(reqBody)
	if err != nil {
		logger.Error().Msg(errors.Wrap(err, "can't read body").Error())

		return http.StatusInternalServerError, ErrCannotReadBody
	}

	if err := validation(body); err != nil {
		if errors.Is(err, ErrJSONError) {
			logger.Warn().Msg(errors.Wrap(err, "try parse body json").Error())

			return http.StatusBadRequest, ErrIncorrectBodyContent
		}

		return http.StatusBadRequest, ErrIncorrectBodyContent
	}

	if err := json.Unmarshal(body, output); err != nil {
		logger.Warn().Msg(errors.Wrap(err, "try request entity").Error())

		return http.StatusBadRequest, ErrIncorrectBodyContent
	}

	return http.StatusOK, nil
}

func ParseQueryParamToInt64(r *http.Request, param string, nullable error,
	incorrectTypeError error, logger *zerolog.Logger) (*int64, error) {
	rawField := r.URL.Query().Get(param)
	if rawField == "" {
		return nil, nullable
	}

	id, err := strconv.ParseInt(rawField, base10, size64)
	if err != nil {
		logger.Error().Msg(errors.Wrapf(err, "can't parse query field %s with value %s", param, rawField).Error())

		return nil, incorrectTypeError
	}

	return &id, nil
}
