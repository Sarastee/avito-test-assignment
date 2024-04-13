package validator

import (
	"errors"

	"github.com/miladibra10/vjson"
)

const (
	errMsgJSONError = "could not parse json input."
)

var (
	ErrJSONError = errors.New(errMsgJSONError) // ErrJSONError is JSON Error Error object
)

// Schema struct
type Schema struct {
	vjson.Schema
}

// NewSchema function which returns Schema struct
func NewSchema(fields ...vjson.Field) Schema {
	return Schema{vjson.NewSchema(fields...)}
}

// ValidateBytes function which validates byte slice
func (s *Schema) ValidateBytes(input []byte) error {
	if err := s.Schema.ValidateBytes(input); err != nil {
		if err.Error() == errMsgJSONError {
			return ErrJSONError
		}

		return err
	}

	return nil
}

// ValidateString function which validates string
func (s *Schema) ValidateString(input string) error {
	if err := s.Schema.ValidateString(input); err != nil {
		if err.Error() == errMsgJSONError {
			return ErrJSONError
		}

		return err
	}

	return nil
}
