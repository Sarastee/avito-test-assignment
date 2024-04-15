package validator

import (
	"errors"
	"net/http"
)

// CheckContentType function which validate request content type
func CheckContentType(r *http.Request) error {
	if len(r.Header.Values("Content-Type")) == 0 {
		return errors.New("empty content-type")
	}

	for contentTypeCurrentIndex, contentType := range r.Header.Values("Content-type") {
		if contentType == "application/json" {
			break
		}

		if contentTypeCurrentIndex == len(r.Header.Values("Content-Type"))-1 {
			return errors.New("wrong content-type")
		}
	}

	return nil
}
