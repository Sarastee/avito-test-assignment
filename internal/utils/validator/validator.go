package validator

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// JSONValidate function which validate HTTP Request as JSON format
func JSONValidate(r *http.Request) error {
	if !CheckContentType(r) {
		return errors.New("wrong or empty Content-Type")
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		return errors.New("something went wrong")
	}

	reader := bytes.NewReader(bytes.Clone(data))

	err = CheckDuplicatesInJSON(json.NewDecoder(reader), nil)
	if err != nil {
		return errors.New("something wrong in json")
	}

	r.Body = io.NopCloser(bytes.NewBuffer(data))

	return nil
}

// CheckContentType function which validate request content type
func CheckContentType(r *http.Request) bool {
	if len(r.Header.Values("Content-Type")) == 0 {
		return false
	}

	for contentTypeCurrentIndex, contentType := range r.Header.Values("Content-type") {
		if contentType == "application/json" {
			break
		}

		if contentTypeCurrentIndex == len(r.Header.Values("Content-Type"))-1 {
			return false
		}
	}

	return true
}

// CheckDuplicatesInJSON function which search request for duplicates
func CheckDuplicatesInJSON(d *json.Decoder, path []string) error {
	// Get next token from JSON
	t, err := d.Token()
	if err != nil {
		return err
	}

	// Is it a delimiter?
	delim, ok := t.(json.Delim)

	// No, nothing more to check
	if !ok {
		// scaler type, nothing to do
		return nil
	}

	switch delim {
	case '{':
		keys := make(map[string]bool)
		for d.More() {

			// Get attribute key

			t, err := d.Token()
			if err != nil {
				return err
			}
			key := t.(string)

			// Check for duplicates

			if keys[key] {
				fmt.Printf("Duplicate %s\n", strings.Join(append(path, key), "/"))
			}
			keys[key] = true

			// Check value

			if err := CheckDuplicatesInJSON(d, append(path, key)); err != nil {
				return err
			}
		}
		// consume trailing }
		if _, err := d.Token(); err != nil {
			return err
		}

	case '[':
		i := 0
		for d.More() {
			if err := CheckDuplicatesInJSON(d, append(path, strconv.Itoa(i))); err != nil {
				return err
			}
			i++
		}
		// consume trailing ]
		if _, err := d.Token(); err != nil {
			return err
		}

	}
	return nil
}
