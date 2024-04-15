package response

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog"
	"github.com/sarastee/avito-test-assignment/internal/model"
)

// SendStatus function which sends status with provided code and data
func SendStatus(w http.ResponseWriter, code int, data any, logger *zerolog.Logger) {
	logger.Info().Msgf("trying to send response with status code: %d", code)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			logger.Warn().Err(err).Msgf("error while sending response with status code: %d", code)
		}
	}
}

// SendError function which sends Error status with provided code and error
func SendError(w http.ResponseWriter, code int, err error, logger *zerolog.Logger) {
	logger.Info().Msgf("trying to send error response %s with status code: %d", err, code)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	if rErr := json.NewEncoder(w).Encode(model.Error{Err: err.Error()}); rErr != nil {
		logger.Warn().Msgf("error while sending error response %s with status code: %d", rErr, code)
	}
}
