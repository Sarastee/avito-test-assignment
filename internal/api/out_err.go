package api

import "errors"

const (
	errMsgInternalError           = "something went wrong, we are already working on it"
	errMsgFieldNotFound           = "field not found"
	errMsgTagOrFeatureNotProvided = "tag_id or feature_id weren't provided"
	errMsgTagIsNotANumber         = "tag_id is not a number"
	errMsgFeatureIsNotANumber     = "feature_id is not a number"

	errMsgIncorrectRole = "provided role not correct, use ADMIN or USER"
)

var (
	ErrInternalError           = errors.New(errMsgInternalError)           // ErrInternalError is Internal Server Error object
	ErrFieldNotFound           = errors.New(errMsgFieldNotFound)           // ErrFieldNotFound is Field Not Found Error object
	ErrTagOrFeatureNotProvided = errors.New(errMsgTagOrFeatureNotProvided) // ErrTagOrFeatureNotProvided is Tag Or Feature Not Provided Error object
	ErrTagIsNotANumber         = errors.New(errMsgTagIsNotANumber)         // ErrTagIsNotANumber is Tag Is Not A Number Error object
	ErrFeatureIDIsNotANumber   = errors.New(errMsgFeatureIsNotANumber)     // ErrFeatureIDIsNotANumber is Feature ID Is Not A Number Error object
	ErrIncorrectRole           = errors.New(errMsgIncorrectRole)           // ErrIncorrectRole is Incorrect Role Error object
)