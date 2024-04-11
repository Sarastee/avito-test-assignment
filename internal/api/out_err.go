package api

import "errors"

const (
	errMsgInternalError = "something went wrong, we are already working on it"
	errMsgFieldNotFound = "field not found"
)

var (
	ErrInternalError = errors.New(errMsgInternalError) // ErrInternalError is Internal Server Error object
	ErrFieldNotFound = errors.New(errMsgFieldNotFound) // ErrFieldNotFound is Field Not Found Error object
)
