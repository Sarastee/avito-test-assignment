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

type Schema struct {
	vjson.Schema
}

func NewSchema(fields ...vjson.Field) Schema {
	return Schema{vjson.NewSchema(fields...)}
}

func (s *Schema) ValidateBytes(input []byte) error {
	if err := s.Schema.ValidateBytes(input); err != nil {
		if err.Error() == errMsgJSONError {
			return ErrJSONError
		}

		return err
	}

	return nil
}

func (s *Schema) ValidateString(input string) error {
	if err := s.Schema.ValidateString(input); err != nil {
		if err.Error() == errMsgJSONError {
			return ErrJSONError
		}

		return err
	}

	return nil
}
