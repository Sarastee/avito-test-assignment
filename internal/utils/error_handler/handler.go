package error_handler

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog"
	"github.com/sarastee/avito-test-assignment/internal/model"
)

// HandleError function which writes HTTP Error in Response
func HandleError(w http.ResponseWriter, logger *zerolog.Logger, err error, statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err = json.NewEncoder(w).Encode(model.Error{Err: err.Error()})
	if err != nil {
		logger.Debug().Msg(err.Error())
	}
}
